// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package Database

const TableNameDaerah = "daerah"

// Daerah mapped from table <daerah>
type Daerah struct {
	ID           int64  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	NamaDaerah   string `gorm:"column:nama_daerah" json:"nama_daerah"`
	AlamatKantor string `gorm:"column:alamat_kantor" json:"alamat_kantor"`
	WilayahID    int64  `gorm:"column:wilayah_id" json:"wilayah_id"`
}

// TableName Daerah's table name
func (*Daerah) TableName() string {
	return TableNameDaerah
}
