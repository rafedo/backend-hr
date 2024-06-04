package Migration

import "gorm.io/datatypes"

type DetailPengajuan struct {
	Id                 uint           `gorm:"primaryKey"`
	PengurusID         Pengurus       `gorm:"column:pengurus_id;foreignKey:PengurusID"`
	JabatanID          Jabatan        `gorm:"column:jabatan_id;foreignKey:JabatanID"`
	PengajuanID        Pengajuan      `gorm:"column:pengajuan_id;foreignKey:PengajuanID"`
	Berkas             string         `gorm:"column:berkas"`
	TahunPengangkatan  datatypes.Date `gorm:"column:tahun_pengangkatan"`
	TahunPemberhentian datatypes.Date `gorm:"column:tahun_pemberhentian"`
}

// TableName returns the table name for Jabatan
func (*DetailPengajuan) TableName() string {
	return "detail_pengajuan" // Ganti dengan nama tabel yang diinginkan
}
