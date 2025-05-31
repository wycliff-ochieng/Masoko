package db

import (
	"database/sql"

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
	description text,
	stock int default 0,
	createdat timestamp,
	updatedat timestamp
	);`

	_, err := p.db.Exec(query)
	return err
}

func (s *Postgrestore) CreateProduct(product *migrate.Product) error {
	query := `INSERT INTO products (name, description, price) VALUES ($1, $2, $3) RETURNING id`
	err := s.db.QueryRow(query, product.Name, product.Description, product.Price).Scan(&product.ID)
	return err
}

//func (P *Postgres) CrreateTransactionTable
