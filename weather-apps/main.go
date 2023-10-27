package main

import (
	"log"
	"net/http"
	"weather-apps/app"
)

func main() {
	port := ":9999"

	http.HandleFunc("/weather", app.GetCurrentWeather)

	log.Printf("Server running at http://localhost%v/weather", port)
	http.ListenAndServe(port, nil)
}