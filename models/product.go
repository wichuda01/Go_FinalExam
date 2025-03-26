// Code generated by sql2gorm. DO NOT EDIT.
package models

import (
	"database/sql"
)

type Product struct {
	ProductID     int          `gorm:"column:product_id;type:int(11);primary_key;AUTO_INCREMENT"`
	ProductName   string       `gorm:"column:product_name;type:varchar(255);NOT NULL"`
	Description   string       `gorm:"column:description;type:text"`
	Price         string       `gorm:"column:price;type:decimal(10,2);NOT NULL"`
	StockQuantity int          `gorm:"column:stock_quantity;type:int(11);NOT NULL"`
	CreatedAt     sql.NullTime `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt     sql.NullTime `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP"`
}

func (m *Product) TableName() string {
	return "product"
}