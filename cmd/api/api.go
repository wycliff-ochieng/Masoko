package api

import (
	"fmt"
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

	store, err := db.NewPostgrestore()
	if err != nil {
		//log.Fatal("couldn tconnect to db")
		fmt.Println("Error connecting to database")
	}

	ph := handlers.NewProduct(l, store)

	router := mux.NewRouter()

	getRouter := router.Methods("GET").Subrouter()
	getRouter.HandleFunc("/products", ph.GetProducts)

	postRouter := router.Methods("POST").Subrouter()
	postRouter.HandleFunc("/products", handlers.MakeHttpHandlerFunc(ph.CreateProduct))

	getSpecificRouter := router.Methods("GET").Subrouter()
	getSpecificRouter.HandleFunc("/products/{id}", handlers.MakeHttpHandlerFunc(ph.GetProductByID))

	//registerUserRouter := router.Methods("GET").Subrouter()
	//egisterUserRouter.HandleFunc("/", uh.RegisterUser)

	//loginUserRouter := router.Methods("POST").Subrouter()
	//loginUserRouter.HandleFunc("/", uh.UserLogin)
	http.ListenAndServe(s.addr, router)
}
