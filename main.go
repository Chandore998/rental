package main

import (
	"fmt"
	"os"

	userService "github.com/Chandore998/rental/pkg/users/service"
	db "github.com/Chandore998/rental/pkg/utils/db"
	logger "github.com/Chandore998/rental/pkg/utils/logger"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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
	app := gin.Default()
	database, err := db.ConfigDb()
	if err != nil {
		logger.ErrorLog.Fatalf("Error connecting to the database: %v", err)
	}
	logger.InfoLog.Printf("Database connection successful: %v", database)

	userService.Init(database, app)

	app.SetTrustedProxies([]string{"localhost"})
	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	app.Run(port)
}
