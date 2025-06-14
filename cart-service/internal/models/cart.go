package models

import "time"




type Cart struct{
	ID string
	UserID string
	SessionId string
	Status string
	Items []Item
	Total float64
	CreatedAt time.Time
	UpdatedAt time.Time
}


type Item struct{
	ProductID string
	Quantity int
	Price float64
	Subtotal float64
}