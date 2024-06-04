package Repository

import (
	"muhammadiyah/Model/Database"

	"gorm.io/gorm"
)

type (
	PWMRepositoryHandler interface {
		CreateWilayah(data *Database.Wilayah) (id int64, err error)
		UpdateWilayah(data *Database.Wilayah) (id int64, err error)
		DeleteWilayah(id int64) error
		FindAllWilayah() (data []Database.Wilayah, err error)
		FindWilayahByID(id int64) (data Database.Wilayah, err error)

		CreateDaerah(data *Database.Daerah) (id int64, err error)
		UpdateDaerah(data *Database.Daerah) (id int64, err error)
		DeleteDaerah(id int64) error
		FindAllDaerah() (data []Database.Daerah, err error)
		FindDaerahByID(id int64) (data Database.Daerah, err error)
		FindDaerahByWilayahID(wilayahID int64) (data []Database.Daerah, err error)

		CreateCabang(data *Database.Cabang) (id int64, err error)
		UpdateCabang(data *Database.Cabang) (id int64, err error)
		DeleteCabang(id int64) error
		FindAllCabang() (data []Database.Cabang, err error)
		FindCabangByID(id int64) (data Database.Cabang, err error)
		FindCabangByDaerahID(daerahID int64) (data []Database.Cabang, err error)

		CreateRanting(data *Database.Ranting) (id int64, err error)
		UpdateRanting(data *Database.Ranting) (id int64, err error)
		DeleteRanting(id int64) error
		FindAllRanting() (data []Database.Ranting, err error)
		FindRantingByID(id int64) (data Database.Ranting, err error)
		FindRantingByCabangID(cabangID int64) (data []Database.Ranting, err error)

		CreateAddress(data *Database.Alamat) (id int64, err error)
		UpdateAddress(data *Database.Alamat) (id int64, err error)
		DeleteAddress(id int64) error
		FindAddressByID(id int64) (data Database.Alamat, err error)
	}

	PWMRepositoryImpl struct {
		DB *gorm.DB
	}
)

func PWMRepositoryControllerProvider(db *gorm.DB) *PWMRepositoryImpl {
	return &PWMRepositoryImpl{
		DB: db,
	}
}

func (h *PWMRepositoryImpl) CreateWilayah(data *Database.Wilayah) (id int64, err error) {

	err = h.DB.Model(&Database.Wilayah{}).Create(&data).Error
	if err != nil {
		return 0, err
	}
	return data.ID, nil
}

