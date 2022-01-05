package image_repository

import (
	"errors"
	"strings"
	"time"

	"github.com/go-sql-driver/mysql"
	repository_intf "github.com/herryg91/dply/dply-server/app/repository"
	"github.com/herryg91/dply/dply-server/entity"

	"gorm.io/gorm"
)

type repository struct {
	db    *gorm.DB
	table string
}

func New(db *gorm.DB) repository_intf.ImageRepository {
	return &repository{db, "image"}
}

func (r *repository) Search(project, repositoryName string, limit, offset int, createdAtDesc bool) ([]entity.Image, error) {
	resp := []entity.Image{}
	orderByStr := "DESC"
	if !createdAtDesc {
		orderByStr = "ASC"
	}

	err := r.db.Table(r.table).Where("project = ? AND repository = ?", project, repositoryName).
		Limit(limit).Offset(offset).
		Order("created_at " + orderByStr).
		Find(&resp).Error
	if err != nil {
		return []entity.Image{}, err
	}
	return resp, nil
}

func (r *repository) GetByDigest(digest string) (*entity.Image, error) {
	resp := &ImageModel{}
	err := r.db.Table(r.table).Where("digest = ?", digest).First(&resp).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, repository_intf.ErrImageNotFound
		}
		return nil, err
	}
	return resp.ToImageEntity(), nil
}

func (r *repository) GetByImage(image string) (*entity.Image, error) {
	resp := &ImageModel{}
	err := r.db.Table(r.table).Where("image = ?", image).First(&resp).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, repository_intf.ErrImageNotFound
		}
		return nil, err
	}
	return resp.ToImageEntity(), nil
}

func (r *repository) Create(req entity.Image) error {
	timeNow := time.Now().UTC()
	cim := ImageModel{}.FromImageEntity(req)
	cim.CreatedAt = &timeNow
	cim.UpdatedAt = &timeNow

	err := r.db.Table(r.table).Create(&cim).Error
	if err != nil {
		mysqlErr := &mysql.MySQLError{}
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 && (strings.Contains(mysqlErr.Message, "digest")) {
			return repository_intf.ErrImageDigestDuplicate
		}
		return err
	}
	return nil
}

func (r *repository) Delete(digest string) error {
	return r.db.Table(r.table).Where("digest = ?", digest).Delete(&ImageModel{}).Error
}
