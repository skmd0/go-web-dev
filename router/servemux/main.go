package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintln(w, "<h1>Welcome!</h1>")
}

func main() {
	mux := &http.ServeMux{}
	mux.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", mux)
}
