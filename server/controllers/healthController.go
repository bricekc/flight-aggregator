package controllers

import "net/http"

func HealthRouter() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("GET /", GetHealth)
	return router
}

func GetHealth(res http.ResponseWriter, _ *http.Request) {
	fmt.Println("Request here")
	fmt.Fprintf(w, "Hell, %q", html.EscapeString((r.URL.Path)))
}