package repository

import "github.com/harukitosa/ddd_sample/domain/model"

type UserRepository interface {
	GetByID(id uint64) (model.User, error)
	Save(user *model.User) (*model.User, error)
	Update(user *model.User) error
	Delete(id uint64) error
}
