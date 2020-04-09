package application

import (
	"time"

	"github.com/harukitosa/ddd_sample/domain/model"
	"github.com/harukitosa/ddd_sample/repository"
)

type UserService struct {
	UserRepository repository.IUserRepository
}

func NewUserService(repository repository.IUserRepository) UserService {
	return UserService{UserRepository: repository}
}

func (userService *UserService) GetUser(userID uint64) (model.User, error) {
	return userService.UserRepository.GetByID(userID)
}

func (UserService *UserService) CreateUser(name string) (*model.User, error) {
	user := model.User{
		Name:       name,
		Createtime: time.Now(),
		Updatetime: time.Now(),
	}
	return UserService.UserRepository.Save(&user)
}

func (UserService *UserService) UpdateUser(user *model.User) error {
	return UserService.UserRepository.Update(user)
}

func (UserService *UserService) DeleteUser(userID uint64) error {
	return UserService.UserRepository.Delete(userID)
}
