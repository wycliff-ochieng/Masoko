package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID
	Email     string
	Password  string
	FirstName string
	LastName  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUser(email, password, firstname, lastname string) (*User, error) {
	return &User{}
}
