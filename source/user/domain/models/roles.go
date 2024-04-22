package models

type (
	Role struct {
		ID    uint   `json:"id" gorm:"primarykey"`
		Name  string `json:"name" gorm:"unique;not null"`
		Users []User `json:"users" gorm:"foreignKey:RoleID"`
	}

	RoleResponse struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
	}
)

type RoleType int

const (
	Admin    RoleType = 1
	Store    RoleType = 2
	Customer RoleType = 3
)
