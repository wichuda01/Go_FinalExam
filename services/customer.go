package services

import (
	"Go_FinalExam/models"
	"Go_FinalExam/repository"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type CustomertService interface {
	GetAllCustomer() (*[]models.Customer, error)
	Login(email string, password string) (*models.Customer, error)
	EditProfile(id uint, address string) *gorm.DB
	ChangePassword(newpassword string, oldpassword string) *gorm.DB
}

type customerDB struct {
	db *gorm.DB
}

func NewCustomerService(gormdb *gorm.DB) CustomertService {
	return customerDB{db: gormdb}
}

func (c customerDB) GetAllCustomer() (*[]models.Customer, error) {
	customerRepo := repository.NewCustomerRePository(c.db)
	customers, err := customerRepo.GetAllCustomer()
	if err != nil {
		return nil, err
	}
	return customers, nil
}

func (c customerDB) Login(email string, password string) (*models.Customer, error) {
	custRepo := repository.NewCustomerRePository(c.db)
	user, err := custRepo.GetEmailCustomer(email)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("user not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid password")
	}

	return user, nil
}

func (c customerDB) EditProfile(id uint, address string) *gorm.DB {
	profileRepo := repository.NewCustomerRePository(c.db)
	res := profileRepo.UdateProfle(id, address)
	if res.Error != nil {
		return res
	}
	return res
}

func (c customerDB) ChangePassword(newpassword string, oldpassword string) *gorm.DB {
	// customerRepo:=repository.NewCustomerRePository(c.db)
	// res:= customerRepo.GetCustomerByPassword(oldpassword)
	// if err !=nil {
	// 	return res
	// }
	return nil
}
