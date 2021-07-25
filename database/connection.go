package database

import (
	"auth-go/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var Db *gorm.DB
func ConnectDb()  {

	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"

	prgUrl:=postgres.Open(dsn)

	db, err := gorm.Open(prgUrl, &gorm.Config{})
	Db=db
	if err!=nil {
		log.Fatalf("Cannot  connect to the databes %s",err.Error())
	}

	db.AutoMigrate(&models.User{},&models.UserRoles{},&models.Roles{})


}
