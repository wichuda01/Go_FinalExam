package services

import (
	"Go_FinalExam/models"
	"Go_FinalExam/repository"

	"gorm.io/gorm"
)

type CartItemService interface {
	AddcartItem(data models.CartItem) *gorm.DB
}

type cartIemDB struct {
	db *gorm.DB
}

func NewCartItemService(gormdb *gorm.DB) CartItemService {
	return cartIemDB{db: gormdb}
}

func (d cartIemDB) AddcartItem(data models.CartItem) *gorm.DB {
	diaryRepo := repository.NewCartItemRepository(d.db)
	res := diaryRepo.InsertCartItem(&data)
	if res.Error != nil {
		return res
	}
	return res
}
