package Migration

type User struct {
	Id         uint     `gorm:"primaryKey"`
	Username   string   `gorm:"column:username"`
	Password   string   `gorm:"column:password"`
	Email      string   `gorm:"column:email"`
	PengurusID Pengurus `gorm:"column:pengurus_id;foreignKey:PengurusID"`
}

// TableName returns the table name for User
func (*User) TableName() string {
	return "user" // Ganti dengan nama tabel yang diinginkan
}
