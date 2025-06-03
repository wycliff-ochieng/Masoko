package db

import "github.com/wycliff-ochieng/cmd/migrate"

type Storage interface {
	CreateProduct(*migrate.Product) error
	//UpdateProduct()
	//DeleteProduct()
	GetProducts() ([]*migrate.Product, error)
}
