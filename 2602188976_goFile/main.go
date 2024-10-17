package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
)

type Response struct {
    Message string   `json:"message"`
    Data    []string `json:"data,omitempty"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
    response := Response{
        Message: "Compare yourself to who you were yesterday, not to who someone else is today ~JP",
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

func dataHandler(w http.ResponseWriter, r *http.Request) {
    response := Response{
        Message: "Data received! Neat quotes :D",
        Data: []string{
            "Quote 1: The best way to predict the future is to invent it.",
            "Quote 2: Life is what happens when you're busy making other plans.",
        },
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

func main() {
    http.HandleFunc("/hello", helloHandler)
    http.HandleFunc("/data", dataHandler)

    fmt.Println("Server running on port 8080...")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
