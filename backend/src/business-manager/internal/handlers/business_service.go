package handlers

import (
	"github.com/MeysamBavi/appointment-scheduler/backend/src/business-manager/internal/models"
	"gorm.io/gorm"
)

func CreateBusinessService(db *gorm.DB, data *models.BusinessService) error {
	return db.Create(data).Error
}

func GetBusinessServices(db *gorm.DB, businessID uint) ([]models.BusinessService, error) {
	var services []models.BusinessService
	result := db.Preload("Business").Find(&services, "business_id=?", businessID)

	return services, result.Error
}

func GetBusinessService(db *gorm.DB, serviceID, businessID uint) (*models.BusinessService, error) {
	var service models.BusinessService
	result := db.Preload("Business").Find(&service, "id=? and business_id=?", serviceID, businessID)

	if result.RowsAffected == 0 {
		return nil, ErrNoRows
	}

	return &service, result.Error
}

func DeleteBusinessService(db *gorm.DB, serviceID, businessID uint) error {
	result := db.Where("id=? and business_id=?", serviceID, businessID).Delete(&models.BusinessService{})
	if result.RowsAffected == 0 {
		return ErrNoRows
	}

	return result.Error
}

func UpdateBusinessService(db *gorm.DB, serviceID uint, data *models.BusinessService) error {
	result := db.Model(&models.BusinessService{}).Where("id=?", serviceID).Updates(*data)
	if result.RowsAffected == 0 {
		return ErrNoRows
	}

	return result.Error
}
