package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

type API struct {
	count int
}

// IndexHandler serves the main vote page
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "page1.html")
}

// P2Handler serves the main vote page
func P2Handler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "page2.html")
}

// P3Handler serves the main vote page
func P3Handler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "page3.html")
}

func (a *API) incrCount(w http.ResponseWriter, r *http.Request) {
	fmt.Println("here")
	a.count++
}

func (a *API) getCount(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(strconv.Itoa(a.count)))
}

func main() {
	a := API{
		count: 0,
	}
	r := chi.NewRouter()
	r.Get("/p1", IndexHandler)
	r.Get("/p2", P2Handler)
	r.Get("/p3", P3Handler)
	r.Post("/tracc", a.incrCount)
	r.Get("/pct", a.getCount)
	r.Method("GET", "/static/*", http.StripPrefix("/static", http.FileServer(http.Dir("static/"))))
	http.ListenAndServe(":3000", r)
}
