package models

import (
	"database/sql"
)

type CartItem struct {
	CartItemID int          `gorm:"column:cart_item_id;type:int(11);primary_key;AUTO_INCREMENT"`
	CartID     int          `gorm:"column:cart_id;type:int(11);NOT NULL"`
	ProductID  int          `gorm:"column:product_id;type:int(11);NOT NULL"`
	Quantity   int          `gorm:"column:quantity;type:int(11);NOT NULL"`
	CreatedAt  sql.NullTime `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt  sql.NullTime `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	Cart       Cart         `gorm:"foreignkey:CartID"`
	Product    Product      `gorm:"foreignkey:ProductID"`
}

func (m *CartItem) TableName() string {
	return "cart_item"
}
