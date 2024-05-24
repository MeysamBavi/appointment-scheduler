package handlers

import (
	"github.com/MeysamBavi/appointment-scheduler/backend/src/business-manager/internal/models"
	"gorm.io/gorm"
)

func CreateEmployee(db *gorm.DB, data *models.Employee) error {
	return db.Create(data).Error
}

func GetEmployees(db *gorm.DB, businessId uint) ([]models.Employee, error) {
	var employees []models.Employee
	result := db.Preload("Business").Find(&employees, "business_id=?", businessId)

	return employees, result.Error
}

func GetEmployee(db *gorm.DB, employeeId, businessID uint) (*models.Employee, error) {
	var business models.Employee
	result := db.Preload("Business").Find(&business, "id=? and business_id=?", employeeId, businessID)

	if result.RowsAffected == 0 {
		return nil, ErrNoRows
	}

	return &business, result.Error
}

func DeleteEmployee(db *gorm.DB, employeeId, businessID uint) error {
	result := db.Where("id=? and business_id=?", employeeId, businessID).Delete(&models.Employee{})
	if result.RowsAffected == 0 {
		return ErrNoRows
	}

	return result.Error
}
