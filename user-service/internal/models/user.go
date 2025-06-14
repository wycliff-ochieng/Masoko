package models

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        uuid.UUID
	FirstName string
	LastName  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserResponse struct {
	ID        uuid.UUID
	FirstName string
	LastName  string
	Email     string
	CreatedAt time.Time
}

func NewUser(firstname, lastname, email, password string) (*User, error) {
	harshedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	return &User{
		ID:        uuid.New(),
		FirstName: firstname,
		LastName:  lastname,
		Email:     email,
		Password:  string(harshedPassword),
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}
