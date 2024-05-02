package Migration

type Ranting struct {
	Id          uint   `gorm:"primaryKey"`
	NamaRanting string `gorm:"column:nama_ranting"`
	AlamatID    Alamat `gorm:"column:alamat_kantor;foreignKey:AlamatID"`
	CabangID    Cabang `gorm:"column:cabang_id;foreignKey:CabangID"`
}

// TableName returns the table name for Ranting
func (*Ranting) TableName() string {
	return "ranting" // Ganti dengan nama tabel yang diinginkan
}
