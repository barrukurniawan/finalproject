package main

import (
	"mygram/db"
	"mygram/server"
	"mygram/server/services"
	"os"
)

/*
Barru Kurniawan
My Gram - Final Project
GLN038MNC007
*/

func main() {
	var PORT = os.Getenv("PORT")
	db := db.ConnectGorm()
	getDB := services.Handler_MyGram(db)
	mygram := server.UserRouther(getDB)
	mygram.Start(":" + PORT)
}
