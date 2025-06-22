package models

import (
	"time"

	"github.com/google/uuid"
)

type Cart struct {
	ID        string
	UserID    uuid.UUID
	SessionId string
	Status    string
	Items     []CartItem
	Total     float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CartItem struct {
	CartID    string
	ProductID string
	Quantity  int
	Price     float64
	Subtotal  float64
}

func (c *Cart) CalculateCartTotal() {

	var total float64

	for _, item := range c.Items {
		total += item.Price * float64(item.Quantity)
	}
	c.Total = total
}
