package main

import (
	"fmt"
	"os"

	"github.com/Chandore998/rental/pkg/utils"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initEnv() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
		panic("Error loading in Env ")
	}
}

func configDb() (*gorm.DB, error) {
	type PgInput struct {
		port     int
		host     string
		user     string
		password string
		dbname   string
		sslmode  string
	}

	var port int
	_, err := fmt.Sscanf(os.Getenv("DB_PORT"), "%d", &port)
	if err != nil {
		return nil, fmt.Errorf("error converting PORT to int: %v", err)
	}

	dbInput := PgInput{port: port, host: os.Getenv("DB_HOST"), user: os.Getenv("DB_USER"), dbname: os.Getenv("DB_NAME"), sslmode: os.Getenv("DB_SSLMODE"), password: os.Getenv("DB_PASSWORD")}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s", dbInput.host, dbInput.user, dbInput.password, dbInput.dbname, dbInput.port, dbInput.sslmode)

	utils.InfoLog.Printf("DSN: %s", dsn)
	dbRes, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	utils.InfoLog.Println("Database connection established successfully")
	return dbRes, nil
}

func main() {
	initEnv()
	db, err := configDb()
	if err != nil {
		utils.ErrorLog.Fatalf("Error connecting to the database: %v", err)
	}
	utils.InfoLog.Printf("Database connection successful: %v", db)
}
