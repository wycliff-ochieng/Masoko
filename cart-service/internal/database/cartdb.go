package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type DBInterface interface {
	QueryRow(query string, args ...interface{}) *sql.Row
	Exec(query string, args ...interface{}) (sql.Result, error)
	Close() error
}

type Postgres struct {
	db *sql.DB
}

func NewPostgres() (*Postgres, error) {
	connStr := "user=postgres dbname=postgres password=gobank sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("could not open db %v", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping %v", err)
	}
	return &Postgres{db: db}, nil
}

func (p *Postgres) Init() error {
	return p.CreateCartTable()
}

func (p *Postgres) Init2() error {
	return p.CreateCartItemTable()
}

func (p *Postgres) CreateCartTable() error {
	query := `CREATE TABLE carts(
	ID SERIAL PRIMARY KEY,
	UserID INT,
	SessionID TEXT,
	Status VARCHAR(20),
	Items TEXT,
	CreatedAT TIMESTAMP,
	UpdatedAT TIMESTAMP,
	CONSTRAINT fk_userID FOREIGN KEY (UserID) REFERENCES users)`

	_, err := p.db.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to create table: %v", err)
	}
	return nil
}

func (p *Postgres) CreateCartItemTable() error {
	query := `CREATE TABLE cart_items(
	id SERIAL PRIMARY KEY,
	cartID INTEGER, 
	productID INTEGER,
	quantity INTEGER,
	createdat TIMESTAMP,
	CONSTRAINT fk_cartID FOREIGN KEY(cartID) REFERENCES carts,
	CONSTRAINT fk_productID FOREIGN KEY(productID) REFERENCES products
	)`

	_, err := p.db.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to crate table cart item: %v", err)
	}
	return nil
}

func (p *Postgres) Close() error {
	if p.db != nil {
		return p.db.Close()
	}
	return nil
}

func (p *Postgres) Exec(query string, args ...interface{}) (sql.Result, error) {
	return p.db.Exec(query, args...)
}

func (p *Postgres) QueryRow(query string, args ...interface{}) *sql.Row {
	return p.db.QueryRow(query, args...)
}
