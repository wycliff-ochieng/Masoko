package db

import (
	"database/sql"
	"fmt"

	"github.com/wycliff-ochieng/cmd/migrate"

	_ "github.com/lib/pq"
)

type Postgrestore struct {
	db *sql.DB
}

func NewPostgrestore() (*Postgrestore, error) {
	connStr := "user=postgres dbname=postgres password=gobank sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &Postgrestore{db: db}, nil

}

func (p *Postgrestore) Init() error {
	return p.CreateProductTable()
}

func (p *Postgrestore) CreateUserTable() error {
	query := "CREATE TABLE IF NOT EXISTS users()"
	_, err := p.db.Exec(query)
	return err
}

func (p *Postgrestore) CreateProductTable() error {
	query := `CREATE TABLE IF NOT EXISTS products(
	id serial primary key,
	name varchar(15) not null,
	price decimal(10,2) not null,
	description text,'
	stock int default 0,
	createdat timestamp,
	updatedat timestamp
	);`

	_, err := p.db.Exec(query)
	return err
}

func (s *Postgrestore) CreateProduct(product *migrate.Product) error {
	query := `INSERT INTO products (name, description, price ,stock , createdat, updatedat) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	err := s.db.QueryRow(query, product.Name, product.Description, product.Price, product.Stock, product.CreatedAt, product.UpdatedAt).Scan(&product.ID)
	return err
}

//func (P *Postgres) CrreateTransactionTable

func (p *Postgrestore) GetProducts() ([]*migrate.Product, error) {
	rows, err := p.db.Query(`SELECT id,name,price,description,stock,createdat,updatedat FROM products`)
	if err != nil {
		return nil, err
	}

	products := []*migrate.Product{}

	for rows.Next() {
		product, err := p.ScanIntoTable(rows)
		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}
	return products, nil

}

func (p *Postgrestore) ScanIntoTable(rows *sql.Rows) (*migrate.Product, error) {
	product := new(migrate.Product)

	err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Description, &product.Stock, &product.CreatedAt, &product.UpdatedAt)

	if err != nil {
		return nil, err
	}
	return product, nil
}

func (p *Postgrestore) GetProductByID(id int) (*migrate.Product, error) {
	rows, err := p.db.Query("SELECT * FROM products WHERE id = $1", id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return p.ScanIntoTable(rows)
	}
	return nil, fmt.Errorf("could not find ID %v", err)

}

func (p *Postgrestore) QueryRow(query string, args ...interface{}) *sql.Row {
	return nil
}

func (p *Postgrestore) Exec(query string, args ...interface{}) (sql.Result, error) {
	return nil, nil
}
