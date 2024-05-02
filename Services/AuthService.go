package Services

import (
	"muhammadiyah/Middleware"
	"muhammadiyah/Model/Database"
	"muhammadiyah/Model/Domain"
	"net/http"

	"muhammadiyah/Model/Web"
	"muhammadiyah/Repository"
)

type (
	AuthServiceHandler interface {
		Login(request Domain.LoginRequest) (response Domain.LoginResponse, serviceErr *Web.ServiceErrorDto)
		Register(request Domain.RegisterRequest) (response Domain.RegisterResponse, serviceErr *Web.ServiceErrorDto)
		FindAllUser() (memberResponse []Domain.UserResponse, serviceErr *Web.ServiceErrorDto)
		GetUserByID(id int64) (response Domain.User, serviceErr *Web.ServiceErrorDto)
	}

	AuthServiceImpl struct {
		authRepo     Repository.AuthRepositoryHandler
		pengurusRepo Repository.PengurusRepositoryHandler
		anggotaRepo  Repository.PWMRepositoryHandler
	}
)

func AuthServiceControllerProvider(authRepo Repository.AuthRepositoryHandler, pengurusRepo Repository.PengurusRepositoryHandler, anggotaRepo Repository.PWMRepositoryHandler) *AuthServiceImpl {
	return &AuthServiceImpl{
		authRepo:     authRepo,
		pengurusRepo: pengurusRepo,
		anggotaRepo:  anggotaRepo,
	}
}

func (h *AuthServiceImpl) Login(request Domain.LoginRequest) (response Domain.LoginResponse, serviceErr *Web.ServiceErrorDto) {
	//response, err := h.authRepo.Login(data)
	//if err != nil {
	//	return response, Web.NewCustomServiceError("Login failed", err, http.StatusUnauthorized)
	//}
	//

	var data Database.User

	// Mencoba mencari pengguna berdasarkan username
	userByUsername, err := h.authRepo.FindUserByUsername(request.Username)
	if err == nil {
		data = userByUsername
	} else {
		// Jika pencarian berdasarkan username tidak berhasil, coba mencari berdasarkan email
		userByEmail, err := h.authRepo.FindUserByEmail(request.Username)
		if err != nil {
			return response, Web.NewCustomServiceError("Username salah", err, http.StatusUnauthorized)
		}
		data = userByEmail
	}

	// Setelah mendapatkan data pengguna, lakukan pengecekan password dan buat token jika berhasil
	if data.Password != request.Password {
		return response, Web.NewCustomServiceError("Login failed", nil, http.StatusUnauthorized)
	}
	user, err := h.authRepo.FindDetailUserByID(data.ID)
	if err != nil {
		return response, Web.NewCustomServiceError("Login failed", nil, http.StatusUnauthorized)
	}
	token, err := Middleware.GenerateToken(Domain.User{
		UserID:     data.ID,
		Jabatan:    user.Jabatan,
		Departemen: user.Departemen,
		Penempatan: user.Penempatan,
	})
	if err != nil {
		return response, Web.NewCustomServiceError("Login failed", err, http.StatusUnauthorized)
	}

	response = Domain.LoginResponse{
		Token: token,
	}

	return response, nil

}
func (h *AuthServiceImpl) Register(request Domain.RegisterRequest) (response Domain.RegisterResponse, serviceErr *Web.ServiceErrorDto) {
	// Cek apakah username atau email sudah digunakan
	anggota, err := h.anggotaRepo.FindMemberByKTA(request.NomorKTA)
	if err != nil {
		return response, Web.NewCustomServiceError("anggota tidak terdaftar", nil, http.StatusBadRequest)
	}
	pengurusID, err := h.pengurusRepo.CreatePengurus(&Database.Penguru{
		AnggotaID: anggota.ID,
		JabatanID: request.JabatanID,
	})
	_, err = h.authRepo.FindUserByUsername(request.Username)
	if err == nil {
		return response, Web.NewCustomServiceError("Username already exists", nil, http.StatusBadRequest)
	}

	_, err = h.authRepo.FindUserByEmail(request.Email)
	if err != nil {
		return response, Web.NewCustomServiceError("Email already exists", nil, http.StatusBadRequest)
	}

	// Simpan pengguna baru ke dalam database
	userID, err := h.authRepo.CreateUser(&Database.User{
		Username:   request.Username,
		Email:      request.Email,
		Password:   request.Password,
		PengurusID: pengurusID,
	})
	if err != nil {
		return response, Web.NewCustomServiceError("Failed to register user", err, http.StatusInternalServerError)
	}

	response = Domain.RegisterResponse{
		UserID:  userID,
		Message: "User registered successfully",
	}
	return response, nil
}
func (h *AuthServiceImpl) FindAllUser() (memberResponse []Domain.UserResponse, serviceErr *Web.ServiceErrorDto) {
	member, err := h.authRepo.FindAllUsers()
	if err != nil {
		return []Domain.UserResponse{}, Web.NewInternalServiceError(err)
	}
	for _, res := range member {
		memberResponse = append(memberResponse, Domain.UserResponse{
			ID:         res.ID,
			Username:   res.Username,
			Password:   res.Password,
			PengurusID: res.PengurusID,
		})

	}
	return memberResponse, nil
}
func (h *AuthServiceImpl) GetUserByID(id int64) (response Domain.User, serviceErr *Web.ServiceErrorDto) {
	user, err := h.authRepo.FindDetailUserByID(id)
	if err != nil {
		return Domain.User{}, Web.NewCustomServiceError("User not found ", err, http.StatusInternalServerError)
	}
	userResponse := Domain.User{
		UserID:     user.UserID,
		Jabatan:    user.Jabatan,
		Departemen: user.Departemen,
		Penempatan: user.Penempatan,
	}
	return userResponse, nil
}
