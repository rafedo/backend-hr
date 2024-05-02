package Domain

type LoginResponse struct {
	Token string `json:"token"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Username  string `json:"username" `
	Email     string `json:"email"`
	Password  string `json:"password" `
	NomorKTA  int64  `json:"nomor_kta" `
	JabatanID int64  `json:"jabatan_id"`
}

type RegisterResponse struct {
	UserID  int64  `json:"userId"`
	Message string `json:"message"`
}

type User struct {
	UserID     int64  `json:"userId"`
	Jabatan    string `json:"jabatan"`
	Departemen string `json:"departemen"`
	Penempatan string `json:"penempatan"`
}

type UserResponse struct {
	ID         int64  `json:"userId"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	PengurusID int64  `json:"pengurusId"`
}
