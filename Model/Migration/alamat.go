package Migration

type Alamat struct {
	Id        uint   `gorm:"primaryKey"`
	Alamat    string `gorm:"column:alamat"`
	Kelurahan string `gorm:"column:kelurahan"`
	Kecamatan string `gorm:"column:kecamatan"`
	KabKota   string `gorm:"column:kab_kota"`
	Propinsi  string `gorm:"column:propinsi"`
	KodePos   string `gorm:"column:kode_pos"`
}

func (*Alamat) TableName() string {
	return "alamat" // Ganti dengan nama tabel yang diinginkan
}
