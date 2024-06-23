package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type TimeResponse struct {
	Time string `json:"time"`
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().Format(time.RFC3339)
	response := TimeResponse{Time: currentTime}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func farewell() string {
    return "Goodbye from Andriy!"
}

func main() {
	http.HandleFunc("/time", timeHandler)

	port := 8795
	serverAddr := fmt.Sprintf(":%d", port)

	fmt.Printf("Server is running at http://localhost:%d/time\n", port)
	err := http.ListenAndServe(serverAddr, nil)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
