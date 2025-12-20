package schema

import (
	"time"

	"github.com/google/uuid"
)

type CreateOrder struct{
	UserID uuid.UUID `json:"userID" binding:"required"`
	OrderName string `json:"orderName" binding:"required"`
	Amount float64 `json:"amount" binding:"required"`
}
type OrderUpdate struct{
	OrderName *string `json:"order_name"`
	Amount *float64 `json:"amount"`
	Status *string `json:"status"`
}
type OrderResponse struct {
	ID uuid.UUID `json:"id"`
	UserID uuid.UUID `json:"userID"`
	OrderName string `json:"order_name"`
	Amount float64 `json:"amount"`
	CreatedAt time.Time `json:"createdAt"`
	Status string `json:"status"`
}

type UserOrderResponse struct{
	ID uuid.UUID `json:"id"`
	OrderName string `json:"order_name"`
	Amount float64 `json:"amount"`
	CreatedAt time.Time `json:"createdAt"`
	Status string `json:"status"`
}