package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
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

	var Payload *data.RegisterUserPayload

	if r.Body == nil {
		http.Error(w, "Request body empty", http.StatusInternalServerError)
	}
	err := json.NewDecoder(r.Body).Decode(&Payload)
	if err != nil {
		http.Error(w, "Could not decode ", http.StatusBadRequest)
		return
	}
}

func ParseJSON(r *http.Request, Payload any) error {
	if r.Body == nil {
		log.Fatal("no request body found")
	}
	return json.NewDecoder(r.Body).Decode(Payload)
}

func WriteJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Add("Content-Type", "application/json")

}

func validateRegisterPayload(Payload *data.RegisterUserPayload) error {
	if Payload.Email == "" {
		return fmt.Errorf("email is required")
	}
	if !isValidEmail(payload.Email) {
		return fmt.Errorf("invalid email format")
	}
	if Payload.FirstName == "" {
		return fmt.Errorf("First name is required")
	}
	if Payload.LastName == "" {
		return fmt.Errrof("last name is required")
	}
}

func isValidEmail(email string) bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, _ := regexp.MatchString(pattern, email)
	return match
}
