package usecase

import (
	"github.com/lfcifuentes/clean-arquitecture/internal/domains"
)

type UserRepository interface {
	GetAll() ([]domains.User, error)
	Save(user domains.User) error
}

type UserUsecase struct {
	UserRepository UserRepository
}

func NewUserUsecase(userRepository UserRepository) *UserUsecase {
	return &UserUsecase{
		UserRepository: userRepository,
	}
}

func (u *UserUsecase) UserList() ([]domains.User, error) {
	return u.UserRepository.GetAll()
}

func (u *UserUsecase) CreateUser(user domains.User) error {
	return u.UserRepository.Save(user)
}
