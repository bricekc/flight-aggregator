package controllers

import (
	"aggregator/repositories"
	"fmt"
	"html"
	"net/http"
)

var travellings1 = &repositories.Travellings1{}

func GetFlights(w http.ResponseWriter, r *http.Request) {
	travellings1.FlightS1()
	fmt.Println("Request here")
	fmt.Fprintf(w, "Hell, %q", html.EscapeString(r.URL.Path))
}
