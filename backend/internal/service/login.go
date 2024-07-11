package services

import (
	"errors"

	"github.com/google/uuid"

	"backend/internal/domain"
	"backend/internal/repository"
)

type LoginService interface {
	Create(username, password string) (string, error)
	GetUser(sessionToken string) (domain.User, error)
}

type loginService struct {
	usersRepo repository.UsersRepository
}

func NewLoginService(usersRepo repository.UsersRepository) LoginService {
	return &loginService{usersRepo: usersRepo}
}

func (s *loginService) Create(username, password string) (string, error) {
	id := uuid.New().String()
	user := &domain.User{
		ID:   id,
		Name: username,
	}

	_, err := s.usersRepo.Create(user)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (s *loginService) GetUser(sessionToken string) (domain.User, error) {
	user, err := s.usersRepo.GetByID(sessionToken)
	if err != nil {
		return domain.User{}, err
	}
	if user == nil {
		return domain.User{}, errors.New("user not found")
	}
	return *user, nil
}
