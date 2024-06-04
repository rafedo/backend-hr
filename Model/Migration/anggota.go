package Migration

import "gorm.io/datatypes"

type Anggota struct {
	Id               uint           `gorm:"primaryKey"`
	NomorKTA         uint           `gorm:"column:nomor_kta;null"`
	CabangID         Cabang         `gorm:"column:cabang;foreignKey:CabangID"`
	NamaLengkap      string         `gorm:"column:nama_lengkap"`
	NIK              uint           `gorm:"column:nik"`
	GelarKesarjanaan string         `gorm:"column:gelar_kesarjanaan;null"`
	GelarLainDepan   string         `gorm:"column:gelar_lain_depan;null"`
	TempatLahir      string         `gorm:"column:tempat_lahir"`
	TanggalLahir     datatypes.Date `gorm:"column:tanggal_lahir"`
	JenisKelamin     string         `gorm:"column:jenis_kelamin"`
	StatusPernikahan string         `gorm:"column:status_pernikahan"`
	Email            string         `gorm:"column:email;null"`
	NoTelp           uint           `gorm:"column:no_telp;null"`
	AlamatID         Alamat         `gorm:"column:alamat;foreignKey:AlamatID"`
	Status           string         `gorm:"column:status"`
	InfoAnggotaID    InfoAnggota    `gorm:"column:info_anggotaID;foreignKey:InfoAnggotaID;"`
}

func (*Anggota) TableName() string {
	return "anggota" // Ganti dengan nama tabel yang diinginkan
}
