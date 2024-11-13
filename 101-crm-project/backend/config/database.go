package config

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Conectar a la base de datos
func ConnectDatabase() {
	var err error
	dsn := "root:root@tcp(127.0.0.1:3306)/projectDB?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error al conectar con la base de datos:", err)
	}
	fmt.Println("Conexión exitosa a la base de datos")
}
