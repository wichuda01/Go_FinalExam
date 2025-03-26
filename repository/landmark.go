package repository

import (
	"Go_FinalExam/models"

	"gorm.io/gorm"
)

type LandmarkRepository interface {
	GetAllLandmark() (*[]models.Landmark, error)
	// GetLandmarkByName(name string) (*[]models.Landmark, error)
	// InsertLandmark(data models.Landmark) *gorm.DB
	// GetLandmarkByID(lid int) (*models.Landmark, error)
	// GetLandmarkByUID(uid int) (*[]models.Landmark, error)
}

type landmarkDB struct {
	db *gorm.DB
}

func NewLandmarkRePository(gormdb *gorm.DB) LandmarkRepository {
	return landmarkDB{db: gormdb}
}

func (l landmarkDB) GetAllLandmark() (*[]models.Landmark, error) {
	landmarks := []models.Landmark{}
	res := l.db.Find(&landmarks)
	if res.Error != nil {
		return nil, res.Error
	}
	return &landmarks, nil
}
