package controllers

import (
	"aggregator/repositories"
	"fmt"
	"net/http"
)

func GetFlights(w http.ResponseWriter, r *http.Request) {
	var travellings1 = &repositories.Travelling1{}
	flights, err := travellings1.GetTravel()
	fmt.Println("Request here")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "%+v", flights)
}
