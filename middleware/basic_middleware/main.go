package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var indexView *template.Template

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := indexView.ExecuteTemplate(w, "bulma", nil)
	if err != nil {
		panic(err)
	}
}

func Apply(next http.Handler) http.Handler {
	return ApplyFn(next.ServeHTTP)
}

func ApplyFn(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			fmt.Println("Root mw")
		} else {
			fmt.Println("Non-root mw")
		}
		next(w, r)
	})
}

type Test struct {
	testView *template.Template
}

func (t Test) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := t.testView.Execute(w, nil); err != nil {
		panic(err)
	}
}

func main() {
	var err error
	indexView, err = template.ParseFiles("base.gohtml", "index.gohtml")
	if err != nil {
		panic(err)
	}
	testView, err := template.ParseFiles("test.gohtml")
	if err != nil {
		panic(err)
	}
	test := Test{testView}
	http.HandleFunc("/", ApplyFn(index))
	http.Handle("/test", Apply(test))
	http.ListenAndServe(":3000", nil)
}
