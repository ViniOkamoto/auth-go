package models

import (
	"gorm.io/gorm"

	"github.com/google/uuid"
)

type (
	User struct {
		gorm.Model
		ID        string `gorm:"type:uuid;primary_key;"`
		FirstName string
		LastName  string
		Password  string
		Email     string `gorm:"unique"`
		RoleID    uint   `gorm:"not null;default:1"`

		Role Role `gorm:"foreignKey:RoleID"`
	}

	UserCreateRequest struct {
		FirstName       string `json:"firstName" validate:"required"`
		LastName        string `json:"lastName" validate:"required"`
		Password        string `json:"password" validate:"required"`
		ConfirmPassword string `json:"confirmPassword" validate:"required,eqfield=Password"`
		Email           string `json:"email" validate:"required,email"`
		RoleID          uint   `json:"roleId" validate:"required"`
	}

	UserUpdateRequest struct {
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		Email     string `json:"email"`
		RoleID    uint   `json:"roleId"`
	}

	UserResponse struct {
		ID        string `json:"id"`
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		Email     string `json:"email"`
		Role      Role   `json:"role"`
	}
)

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.ID = uuid.New().String()
	return
}
