/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"controllers"
	"log"
	"net/http"
)

func main() {

	router := http.NewServeMux()
	router.Handle("/api/health/", http.StripPrefix("/api/health", controllers.HealthRouter()))
	router.Handle("/api/flight/", http.StripPrefix("/api/flight", controllers.Router()))

	server := http.Server{
		Addr: ":3001",
	}

	log.Println("Server start on : http://localhost:3001")

	err := server.ListenAndServe()

	log.Fatalln(err)
}
