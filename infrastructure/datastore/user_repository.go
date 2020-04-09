package datastore

import (
	"github.com/harukitosa/ddd_sample/domain/model"
	"github.com/harukitosa/ddd_sample/repository"
	"github.com/jinzhu/gorm"
)

type UserRepositoryImpliment struct {
	DB *gorm.DB
}

func NewUserRepositoryImpliment(DB *gorm.DB) repository.IUserRepository {
	return &UserRepositoryImpliment{
		DB: DB,
	}
}

func (r *UserRepositoryImpliment) GetByID(id uint64) (model.User, error) {
	var user model.User
	r.DB.Where("id = ?", id).Find(&user)
	return user, nil
}

func (r *UserRepositoryImpliment) Save(user *model.User) (*model.User, error) {
	r.DB.Create(&user)
	return user, nil
}

func (r *UserRepositoryImpliment) Update(user *model.User) error {
	r.DB.Save(&user)
	return nil
}

func (r *UserRepositoryImpliment) Delete(id uint64) error {
	r.DB.Where("id = ?", id).Delete(&model.User{})
	return nil
}
