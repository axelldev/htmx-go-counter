package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Counter struct {
	Count int
}

func main() {
	counter := Counter{Count: 0}
	server := http.Server{
		Addr: ":4321",
	}

	http.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("templates/index.html")
		if err != nil {
			log.Fatal(err)
		}
		tmpl.Execute(w, nil)
	})

	http.HandleFunc("GET /counter", func(w http.ResponseWriter, r *http.Request) {
		if err := renderCounterTemplate(w, counter); err != nil {
			log.Fatal(err)
		}
	})

	http.HandleFunc("GET /increment", func(w http.ResponseWriter, r *http.Request) {
		counter.Count++
		renderCounterTemplate(w, counter)
	})

	http.HandleFunc("GET /decrement", func(w http.ResponseWriter, r *http.Request) {
		counter.Count--
		if err := renderCounterTemplate(w, counter); err != nil {
			log.Fatal(err)
		}
	})

	fmt.Println("Server is running on port http://localhost:4321")
	log.Fatal(server.ListenAndServe())
}

func renderCounterTemplate(w http.ResponseWriter, counter Counter) error {
	tmpl, err := template.ParseFiles("templates/counter.html")
	if err != nil {
		return err
	}
	return tmpl.Execute(w, counter)
}
