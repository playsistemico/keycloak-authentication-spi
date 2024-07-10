package defines

type LoginData struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AuthData struct {
	Token string `json:"token" binding:"required"`
}
