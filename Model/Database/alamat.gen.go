// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package Database

const TableNameAlamat = "alamat"

// Alamat mapped from table <alamat>
type Alamat struct {
	ID        int64  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Alamat    string `gorm:"column:alamat" json:"alamat"`
	Kelurahan string `gorm:"column:kelurahan" json:"kelurahan"`
	Kecamatan string `gorm:"column:kecamatan" json:"kecamatan"`
	KabKota   string `gorm:"column:kab_kota" json:"kab_kota"`
	Propinsi  string `gorm:"column:propinsi" json:"propinsi"`
	KodePos   string `gorm:"column:kode_pos" json:"kode_pos"`
}

// TableName Alamat's table name
func (*Alamat) TableName() string {
	return TableNameAlamat
}
