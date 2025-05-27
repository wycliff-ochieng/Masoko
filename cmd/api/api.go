package api

import (
	"database/sql"
	"log"
	"net/http"

	"os"

	"github.com/gorilla/mux"
	"github.com/wycliff-ochieng/handlers"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIserver(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() {

	l := log.New(os.Stdout, "ECOMMERCE APPLICATION", log.LstdFlags)

	uh := handlers.NewUser(l)

	router := mux.NewRouter()

	registerUserRouter := router.Methods("GET").Subrouter()
	registerUserRouter.HandleFunc("/", uh.RegisterUser)

	loginUserRouter := router.Methods("POST").Subrouter()
	loginUserRouter.HandleFunc("/", uh.UserLogin)
	http.ListenAndServe("/3000", nil)
}
