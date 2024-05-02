package Migration

type Anggota struct {
	Id               uint    `gorm:"primaryKey"`
	NomorKTA         uint    `gorm:"column:nomor_kta"`
	RantingID        Ranting `gorm:"column:ranting;foreignKey:RantingID"`
	NamaLengkap      string  `gorm:"column:nama_lengkap"`
	GelarKesarjanaan string  `gorm:"column:gelar_kesarjanaan"`
	GelarLainDepan   string  `gorm:"column:gelar_lain_depan"`
	TempatLahir      string  `gorm:"column:tempat_lahir"`
	TanggalLahir     string  `gorm:"column:tanggal_lahir"`
	JenisKelamin     string  `gorm:"column:jenis_kelamin"`
	AlamatID         Alamat  `gorm:"column:alamat;foreignKey:AlamatID"`
	Status           string  `gorm:"column:status"`
}

func (*Anggota) TableName() string {
	return "anggota" // Ganti dengan nama tabel yang diinginkan
}
