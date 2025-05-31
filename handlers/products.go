package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/wycliff-ochieng/cmd/migrate"
	"github.com/wycliff-ochieng/db"
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

	pl := migrate.GetProducts()

	d, err := json.Marshal(pl)
	if err != nil {
		http.Error(w, "unable to decode to json", http.StatusInternalServerError)
	}
	w.Write(d)
}

func (p *Product) CreateProduct(w http.ResponseWriter, r *http.Request) error {
	p.l.Println("Handle GetPRoduct from db")

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

//func (p *Product) CreateProduct(w http.ResponseWriter, r *http.Request) {
//
//	p.l.Println("Handle POST products")
//}
