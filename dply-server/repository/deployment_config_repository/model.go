package deployment_config_repository

import (
	"encoding/json"
	"time"

	"github.com/herryg91/dply/dply-server/entity"
)

type DeploymentConfigModel struct {
	Id        int        `gorm:"column:id"`
	Project   string     `gorm:"column:project"`
	Env       string     `gorm:"column:env"`
	Name      string     `gorm:"column:name"`
	Config    []byte     `gorm:"column:config"`
	CreatedBy int        `gorm:"column:created_by"`
	CreatedAt *time.Time `gorm:"column:created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at"`
}

func (am *DeploymentConfigModel) ToDeploymentConfigEntity() *entity.DeploymentConfig {
	if am == nil {
		return nil
	}

	dc := &entity.DeploymentConfig{
		Project:   am.Project,
		Env:       am.Env,
		Name:      am.Name,
		CreatedBy: am.CreatedBy,
	}
	json.Unmarshal(am.Config, &dc)
	return dc
}

func (DeploymentConfigModel) FromDeploymentConfigEntity(a entity.DeploymentConfig) *DeploymentConfigModel {
	dc := map[string]interface{}{
		"liveness_probe":  a.LivenessProbe,
		"readiness_probe": a.ReadinessProbe,
		"startup_probe":   a.StartupProbe,
	}
	dcJson, _ := json.Marshal(&dc)

	return &DeploymentConfigModel{
		Project:   a.Project,
		Env:       a.Env,
		Name:      a.Name,
		Config:    dcJson,
		CreatedBy: a.CreatedBy,
	}
}
