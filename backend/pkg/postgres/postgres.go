package postgres

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Config struct {
	Host               string `config:"host"`
	Port               string `config:"port"`
	User               string `config:"user"`
	Password           string `config:"password"`
	DBName             string `config:"db_name"`
	MaxIdleConnections int    `config:"max_idle_connection"`
	MaxConnections     int    `config:"max_connections"`
}

func Connect(config Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		config.Host, config.User, config.Password, config.DBName, config.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})
	if err != nil {
		return nil, err
	}

	sqlDb, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDb.SetMaxIdleConns(config.MaxIdleConnections)
	sqlDb.SetMaxOpenConns(config.MaxConnections)

	return db, nil
}
