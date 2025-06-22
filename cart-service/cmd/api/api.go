package api

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/wycliff-ochieng/cart-service/internal/database"
	"github.com/wycliff-ochieng/cart-service/internal/handler"
	service "github.com/wycliff-ochieng/cart-service/internal/services"
)

type APIServer struct {
	addr string
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{addr: addr}
}

func (s *APIServer) Run() {

	l := log.New(os.Stdout, "customer user router", log.LstdFlags)

	db, err := database.NewPostgres()
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}

	if err := db.Init(); err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}

	if err := db.Init2(); err != nil {
		log.Fatalf("failed to initialize cart item table: %v", err)
	}

	u := service.NewCartService(db)

	uh := handler.NewCartHandler(l, u)

	router := mux.NewRouter()

	registerRouter := router.Methods("POST").Subrouter()
	registerRouter.HandleFunc("/register", uh.GetCart)

	if err := http.ListenAndServe(s.addr, router); err != nil {
		log.Fatalf("Server startup failed: %v", err)
	}
}
