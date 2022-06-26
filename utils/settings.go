package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Settings interface{
	DatabaseSettings() (AppConfig,DBConfig)
}

type AppConfig struct {
	AppName string
	AppEnv  string
	AppPort string
}

type DBConfig struct {
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
	DBDriver   string
}

func Initialize(appConfig AppConfig)  {
	fmt.Println("welcome to" + appConfig.AppName)
}

func getEnv(key, fallback string) string  {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func DatabaseSettings() (AppConfig,DBConfig) {
	var appConfig = AppConfig{}
	var dbConfig = DBConfig{}

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error on loading .env file")
	}

	appConfig.AppName = getEnv("APP_NAME", "Momen")
	appConfig.AppEnv = getEnv("APP_ENV", "development")
	appConfig.AppPort = getEnv("APP_PORT", "9000")

	dbConfig.DBHost = getEnv("DB_HOST", "localhost")
	dbConfig.DBUser = getEnv("DB_USER", "root")
	dbConfig.DBPassword = getEnv("DB_PASSWORD", "")
	dbConfig.DBName = getEnv("DB_NAME", "momen")
	dbConfig.DBPort = getEnv("DB_PORT", "3306")
	dbConfig.DBDriver = getEnv("DB_DRIVER", "mysql")

	return appConfig, dbConfig
}