func (h *PWMRepositoryImpl) UpdateWilayah(data *Database.Wilayah) (id int64, err error) {

	err = h.DB.Model(&Database.Wilayah{}).
		Where("id = ?", data.ID).
		Updates(&data).Error
	if err != nil {
		return 0, err
	}
	return data.ID, nil
}
func (h *PWMRepositoryImpl) DeleteWilayah(id int64) error {

	err := h.DB.Delete(&Database.Wilayah{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
func (h *PWMRepositoryImpl) FindAllWilayah() (data []Database.Wilayah, err error) {
	err = h.DB.Model(&Database.Wilayah{}).
		Select("id", "nama_wilayah", "alamat_kantor").
		Order("id asc").
		Find(&data).
		Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (h *PWMRepositoryImpl) FindWilayahByID(id int64) (data Database.Wilayah, err error) {
	err = h.DB.Model(&Database.Wilayah{}).
		Where("id = ?", id).
		First(&data).
		Error

	return data, err
}

func (h *PWMRepositoryImpl) CreateDaerah(data *Database.Daerah) (id int64, err error) {

	err = h.DB.Model(&Database.Daerah{}).
		Create(&data).Error
	if err != nil {
		return 0, err
	}
	return data.ID, nil
}
func (h *PWMRepositoryImpl) UpdateDaerah(data *Database.Daerah) (id int64, err error) {

	err = h.DB.Model(&Database.Daerah{}).
		Where("id = ?", data.ID).
		Updates(&data).Error
	if err != nil {
		return 0, err
	}
	return data.ID, nil
}
func (h *PWMRepositoryImpl) DeleteDaerah(id int64) error {

	err := h.DB.Delete(&Database.Daerah{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
func (h *PWMRepositoryImpl) FindAllDaerah() (data []Database.Daerah, err error) {
	err = h.DB.Model(&Database.Daerah{}).
		Select("id", "nama_daerah", "alamat_kantor", "wilayah_id").
		Order("id asc").
		Find(&data).
		Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (h *PWMRepositoryImpl) FindDaerahByID(id int64) (data Database.Daerah, err error) {
	err = h.DB.Model(&Database.Daerah{}).
		Where("id = ?", id).
		First(&data).
		Error

	return data, err
}

func (h *PWMRepositoryImpl) FindDaerahByWilayahID(wilayahID int64) (data []Database.Daerah, err error) {
	err = h.DB.Model(&Database.Daerah{}).
		Where("wilayah_id = ?", wilayahID).
		Find(&data).
		Error

	return data, err
}

func (h *PWMRepositoryImpl) CreateCabang(data *Database.Cabang) (id int64, err error) {

	err = h.DB.Model(&Database.Cabang{}).
		Save(&data).Error
	if err != nil {
		return 0, err
	}
	return data.ID, err
}
func (h *PWMRepositoryImpl) UpdateCabang(data *Database.Cabang) (id int64, err error) {

	err = h.DB.Model(&Database.Cabang{}).
		Where("id = ?", data.ID).
		Updates(&data).Error
	if err != nil {
		return 0, err
	}
	return data.ID, err
}
func (h *PWMRepositoryImpl) DeleteCabang(id int64) error {
	err := h.DB.Delete(&Database.Cabang{}, id).Error

	return err
}
func (h *PWMRepositoryImpl) FindAllCabang() (data []Database.Cabang, err error) {
	err = h.DB.Model(&Database.Cabang{}).
		Select("id", "nama_cabang", "alamat_kantor", "daerah_id").
		Order("id asc").
		Find(&data).
		Error
	if err != nil {
		return nil, err
	}
	return data, err
}

func (h *PWMRepositoryImpl) FindCabangByID(id int64) (data Database.Cabang, err error) {
	err = h.DB.Model(&Database.Cabang{}).
		Where("id = ?", id).
		First(&data).
		Error

	return data, err
}

func (h *PWMRepositoryImpl) FindCabangByDaerahID(daerahID int64) (data []Database.Cabang, err error) {
	err = h.DB.Model(&Database.Cabang{}).
		Where("daerah_id = ?", daerahID).
		Find(&data).
		Error

	return data, err
}

func (h *PWMRepositoryImpl) CreateRanting(data *Database.Ranting) (id int64, err error) {

	err = h.DB.Model(&Database.Ranting{}).
		Save(&data).Error

	return data.ID, err
}

func (h *PWMRepositoryImpl) UpdateRanting(data *Database.Ranting) (id int64, err error) {

	err = h.DB.Model(&Database.Ranting{}).
		Where("id = ?", data.ID).
		Updates(&data).Error
	if err != nil {
		return 0, err
	}
	return data.ID, err
}
func (h *PWMRepositoryImpl) DeleteRanting(id int64) error {

	err := h.DB.Delete(&Database.Ranting{}, id).Error

	return err

}
func (h *PWMRepositoryImpl) FindAllRanting() (data []Database.Ranting, err error) {
	err = h.DB.Model(&Database.Ranting{}).
		Select("id", "nama_ranting", "alamat_kantor", "cabang_id").
		Order("id asc").
		Find(&data).
		Error

	return data, err
}
func (h *PWMRepositoryImpl) FindRantingByID(id int64) (data Database.Ranting, err error) {
	err = h.DB.Model(&Database.Ranting{}).
		Where("id = ?", id).
		First(&data).
		Error
	if err != nil {
		return Database.Ranting{}, err
	}
	return data, err
}

func (h *PWMRepositoryImpl) FindRantingByCabangID(cabangID int64) (data []Database.Ranting, err error) {
	err = h.DB.Model(&Database.Ranting{}).
		Where("cabang_id = ?", cabangID).
		Find(&data).
		Error
	if err != nil {
		return []Database.Ranting{}, err
	}
	return data, err
}

func (h *PWMRepositoryImpl) CreateAddress(data *Database.Alamat) (id int64, err error) {
	err = h.DB.Model(&Database.Alamat{}).Create(&data).Error
	if err != nil {
		return 0, err
	}
	return data.ID, err
}

func (h *PWMRepositoryImpl) UpdateAddress(data *Database.Alamat) (id int64, err error) {
	err = h.DB.Model(&Database.Alamat{}).Where("id = ?", data.ID).Updates(&data).Error
	return data.ID, err
}

func (h *PWMRepositoryImpl) DeleteAddress(id int64) error {
	err := h.DB.Delete(&Database.Alamat{}, id).Error
	return err
}

func (h *PWMRepositoryImpl) FindAddressByID(id int64) (data Database.Alamat, err error) {
	err = h.DB.Model(&Database.Alamat{}).
		Where("id = ?", id).
		First(&data).
		Error

	return data, err
}
