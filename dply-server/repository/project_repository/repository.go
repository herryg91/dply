package project_repository

import (
	"errors"

	"github.com/go-sql-driver/mysql"
	repository_intf "github.com/herryg91/dply/dply-server/app/repository"
	"github.com/herryg91/dply/dply-server/entity"

	"gorm.io/gorm"
)

type repository struct {
	db    *gorm.DB
	table string
}

func New(db *gorm.DB) repository_intf.ProjectRepository {
	return &repository{db, "project"}
}

func (r *repository) GetAll() ([]entity.Project, error) {
	projectModels := []ProjectModel{}
	err := r.db.Table(r.table).Find(&projectModels).Error
	if err != nil {
		return nil, err
	}

	res := []entity.Project{}
	for _, pm := range projectModels {
		res = append(res, *pm.ToProjectEntity())
	}
	return res, nil
}
func (r *repository) Create(p entity.Project) error {
	err := r.db.Table(r.table).Create(&ProjectModel{
		Name:        p.Name,
		Description: p.Description,
	}).Error
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			return repository_intf.ErrProjectDuplicate
		}

		return err
	}

	return nil
}
func (r *repository) DeleteByName(name string) error {
	err := r.db.Table(r.table).Where("name = ?", name).Delete(&ProjectModel{}).Error
	if err != nil {
		return err
	}

	return nil
}
