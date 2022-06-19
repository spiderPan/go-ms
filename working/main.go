package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Hello World!")

		d, err := ioutil.ReadAll(r.Body)

		if err != nil {
			http.Error(w, "Ooops", http.StatusBadRequest)
			return
		}
		// log.Printf("Data %s Error %s", d, err)

		fmt.Fprintf(w, "Hello %s with error %s", d, err)
	})

	http.HandleFunc("/goodbye", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Goodbye World!")
	})
	http.ListenAndServe(":9090", nil)
}
