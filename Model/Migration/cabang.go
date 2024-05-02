package Migration

type Cabang struct {
	Id         uint   `gorm:"primaryKey"`
	NamaCabang string `gorm:"column:nama_cabang"`
	AlamatID   Alamat `gorm:"column:alamat_kantor;foreignKey:AlamatID"`
	DaerahID   Daerah `gorm:"column:daerah_id;foreignKey:DaerahID"`
}

// TableName returns the table name for Cabang
func (*Cabang) TableName() string {
	return "cabang" // Ganti dengan nama tabel yang diinginkan
}
