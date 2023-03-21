package affinity_repository

import (
	"encoding/json"
	"time"

	"github.com/herryg91/dply/dply-server/entity"
)

type AffinityModel struct {
	Id        int        `gorm:"column:id"`
	Project   string     `gorm:"column:project"`
	Env       string     `gorm:"column:env"`
	Name      string     `gorm:"column:name"`
	Affinity  []byte     `gorm:"column:affinity"`
	CreatedBy int        `gorm:"column:created_by"`
	CreatedAt *time.Time `gorm:"column:created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at"`
}

func (am *AffinityModel) ToAffinityEntity() *entity.Affinity {
	if am == nil {
		return nil
	}

	affinity := &entity.Affinity{
		Project:   am.Project,
		Env:       am.Env,
		Name:      am.Name,
		CreatedBy: am.CreatedBy,
	}
	json.Unmarshal(am.Affinity, &affinity)
	return affinity
}

func (AffinityModel) FromAffinityEntity(a entity.Affinity) *AffinityModel {
	affinity := map[string]interface{}{
		"node_affinity":     a.NodeAffinity,
		"pod_affinity":      a.PodAffinity,
		"pod_anti_affinity": a.PodAntiAffinity,
		"tolerations":       a.Tolerations,
	}
	affinityJson, _ := json.Marshal(&affinity)

	return &AffinityModel{
		Project:   a.Project,
		Env:       a.Env,
		Name:      a.Name,
		Affinity:  affinityJson,
		CreatedBy: a.CreatedBy,
	}
}

type AffinityTemplateModel struct {
	Id           int        `gorm:"column:id"`
	TemplateName string     `gorm:"column:template_name"`
	Affinity     []byte     `gorm:"column:affinity"`
	CreatedAt    *time.Time `gorm:"column:created_at"`
	UpdatedAt    *time.Time `gorm:"column:updated_at"`
}

func (am *AffinityTemplateModel) ToAffinityTemplateEntity() *entity.AffinityTemplate {
	if am == nil {
		return nil
	}

	resp := &entity.AffinityTemplate{
		TemplateName:    am.TemplateName,
		NodeAffinity:    []entity.AffinityTerm{},
		PodAffinity:     []entity.AffinityTerm{},
		PodAntiAffinity: []entity.AffinityTerm{},
		Tolerations:     []entity.AffinityToleration{},
	}

	json.Unmarshal(am.Affinity, &resp)

	return resp
}
func (AffinityTemplateModel) FromAffinityTemplateEntity(a entity.AffinityTemplate) *AffinityTemplateModel {
	affinity := map[string]interface{}{
		"node_affinity":     a.NodeAffinity,
		"pod_affinity":      a.PodAffinity,
		"pod_anti_affinity": a.PodAntiAffinity,
		"tolerations":       a.Tolerations,
	}
	affinityJson, _ := json.Marshal(&affinity)
	return &AffinityTemplateModel{
		TemplateName: a.TemplateName,
		Affinity:     affinityJson,
	}
}
