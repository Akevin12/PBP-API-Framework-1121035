package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-martini/martini"
	"github.com/gorilla/mux"
)

func main() {

	// Martini
	// Menggunakan Martini untuk membuat instance.
	m := martini.Classic()

	// Routing (Martini)
	// Menambahkan route di Martini.
	m.Get("/hello/:name", func(params martini.Params) string {
		return "Hello " + params["name"]
	})

	// Services (Martini)
	// Menetapkan layanan di Martini.
	db := &MyDatabase{}
	m.Map(db)

	// Serving Static Files (Martini)
	m.Use(martini.Static("assets"))

	// Middleware Handlers (Martini)
	// Menambahkan middleware di Martini.
	m.Use(func() {
		// logic
	})

	// Next() (Martini)
	// Menggunakan Context.Next() di Martini untuk melanjutkan ke middleware atau handler berikutnya.
	m.Use(func(c martini.Context, log *log.Logger) {
		log.Println("before a request")
		c.Next()
		log.Println("after a request")
	})

	// Run Martini
	// Menjalankan Martini server.
	m.Run()

	// Gorilla Mux
	// Membuat instance router menggunakan Gorilla Mux.
	r := mux.NewRouter()

	// Routing (Gorilla Mux)
	// Menambahkan route di Gorilla Mux.
	r.HandleFunc("/hello/{name}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		name := vars["name"]
		fmt.Fprintf(w, "Hello %s", name)
	}).Methods("GET")

	// Services (no direct equivalent in Gorilla Mux)
	// Tidak ada peta layanan global atau permintaan di Gorilla Mux.

	// Serving Static Files (Gorilla Mux)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("assets/"))))

	// Middleware Handlers (Gorilla Mux)
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// logic
			next.ServeHTTP(w, r)
		})
	})

	// Next() (no direct equivalent in Gorilla Mux)
	// Tidak ada fungsi Next() yang setara di Gorilla Mux.

	// Run Gorilla Mux
	// Menjalankan server menggunakan Gorilla Mux.
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)

}

type MyDatabase struct {
	// Database fields
}
