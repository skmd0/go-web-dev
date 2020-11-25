package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

const (
	userID = "user_id"
)

var router *mux.Router

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	url, err := router.Get(userID).URL("id", "1")
	if err != nil {
		http.Error(w, "Something went wrong!", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, url.Path, http.StatusFound)
}

func user(w http.ResponseWriter, r *http.Request) {
	_ = r
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintln(w, "Successful URL redirect with named URL path generation.")
}

func main() {
	router = mux.NewRouter()
	router.HandleFunc("/", index)
	router.HandleFunc("/user/{id:[0-9]+}", user).Name(userID)
	http.ListenAndServe(":3000", router)
}
