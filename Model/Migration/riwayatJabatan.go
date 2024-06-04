package Migration

import "gorm.io/datatypes"

type RiwayatJabatan struct {
	Id                 uint           `gorm:"primaryKey"`
	PengurusID         Pengurus       `gorm:"column:pengurus_id;foreignKey:PengurusID"`
	JabatanID          Jabatan        `gorm:"column:jabatan_id;foreignKey:JabatanID"`
	TahunPengangkatan  datatypes.Date `gorm:"column:tahun_pengangkatan"`
	TahunPemberhentian datatypes.Date `gorm:"column:tahun_pemberhentian"`
}

// TableName returns the table name for Jabatan
func (*RiwayatJabatan) TableName() string {
	return "riwayat_jabatan" // Ganti dengan nama tabel yang diinginkan
}
