package scale_repository

import (
	"time"

	"github.com/herryg91/dply/dply-server/entity"
)

type ScaleModel struct {
	Id                   int        `gorm:"column:id"`
	Project              string     `gorm:"column:project"`
	Env                  string     `gorm:"column:env"`
	Name                 string     `gorm:"column:name"`
	MinReplica           int        `gorm:"column:min_replica"`
	MaxReplica           int        `gorm:"column:max_replica"`
	MinCpu               int        `gorm:"column:min_cpu"`
	MaxCpu               int        `gorm:"column:max_cpu"`
	MinMemory            int        `gorm:"column:min_memory"`
	MaxMemory            int        `gorm:"column:max_memory"`
	TargetCPUUtilization int        `gorm:"column:target_cpu"`
	CreatedBy            int        `gorm:"column:created_by"`
	CreatedAt            *time.Time `gorm:"column:created_at"`
	UpdatedAt            *time.Time `gorm:"column:updated_at"`
}

func (sm *ScaleModel) ToScaleEntity() *entity.Scale {
	if sm == nil {
		return nil
	}
	return &entity.Scale{
		Project:              sm.Project,
		Env:                  sm.Env,
		Name:                 sm.Name,
		MinReplica:           sm.MinReplica,
		MaxReplica:           sm.MaxReplica,
		MinCpu:               sm.MinCpu,
		MaxCpu:               sm.MaxCpu,
		MinMemory:            sm.MinMemory,
		MaxMemory:            sm.MaxMemory,
		TargetCPUUtilization: sm.TargetCPUUtilization,
		CreatedBy:            sm.CreatedBy,
	}
}

func (ScaleModel) FromScaleEntity(s entity.Scale) *ScaleModel {
	return &ScaleModel{
		Project:              s.Project,
		Env:                  s.Env,
		Name:                 s.Name,
		MinReplica:           s.MinReplica,
		MaxReplica:           s.MaxReplica,
		MinCpu:               s.MinCpu,
		MaxCpu:               s.MaxCpu,
		MinMemory:            s.MinMemory,
		MaxMemory:            s.MaxMemory,
		TargetCPUUtilization: s.TargetCPUUtilization,
		CreatedBy:            s.CreatedBy,
	}
}
