package controllers

import (
	"fmt"
	"html"
	"net/http"
)

func FlightRouter() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("GET /", GetFlights)
	return router
}

func GetFlights(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request here")
	fmt.Fprintf(w, "Hell, %q", html.EscapeString(r.URL.Path))
}