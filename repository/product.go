package repository

import (
	"Go_FinalExam/models"

	"gorm.io/gorm"
)

type ProductRepository interface {
	GetAllProduct() (*[]models.Product, error)
	// GetLandmarkByName(name string) (*[]models.Landmark, error)
	// InsertLandmark(data models.Landmark) *gorm.DB
	// GetLandmarkByID(lid int) (*models.Landmark, error)
	// GetLandmarkByUID(uid int) (*[]models.Landmark, error)
}

type productkDB struct {
	db *gorm.DB
}

func NewProductRePository(gormdb *gorm.DB) ProductRepository {
	return customerDB{db: gormdb}
}

func (l customerDB) GetAllProduct() (*[]models.Product, error) {
	products := []models.Product{}
	res := l.db.Find(&products)
	if res.Error != nil {
		return nil, res.Error
	}
	return &products, nil
}
