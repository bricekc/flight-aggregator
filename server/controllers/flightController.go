package controllers

import (
	"aggregator/repositories"
	"aggregator/services"
	"encoding/json"
	"net/http"
)

func GetFlights(res http.ResponseWriter, req *http.Request) {
	// Initiate the repos
	var travellings1 = &repositories.Travelling1{}
	var travellings2 = &repositories.Travelling2{}

	// Initiate service and get datas
	flightService := services.NewFlightService(travellings1, travellings2)
	allTravellings := flightService.GetAllTravel()

	// Transform the data in json
	jsonStrAllTravellings, err := json.Marshal(allTravellings)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	res.Header().Set("Content-Type", "application/json")

	res.Write(jsonStrAllTravellings)
}
