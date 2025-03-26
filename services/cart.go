package services

import (
	"gorm.io/gorm"
)

type CarttService interface {
}

type cartDB struct {
	db *gorm.DB
}

func NewCartService(gormdb *gorm.DB) CarttService {
	return cartDB{db: gormdb}
}
