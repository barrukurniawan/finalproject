package main

import (
	"mygram/db"
	_ "mygram/docs"
	"mygram/server"
	"mygram/server/services"
	"os"
)

// @title Barru Final Project
// @description Final Project Golang
// @version v1.0
// @termsOfService http://swagger.io/terms/
// @BasePath /
// @host finalproject-production-2281.up.railway.app
// @contact.name Barru Kurniawan
// @contact.email barru.kurniawan@gmail.com

func main() {
	var PORT = os.Getenv("PORT")
	db := db.ConnectGorm()
	getDB := services.Handler_MyGram(db)
	mygram := server.UserRouther(getDB)
	mygram.Start(":" + PORT)
}
