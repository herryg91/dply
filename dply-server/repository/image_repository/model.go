package image_repository

import (
	"time"

	"github.com/herryg91/dply/dply-server/entity"
)

type ImageModel struct {
	Id          int    `gorm:"column:id"`
	Digest      string `gorm:"column:digest"`
	Image       string `gorm:"column:image"`
	Repository  string `gorm:"column:repository"`
	Description string `gorm:"column:description"`
	CreatedBy   int    `gorm:"column:created_by"`

	CreatedAt *time.Time `gorm:"column:created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at"`
}

func (im *ImageModel) ToImageEntity() *entity.Image {
	if im == nil {
		return nil
	}
	return &entity.Image{
		Id:          im.Id,
		Digest:      im.Digest,
		Image:       im.Image,
		Repository:  im.Repository,
		Description: im.Description,
		CreatedBy:   im.CreatedBy,
		CreatedAt:   im.CreatedAt,
	}
}

func (ImageModel) FromImageEntity(i entity.Image) *ImageModel {
	return &ImageModel{
		Id:          i.Id,
		Digest:      i.Digest,
		Image:       i.Image,
		Repository:  i.Repository,
		Description: i.Description,
		CreatedBy:   i.CreatedBy,
	}
}
