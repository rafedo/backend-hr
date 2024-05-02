package Migration

type Wilayah struct {
	Id          uint   `gorm:"primaryKey"`
	NamaWilayah string `gorm:"column:nama_wilayah"`
	AlamatID    Alamat `gorm:"column:alamat_kantor;foreignKey:AlamatID"`
}

// TableName returns the table name for Wilayah
func (*Wilayah) TableName() string {
	return "wilayah" // Ganti dengan nama tabel yang diinginkan
}
