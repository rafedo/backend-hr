package Domain

type Arsip struct {
	Id       uint   `gorm:"primaryKey"`
	Judul    string `gorm:"column:judul"`
	Subject  string `gorm:"column:subject"`
	Isi      string `gorm:"column:isi"`
	Lampiran string `gorm:"column:lampiran"`
}

// TableName returns the table name for Arsip
func (*Arsip) TableName() string {
	return "arsip" // Ganti dengan nama tabel yang diinginkan
}
