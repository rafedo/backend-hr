package Migration

type Departemen struct {
	Id           uint       `gorm:"primaryKey"`
	Nama         string     `gorm:"column:nama"`
	PenempatanID Penempatan `gorm:"column:penempatan_id;foreignKey:PenempatanID"`
	Bagian       string
}

// TableName returns the table name for Departemen
func (*Departemen) TableName() string {
	return "departemen" // Ganti dengan nama tabel yang diinginkan
}
