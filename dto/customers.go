package dto

import "time"

type CustomerDto []CustomerDtoElement

type CustomerDtoElement struct {
	CustomerID  int64  `json:"CustomerID"`
	FirstName   string `json:"FirstName"`
	LastName    string `json:"LastName"`
	Email       string `json:"Email"`
	PhoneNumber string `json:"PhoneNumber"`
	Address     string `json:"Address"`
	// Password    string `json:"Password"`
	CreatedAt AtedAt `json:"CreatedAt"`
	UpdatedAt AtedAt `json:"UpdatedAt"`
}

type AtedAt struct {
	Time  time.Time `json:"Time"`
	Valid bool      `json:"Valid"`
}
