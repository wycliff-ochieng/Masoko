package service

import (
	"fmt"

	"github.com/wycliff-ochieng/db"
	"github.com/wycliff-ochieng/models"
)

type AuthService struct {
	db db.Storage
}

func NewAuthService(db db.Storage) *AuthService {
	return &AuthService{db: db}
}

func (u *AuthService) Register(firstname string, lastname string, email string, password string) (*models.UserResponse, error) {
	var exists bool

	err := u.db.QueryRow("SELECT id,firstname,lastname,email,createdat,updatedat FROM users WHERE email = $1", email).Scan(&exists)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, fmt.Errorf("email already exists")
	}
	//create user if check is passed

	user, err := models.NewUser(firstname, lastname, email, password)
	if err != nil {
		return nil, err
	}

	//insert into database
	query := `INSERT INTO users (id,firstname,lastname,email,createdat,updatedat) VALUES($1,$2,$3,$4,$5,$6)`

	_, err = u.db.Exec(query, user.ID, user.FirstName, user.LastName, user.Email, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("could not insert due to %v", err)
	}

	return &models.UserResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}, nil
}
