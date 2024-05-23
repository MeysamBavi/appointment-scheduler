package handlers

import (
	"errors"

	"github.com/MeysamBavi/appointment-scheduler/backend/src/business-manager/internal/models"
	"gorm.io/gorm"
)

var (
	ErrNoRows = errors.New("no rows affected")
)

func CreateBusiness(db *gorm.DB, data *models.Business) error {
	return db.Create(data).Error
}

func GetBusinesses(db *gorm.DB, userID uint) ([]models.Business, error) {
	var businesses []models.Business
	result := db.Preload("ServiceType").Find(&businesses, "user_id=?", userID)

	return businesses, result.Error
}

func GetBusiness(db *gorm.DB, businessID uint) (*models.Business, error) {
	var business models.Business
	result := db.Preload("ServiceType").Find(&business, "id=?", businessID)

	if result.RowsAffected == 0 {
		return nil, ErrNoRows
	}

	return &business, result.Error
}

func UpdateBusiness(db *gorm.DB, businessID uint, data *models.Business) error {
	result := db.Model(&models.Business{}).Where("id=?", businessID).Updates(*data)
	if result.RowsAffected == 0 {
		return ErrNoRows
	}

	return result.Error
}

func DeleteBusiness(db *gorm.DB, businessID uint) error {
	result := db.Delete(&models.Business{}, businessID)
	if result.RowsAffected == 0 {
		return ErrNoRows
	}

	return result.Error
}
