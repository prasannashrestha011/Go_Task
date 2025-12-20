package schema

import (
	"time"

	"github.com/google/uuid"
)

type CreateOrder struct{
	UserID uuid.UUID `json:"userID" binding:"required"`
	OrderName string `json:"orderName" binding:"required"`
	Price float64 `json:"price" binding:"required"`
	Quantity int `json:"quantity" binding:"required"`
}
type OrderUpdate struct{
	OrderName *string `json:"order_name"`
	Price *float64 `json:"price" binding:"required"`
	Quantity *int `json:"quantity" biniding:"required"`
	Status *string `json:"status"`
}
type OrderResponse struct {
	ID uuid.UUID `json:"id"`
	UserID uuid.UUID `json:"userID"`
	OrderName string `json:"order_name"`
	Quantity int `json:"quantity" binding:"required"`
	Price float64 `json:"price" binding:"required"`
	Amount float64 `json:"amount"`
	CreatedAt time.Time `json:"createdAt"`
	Status string `json:"status"`
}

type UserOrderResponse struct{
	ID uuid.UUID `json:"id"`
	OrderName string `json:"order_name"`
	Quantity int `json:"quantity" binding:"required"`
	Price float64 `json:"price" binding:"required"`
	Amount float64 `json:"amount"`
	CreatedAt time.Time `json:"createdAt"`
	Status string `json:"status"`
}