package database

import (
	"github.com/kwarabei/Discoverish/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Setup() *gorm.DB {
	dsn := "host=localhost user=stub_user dbname=discoverish_db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.User{})

	return db
}
