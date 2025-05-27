package db

import "database/sql"


type Postgrestore struct{
	db *sql.DB
}

func NewPostgrestore() (*Postgrestore, error){
	connStr := 
	
}