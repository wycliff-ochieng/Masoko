package handler

import (
	"log"

	service "github.com/wycliff-ochieng/cart-service/internal/services"

	"net/http"
)

type Cart struct {
	l *log.Logger
	k *service.CartService
}

func NewCartHandler(l *log.Logger, k *service.CartService) *Cart {
	return &Cart{
		l: l,
		k: k,
	}
}

func (c *Cart) CreateCart(w http.ResponseWriter, r *http.Request) error {
	//c.store.CreateCart()
	return nil
}

func (c *Cart) GetCart(w http.ResponseWriter, r *http.Request) {
	return
}
