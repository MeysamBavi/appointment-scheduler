package models

import "gorm.io/gorm"

type Employee struct {
	gorm.Model
	UserId   uint
	Business Business `gorm:"embedded"`
	// TODO: permissions
}
