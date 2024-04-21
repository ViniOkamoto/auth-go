package models

import (
	"gorm.io/gorm"
)

type (
	User struct {
		gorm.Model
		FirstName string
		LastName  string
		Username  string
		Password  string
		Email     string `gorm:"unique"`
		RoleID    uint   `gorm:"not null;default:1  " `
	}

	UserRequest struct {
		FirstName       string `json:"first_name"`
		LastName        string `json:"last_name"`
		Username        string `json:"username"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirm_password"`
		Email           string `json:"email"`
		RoleID          uint   `json:"role_id"`
	}

	UserResponse struct {
		ID       uint   `json:"id"`
		Username string `json:"username"`
		Email    string `json:"email"`
		Role     string `json:"role"`
	}
)
