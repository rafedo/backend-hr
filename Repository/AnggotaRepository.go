package Repository

import (
	"muhammadiyah/Model/Database"

	"gorm.io/gorm"
)

type (
	AnggotaRepositoryHandler interface {
		CreateMember(data *Database.Anggotum) (id int64, err error)
		UpdateMember(data *Database.Anggotum) (id int64, err error)
		DeleteMember(id int64) error
		FindAllMembers() (data []Database.Anggotum, err error)
		FindMemberByID(id int64) (data Database.Anggotum, err error)
		FindMemberByKTA(nomorKTA int64) (data Database.Anggotum, err error)

		FindMembersByRantingID(rantingID int64) (data []Database.Anggotum, err error)
		FindMembersByCabangID(cabangID int64) (data []Database.Anggotum, err error)
		FindMembersByDaerahID(daerahID int64) (data []Database.Anggotum, err error)
		FindMembersByWilayahID(wilayahID int64) (data []Database.Anggotum, err error)
		CountMembers() (count int64, err error)

		CreateInfoAnggota(data *Database.InfoAnggotum) (id int64, err error)
		UpdateInfoAnggota(data *Database.InfoAnggotum) (id int64, err error)
		DeleteInfoAnggota(id int64) error
		FindInfoAnggotaByID(id int64) (data Database.InfoAnggotum, err error)
	}

	AnggotaRepositoryImpl struct {
		DB *gorm.DB
	}
)

func AnggotaRepositoryControllerProvider(db *gorm.DB) *AnggotaRepositoryImpl {
	return &AnggotaRepositoryImpl{
		DB: db,
	}
}

func (h *AnggotaRepositoryImpl) FindMemberByKTA(nomorKTA int64) (data Database.Anggotum, err error) {
	err = h.DB.Model(&Database.Anggotum{}).
		Where("nomor_kta = ?", nomorKTA).
		First(&data).
		Error

	return data, err
}
func (h *AnggotaRepositoryImpl) CreateMember(data *Database.Anggotum) (id int64, err error) {
	err = h.DB.Model(&Database.Anggotum{}).Save(&data).Error
	return data.ID, err
}
func (h *AnggotaRepositoryImpl) UpdateMember(data *Database.Anggotum) (id int64, err error) {
	err = h.DB.Model(&Database.Anggotum{}).Where("id = ?", data.ID).Updates(&data).Error
	return data.ID, err
}
func (h *AnggotaRepositoryImpl) DeleteMember(id int64) error {
	err := h.DB.Delete(&Database.Anggotum{}, id).Error
	return err
}
func (h *AnggotaRepositoryImpl) FindAllMembers() (data []Database.Anggotum, err error) {
	err = h.DB.Model(&Database.Anggotum{}).
		Select("id", "nomor_kta", "cabang", "nama_lengkap", "gelar_kesarjanaan", "gelar_lain_depan", "tempat_lahir", "tanggal_lahir", "jenis_kelamin", "alamat", "status", "info_anggotaID").
		Order("id asc").
		Find(&data).
		Error

	return data, err
}
func (h *AnggotaRepositoryImpl) FindMemberByID(id int64) (data Database.Anggotum, err error) {
	err = h.DB.Model(&Database.Anggotum{}).
		Where("id = ?", id).
		First(&data).
		Error

	return data, err
}
func (h *AnggotaRepositoryImpl) FindMembersByRantingID(rantingID int64) (data []Database.Anggotum, err error) {
	err = h.DB.Model(&Database.Anggotum{}).
		Where("ranting_id = ?", rantingID).
		Find(&data).
		Error

	return data, err
}
func (h *AnggotaRepositoryImpl) FindMembersByCabangID(cabangID int64) (data []Database.Anggotum, err error) {
	err = h.DB.Model(&Database.Anggotum{}).
		Joins("JOIN rantings ON anggota.ranting_id = rantings.id").
		Where("rantings.cabang_id = ?", cabangID).
		Find(&data).
		Error

	return data, err
}
func (h *AnggotaRepositoryImpl) FindMembersByDaerahID(daerahID int64) (data []Database.Anggotum, err error) {
	err = h.DB.
		Joins("JOIN rantings ON anggota.ranting_id = rantings.id").
		Joins("JOIN cabangs ON rantings.cabang_id = cabangs.id").
		Where("cabangs.daerah_id = ?", daerahID).
		Find(&data).
		Error

	return data, err
}
func (h *AnggotaRepositoryImpl) FindMembersByWilayahID(wilayahID int64) (data []Database.Anggotum, err error) {
	err = h.DB.
		Joins("JOIN ranting ON anggota.ranting_id = ranting.id").
		Joins("JOIN cabang ON ranting.cabang_id = cabang.id").
		Joins("JOIN daerah ON cabang.daerah_id = daerah.id").
		Where("daerah.wilayah_id = ?", wilayahID).
		Find(&data).
		Error

	return data, err
}
func (h *AnggotaRepositoryImpl) CountMembers() (count int64, err error) {
	err = h.DB.Model(&Database.Anggotum{}).Count(&count).Error
	return count, err
}
func (h *AnggotaRepositoryImpl) CreateInfoAnggota(data *Database.InfoAnggotum) (id int64, err error) {
	err = h.DB.Model(&Database.InfoAnggotum{}).Save(&data).Error
	return data.ID, err
}
func (h *AnggotaRepositoryImpl) UpdateInfoAnggota(data *Database.InfoAnggotum) (id int64, err error) {
	err = h.DB.Model(&Database.InfoAnggotum{}).Where("id = ?", data.ID).Updates(&data).Error
	return data.ID, err
}
func (h *AnggotaRepositoryImpl) DeleteInfoAnggota(id int64) error {
	err := h.DB.Delete(&Database.InfoAnggotum{}, id).Error
	return err
}
func (h *AnggotaRepositoryImpl) FindInfoAnggotaByID(id int64) (data Database.InfoAnggotum, err error) {
	err = h.DB.Model(&Database.InfoAnggotum{}).
		Where("id = ?", id).
		First(&data).
		Error

	return data, err
}
