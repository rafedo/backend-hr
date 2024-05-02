package Repository

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"muhammadiyah/Model/Database"
	"muhammadiyah/Model/Domain"
)

type (
	AuthRepositoryHandler interface {
		Login(data Domain.LoginRequest) (response Database.User, err error)
		CreateUser(data *Database.User) (id int64, err error)
		UpdateUser(data Database.User) (id int64, err error)
		DeleteUser(id int64) error
		FindAllUsers() (data []Database.User, err error)
		FindUserByID(id int64) (data Database.User, err error)
		FindUserByUsername(username string) (data Database.User, err error)
		FindUserByEmail(email string) (data Database.User, err error)
		FindDetailUserByID(id int64) (data Domain.User, err error)
	}

	AuthRepositoryImpl struct {
		DB *gorm.DB
	}
)

func AuthRepositoryControllerProvider(db *gorm.DB) *AuthRepositoryImpl {
	return &AuthRepositoryImpl{
		DB: db,
	}
}

func (h *AuthRepositoryImpl) Login(data Domain.LoginRequest) (response Database.User, err error) {

	err = h.DB.Model(&Database.Alamat{}).
		Where("username = ?", data.Username).
		First(&response).
		Error

	//err = h.DB.Where("username = ?", data.Username).First(&response).Error
	//if err != nil {
	//	return nil, err
	//}
	//
	//// Verifikasi password
	//if response.Password != data.Password {
	//	return nil, err
	//}

	// Inisialisasi response sebelum menggunakannya
	//response = &Domain.LoginResponse{}
	//
	//// Buat token JWT setelah verifikasi berhasil
	//response.JWT, err = GenerateToken(&user)
	//if err != nil {
	//	return nil, err
	//}
	//
	return response, nil
}

//	func (h *AuthRepositoryImpl) Register(data Database.User) (response *Domain.RegisterResponse, err error) {
//		// Membuat objek user baru dari data yang diterima
//
//		// Menyimpan data user baru ke dalam database
//		if err := h.DB.Create(&newUser).Error; err != nil {
//			return nil, err
//		}
//
//		// Inisialisasi response sebelum menggunakannya
//
//		return response, nil
//	}

func (h *AuthRepositoryImpl) CreateUser(data *Database.User) (id int64, err error) {
	if err := h.DB.Model(&Database.User{}).Create(&data).Error; err != nil {
		return 0, fmt.Errorf("failed to create user: %v", err)
	}
	return data.ID, nil

}
func (h *AuthRepositoryImpl) UpdateUser(data Database.User) (id int64, err error) {
	if err := h.DB.Save(data).Error; err != nil {
		return 0, fmt.Errorf("failed to update user: %v", err)
	}
	return data.ID, nil
}
func (h *AuthRepositoryImpl) DeleteUser(id int64) error {
	if err := h.DB.Delete(&Database.User{}, id).Error; err != nil {
		return fmt.Errorf("failed to delete user: %v", err)
	}
	return nil
}
func (h *AuthRepositoryImpl) FindAllUsers() (data []Database.User, err error) {
	if err := h.DB.Find(&data).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch users: %v", err)
	}
	return data, nil
}
func (h *AuthRepositoryImpl) FindUserByID(id int64) (data Database.User, err error) {
	if err := h.DB.First(&data, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return Database.User{}, fmt.Errorf("user not found")
		}
		return Database.User{}, fmt.Errorf("failed to fetch user: %v", err)
	}
	return data, nil
}
func (h *AuthRepositoryImpl) FindUserByUsername(username string) (data Database.User, err error) {
	if err := h.DB.Model(&Database.User{}).Where("username = ?", username).First(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return Database.User{}, fmt.Errorf("user not found")
		}
		return Database.User{}, fmt.Errorf("failed to fetch user by username: %v", err)
	}
	return data, nil
}
func (h *AuthRepositoryImpl) FindUserByEmail(email string) (data Database.User, err error) {
	if err := h.DB.Model(&Database.User{}).Where("email = ?", email).First(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return Database.User{}, fmt.Errorf("user not found")
		}
		return Database.User{}, fmt.Errorf("failed to fetch user by email: %v", err)
	}
	return data, nil
}
func (h *AuthRepositoryImpl) FindDetailUserByID(id int64) (data Domain.User, err error) {
	err = h.DB.Model(&Database.User{}).
		Select("public.user.id as userID", "jabatan.nama as jabatan", "departemen.nama as departemen", "penempatan.lokasi_type as penempatan").
		Joins("JOIN pengurus ON public.user.pengurus_id = pengurus.id").
		Joins("JOIN jabatan ON pengurus.jabatan_id = jabatan.id").
		Joins("JOIN departemen ON jabatan.departemen_id = departemen.id").
		Joins("JOIN penempatan ON penempatan.id = departemen.penempatan_id").
		Where("public.user.id = ?", id).
		Order("userID").
		Limit(1).
		Scan(&data).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return Domain.User{}, fmt.Errorf("user not found")
	} else if err != nil {
		return Domain.User{}, fmt.Errorf("failed to fetch user: %v", err)
	}

	return data, nil
}
