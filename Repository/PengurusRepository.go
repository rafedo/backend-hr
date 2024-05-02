package Repository

import (
	"muhammadiyah/Model/Database"

	"gorm.io/gorm"
)

type (
	PengurusRepositoryHandler interface {
		CreatePengurus(data *Database.Penguru) (id int64, err error)
		UpdatePengurus(data *Database.Penguru) (id int64, err error)
		DeletePengurus(id int64) error
		FindAllPengurus() (data []Database.Penguru, err error)
		FindPengurusByID(id int64) (Database.Penguru, error)
	}

	PengurusRepositoryImpl struct {
		DB *gorm.DB
	}
)

func PengurusRepositoryControllerProvider(db *gorm.DB) *PengurusRepositoryImpl {
	return &PengurusRepositoryImpl{
		DB: db,
	}
}

func (repo *PengurusRepositoryImpl) CreatePengurus(data *Database.Penguru) (id int64, err error) {
	err = repo.DB.Model(&Database.Penguru{}).Save(data).Error
	return data.ID, err
}

func (repo *PengurusRepositoryImpl) UpdatePengurus(data *Database.Penguru) (id int64, err error) {
	err = repo.DB.Model(&Database.Penguru{}).Where("id = ?", data.ID).Updates(data).Error
	return data.ID, err
}

func (repo *PengurusRepositoryImpl) DeletePengurus(id int64) error {
	err := repo.DB.Delete(&Database.Penguru{}, id).Error
	return err
}

func (repo *PengurusRepositoryImpl) FindAllPengurus() (data []Database.Penguru, err error) {
	err = repo.DB.Model(&Database.Penguru{}).Find(&data).Error
	return data, err
}

func (repo *PengurusRepositoryImpl) FindPengurusByID(id int64) (Database.Penguru, error) {
	var pengurus Database.Penguru
	err := repo.DB.Model(&Database.Penguru{}).Where("id = ?", id).First(&pengurus).Error
	return pengurus, err
}

//func (repo *PengurusRepositoryImpl) FindPengurusByKTA(nomorKTA string) (Database.Penguru, error) {
//	var pengurus Database.Penguru
//	err := repo.DB.Model(&Database.Penguru{}).Where("nomorKTA = ?", nomorKTA).First(&pengurus).Error
//	return pengurus, err
//}
