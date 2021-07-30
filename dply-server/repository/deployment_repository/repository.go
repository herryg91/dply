package deployment_repository

import (
	"encoding/json"
	"errors"
	"time"

	repository_intf "github.com/herryg91/dply/dply-server/app/repository"
	"github.com/herryg91/dply/dply-server/entity"

	"gorm.io/gorm"
)

type repository struct {
	db    *gorm.DB
	table string
}

func New(db *gorm.DB) repository_intf.DeploymentRepository {
	return &repository{db, "deployment"}
}

func (r *repository) Get(env, name string) (*entity.Deployment, error) {
	deploymentModel := &DeploymentModel{}
	err := r.db.Table(r.table).Where("env = ? AND name = ?", env, name).Order("id desc").Limit(1).First(&deploymentModel).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, repository_intf.ErrDeploymentNotFound
		}
		return nil, err
	}
	return deploymentModel.ToDeploymentEntity(), nil
}

func (r *repository) Create(in entity.Deployment) error {
	variables, _ := json.Marshal(in.Envar.Variables)
	ports, _ := json.Marshal(in.Port.Ports)

	timeNow := time.Now().UTC()
	deploymentModel := &DeploymentModel{
		Env:         in.Env,
		Name:        in.Name,
		ImageDigest: in.DeploymentImage.Digest,
		Variables:   variables,
		Ports:       ports,
		CreatedBy:   in.CreatedBy,
		CreatedAt:   &timeNow,
		UpdatedAt:   &timeNow,
	}
	err := r.db.Table(r.table).Create(&deploymentModel).Error
	if err != nil {
		return err
	}
	return nil
}
