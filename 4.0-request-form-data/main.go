package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

var homeView *View

func NewView(layout string, files ...string) (*View, error) {
	files = append(files, "base.gohtml")
	t, err := template.ParseFiles(files...)
	if err != nil {
		return nil, err
	}
	return &View{Template: t, Layout: layout}, nil
}

type View struct {
	Template *template.Template
	Layout   string
}

func (v *View) Render(w http.ResponseWriter, data interface{}) error {
	return v.Template.ExecuteTemplate(w, v.Layout, data)
}

func homeGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := homeView.Render(w, nil)
	if err != nil {
		panic(err)
	}
}

// This only gets executed when POST action is done on / path
func homePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	// form needs to be parsed first before you can access the data
	if err := r.ParseForm(); err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "%v\n", r.PostForm["email"])
}

func main() {
	var err error
	homeView, err = NewView("base", "home.gohtml")
	if err != nil {
		panic(err)
	}
	r := mux.NewRouter()
	// when page gets loaded homeGet gets executed
	r.HandleFunc("/", homeGet).Methods("GET")
	// when form is submitted homePost gets executed
	r.HandleFunc("/", homePost).Methods("POST")
	if err := http.ListenAndServe(":3000", r); err != nil {
		panic(err)
	}
}
