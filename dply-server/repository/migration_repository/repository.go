package migration_repository

import (
	"errors"
	"time"

	repository_intf "github.com/herryg91/dply/dply-server/app/repository"
	"github.com/herryg91/dply/dply-server/entity"

	"gorm.io/gorm"
)

type repository struct {
	db    *gorm.DB
	table string
}

func New(db *gorm.DB) repository_intf.MigrationRepository {
	return &repository{db, "migrations"}
}

func (r *repository) CreateTable() error {
	err := r.db.Exec(`
		CREATE TABLE `+r.table+` (
			id INT NOT NULL AUTO_INCREMENT,
			name VARCHAR(128) NOT NULL UNIQUE,
			created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			PRIMARY KEY ( id )
		);
	`, r.table).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) IsTableExist() (bool, error) {
	var count int64
	err := r.db.Raw("SELECT count(*) FROM information_schema.tables  WHERE table_schema = 'dply' AND table_name = '" + r.table + "'").Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *repository) Get() ([]entity.Migration, error) {
	mm := []*MigrationModel{}
	err := r.db.Table(r.table).Find(&mm).Error
	if err != nil {
		return []entity.Migration{}, err
	}
	resp := []entity.Migration{}
	for _, m := range mm {
		resp = append(resp, *m.ToMigrationEntity())
	}

	return resp, nil
}

func (r *repository) GetLast() (*entity.Migration, error) {
	mm := MigrationModel{}

	err := r.db.Table(r.table).Order("name DESC").First(&mm).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return mm.ToMigrationEntity(), nil
}

func (r *repository) Create(req entity.Migration) error {
	mm := MigrationModel{}.FromMigrationEntity(req)
	timeNow := time.Now().UTC()
	mm.CreatedAt = &timeNow
	mm.UpdatedAt = &timeNow
	err := r.db.Table(r.table).Create(&mm).Error
	if err != nil {
		return err
	}
	return nil
}
