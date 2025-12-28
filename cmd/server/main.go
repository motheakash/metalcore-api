package main

import (
	"metalcore-api/internal/config"
	"metalcore-api/internal/database"
	"metalcore-api/internal/router"
	"os"
)

func main() {
	config.LoadEnv()
	database.ConnectDB()

	r := router.SetupRouter(database.DB)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8090"
	}

	r.Run(":" + port)
}
