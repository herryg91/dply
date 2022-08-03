package deployment_config_repository

import (
	"errors"
	"time"

	repository_intf "github.com/herryg91/dply/dply-server/app/repository"
	"github.com/herryg91/dply/dply-server/entity"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) repository_intf.DeploymentConfigRepository {
	return &repository{db}
}

func (r *repository) Get(project, env, name string) (*entity.DeploymentConfig, error) {
	dcModel := &DeploymentConfigModel{}
	err := r.db.Table("deployment_config").Where("project = ? AND env = ? AND name = ?", project, env, name).First(&dcModel).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, repository_intf.ErrDeploymentConfigNotFound
		}
		return nil, err
	}
	return dcModel.ToDeploymentConfigEntity(), nil
}

func (r *repository) Upsert(data entity.DeploymentConfig) error {
	timeNow := time.Now().UTC()

	dcModel := DeploymentConfigModel{}.FromDeploymentConfigEntity(data)
	dcModel.UpdatedAt = &timeNow
	dcModel.CreatedAt = &timeNow

	err := r.db.Table("deployment_config").Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{
			"config", "updated_at", "created_by",
		}),
	}).Create(&dcModel).Error
	if err != nil {
		return err
	}

	return nil
}
