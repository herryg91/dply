package config_repository

import (
	"encoding/json"
	"time"

	"github.com/herryg91/dply/dply-server/entity"
)

type ConfigModel struct {
	Id        int        `gorm:"column:id"`
	Env       string     `gorm:"column:env"`
	Name      string     `gorm:"column:name"`
	Variables []byte     `gorm:"column:variables"`
	Ports     []byte     `gorm:"column:ports"`
	Affinity  []byte     `gorm:"column:affinity"`
	CreatedBy int        `gorm:"column:created_by"`
	CreatedAt *time.Time `gorm:"column:created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at"`
}

func (c *ConfigModel) ToConfigEntity() *entity.Config {
	if c == nil {
		return nil
	}

	out := &entity.Config{
		Id:        c.Id,
		Env:       c.Env,
		Name:      c.Name,
		CreatedBy: c.CreatedBy,
	}
	json.Unmarshal(c.Variables, &out.Variables)
	json.Unmarshal(c.Ports, &out.Ports)
	json.Unmarshal(c.Affinity, &out.Affinity)
	return out
}
