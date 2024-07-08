package services

type AuthService interface {
	Authenticate(username, password string) (string, error)
	ValidateToken(token string) bool
}
