package repository

import (
	"Go_FinalExam/models"

	"gorm.io/gorm"
)

type CarttRepository interface {
	InsertCart(*models.Cart) *gorm.DB
}

type cartDB struct {
	db *gorm.DB
}

func NewCartCarttRepository(gormdb *gorm.DB) CarttRepository {
	return cartDB{db: gormdb}
}
func (d cartDB) InsertCart(*models.Cart) *gorm.DB {
	return nil
}
