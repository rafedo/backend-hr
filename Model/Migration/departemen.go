package Domain

type Departemen struct {
	Id           uint   `gorm:"primaryKey"`
	Nama         string `gorm:"column:nama"`
	PenempatanID uint   `gorm:"column:penempatan_id;foreignKey:PenempatanID"`
}

// TableName returns the table name for Departemen
func (*Departemen) TableName() string {
	return "departemen" // Ganti dengan nama tabel yang diinginkan
}
