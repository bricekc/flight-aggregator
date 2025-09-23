package controllers

import (
	"fmt"
	"html"
	"net/http"
)

func GetHealth(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request here")
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}