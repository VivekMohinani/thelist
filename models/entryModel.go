package models

import(
	"gorm.io/gorm"
)

type Entry struct {
	gorm.Model
	Fname 	string
	Lname 	string
	State 	string
	Phone 	string
	UserId 	uint `gorm:"user_id"`
}