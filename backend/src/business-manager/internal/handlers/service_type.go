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

func GetServiceType(db *gorm.DB, serviceID uint) (*models.ServiceType, error) {
	var serviceType models.ServiceType
	result := db.Find(&serviceType, "id=?", serviceID)

	if result.RowsAffected == 0 {
		return nil, ErrNoRows
	}

	return &serviceType, result.Error
}
