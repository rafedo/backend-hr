// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package Database

const TableNamePenempatan = "penempatan"

// Penempatan mapped from table <penempatan>
type Penempatan struct {
	ID        int64 `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	WilayahID int64 `gorm:"column:wilayah_id" json:"wilayah_id"`
	DaerahID  int64 `gorm:"column:daerah_id" json:"daerah_id"`
	CabangID  int64 `gorm:"column:cabang_id" json:"cabang_id"`
	RantingID int64 `gorm:"column:ranting_id" json:"ranting_id"`
}

// TableName Penempatan's table name
func (*Penempatan) TableName() string {
	return TableNamePenempatan
}
