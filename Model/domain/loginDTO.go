package Domain

type LoginResponse struct {
	JWT string `json:"token"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
