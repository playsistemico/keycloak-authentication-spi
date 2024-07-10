package services

type LoginService interface {
	Login(username, password string) (string, error)
	ValidateSession(token string) (bool, error)
}
