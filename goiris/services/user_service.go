package services

import (
	"gosummary/goiris/datamodels"
	"gosummary/goiris/repositories"
)

type IUserService interface {
	IsPwdSuccess(userName string, pwd string) (user *datamodels.User, isOk bool)
	AddUser(user *datamodels.User) (userId int64, err error)
}

type UserService struct {
	UserService repositories.IUserRepository
}

func NewUserService(repository repositories.IUserRepository) IUserService {
	return &UserService{
		UserService: repository,
	}
}

func (u *UserService) IsPwdSuccess(userName string, pwd string) (user *datamodels.User, isOk bool) {
	// todo

	return
}
func (u *UserService) AddUser(user *datamodels.User) (userId int64, err error) {
	// todo

	return
}