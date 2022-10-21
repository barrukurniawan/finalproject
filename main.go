package main

import (
	"mygram/db"
	"mygram/server"
	"mygram/server/services"
)

/*
Barru Kurniawan
My Gram - Final Project
GLN038MNC007
*/

func main() {
	db := db.ConnectGorm()
	getDB := services.Handler_MyGram(db)
	mygram := server.UserRouther(getDB)
	mygram.Start(":5000")
}
