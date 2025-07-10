package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username     string `json:"username" gorm:"uniqueIndex" binding:"required"`
	Email        string `json:"email" gorm:"unique;not null"`
	Password     string `json:"password" binding:"required"`
	Role         string `json:"role" binding:"required"`
	RefreshToken string `json:"-"`
}
