package db

import (
	"log"

	"github.com/posteris/client-service/models"
	"github.com/posteris/database"
	"gorm.io/gorm"
)

//variable to store the databases instance
var databaseInstance *gorm.DB

//defaultInstance that defines database instance name. It was created to
//simplify the creation and recovery default database instance
// const defaultInstance string = "default"

//InitDatabase function to start the database conection
func InitDatabase() {
	gormInstance, err := database.NewFromEnv(nil)
	if err != nil {
		log.Fatal("Database connection error!")
	}

	databaseInstance = gormInstance

	Automigrate()
}

func Automigrate() {
	databaseInstance.AutoMigrate(&models.Client{})
}

//GetInstance function to recovery the datault migration. Note that this function
//uses the defaultInstance variabel to return the default instance.
//if you want to have mor than one *gorm.DB instance you should to expose the
//factory variable to allows to requeste the instance directly.
func GetInstance() *gorm.DB {
	return databaseInstance
}
