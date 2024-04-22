package models

type (
	RegisterRequest struct {
		FirstName       string `json:"firstName" validate:"required"`
		LastName        string `json:"lastName" validate:"required"`
		Password        string `json:"password" validate:"required"`
		ConfirmPassword string `json:"confirmPassword" validate:"required,eqfield=Password"`
		Email           string `json:"email" validate:"required,email"`
		RoleID          uint   `json:"roleId" validate:"required"`
	}
)
