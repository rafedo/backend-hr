package Migration

type Jabatan struct {
	Id           uint       `gorm:"primaryKey"`
	Nama         string     `gorm:"column:nama"`
	DepartemenID Departemen `gorm:"column:departemen_id;foreignKey:DepartemenID"`
}

// TableName returns the table name for Jabatan
func (*Jabatan) TableName() string {
	return "jabatan" // Ganti dengan nama tabel yang diinginkan
}
