package main

import (
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

var homeView *View

func NewView(layouts ...string) (*View, error) {
	layouts = append(layouts, "footer.gohtml")
	t, err := template.ParseFiles(layouts...)
	if err != nil {
		return nil, err
	}
	return &View{Template: t}, nil
}

type View struct {
	Template *template.Template
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := homeView.Template.Execute(w, nil)
	if err != nil {
		panic(err)
	}
}

func main() {
	var err error
	homeView, err = NewView("home.gohtml")
	if err != nil {
		panic(err)
	}
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	http.ListenAndServe(":3000", r)
}
