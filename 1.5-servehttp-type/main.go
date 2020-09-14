package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type View struct{}

func (v *View) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Hello world!</h1>")
}

func main() {
	r := mux.NewRouter()
	// Using Handle() instead of HandleFun() because View struct implements http.Handler interface
	r.Handle("/", &View{})
	http.ListenAndServe(":3000", r)
}
