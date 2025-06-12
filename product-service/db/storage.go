package db

import (
	"github.com/wycliff-ochieng/cmd/migrate"
	"database/sql"
)

type Storage interface {
	CreateProduct(*migrate.Product) error
	//UpdateProduct()
	//DeleteProduct()
	GetProductByID(int) (*migrate.Product, error)
	GetProducts() ([]*migrate.Product, error)
	QueryRow(query string, args...interface{}) *sql.Row
	Exec(query string, args...interface{}) (sql.Result, error)
} 