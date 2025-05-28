package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/wycliff-ochieng/cmd/migrate"
)

type Product struct {
	l *log.Logger
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

func (p *Product) CreateProduct(w http.ResponseWriter, r *http.Request) {

	p.l.Println("Handle POST products")
}
