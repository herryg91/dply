package user_repository

import (
	"time"

	"github.com/herryg91/dply/dply-server/entity"
)

type UserModel struct {
	Id        int        `gorm:"column:id"`
	Email     string     `gorm:"column:email"`
	Password  string     `gorm:"column:password"`
	UserType  string     `gorm:"column:usertype"`
	Name      string     `gorm:"column:name"`
	Token     string     `gorm:"column:token"`
	Status    bool       `gorm:"column:status"`
	CreatedAt *time.Time `gorm:"column:created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at"`
}

func (um *UserModel) ToUserEntity() *entity.User {
	if um == nil {
		return nil
	}
	return &entity.User{
		Id:       um.Id,
		Email:    um.Email,
		Password: um.Password,
		UserType: entity.UserType(um.UserType),
		Name:     um.Name,
		Token:    um.Token,
	}
}

func (UserModel) FromUserEntity(u entity.User) *UserModel {
	return &UserModel{
		Id:       u.Id,
		Email:    u.Email,
		Password: u.Password,
		UserType: string(u.UserType),
		Name:     u.Name,
		Token:    u.Token,
	}
}
