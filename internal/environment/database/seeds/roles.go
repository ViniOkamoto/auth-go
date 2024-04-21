package seeds

import (
	"github.com/viniokamoto/go-store/source/user/domain/models"
	"gorm.io/gorm"
)

func createRole(db *gorm.DB, name string) error {
	return db.Create(&models.Role{
		Name: name,
	}).Error
}

func RoleSeed() Seed {
	return Seed{
		Name: "Seeding roles",
		Run: func(db *gorm.DB) error {
			err := createRole(db, "admin")
			if err != nil {
				return err
			}
			err = createRole(db, "store")
			if err != nil {
				return err
			}
			err = createRole(db, "customer")
			if err != nil {
				return err
			}
			return err
		},
	}
}
