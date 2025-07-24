package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID           uint   `json:"id"`
	Username     string `json:"username" gorm:"uniqueIndex" binding:"required"`
	Email        string `json:"email" gorm:"unique;not null"`
	Password     string `json:"password" binding:"required"`
	Role         string `json:"role" binding:"required"`
	RefreshToken string `json:"-"`
}

type UpdateUserInput struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Role     string `json:"role" binding:"required"`
}

type RegisterUserInput struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role" binding:"required"`
}
