package main

import (
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

var homeView *template.Template

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := homeView.Execute(w, nil)
	if err != nil {
		panic(err)
	}
}

func main() {
	var err error
	homeView, err = template.ParseFiles("home.gohtml")
	if err != nil {
		panic(err)
	}
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	http.ListenAndServe(":3000", r)
}
