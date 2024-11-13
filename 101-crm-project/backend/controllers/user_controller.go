package controllers

import (
	"backend/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type UserController struct {
	DB *gorm.DB
}

// Crear un nuevo usuario
func (controller *UserController) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Datos no válidos"})
		return
	}

	// Guardar el usuario en la base de datos
	if err := controller.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error al guardar el usuario"})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// Obtener todos los usuarios
func (controller *UserController) GetUsers(c *gin.Context) {
	var users []models.User
	if err := controller.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error al obtener los usuarios"})
		return
	}

	c.JSON(http.StatusOK, users)
}
