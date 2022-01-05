package envar_repository

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

func New(db *gorm.DB) repository_intf.EnvarRepository {
	return &repository{db, "envar"}
}

func (r *repository) Get(project, env, name string) (*entity.Envar, error) {
	envarModel := &EnvarModel{}
	err := r.db.Table(r.table).Where("project = ? AND env = ? AND name = ?", project, env, name).First(&envarModel).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, repository_intf.ErrEnvarNotFound
		}
		return nil, err
	}
	return envarModel.ToEnvarEntity(), nil
}

func (r *repository) Upsert(data entity.Envar) error {
	timeNow := time.Now().UTC()

	envarModel := EnvarModel{}.FromEnvarEntity(data)
	envarModel.UpdatedAt = &timeNow
	envarModel.CreatedAt = &timeNow

	err := r.db.Table(r.table).Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{
			"variables", "updated_at", "created_by",
		}),
	}).Create(&envarModel).Error
	if err != nil {
		return err
	}

	return nil
}
