package handlers

import (
	"go-ms/data"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) GetProducts(w http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()

	err := lp.ToJSON(w)

	if err != nil {
		http.Error(w, "Unable to marshal JSON", http.StatusInternalServerError)
	}
}

func (p *Products) AddProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("About to add product")

	prod := &data.Product{}

	err := prod.FromJSON(r.Body)

	if err != nil {
		http.Error(w, "Unable to add JSON", http.StatusBadRequest)
	}

	p.l.Printf("Data %#v", prod)
	data.AddProduct(prod)
}

func (p *Products) UpdateProducts(w http.ResponseWriter, r *http.Request) {
	prod := &data.Product{}
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	err := prod.FromJSON(r.Body)

	if err != nil {
		http.Error(w, "Unable to add JSON", http.StatusBadRequest)
	}

	p.l.Printf("Data %#v", prod)
	err = data.UpdateProduct(id, prod)

	if err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
	}
}
