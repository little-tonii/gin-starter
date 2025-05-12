package config

import (
	"fmt"
	"health-care-system/internal/infrastructure/model"
	"health-care-system/internal/shared/constant"
	"sync"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	database *gorm.DB
	once     sync.Once
)

func seedRolesRawSQL(db *gorm.DB) error {
	rolesToSeed := []string{
		"patient",
		"doctor",
		"nurse",
		"admin",
		"pharmacist",
		"lab_technician",
		"insurance_provider",
	}
	for _, roleName := range rolesToSeed {
		sql := fmt.Sprintf("INSERT INTO roles (name) VALUES ('%s') ON CONFLICT (name) DO NOTHING;", roleName)
		result := db.Exec(sql)
		if result.Error != nil {
			return fmt.Errorf("Thất bại khi insert role '%s' với câu truy vấn thô: %w", roleName, result.Error)
		}
	}
	return nil
}

func InitializeDatabase() error {
	var error error
	once.Do(func() {
		destination := fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=%s",
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

		if err := connection.AutoMigrate(&model.RoleModel{}, &model.UserModel{}); err != nil {
			error = err
			return
		}

		if err := seedRolesRawSQL(connection); err != nil {
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
