package Domain

type Pengurus struct {
	Id        uint `gorm:"primaryKey"`
	AnggotaID uint `gorm:"column:anggota_id;foreignKey:AnggotaID"`
	JabatanID uint `gorm:"column:jabatan_id;foreignKey:JabatanID"`
}

// TableName returns the table name for Pengurus
func (*Pengurus) TableName() string {
	return "pengurus" // Ganti dengan nama tabel yang diinginkan
}
