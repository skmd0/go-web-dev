package main

import (
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
	// Added Layout field which is used for selecting which template should be rendered first
	Layout string
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	// ExecuteTemplate renders the named template first, which is provided as second parameter
	err := homeView.Template.ExecuteTemplate(w, homeView.Layout, nil)
	if err != nil {
		panic(err)
	}
}

func main() {
	var err error
	// Added "base" parameter which is same as "base" definition in base.gohtml
	homeView, err = NewView("base", "home.gohtml")
	if err != nil {
		panic(err)
	}
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	if err := http.ListenAndServe(":3000", r); err != nil {
		panic(err)
	}
}
