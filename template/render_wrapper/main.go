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
	Layout   string
}

// Render method executes the named templated specified in the View struct
// This way we can simplify the logic in the handleFunc function
func (v *View) Render(w http.ResponseWriter, data interface{}) error {
	return v.Template.ExecuteTemplate(w, v.Layout, data)
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	// Call Render() method instead of directly calling a function on a field
	err := homeView.Render(w, nil)
	if err != nil {
		panic(err)
	}
}

func main() {
	var err error
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
