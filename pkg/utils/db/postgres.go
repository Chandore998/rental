package db

import (
	"fmt"
	"os"

	logger "github.com/Chandore998/rental/pkg/utils/logger"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// postgres database struct
type PgInput struct {
	port     int
	host     string
	user     string
	password string
	dbname   string
	sslmode  string
}

func initEnv() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
		panic("Error loading in Env ")
	}
}

func ConfigDb() (*gorm.DB, error) {
	// load env file
	initEnv()

	dbInput := PgInput{port: getEnvInt("DB_PORT"), host: os.Getenv("DB_HOST"), user: os.Getenv("DB_USER"), dbname: os.Getenv("DB_NAME"), sslmode: os.Getenv("DB_SSLMODE"), password: os.Getenv("DB_PASSWORD")}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s", dbInput.host, dbInput.user, dbInput.password, dbInput.dbname, dbInput.port, dbInput.sslmode)

	logger.InfoLog.Printf("DSN: %s", dsn)
	dbRes, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	logger.InfoLog.Println("Database connection established successfully")
	return dbRes, nil
}

func getEnvInt(key string) int {
	valStr := os.Getenv(key)
	var valInt int
	_, err := fmt.Sscanf(valStr, "%d", &valInt)
	if err != nil {
		return 0
	}
	return valInt
}
