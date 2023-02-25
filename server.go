package main

import (
    "net/http"
)

func main() {
    
    http.HandleFunc("/time", timeService.GetTimeHandler)

    http.ListenAndServe(":8080", nil)
}