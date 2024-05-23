package models

import "gorm.io/gorm"

type BusinessService struct {
	gorm.Model
	Business Business `gorm:"embedded"`
	Name     string
}
