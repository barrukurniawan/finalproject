package db

import (
	"fmt"
	"log"
	"mygram/server/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host   = "localhost"
	port   = "5432"
	user   = "postgres"
	pass   = "sman60jakarta"
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
