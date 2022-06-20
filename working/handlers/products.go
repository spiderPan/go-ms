package handlers

import (
	"go-ms/data"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		p.getProducts(w, r)

	case http.MethodPost:
		p.addProduct(w, r)

	case http.MethodPut:
		// expect the ID in the URI
		reg := regexp.MustCompile(`/([0-9]+)`)
		g := reg.FindAllStringSubmatch(r.URL.Path, -1)
		if len(g) != 1 || len(g[0]) != 2 {
			http.Error(w, "Invalid URI", http.StatusBadRequest)
			return
		}

		idString := g[0][1]
		id, err := strconv.Atoi(idString)
		if err != nil {
			http.Error(w, "Invalid URI", http.StatusBadRequest)
			return
		}

		p.updateProducts(id, w, r)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

}

func (p *Products) getProducts(w http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()

	err := lp.ToJSON(w)

	if err != nil {
		http.Error(w, "Unable to marshal JSON", http.StatusInternalServerError)
	}
}

func (p *Products) addProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("About to add product")

	prod := &data.Product{}

	err := prod.FromJSON(r.Body)

	if err != nil {
		http.Error(w, "Unable to add JSON", http.StatusBadRequest)
	}

	p.l.Printf("Data %#v", prod)
	data.AddProduct(prod)
}

func (p *Products) updateProducts(id int, w http.ResponseWriter, r *http.Request) {
	prod := &data.Product{}

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
