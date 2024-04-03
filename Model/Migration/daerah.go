package Domain

type Daerah struct {
	Id           uint   `gorm:"primaryKey"`
	NamaDaerah   string `gorm:"column:nama_daerah"`
	AlamatKantor string `gorm:"column:alamat_kantor"`
	WilayahID    uint   `gorm:"column:wilayah_id;foreignKey:WilayahID"`
}

// TableName returns the table name for Daerah
func (*Daerah) TableName() string {
	return "daerah" // Ganti dengan nama tabel yang diinginkan
}
