package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// HomeHandler help hanle home request
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	http.Handle("/", r)
}
