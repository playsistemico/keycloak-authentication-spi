package services

import (
	"errors"
	"net/http"
)

type MockAuthService struct{}

func (s *MockAuthService) Authenticate(username, password string) (string, error) {
	if username == "user" && password == "password" {
		return "session-token", nil
	}
	return "", errors.New(http.StatusText(http.StatusInternalServerError))
}

func (s *MockAuthService) ValidateToken(token string) bool {
	return token == "session-token"
}
