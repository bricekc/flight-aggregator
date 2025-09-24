package controllers

import (
	"fmt"
	"net/http"
)

func GetHealth(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request here")
	w.WriteHeader(http.StatusCreated)
}
