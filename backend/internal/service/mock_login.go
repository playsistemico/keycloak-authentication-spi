package services

import (
	"errors"
	"net/http"
)

type MockLoginService struct{}

func (s *MockLoginService) Login(username, password string) (string, error) {
	if username == "user" && password == "password" {
		return "session-token", nil
	}
	return "", errors.New(http.StatusText(http.StatusInternalServerError))
}

func (s *MockLoginService) ValidateSession(token string) (bool, error) {
	if token == "session-token" {
		return true, nil
	}
	return false, errors.New(http.StatusText(http.StatusInternalServerError))
}
