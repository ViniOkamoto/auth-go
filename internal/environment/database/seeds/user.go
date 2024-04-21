package seeds

import (
	"github.com/viniokamoto/go-store/source/user/domain/models"
	"gorm.io/gorm"
)

func createUser(db *gorm.DB, firstName string, lastName, username string, email string, password string, roleID uint) error {
	return db.Create(&models.User{
		FirstName: firstName,
		LastName:  lastName,
		Username:  username,
		Email:     email,
		Password:  password,
		RoleID:    roleID,
	}).Error
}

func UserSeed() Seed {
	return Seed{
		Name: "Seeding users",
		Run: func(db *gorm.DB) error {
			err := createUser(db, "Vinicius", "Okamoto", "viniokamoto", "leo_kamoto@hotmail.com", "123456", 3)
			if err != nil {
				return err
			}
			return err
		},
	}
}
