package models

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	ID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserID     uuid.UUID
	OrderName string
	Price     float64
	Quantity int
	Amount    float64
	Status     string
	CreatedAt time.Time
}

func (Order) TableName() string{
	return `order`
}