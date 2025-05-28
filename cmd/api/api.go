package api

import (
	"log"
	"net/http"

	"os"

	"github.com/gorilla/mux"
	"github.com/wycliff-ochieng/db"
	"github.com/wycliff-ochieng/handlers"
)

type APIServer struct {
	addr string
	db   db.Storage
}

func NewAPIserver(addr string, db db.Storage) *APIServer {
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
	http.ListenAndServe(s.addr, router)
}
