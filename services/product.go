package services

import (
	"Go_FinalExam/models"
	"Go_FinalExam/repository"

	"gorm.io/gorm"
)

type ProductService interface {
	GetAllProducts() (*[]models.Product, error)
	// GetLandmarkByName(name string) (*[]models.Landmark, error)
	// AddLandmark(data models.Landmark) *gorm.DB
	// SearchLandmarkByID(lid int) (*models.Landmark, error)
	// SearchLandmarkByUID(uid int) (*[]models.Landmark, error)
}

type productDB struct {
	db *gorm.DB
}

func NewProductService(gormdb *gorm.DB) ProductService {
	return productDB{db: gormdb}
}

func (l productDB) GetAllProducts() (*[]models.Product, error) {
	productRepo := repository.NewProductRePository(l.db)
	products, err := productRepo.GetAllProduct()
	if err != nil {
		return nil, err
	}
	return products, nil
}
