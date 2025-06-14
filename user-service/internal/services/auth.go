package service

import (
	"errors"
	"fmt"

	"github.com/wycliff-ochieng/user-service/internal/database"
	"github.com/wycliff-ochieng/user-service/internal/models"
)

type AuthService struct {
	db database.DBInterface
}

var (
	ErrEmailExists = errors.New("email already exists")
)

func NewAuthService(db database.DBInterface) *AuthService {
	return &AuthService{db: db}
}

func (u *AuthService) Register(firstname string, lastname string, email string, password string) (*models.UserResponse, error) {
	var exists bool

	err := u.db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)", email).Scan(&exists)
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
	query := `INSERT INTO customers (id,firstname,lastname,email,password,createdat,updatedat) VALUES($1,$2,$3,$4,$5,$6,$7)`

	_, err = u.db.Exec(query, user.ID, user.FirstName, user.LastName, user.Email, user.Password, user.CreatedAt, user.UpdatedAt)
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
