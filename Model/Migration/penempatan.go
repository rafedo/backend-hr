package Domain

type Penempatan struct {
	Id        uint `gorm:"primaryKey"`
	WilayahID uint `gorm:"column:wilayah_id;polymorphic:Owner;"`
	DaerahID  uint `gorm:"column:daerah_id;polymorphic:Owner;"`
	CabangID  uint `gorm:"column:cabang_id;polymorphic:Owner;"`
	RantingID uint `gorm:"column:ranting_id;polymorphic:Owner;"`
}

// TableName returns the table name for Penempatan
func (*Penempatan) TableName() string {
	return "penempatan" // Ganti dengan nama tabel yang diinginkan
}
