package db

import "github.com/wycliff-ochieng/cmd/migrate"

type Storage interface {
	CreateProduct(*migrate.Product) error
	//UpdateProduct()
	//DeleteProduct()
	GetProductByID(int) (*migrate.Product, error)
	GetProducts() ([]*migrate.Product, error)
}
