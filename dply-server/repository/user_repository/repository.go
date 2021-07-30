package user_repository

import (
	"errors"
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

func New(db *gorm.DB) repository_intf.UserRepository {
	return &repository{db, "user"}
}

func (r *repository) GetById(id int) (*entity.User, error) {
	userModel := &UserModel{}
	err := r.db.Table(r.table).Where("id = ?", id).First(&userModel).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, repository_intf.ErrUserNotFound
		}
		return nil, err
	}

	if !userModel.Status {
		return nil, repository_intf.ErrUserInactive
	}
	return userModel.ToUserEntity(), nil
}

func (r *repository) GetByEmail(email string) (*entity.User, error) {
	userModel := &UserModel{}
	err := r.db.Table(r.table).Where("email = ?", email).First(&userModel).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, repository_intf.ErrUserNotFound
		}
		return nil, err
	}

	if !userModel.Status {
		return nil, repository_intf.ErrUserInactive
	}
	return userModel.ToUserEntity(), nil
}
func (r *repository) GetByToken(token string) (*entity.User, error) {
	userModel := &UserModel{}
	err := r.db.Table(r.table).Where("token = ?", token).First(&userModel).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, repository_intf.ErrUserNotFound
		}
		return nil, err
	}
	if !userModel.Status {
		return nil, repository_intf.ErrUserInactive
	}
	return userModel.ToUserEntity(), nil
}
func (r *repository) Create(req entity.User) error {
	timeNow := time.Now().UTC()
	data := UserModel{}.FromUserEntity(req)
	data.CreatedAt = &timeNow
	data.UpdatedAt = &timeNow
	data.Status = true

	err := r.db.Table(r.table).Create(&data).Error
	if err != nil {
		mysqlErr := &mysql.MySQLError{}
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			return repository_intf.ErrUserDuplicate
		}
		return err
	}
	return nil
}

func (r *repository) Edit(req entity.User) error {
	timeNow := time.Now().UTC()
	data := UserModel{}.FromUserEntity(req)
	data.UpdatedAt = &timeNow
	data.Id = 0
	err := r.db.Table(r.table).Where("email = ?", req.Email).Updates(&data).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) EditStatus(email string, status bool) error {
	timeNow := time.Now().UTC()

	err := r.db.Table(r.table).Where("email = ?", email).Updates(map[string]interface{}{
		"status":     status,
		"updated_at": &timeNow,
	}).Error
	if err != nil {
		return err
	}
	return nil
}
func (r *repository) EditPassword(email, password, newToken string) error {
	timeNow := time.Now().UTC()

	err := r.db.Table(r.table).Where("email = ?", email).Updates(map[string]interface{}{
		"password":   password,
		"token":      newToken,
		"updated_at": &timeNow,
	}).Error
	if err != nil {
		return err
	}
	return nil
}
