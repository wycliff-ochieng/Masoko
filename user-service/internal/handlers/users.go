package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	service "github.com/wycliff-ochieng/user-service/internal/services"
)

type AuthHandler struct {
	l *log.Logger
	u *service.AuthService
}

type RegisterReq struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func NewAuthHandler(l *log.Logger, u *service.AuthService) *AuthHandler {
	return &AuthHandler{
		l: l,
		u: u,
	}
}

func (h *AuthHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {

	h.l.Println("REGISTERING USER NOW.....")

	var req *RegisterReq

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "unable to Decode", http.StatusInternalServerError)
		return
	}

	if req.FirstName == "" || req.LastName == "" || req.Email == "" || req.Password == "" {
		http.Error(w, "please enter characters", http.StatusExpectationFailed)
		return
	}

	//register user
	user, err := h.u.Register(req.FirstName, req.LastName, req.Email, req.Password)
	if err == service.ErrEmailExists {
		http.Error(w, "email already exists", http.StatusExpectationFailed)
		return
	}
	if err != nil {
		h.l.Printf("Error:%v", err)
		http.Error(w, "failed to register user", http.StatusFailedDependency)
		return
	}

	w.Header().Add("Content-type", "application/json")
	json.NewEncoder(w).Encode(&user)

}

func (h *AuthHandler) LoginUser(w http.ResponseWriter, r *http.Request) {}
