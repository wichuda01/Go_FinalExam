package models

import (
	"database/sql"
)

type Cart struct {
	CartID     int          `gorm:"column:cart_id;type:int(11);primary_key;AUTO_INCREMENT"`
	CustomerID int          `gorm:"column:customer_id;type:int(11);NOT NULL"`
	CartName   string       `gorm:"column:cart_name;type:varchar(255)"`
	CreatedAt  sql.NullTime `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt  sql.NullTime `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP"`
}

func (m *Cart) TableName() string {
	return "cart"
}
