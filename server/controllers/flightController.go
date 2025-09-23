package controllers

import (
	"fmt"
	"html"
	"net/http"
)

func GetFlights(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request here")
	fmt.Fprintf(w, "Hell, %q", html.EscapeString(r.URL.Path))
}