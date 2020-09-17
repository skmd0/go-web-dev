package main

import (
	"fmt"
	"net/http"
)

func NewHome() (*Home, error) {
	homeView, err := NewView("base", "home.gohtml")
	if err != nil {
		return nil, err
	}
	return &Home{HomeView: homeView}, nil
}

type Home struct {
	HomeView *View
}

func (h *Home) ShowHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := h.HomeView.Render(w, nil)
	if err != nil {
		fmt.Println("ERR: failed to render HomeView")
	}
}
