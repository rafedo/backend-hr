package Repository

import (
	"muhammadiyah/Model/Database"
	"muhammadiyah/Model/Domain"

	"gorm.io/gorm"
)

type (
	PengurusRepositoryHandler interface {
		CreatePengurus(data *Database.Penguru) (id int64, err error)
		UpdatePengurus(data *Database.Penguru) (id int64, err error)
		DeletePengurus(id int64) error
		FindAllPengurus() (data []Database.Penguru, err error)
		FindPengurusByID(id int64) (Domain.PengurusInfoResponse, error)
		FindPengurusInfoByDepartementID(departementID int64) ([]Domain.PengurusInfoResponse, error)
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

func (repo *PengurusRepositoryImpl) FindPengurusByID(id int64) (Domain.PengurusInfoResponse, error) {
	var pengurus Domain.PengurusInfoResponse
	err := repo.DB.Table("pengurus").
		Select("pengurus.id, pengurus.anggota_id, anggota.nomor_kta, anggota.nama_lengkap, anggota.gelar_kesarjanaan, anggota.gelar_lain_depan, anggota.status, pengurus.jabatan_id, jabatan.nama as nama_jabatan, departemen.id as departemen_id, departemen.nama as nama_departemen, penempatan.id as penempatan_id, penempatan.lokasi_id, penempatan.lokasi_type, penempatan.jenis").
		Joins("join anggota on pengurus.anggota_id = anggota.id").
		Joins("join jabatan on pengurus.jabatan_id = jabatan.id").
		Joins("join departemen on jabatan.departemen_id = departemen.id").
		Joins("join penempatan on departemen.penempatan_id = penempatan.id").
		Where("pengurus.id = ?", id).First(&pengurus).Error
	return pengurus, err
}
func (r *PengurusRepositoryImpl) FindPengurusInfoByDepartementID(departementID int64) ([]Domain.PengurusInfoResponse, error) {
	var pengurusInfoList []Domain.PengurusInfoResponse

	err := r.DB.Table("pengurus").
		Select("pengurus.id, pengurus.anggota_id, anggota.nomor_kta, anggota.nama_lengkap, anggota.gelar_kesarjanaan, anggota.gelar_lain_depan, anggota.status, pengurus.jabatan_id, jabatan.nama as nama_jabatan, departemen.id as departemen_id, departemen.nama as nama_departemen, penempatan.id as penempatan_id, penempatan.lokasi_id, penempatan.lokasi_type, penempatan.jenis").
		Joins("join anggota on pengurus.anggota_id = anggota.id").
		Joins("join jabatan on pengurus.jabatan_id = jabatan.id").
		Joins("join departemen on jabatan.departemen_id = departemen.id").
		Joins("join penempatan on departemen.penempatan_id = penempatan.id").
		Where("departemen.id = ?", departementID).
		Scan(&pengurusInfoList).Error

	return pengurusInfoList, err
}

//func (repo *PengurusRepositoryImpl) FindPengurusByKTA(nomorKTA string) (Database.Penguru, error) {
//	var pengurus Database.Penguru
//	err := repo.DB.Model(&Database.Penguru{}).Where("nomorKTA = ?", nomorKTA).First(&pengurus).Error
//	return pengurus, err
//}
