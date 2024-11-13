package models

import "gorm.io/gorm"

// Modelo de usuario
type User struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email" gorm:"unique"`
}
