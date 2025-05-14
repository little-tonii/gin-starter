package config

import (
	"fmt"
	"gin-starter/internal/infrastructure/model"
	"gin-starter/internal/shared/constant"
	"sync"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	database     *gorm.DB
	databaseOnce sync.Once
)

func InitializeDatabase() error {
	var error error
	databaseOnce.Do(func() {
		destination := fmt.Sprintf(
			"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s TimeZone=%s",
			constant.Environment.POSTGRES_HOST,
			constant.Environment.POSTGRES_PORT,
			constant.Environment.POSTGRES_USER,
			constant.Environment.POSTGRES_PASSWORD,
			constant.Environment.POSTGRES_DB,
			constant.Environment.POSTGRES_SSL_MODE,
			constant.Environment.POSTGRES_TIME_ZONE,
		)

		connection, err := gorm.Open(postgres.Open(destination), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})

		if err != nil {
			error = err
			return
		}

		sqlDatabase, err := connection.DB()

		if err != nil {
			error = err
			return
		}

		sqlDatabase.SetMaxIdleConns(5)
		sqlDatabase.SetMaxOpenConns(20)
		sqlDatabase.SetConnMaxLifetime(10 * time.Minute)

		if err := connection.AutoMigrate(&model.UserModel{}); err != nil {
			error = err
			return
		}

		database = connection
	})
	return error
}

func CloseDatabase() error {
	sqlDatabase, error := database.DB()
	if error != nil {
		return error
	}
	return sqlDatabase.Close()
}

func GetDatabase() *gorm.DB {
	return database
}
