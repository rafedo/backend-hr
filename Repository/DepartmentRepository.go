package Repository

import (
	"gorm.io/gorm"
	"muhammadiyah/Model/Database"
)

type (
	DepartmentRepositoryHandler interface {
		CreateDepartment(data *Database.Departeman) (id int64, err error)
		UpdateDepartment(data *Database.Departeman) (id int64, err error)
		DeleteDepartment(id int64) error
		FindAllDepartments() (data []Database.Departeman, err error)

		CreatePlacement(data *Database.Penempatan) (id int64, err error)
		UpdatePlacement(data *Database.Penempatan) (id int64, err error)
		DeletePlacement(id int64) error
		FindAllPlacements() (data []Database.Penempatan, err error)

		CreatePosition(data *Database.Jabatan) (id int64, err error)
		UpdatePosition(data *Database.Jabatan) (id int64, err error)
		DeletePosition(id int64) error
		FindAllPositions() (data []Database.Jabatan, err error)
	}

	DepartmentRepositoryImpl struct {
		DB *gorm.DB
	}
)

func DepartmentRepositoryControllerProvider(db *gorm.DB) *DepartmentRepositoryImpl {
	return &DepartmentRepositoryImpl{
		DB: db,
	}
}

func (h *DepartmentRepositoryImpl) CreateDepartment(data *Database.Departeman) (id int64, err error) {
	err = h.DB.Model(&Database.Departeman{}).Save(&data).Error
	return data.ID, err
}

func (h *DepartmentRepositoryImpl) UpdateDepartment(data *Database.Departeman) (id int64, err error) {
	err = h.DB.Model(&Database.Departeman{}).Where("id = ?", data.ID).Updates(&data).Error
	return data.ID, err
}

func (h *DepartmentRepositoryImpl) DeleteDepartment(id int64) error {
	err := h.DB.Delete(&Database.Departeman{}, id).Error
	return err
}

func (h *DepartmentRepositoryImpl) FindAllDepartments() (data []Database.Departeman, err error) {
	err = h.DB.Model(&Database.Departeman{}).
		Select("id", "nama", "penempatan_id").
		Order("id asc").
		Find(&data).
		Error

	return data, err
}

func (h *DepartmentRepositoryImpl) CreatePlacement(data *Database.Penempatan) (id int64, err error) {
	err = h.DB.Model(&Database.Penempatan{}).Save(&data).Error
	return data.ID, err
}

func (h *DepartmentRepositoryImpl) UpdatePlacement(data *Database.Penempatan) (id int64, err error) {
	err = h.DB.Model(&Database.Penempatan{}).Where("id = ?", data.ID).Updates(&data).Error
	return data.ID, err
}

func (h *DepartmentRepositoryImpl) DeletePlacement(id int64) error {
	err := h.DB.Delete(&Database.Penempatan{}, id).Error
	return err
}

func (h *DepartmentRepositoryImpl) FindAllPlacements() (data []Database.Penempatan, err error) {
	err = h.DB.Model(&Database.Penempatan{}).
		Select("id", "lokasi_id", "lokasi_type", "jenis").
		Order("id asc").
		Find(&data).
		Error

	return data, err
}

func (h *DepartmentRepositoryImpl) CreatePosition(data *Database.Jabatan) (id int64, err error) {
	err = h.DB.Model(&Database.Jabatan{}).Save(&data).Error
	return data.ID, err
}

func (h *DepartmentRepositoryImpl) UpdatePosition(data *Database.Jabatan) (id int64, err error) {
	err = h.DB.Model(&Database.Jabatan{}).Where("id = ?", data.ID).Updates(&data).Error
	return data.ID, err
}

func (h *DepartmentRepositoryImpl) DeletePosition(id int64) error {
	err := h.DB.Delete(&Database.Jabatan{}, id).Error
	return err
}

func (h *DepartmentRepositoryImpl) FindAllPositions() (data []Database.Jabatan, err error) {
	err = h.DB.Model(&Database.Jabatan{}).
		Select("id", "nama", "departemen_id").
		Order("id asc").
		Find(&data).
		Error

	return data, err
}
