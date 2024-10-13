// models/order.go

package models

import (
	"time"
)

// Order represents an order
type Order struct {
	ID         uint32    `json:"id" gorm:"primary_key;not null;type:int;autoIncrement"`
	CustomerID uint32    `json:"customer_id" validate:"required" gorm:"not null;type:int"`
	OrderDate  time.Time `json:"order_date" validate:"required" gorm:"type:date;not null"`
	Total      float32   `json:"total" validate:"required,gt=0" gorm:"type:decimal(10,4);not null"`
	Status     string    `json:"status" validate:"required,oneof=pending shipped delivered" gorm:"type:varchar(50);not null"`
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime;type:datetime;not null"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"autoUpdateTime;type:datetime;not null"`
}

// TableName returns the table name for the Order model
func (o *Order) TableName() string {
	return "orders"
}
