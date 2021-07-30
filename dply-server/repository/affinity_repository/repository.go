package affinity_repository

import (
	"errors"
	"time"

	repository_intf "github.com/herryg91/dply/dply-server/app/repository"
	"github.com/herryg91/dply/dply-server/entity"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type repository struct {
	db                      *gorm.DB
	table_affinity          string
	table_affinity_template string
}

func New(db *gorm.DB) repository_intf.AffinityRepository {
	return &repository{db, "affinity", "affinity_template"}
}

func (r *repository) Get(env, name string) (*entity.Affinity, error) {
	affinityModel := &AffinityModel{}
	err := r.db.Table(r.table_affinity).Where("env = ? AND name = ?", env, name).First(&affinityModel).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, repository_intf.ErrAffinityNotFound
		}
		return nil, err
	}
	return affinityModel.ToAffinityEntity(), nil
}

func (r *repository) Upsert(data entity.Affinity) error {
	timeNow := time.Now().UTC()

	affinityModel := AffinityModel{}.FromAffinityEntity(data)
	affinityModel.UpdatedAt = &timeNow
	affinityModel.CreatedAt = &timeNow

	err := r.db.Table(r.table_affinity).Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{
			"affinity", "updated_at", "created_by",
		}),
	}).Create(&affinityModel).Error
	if err != nil {
		return err
	}

	return nil
}
func (r *repository) GetAffinityByTemplate(templateName string) (*entity.AffinityTemplate, error) {
	affinityTemplateModel := &AffinityTemplateModel{}
	err := r.db.Table(r.table_affinity_template).Where("template_name = ?", templateName).First(&affinityTemplateModel).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, repository_intf.ErrAffinityTemplateNotFound
		}
		return nil, err
	}
	return affinityTemplateModel.ToAffinityTemplateEntity(), nil
}

func (r *repository) UpsertAffinityByTemplate(data entity.AffinityTemplate) error {
	timeNow := time.Now().UTC()

	affinityTemplateModel := AffinityTemplateModel{}.FromAffinityTemplateEntity(data)
	affinityTemplateModel.UpdatedAt = &timeNow
	affinityTemplateModel.CreatedAt = &timeNow

	err := r.db.Table(r.table_affinity_template).Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{
			"affinity", "updated_at",
		}),
	}).Create(&affinityTemplateModel).Error
	if err != nil {
		return err
	}

	return nil
}
