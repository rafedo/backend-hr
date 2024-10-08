// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package Database

import (
	"time"
)

const TableNameDetailPengajuan = "detail_pengajuan"

// DetailPengajuan mapped from table <detail_pengajuan>
type DetailPengajuan struct {
	ID                 int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	PengurusID         int64     `gorm:"column:pengurus_id" json:"pengurus_id"`
	JabatanID          int64     `gorm:"column:jabatan_id" json:"jabatan_id"`
	PengajuanID        int64     `gorm:"column:pengajuan_id" json:"pengajuan_id"`
	Berkas             string    `gorm:"column:berkas" json:"berkas"`
	TahunPengangkatan  time.Time `gorm:"column:tahun_pengangkatan" json:"tahun_pengangkatan"`
	TahunPemberhentian time.Time `gorm:"column:tahun_pemberhentian" json:"tahun_pemberhentian"`
}

// TableName DetailPengajuan's table name
func (*DetailPengajuan) TableName() string {
	return TableNameDetailPengajuan
}
