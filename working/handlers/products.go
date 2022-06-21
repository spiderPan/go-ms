package handlers

import (
	"context"
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

	prod := r.Context().Value(KeyProducts{}).(data.Product)
	data.AddProduct(&prod)
}

func (p *Products) UpdateProducts(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	prod := r.Context().Value(KeyProducts{}).(data.Product)
	err := data.UpdateProduct(id, &prod)

	if err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
	}
}

type KeyProducts struct{}

func (p *Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		prod := data.Product{}

		err := prod.FromJSON(r.Body)

		if err != nil {
			p.l.Println("[ERROR] deserializaing product", err)
			http.Error(w, "Unable to read product", http.StatusBadRequest)
		}

		// add the product to the context
		ctx := context.WithValue(r.Context(), KeyProducts{}, prod)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
