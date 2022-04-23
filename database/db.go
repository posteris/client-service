package database

import (
	"log"

	"github.com/posteris/client-service/models"
	"github.com/posteris/database"
	"gorm.io/gorm"
)

var factory *database.Factory

//
const defaultInstance string = "default"

//InitDatabase function to start the database conection
func InitDatabase() {
	factory = database.GetFactory()

	created := factory.AddFromEnv(defaultInstance, nil)

	if !created {
		log.Fatal("Database connection error!")
	}

	automigrate()
}

//automigrate function make migrations
func automigrate() {
	instance := factory.GetInstance(defaultInstance)

	err := instance.AutoMigrate(&models.Client{})
	if err != nil {
		log.Fatalf("Migration error: %v", err)
	}
}

// obtains the database instance
func GetInstance() *gorm.DB {
	return factory.GetInstance(defaultInstance)
}
