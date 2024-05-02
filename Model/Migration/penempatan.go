package Migration

type Penempatan struct {
	ID         int64  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	LokasiID   int64  `gorm:"column:lokasi_id" json:"lokasi_id"`
	LokasiType string `gorm:"column:lokasi_type" json:"lokasi_type"`
	Jenis      string `gorm:"column:jenis" json:"jenis"`
	// Kolom-kolom lain yang diperlukan
}

// TableName Penempatan's table name
func (*Penempatan) TableName() string {
	return "penempatan"
}
