package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func getView(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintln(w, "<h1>GET!</h1>")
}

func postView(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintln(w, "<h1>POST!</h1>")
}

func main() {
	r := mux.NewRouter()
	// Specify GET HTTP action
	r.HandleFunc("/", getView).Methods("GET")
	// Specify POST HTTP action
	r.HandleFunc("/", postView).Methods("POST")
	http.ListenAndServe(":3000", r)
}
