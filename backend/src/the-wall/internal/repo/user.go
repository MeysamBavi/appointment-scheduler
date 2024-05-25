package repo

import (
	"context"
	"fmt"
	"github.com/MeysamBavi/appointment-scheduler/backend/src/the-wall/internal/models"
	"gorm.io/gorm"
)

type User interface {
	Create(ctx context.Context, user *models.User) error
	Get(ctx context.Context, id uint) (*models.User, error)
	GetByPhoneNumber(ctx context.Context, phoneNumber string) (*models.User, error)
}

type userImpl struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) (User, error) {
	if err := db.AutoMigrate(&models.User{}); err != nil {
		return nil, fmt.Errorf("could not auto migrate users: %w", err)
	}

	return &userImpl{
		db: db,
	}, nil
}

func (u userImpl) Create(ctx context.Context, user *models.User) error {
	return u.db.WithContext(ctx).Create(user).Error
}

func (u userImpl) Get(ctx context.Context, id uint) (*models.User, error) {
	m := models.User{
		Model: gorm.Model{ID: id},
	}
	err := u.db.WithContext(ctx).First(&m).Error
	return &m, err
}

func (u userImpl) GetByPhoneNumber(ctx context.Context, phoneNumber string) (*models.User, error) {
	var m models.User
	err := u.db.WithContext(ctx).Where("phone_number = ?", phoneNumber).First(&m).Error
	return &m, err
}
