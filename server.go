package main

import (
    "net/http"
    "labik-1/service"
)

func main() {
    timeService := service.NewTimeService()

    http.HandleFunc("/time", timeService.GetTimeHandler)

    http.ListenAndServe(":8080", nil)
}