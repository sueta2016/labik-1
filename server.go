package main

import (
	"labik-1/service"
	"net/http"
)

func main() {
	timeService := service.NewTimeService()

	http.HandleFunc("/time", timeService.GetTimeHandler)

	http.ListenAndServe(":8795", nil)
}
