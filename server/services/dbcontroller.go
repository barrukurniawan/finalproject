package services

import (
	"gorm.io/gorm"
)

type HandlersController struct {
	db *gorm.DB
}

func Handler_MyGram(db *gorm.DB) *HandlersController {
	return &HandlersController{
		db: db,
	}
}
