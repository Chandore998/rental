package main

import (
	"fmt"

	"github.com/joho/godotenv"

	db "github.com/Chandore998/rental/pkg/utils/db"
	logger "github.com/Chandore998/rental/pkg/utils/logger"
)

func initEnv() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
		panic("Error loading in Env ")
	}
}

func main() {
	initEnv()
	database, err := db.ConfigDb()
	if err != nil {
		logger.ErrorLog.Fatalf("Error connecting to the database: %v", err)
	}
	logger.InfoLog.Printf("Database connection successful: %v", database)
}
