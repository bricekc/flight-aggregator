/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {

	router := http.NewServeMux()
	router.HandleFunc("GET /api/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Request here")
		fmt.Fprintf(w, "Hello, %q", html.EscapeString((r.URL.Path)))
	})

	server := http.Server{
		Addr: ":3001",
	}

	log.Println("Server start on : http://localhost:3001")

	err := server.ListenAndServe()

	log.Fatalln(err)
}
