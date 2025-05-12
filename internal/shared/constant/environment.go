package constant

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type environment struct {
	POSTGRES_HOST      string
	POSTGRES_PORT      string
	POSTGRES_SSL_MODE  string
	POSTGRES_TIME_ZONE string
	POSTGRES_USER      string
	POSTGRES_PASSWORD  string
	POSTGRES_DB        string
	JWT_SECRET_KEY     string
}

var Environment *environment

func LoadEnvironment() []error {
	errorList := make([]error, 0)
	env, _ := os.LookupEnv("ENVIRONMENT")
	if env != "production" {
		error := godotenv.Load(".env")
		if error != nil {
			errorList = append(errorList, error)
			return errorList
		}
	}
	databaseHost, exists := os.LookupEnv("POSTGRES_HOST")
	if !exists {
		errorList = append(errorList, errors.New("Biến môi trường POSTGRES_HOST chưa được thiết lập"))
	}
	if env != "production" {
		databaseHost = "localhost"
	}
	databasePort, exists := os.LookupEnv("POSTGRES_PORT")
	if !exists {
		errorList = append(errorList, errors.New("Biến môi trường POSTGRES_PORT chưa được thiết lập"))
	}
	databaseUser, exists := os.LookupEnv("POSTGRES_USER")
	if !exists {
		errorList = append(errorList, errors.New("Biến môi trường POSTGRES_USER chưa được thiết lập"))
	}
	databasePassword, exists := os.LookupEnv("POSTGRES_PASSWORD")
	if !exists {
		errorList = append(errorList, errors.New("Biến môi trường POSTGRES_PASSWORD chưa được thiết lập"))
	}
	databaseName, exists := os.LookupEnv("POSTGRES_DB")
	if !exists {
		errorList = append(errorList, errors.New("Biến môi trường POSTGRES_DB chưa được thiết lập"))
	}
	databaseSSLMode, exists := os.LookupEnv("POSTGRES_SSL_MODE")
	if !exists {
		errorList = append(errorList, errors.New("Biến môi trường POSTGRES_SSL_MODE chưa được thiết lập"))
	}
	databaseTimeZone, exists := os.LookupEnv("POSTGRES_TIME_ZONE")
	if !exists {
		errorList = append(errorList, errors.New("Biến môi trường POSTGRES_TIME_ZONE chưa được thiết lập"))
	}
	jwtSecretKey, exists := os.LookupEnv("JWT_SECRET_KEY")
	if !exists {
		errorList = append(errorList, errors.New("Biến môi trường JWT_SECRET_KEY chưa được thiết lập"))
	}
	Environment = &environment{
		POSTGRES_HOST:      databaseHost,
		POSTGRES_PORT:      databasePort,
		POSTGRES_USER:      databaseUser,
		POSTGRES_PASSWORD:  databasePassword,
		POSTGRES_DB:        databaseName,
		POSTGRES_SSL_MODE:  databaseSSLMode,
		POSTGRES_TIME_ZONE: databaseTimeZone,
		JWT_SECRET_KEY:     jwtSecretKey,
	}
	return errorList
}
