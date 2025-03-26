package repository

import (
	"Go_FinalExam/models"

	"gorm.io/gorm"
)

type CartItemRepository interface {
	InsertCartItem(data *models.CartItem) *gorm.DB
}

type cartitemDB struct {
	db *gorm.DB
}

func NewCartItemRepository(gormdb *gorm.DB) CartItemRepository {
	return cartitemDB{db: gormdb}
}
func (d cartitemDB) InsertCartItem(data *models.CartItem) *gorm.DB {
	res := d.db.Create(&data)
	if res.Error != nil {
		return res
	}
	return res

}
