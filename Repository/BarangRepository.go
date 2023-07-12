package Repository

import (
	"crud-barang/Model/Database"
	"gorm.io/gorm"
)

type (
	BarangRepositoryHandler interface {
		Save(data *Database.Barang) (id int64, err error)
		Update(data *Database.Barang) (id int64, err error)
		Delete(id int64) error
		FindById(id int64) (data Database.Barang, err error)
		FindAll() (data []Database.Barang, err error)
	}

	BarangRepositoryImpl struct {
		DB *gorm.DB
	}
)

func BarangRepositoryControllerProvider(db *gorm.DB) *BarangRepositoryImpl {
	return &BarangRepositoryImpl{
		DB: db,
	}
}

func (h *BarangRepositoryImpl) Save(data *Database.Barang) (id int64, err error) {
	err = h.DB.Model(&Database.Barang{}).
		Save(&data).Error

	return data.ID, err
}

func (h *BarangRepositoryImpl) Update(data *Database.Barang) (id int64, err error) {
	err = h.DB.Model(&Database.Barang{}).
		Where("id = ?", data.ID).
		Updates(&data).Error

	return data.ID, err
}

func (h *BarangRepositoryImpl) Delete(id int64) error {
	err := h.DB.Delete(&Database.Barang{}, id).Error
	return err
}

func (h *BarangRepositoryImpl) FindById(id int64) (data Database.Barang, err error) {
	err = h.DB.Model(&Database.Barang{}).
		Select("name").
		Where("id = ?", id).
		Take(&data).
		Error
	return data, err
}

func (h *BarangRepositoryImpl) FindAll() (data []Database.Barang, err error) {
	err = h.DB.Model(&Database.Barang{}).
		Select("id", "name", "category", "price").
		Order("id asc").
		Find(&data).
		Error

	return data, nil
}
