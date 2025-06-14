package database

import (
	"database/sql"
	"fmt"
)

type Postgres struct {
	db *sql.DB
}

func NewPostgrestore() (*Postgres, error) {
	connStr := "user=postgres dbname=cart_service password=gobank sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil,fmt.Errorf("could not open db %v", err)
	}

	if err = db.Ping(); err != nil {
		return nil,fmt.Errorf("Failed to ping %v", err)
	}
	return &Postgres{db: db}, nil
}


func(p *Postgres) Init() error{
	return p.CreateCartTable()
}


func(p *Postgres) CreateCartTable() error{
	query := `CREATE TABLE cart(
	ID SERIAL PRIMARY KEY,
	UserID INT
	SessionID TEXT,
	Status VARCHAR(20),
	Items TEXT,
	CreatedAT TIMESTAMP,
	UpdatedAT TIMESTAMP
	CONSTRAINT fk_userID FOREIGN KEY (UserID)
	REFERENCES users)`

	_,err := p.db.Exec(query)
	if err != nil{
		return fmt.Errorf("failed to create table: %v",err)
	}
	return nil
}
