// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package Database

const TableNameUser = "user"

// User mapped from table <user>
type User struct {
	ID         int64  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Username   string `gorm:"column:username" json:"username"`
	Password   string `gorm:"column:password" json:"password"`
	Email      string `gorm:"column:email" json:"email"`
	PengurusID int64  `gorm:"column:pengurus_id" json:"pengurus_id"`
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
