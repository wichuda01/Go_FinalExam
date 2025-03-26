package services

import (
	"Go_FinalExam/models"
	"Go_FinalExam/repository"

	"gorm.io/gorm"
)

type landmarkService interface {
	GetAllLandmark() (*[]models.Landmark, error)
	// GetLandmarkByName(name string) (*[]models.Landmark, error)
	// AddLandmark(data models.Landmark) *gorm.DB
	// SearchLandmarkByID(lid int) (*models.Landmark, error)
	// SearchLandmarkByUID(uid int) (*[]models.Landmark, error)
}

type landmarkDB struct {
	db *gorm.DB
}

func NewLandmarkService(gormdb *gorm.DB) landmarkService {
	return landmarkDB{db: gormdb}
}

func (l landmarkDB) GetAllLandmark() (*[]models.Landmark, error) {
	landmarkRepo := repository.NewLandmarkRePository(l.db)
	landmarks, err := landmarkRepo.GetAllLandmark()
	if err != nil {
		return nil, err
	}
	return landmarks, nil
}
