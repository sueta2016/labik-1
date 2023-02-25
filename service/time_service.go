package service

import (
    "encoding/json"
    "net/http"
    "time"
)

type TimeService struct{}

func NewTimeService() *TimeService {
    return &TimeService{}
}

func (ts *TimeService) GetTimeHandler(w http.ResponseWriter, r *http.Request) {
    currentTime := time.Now()

    // Create a map to hold the time data
    timeData := make(map[string]interface{})
    timeData["time"] = currentTime.Format(time.RFC3339)

    // Convert the map to JSON
    jsonData, err := json.Marshal(timeData)
    if err != nil {
        http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
        return
    }

    // Set the response headers and send the JSON data
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(jsonData)
}