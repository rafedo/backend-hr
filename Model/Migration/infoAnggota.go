package Migration

type InfoAnggota struct {
	Id                     uint   `gorm:"primaryKey"`
	Profesi                string `gorm:"column:profesi;null"`
	ProfesiLainnya         string `gorm:"column:profesi_lainnya;null"`
	Pekerjaan              string `gorm:"column:pekerjaan;null"`
	Instansi               string `gorm:"column:instansi;null"`
	PendidikanTerakhir     string `gorm:"column:pendidikan_terakhir;null"`
	PernahBelajarPesantren bool   `gorm:"column:pernah_belajar_pesantren;null"`
	Bahasa                 string `gorm:"column:bahasa;null"`
	Organisasi             string `gorm:"column:organisasi;null"`
}

func (*InfoAnggota) TableName() string {
	return "info_anggota" // Ganti dengan nama tabel yang diinginkan
}
