package Migration

type Daerah struct {
	Id         uint    `gorm:"primaryKey"`
	NamaDaerah string  `gorm:"column:nama_daerah"`
	AlamatID   Alamat  `gorm:"column:alamat_kantor;foreignKey:AlamatID"`
	WilayahID  Wilayah `gorm:"column:wilayah_id;foreignKey:WilayahID"`
}

// TableName returns the table name for Daerah
func (*Daerah) TableName() string {
	return "daerah" // Ganti dengan nama tabel yang diinginkan
}
