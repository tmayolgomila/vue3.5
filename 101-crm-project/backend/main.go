package main

import (
	"backend/config"
	"backend/routes"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	//Load configuration and connect to the database
	db, err := config.SetupDatabase()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer db.Close()
	//Router config
	router := gin.Default()
	routes.RegisterRoutes(router)

	//Start the server
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Start server error: %v", err)
	}
}
