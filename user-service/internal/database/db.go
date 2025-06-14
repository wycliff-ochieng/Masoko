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
		return nil, fmt.Errorf("failed to connect to db %v", err)
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &Postgres{db: db}, nil
}

func (p *Postgres) Init() error {
	return p.CreateUserTable()
}

func (p *Postgres) CreateUserTable() error {
	query := `CREATE TABLE IF NOT EXISTS customers(
	id uuid primary key,
	firstname varchar(50)not null,
	lastname varchar(50) not null,
	email varchar(100)not null,
	password text not null,
	createdat timestamp not null default current_timestamp,
	updatedat timestamp not null default current_timestamp)`

	_, err := p.db.Exec(query)
	return err
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
