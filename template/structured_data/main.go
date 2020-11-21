package main

import (
	"html/template"
	"net/http"
)

const (
	AlertLvlError = "danger"
)

type Alert struct {
	Level   string
	Message string
}

type Data struct {
	Alert *Alert
	Yield interface{}
}

var indexView *template.Template

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	a := &Alert{
		Level:   AlertLvlError,
		Message: "Something went wrong!",
	}
	d := Data{
		Alert: a,
		Yield: "Hello world!",
	}
	// d must always be type Data because the template expects .Alert which is part of Data struct
	// you can do a type switch to convert the passed data into Data struct with passed data as yield field
	err := indexView.ExecuteTemplate(w, "bulma", d)
	if err != nil {
		panic(err)
	}
}

func main() {
	var err error
	indexView, err = template.ParseFiles("base.gohtml", "structured_data.gohtml", "index.gohtml")
	if err != nil {
		panic(err)
	}
	http.HandleFunc("/", index)
	http.ListenAndServe(":3000", nil)
}
