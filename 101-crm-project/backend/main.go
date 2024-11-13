package main

import (
	"backend/config"
	"backend/routes"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Conectar a la base de datos
	config.ConnectDatabase()

	// Configurar el router
	router := gin.Default()

	// Configurar el middleware CORS (Antes de registrar rutas)
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Cambia esto si el frontend cambia de puerto/origen
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Registrar rutas
	routes.RegisterRoutes(router, config.DB)

	// Ruta de prueba adicional
	router.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "La API está funcionando correctamente",
		})
	})

	// Iniciar el servidor
	log.Println("Iniciando el servidor en http://localhost:8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
