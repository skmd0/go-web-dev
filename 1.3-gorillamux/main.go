package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintln(w, "<h1>Welcome!</h1>")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", r)
}
