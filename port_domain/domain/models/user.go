package models

type User struct {
	Id   int    `gorm:"column:id"`
	Name string `gorm:"not null"`
	Timestamp `gorm:"embedded"`
}
