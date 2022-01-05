package port_repository

import (
	"errors"
	"time"

	repository_intf "github.com/herryg91/dply/dply-server/app/repository"
	"github.com/herryg91/dply/dply-server/entity"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type repository struct {
	db                  *gorm.DB
	table_port          string
	table_port_template string
}

func New(db *gorm.DB) repository_intf.PortRepository {
	return &repository{db, "port", "port_template"}
}

func (r *repository) Get(project, env, name string) (*entity.Port, error) {
	portModel := &PortModel{}
	err := r.db.Table(r.table_port).Where("project = ? AND env = ? AND name = ?", project, env, name).First(&portModel).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, repository_intf.ErrPortNotFound
		}
		return nil, err
	}
	return portModel.ToPortEntity(), nil
}

func (r *repository) Upsert(data entity.Port) error {
	timeNow := time.Now().UTC()

	portModel := PortModel{}.FromPortEntity(data)
	portModel.UpdatedAt = &timeNow
	portModel.CreatedAt = &timeNow

	err := r.db.Table(r.table_port).Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{
			"ports", "updated_at", "created_by",
		}),
	}).Create(&portModel).Error
	if err != nil {
		return err
	}

	return nil
}
func (r *repository) GetPortByTemplate(templateName string) (*entity.PortTemplate, error) {
	portTemplateModel := &PortTemplateModel{}
	err := r.db.Table(r.table_port_template).Where("template_name = ?", templateName).First(&portTemplateModel).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, repository_intf.ErrPortTemplateNotFound
		}
		return nil, err
	}
	return portTemplateModel.ToPortTemplateEntity(), nil
}
func (r *repository) UpsertPortByTemplate(data entity.PortTemplate) error {
	timeNow := time.Now().UTC()

	portTemplateModel := PortTemplateModel{}.FromPortTemplateEntity(data)
	portTemplateModel.UpdatedAt = &timeNow
	portTemplateModel.CreatedAt = &timeNow

	err := r.db.Table(r.table_port_template).Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{
			"ports", "updated_at",
		}),
	}).Create(&portTemplateModel).Error
	if err != nil {
		return err
	}

	return nil
}
