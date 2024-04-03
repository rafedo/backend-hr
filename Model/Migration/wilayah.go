package Domain

type Wilayah struct {
	Id           uint   `gorm:"primaryKey"`
	NamaWilayah  string `gorm:"column:nama_wilayah"`
	AlamatKantor string `gorm:"column:alamat_kantor"`
}

// TableName returns the table name for Wilayah
func (*Wilayah) TableName() string {
	return "wilayah" // Ganti dengan nama tabel yang diinginkan
}
