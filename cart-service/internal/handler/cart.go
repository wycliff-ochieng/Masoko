package handler

import (
	"log"
	"net/http"
)

type Cart struct {
	l *log.Logger
}

func NewCart(l *log.Logger) *Cart {
	return &Cart{l}
}

func (c *Cart) CreateCart(w http.ResponseWriter, r *http.Request) error {
	//c.store.CreateCart()
	return nil
}
