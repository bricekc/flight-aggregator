package main

import (
	"aggregator/controllers"
	"log"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	router.Handle("/api/health/", http.StripPrefix("/api/health", controllers.HealthRouter()))
	router.Handle("/api/flight/", http.StripPrefix("/api/flight", controllers.FlightRouter()))

	server := http.Server{
		Addr:    ":3001",
		Handler: router,
	}
	log.Println("Server start on : http://localhost:3001")
	log.Fatal(server.ListenAndServe())
}
