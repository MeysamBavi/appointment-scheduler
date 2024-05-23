package models

import "gorm.io/gorm"

type Employee struct {
	gorm.Model
	Business   Business
	UserID     uint `gorm:"index:idx_user_business_unique,unique"`
	BusinessID uint `gorm:"index:idx_user_business_unique,unique"`
	// TODO: permissions
}
