package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/wycliff-ochieng/product-service/migrate"
	"github.com/wycliff-ochieng/product-service/db"
)

type Product struct {
	l     *log.Logger
	store db.Storage
}

func NewProduct(l *log.Logger, store db.Storage) *Product {
	return &Product{
		l:     l,
		store: store,
	}
}

type APIFunc func(w http.ResponseWriter, r *http.Request) error

type APIError struct {
	Error string
}

func MakeHttpHandlerFunc(a APIFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := a(w, r); err != nil {
			WriteJSON(w, http.StatusRequestTimeout, APIError{Error: err.Error()})
		}
	}
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Set("Centent-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

func (p *Product) GetProducts(w http.ResponseWriter, r *http.Request) {

	p.l.Println("Handle GET products requests...")

	product, err := p.store.GetProducts()
	if err != nil {
		fmt.Errorf("cant retrieve products %v", err)
	}
	WriteJSON(w, http.StatusOK, product)

}

func (p *Product) CreateProduct(w http.ResponseWriter, r *http.Request) error {
	p.l.Println("Handle POST PRoduct to db")

	createproductreq := new(migrate.CreateProductReq)

	if err := json.NewDecoder(r.Body).Decode(createproductreq); err != nil {
		return err
	}

	product := migrate.NewProduct(createproductreq.Name, createproductreq.Price, createproductreq.Description, createproductreq.Stock)

	if err := p.store.CreateProduct(product); err != nil {
		return err
	}
	return WriteJSON(w, http.StatusOK, product)
}

func (p *Product) GetProductByID(w http.ResponseWriter, r *http.Request) error {
	p.l.Println("Handling GET product by ID method.....")

	vars := mux.Vars(r)["id"]

	id, err := strconv.Atoi(vars)
	if err != nil {
		return err
	}

	product, err := p.store.GetProductByID(id)
	if err != nil {
		return err
	}
	return WriteJSON(w, http.StatusOK, product)
}
