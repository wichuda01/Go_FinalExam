package repository

import (
	"Go_FinalExam/models"
	"errors"
	"time"

	"gorm.io/gorm"
)

type CustomerRepository interface {
	GetAllCustomer() (*[]models.Customer, error)
	CheckEmailCustomer(email string) (bool, error)
	GetEmailCustomer(email string) (*models.Customer, error)
	UdateProfle(id uint, address string) *gorm.DB
	GetCustomerByPassword(password string) (*models.Customer, error)
	// GetLandmarkByName(name string) (*[]models.Landmark, error)
	// InsertLandmark(data models.Landmark) *gorm.DB
	// GetLandmarkByID(lid int) (*models.Landmark, error)
	// GetLandmarkByUID(uid int) (*[]models.Landmark, error)
}

type customerDB struct {
	db *gorm.DB
}

func NewCustomerRePository(gormdb *gorm.DB) CustomerRepository {
	return customerDB{db: gormdb}
}

func (c customerDB) GetAllCustomer() (*[]models.Customer, error) {
	customers := []models.Customer{}
	res := c.db.Find(&customers)
	if res.Error != nil {
		return nil, res.Error
	}
	return &customers, nil
}

func (c customerDB) CheckEmailCustomer(email string) (bool, error) {
	var count int64
	customer := models.Customer{}
	// var customer models.Customer
	res := c.db.Model(&customer).Where("email=?", email).Count(&count)
	if res.Error != nil {
		return false, res.Error
	}
	// true == have
	if count > 0 {
		return true, nil
	}
	return false, nil
}

func (c customerDB) GetEmailCustomer(email string) (*models.Customer, error) {
	var customer models.Customer
	res := c.db.Where("email = ?", email).First(&customer)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, res.Error
	}
	return &customer, nil
}
func (c customerDB) GetCustomerByPassword(password string) (*models.Customer, error) {
	customer := models.Customer{}
	res := c.db.Where("password=?", password).Find(&customer)
	if res.Error != nil {
		return nil, res.Error
	}
	return &customer, nil
}

func (c customerDB) UdateProfle(id uint, address string) *gorm.DB {
	profile := models.Customer{}
	res := c.db.Find(&profile, id)
	if res.Error != nil {
		return res
	}
	profile.Address = address
	profile.UpdatedAt.Time = time.Time{}
	res = c.db.Save(&profile)
	if res.Error != nil {
		return res
	}
	return res
}
