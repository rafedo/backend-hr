package Domain

type Ranting struct {
	Id           uint   `gorm:"primaryKey"`
	NamaRanting  string `gorm:"column:nama_ranting"`
	AlamatKantor string `gorm:"column:alamat_kantor"`
	RantingID    uint   `gorm:"column:ranting_id;foreignKey:RantingID"`
}

// TableName returns the table name for Ranting
func (*Ranting) TableName() string {
	return "ranting" // Ganti dengan nama tabel yang diinginkan
}
