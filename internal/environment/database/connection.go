package database

import (
	"log"

	"github.com/viniokamoto/go-store/internal/environment"
	"github.com/viniokamoto/go-store/internal/environment/database/seeds"
	"github.com/viniokamoto/go-store/source/user/domain/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDBConnection() {
	var err error
	database, err := gorm.Open(postgres.New(
		postgres.Config{
			DSN:                  environment.Config.DatabaseURL,
			PreferSimpleProtocol: true,
		},
	), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	DB = database

	DB.AutoMigrate(retrieveAll()...)

	RunSeed(DB)
	// db.Migrator().DropTable(retrieveAll()...)

}

func retrieveAll() []interface{} {
	return []interface{}{
		models.Role{},
		models.User{},
	}
}

func RunSeed(db *gorm.DB) error {
	for _, seed := range seeds.All() {
		err := seed.Run(db)
		if err != nil {
			return err
		}
	}
	return nil
}
