package config

import (
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() {
	dsn := AppConfig.Database.Dsn
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to initialize database, error: %v", err)
	}
	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(AppConfig.Database.MaxIdleConns)
	sqlDB.SetMaxOpenConns(AppConfig.Database.MaxOpenConns)
	// sqlDB.SetConnMaxIdleTime(5)
	sqlDB.SetConnMaxLifetime(time.Hour)

	if err != nil {
		log.Fatalf("Failed to config database, error: %v", err)
	}
}
