package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	PhoneNumber string `gorm:"uniqueIndex;not null"`
	Firstname   string
	Lastname    string
}
