package scale_repository

import (
	"errors"
	"time"

	repository_intf "github.com/herryg91/dply/dply-server/app/repository"
	"github.com/herryg91/dply/dply-server/entity"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type repository struct {
	db    *gorm.DB
	table string
}

func New(db *gorm.DB) repository_intf.ScaleRepository {
	return &repository{db, "scale"}
}

func (r *repository) Get(env, name string) (*entity.Scale, error) {
	scaleModel := &ScaleModel{}
	err := r.db.Table(r.table).Where("env = ? AND name = ?", env, name).First(&scaleModel).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, repository_intf.ErrScaleNotFound
		}
		return nil, err
	}
	return scaleModel.ToScaleEntity(), nil
}

func (r *repository) Upsert(data entity.Scale) error {
	timeNow := time.Now().UTC()

	scaleModel := ScaleModel{}.FromScaleEntity(data)
	scaleModel.UpdatedAt = &timeNow
	scaleModel.CreatedAt = &timeNow

	err := r.db.Table(r.table).Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{
			"min_replica", "max_replica", "min_cpu", "max_cpu", "min_memory", "max_memory", "target_cpu", "updated_at", "created_by",
		}),
	}).Create(&scaleModel).Error
	if err != nil {
		return err
	}

	return nil
}
