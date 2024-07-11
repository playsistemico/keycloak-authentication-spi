package services

import (
	"errors"
	"net/http"

	"backend/internal/domain"
)

type MockLoginService struct{}

func (s *MockLoginService) Create(username, password string) (string, error) {
	if username == "user" && password == "password" {
		return "session-token", nil
	}
	return "", errors.New(http.StatusText(http.StatusInternalServerError))
}

func (s *MockLoginService) GetUser(sessionToken string) (domain.User, error) {
	if sessionToken == "session-token" {
		return domain.User{
			ID: "1111-1111-1111",
		}, nil
	}
	return domain.User{}, errors.New(http.StatusText(http.StatusInternalServerError))
}
