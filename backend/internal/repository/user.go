package repository

import (
	"backend/internal/domain"
	"errors"
)

type UsersRepository interface {
	Create(user *domain.User) (*domain.User, error)
	GetByID(id string) (*domain.User, error)
}

type usersRepository struct {
	db map[string]domain.User
}

func NewUsersRepository() UsersRepository {
	return &usersRepository{
		db: make(map[string]domain.User),
	}
}

func (r *usersRepository) Create(u *domain.User) (*domain.User, error) {
	r.db[u.ID] = *u
	return u, nil
}

func (r *usersRepository) GetByID(id string) (*domain.User, error) {
	user, exists := r.db[id]

	if !exists {
		return nil, errors.New("user not found")
	}

	return &user, nil
}
