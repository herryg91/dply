package deployment_repository

import (
	"encoding/json"
	"time"

	"github.com/herryg91/dply/dply-server/entity"
)

type DeploymentModel struct {
	Id          int        `gorm:"column:id"`
	Project     string     `gorm:"column:project"`
	Env         string     `gorm:"column:env"`
	Name        string     `gorm:"column:name"`
	ImageDigest string     `gorm:"column:image_digest"`
	Variables   []byte     `gorm:"column:variables"`
	Ports       []byte     `gorm:"column:ports"`
	CreatedBy   int        `gorm:"column:created_by"`
	CreatedAt   *time.Time `gorm:"column:created_at"`
	UpdatedAt   *time.Time `gorm:"column:updated_at"`
}

func (dm *DeploymentModel) ToDeploymentEntity() *entity.Deployment {
	if dm == nil {
		return nil
	}
	variables := map[string]interface{}{}
	ports := &entity.Port{
		Project: dm.Project,
		Env:     dm.Env,
		Name:    dm.Name,
	}
	json.Unmarshal(dm.Variables, &variables)
	json.Unmarshal(dm.Ports, &ports)
	return &entity.Deployment{
		Id:      dm.Id,
		Project: dm.Project,
		Env:     dm.Env,
		Name:    dm.Name,
		DeploymentImage: entity.Image{
			Project: dm.Project,
			Digest:  dm.ImageDigest,
		},
		Envar: entity.Envar{
			Project:   dm.Project,
			Env:       dm.Env,
			Name:      dm.Name,
			Variables: variables,
			CreatedBy: dm.CreatedBy,
		},
		Port:      *ports,
		Scale:     entity.Scale{},
		Affinity:  entity.Affinity{},
		CreatedBy: dm.CreatedBy,
	}
}
