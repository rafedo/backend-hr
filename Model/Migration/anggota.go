package Domain

type Anggota struct {
	Id     uint   `gorm:"primaryKey"`
	KTA    string `gorm:"column:kta"`
	Nama   string `gorm:"column:nama"`
	Status string `gorm:"column:status"`
}

func (*Anggota) TableName() string {
	return "anggota" // Ganti dengan nama tabel yang diinginkan
}
