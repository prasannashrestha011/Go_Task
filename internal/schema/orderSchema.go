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
type OrderResponse struct {
	ID uuid.UUID `json:"id"`
	UserID uuid.UUID `json:"userID"`
	Amount float64 `json:"amount"`
	CreatedAT time.Time `json:"createdAt"`
	Status string `json:"status"`
}

type UserOrderResponse struct{
	ID uuid.UUID `json:"id"`
	Amount float64 `json:"amount"`
	CreatedAt time.Time `json:"createdAt"`
	Status string `json:"status"`
}