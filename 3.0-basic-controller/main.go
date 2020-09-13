package main

import (
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

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

func main() {
	var err error
	controller, err := NewHome()
	if err != nil {
		panic(err)
	}
	r := mux.NewRouter()
	r.HandleFunc("/", controller.ShowHome)
	if err := http.ListenAndServe(":3000", r); err != nil {
		panic(err)
	}
}
