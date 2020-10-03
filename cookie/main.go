package main

import (
	"fmt"
	"net/http"
)

func createCookie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	cookie := http.Cookie{
		Name:  "email",
		Value: "domenko@skamlici.com",
	}
	http.SetCookie(w, &cookie)
	fmt.Fprintln(w, "Cookie created! <a href='/read'>See the cookie</a>")
}

func readCookie(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("email")
	if err != nil {
		fmt.Fprintln(w, "Unable to read cookie.")
		return
	}
	fmt.Fprintln(w, cookie)
}

func main() {
	http.HandleFunc("/", createCookie)
	http.HandleFunc("/read", readCookie)
	http.ListenAndServe(":3000", nil)
}
