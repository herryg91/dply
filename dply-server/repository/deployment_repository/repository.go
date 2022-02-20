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

func (r *repository) Get(project, env, name string) (*entity.Deployment, error) {
	deploymentModel := &DeploymentModel{}
	err := r.db.Table(r.table).Where("project = ? AND env = ? AND name = ?", project, env, name).Order("id desc").Limit(1).First(&deploymentModel).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, repository_intf.ErrDeploymentNotFound
		}
		return nil, err
	}
	return deploymentModel.ToDeploymentEntity(), nil
}

func (r *repository) GetLatestDeploymentGroupByName(project, name string) ([]*entity.Deployment, error) {
	deployments := []*DeploymentModel{}
	err := r.db.Raw(`
		SELECT d1.*
		FROM deployment d1
		LEFT JOIN deployment d2 ON (d1.env = d2.env AND d1.name = d2.name AND d1.id < d2.id)
		WHERE d1.project = ?
		AND d1.name = ?
		AND d2.id IS NULL
	`, project, name).Find(&deployments).Error
	if err != nil {
		return nil, err
	}
	resp := []*entity.Deployment{}
	for _, d := range deployments {
		resp = append(resp, d.ToDeploymentEntity())
	}

	return resp, nil
}

func (r *repository) Create(in entity.Deployment) error {
	variables, _ := json.Marshal(in.Envar.Variables)
	portsJson := map[string]interface{}{
		"ports":       in.Port.Ports,
		"access_type": in.Port.AccessType,
		"external_ip": in.Port.ExternalIP,
	}
	portsJsonMarshalled, _ := json.Marshal(&portsJson)

	timeNow := time.Now().UTC()
	deploymentModel := &DeploymentModel{
		Project:     in.Project,
		Env:         in.Env,
		Name:        in.Name,
		ImageDigest: in.DeploymentImage.Digest,
		Variables:   variables,
		Ports:       portsJsonMarshalled,
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
