package mysql

import (
	"errors"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMysql() (*gorm.DB, error) {
	dsn := "claran:chr070309@tcp(localhost:3306)/Demo_logging_and_registering?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, errors.New("failed to connect to database")
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, errors.New("failed to connect to database")
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db, nil
}
