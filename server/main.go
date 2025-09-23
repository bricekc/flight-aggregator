package main

import (
	"aggregator/controllers"
	"log"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/api/health/", controllers.GetHealth)
	router.HandleFunc("/api/flight/", controllers.GetFlights)

	server := http.Server{
		Addr:    ":3001",
		Handler: router,
	}
	log.Println("Server start on : http://localhost:3001")
	log.Fatal(server.ListenAndServe())
}
