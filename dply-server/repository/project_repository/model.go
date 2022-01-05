package project_repository

import (
	"time"

	"github.com/herryg91/dply/dply-server/entity"
)

type ProjectModel struct {
	Id          int        `gorm:"column:id"`
	Name        string     `gorm:"column:name"`
	Description string     `gorm:"column:description"`
	CreatedAt   *time.Time `gorm:"column:created_at"`
	UpdatedAt   *time.Time `gorm:"column:updated_at"`
}

func (pm *ProjectModel) ToProjectEntity() *entity.Project {
	if pm == nil {
		return nil
	}

	return &entity.Project{
		Id:          pm.Id,
		Name:        pm.Name,
		Description: pm.Description,
	}
}

func (ProjectModel) FromProjectEntity(p entity.Project) *ProjectModel {
	return &ProjectModel{
		Id:          p.Id,
		Name:        p.Name,
		Description: p.Description,
	}
}
