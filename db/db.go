package db

import (
	"database/sql"

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
	return p.CreateUserTable()
}

func (p *Postgrestore) CreateUserTable() error {
	query := "CREATE TABLE IF NOT EXISTS users()"
	_, err := p.db.Exec(query)
	return err
}

//func (P *Postgres) CrreateTransactionTable
