package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
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

type FormData struct {
	Email string `schema:"email"`
}

func homePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := r.ParseForm(); err != nil {
		panic(err)
	}
	// create new gorilla/schema decoder
	dec := schema.NewDecoder()
	var formData FormData
	// map submitted form data into struct
	err := dec.Decode(&formData, r.PostForm)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "%v\n", formData)
}

func main() {
	var err error
	homeView, err = NewView("base", "home.gohtml")
	if err != nil {
		panic(err)
	}
	r := mux.NewRouter()
	r.HandleFunc("/", homeGet).Methods("GET")
	r.HandleFunc("/", homePost).Methods("POST")
	if err := http.ListenAndServe(":3000", r); err != nil {
		panic(err)
	}
}
