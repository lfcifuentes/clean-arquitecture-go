package repository

import (
	"errors"

	"github.com/lfcifuentes/clean-arquitecture/internal/domains"
)

type UserMemoryRepository struct {
	users []domains.User
}

func NewUserMemoryRepository() *UserMemoryRepository {
	return &UserMemoryRepository{
		users: []domains.User{},
	}
}

func (r *UserMemoryRepository) GetAll() ([]domains.User, error) {
	return r.users, nil
}

func (r *UserMemoryRepository) Save(user domains.User) error {
	// check if the user already exists by id
	for _, u := range r.users {
		if u.ID == user.ID {
			return errors.New("user already exists")
		}
	}
	r.users = append(r.users, user)
	return nil
}
