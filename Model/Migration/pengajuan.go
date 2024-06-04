package Migration

type Pengajuan struct {
	Id              uint   `gorm:"primaryKey"`
	AsalKantorID    uint   `gorm:"column:asal_kantor_id;not null"`
	TujuanKantorID  uint   `gorm:"column:tujuan_kantor_id;not null"`
	StatusPengajuan string `gorm:"type:column:status_pengajuan;type: varchar(50)"`
	Keterangan      string `gorm:"column:keterangan;type:text"`
	File            string `gorm:"column:file;type:varchar(50)"`
}

// TableName returns the table name for Pengajuan
func (*Pengajuan) TableName() string {
	return "pengajuan" // Ganti dengan nama tabel yang diinginkan
}
