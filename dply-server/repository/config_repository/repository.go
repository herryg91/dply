package config_repository

import (
	"encoding/json"
	"errors"
	"time"

	repository_intf "github.com/herryg91/dply/dply-server/app/repository"
	"github.com/herryg91/dply/dply-server/entity"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type repository struct {
	db    *gorm.DB
	table string
}

func New(db *gorm.DB) repository_intf.ConfigRepository {
	return &repository{db, "config"}
}

func (r *repository) Get(env, name string) (*entity.Config, error) {
	c := &ConfigModel{}
	err := r.db.Table(r.table).Where("env = ? AND name = ?", env, name).First(&c).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, repository_intf.ErrConfigNotFound
		}
		return nil, err
	}

	return c.ToConfigEntity(), nil
}
func (r *repository) Upsert(param entity.Config) error {
	timeNow := time.Now().UTC()
	variablesJson, _ := json.Marshal(&param.Variables)
	portsJson, _ := json.Marshal(&param.Ports)
	affinityJson, _ := json.Marshal(&param.Affinity)

	c := ConfigModel{
		Env:       param.Env,
		Name:      param.Name,
		Variables: variablesJson,
		Ports:     portsJson,
		Affinity:  affinityJson,
		CreatedBy: param.CreatedBy,
		CreatedAt: &timeNow,
		UpdatedAt: &timeNow,
	}

	err := r.db.Table(r.table).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"variables", "ports", "affinity", "created_by"}),
	}).Create(&c).Error
	if err != nil {
		return err
	}
	return nil
}
