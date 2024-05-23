package handlers

import (
	"github.com/MeysamBavi/appointment-scheduler/backend/src/business-manager/internal/models"
	"gorm.io/gorm"
)

func GetServiceTypes(db *gorm.DB, query string) ([]models.ServiceType, error) {
	var serviceTypes []models.ServiceType
	result := db.Find(&serviceTypes, "name LIKE ?", "%"+query+"%")

	return serviceTypes, result.Error
}
