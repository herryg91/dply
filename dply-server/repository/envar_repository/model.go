package envar_repository

import (
	"encoding/json"
	"time"

	"github.com/herryg91/dply/dply-server/entity"
)

type EnvarModel struct {
	Id        int        `gorm:"column:id"`
	Project   string     `gorm:"column:project"`
	Env       string     `gorm:"column:env"`
	Name      string     `gorm:"column:name"`
	Variables []byte     `gorm:"column:variables"`
	CreatedBy int        `gorm:"column:created_by"`
	CreatedAt *time.Time `gorm:"column:created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at"`
}

func (em *EnvarModel) ToEnvarEntity() *entity.Envar {
	if em == nil {
		return nil
	}

	variables := map[string]interface{}{}
	json.Unmarshal(em.Variables, &variables)
	return &entity.Envar{
		Project:   em.Project,
		Env:       em.Env,
		Name:      em.Name,
		Variables: variables,
		CreatedBy: em.CreatedBy,
	}
}

func (EnvarModel) FromEnvarEntity(e entity.Envar) *EnvarModel {
	variables, _ := json.Marshal(&e.Variables)
	return &EnvarModel{
		Project:   e.Project,
		Env:       e.Env,
		Name:      e.Name,
		Variables: variables,
		CreatedBy: e.CreatedBy,
	}
}
