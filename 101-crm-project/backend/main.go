package main

import (
	"backend/config"
	"backend/routes"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	// Conectar a la base de datos
	config.ConnectDatabase()

	// Configurar el router
	router := gin.Default()
	routes.RegisterRoutes(router, config.DB) // Pasamos la conexión de la base de datos

	// Iniciar el servidor
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
