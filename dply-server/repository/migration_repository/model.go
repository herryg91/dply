package migration_repository

import (
	"time"

	"github.com/herryg91/dply/dply-server/entity"
)

type MigrationModel struct {
	Id        int        `gorm:"column:id"`
	Name      string     `gorm:"column:name"`
	CreatedAt *time.Time `gorm:"column:created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at"`
}

func (mm *MigrationModel) ToMigrationEntity() *entity.Migration {
	if mm == nil {
		return nil
	}
	return &entity.Migration{
		Id:        mm.Id,
		Name:      mm.Name,
		CreatedAt: mm.CreatedAt,
		UpdatedAt: mm.UpdatedAt,
	}
}

func (MigrationModel) FromMigrationEntity(m entity.Migration) *MigrationModel {
	return &MigrationModel{
		Id:        m.Id,
		Name:      m.Name,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}
