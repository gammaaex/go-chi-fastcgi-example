package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/fcgi"

	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("This Go app work on FastCGI"))
	})
	r.Route("/route1", func(r chi.Router) {
		r.Get("/", handleRoute1)
	})
	r.Route("/route2", func(r chi.Router) {
		r.Get("/", handleRoute2)
	})

	err := fcgi.Serve(nil, r)
	if err != nil {
		log.Fatal(err)
	}
}

func handleRoute1(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	fmt.Fprintln(w, "-------- handleRoute1 on FastCGI ---------")
	fmt.Fprintln(w, "Method:", r.Method)
	fmt.Fprintln(w, "URL:", r.URL.String())
	fmt.Fprintln(w, "URL.Path:", r.URL.Path)
}

func handleRoute2(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	fmt.Fprintln(w, "-------- handleRoute2 on FastCGI ---------")
	fmt.Fprintln(w, "Method:", r.Method)
	fmt.Fprintln(w, "URL:", r.URL.String())
	fmt.Fprintln(w, "URL.Path:", r.URL.Path)
}
