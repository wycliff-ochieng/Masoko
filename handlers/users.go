package handlers

import (
	"log"
	"net/http"
)

type User struct {
	l *log.Logger
}

func NewUser(l *log.Logger) *User {
	return &User{l}
}

func (u *User) UserLogin(w http.ResponseWriter, r *http.Request) {
	return
}

func (u *User) RegisterUser(w http.ResponseWriter, r *http.Request) {
	return
}
