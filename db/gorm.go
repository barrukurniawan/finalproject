package db

import (
	"fmt"
	"log"
	"mygram/server/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host   = "containers-us-west-101.railway.app"
	port   = "5901"
	user   = "postgres"
	pass   = "dDGbTqfhh8DIfxrss7kT"
	dbname = "railway"
)

func ConnectGorm() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, pass, dbname,
	)

	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		panic("failed to connect database")
	} else {
		fmt.Println("Successfully connect to database")
	}

	errors := db.AutoMigrate(models.User{}, models.Photo{}, models.Comment{}, models.SocialMedia{})
	if errors != nil {
		log.Println(errors.Error())
	}

	return db
}
