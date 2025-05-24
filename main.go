package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/calculate", func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers for React client
		w.Header().Set("Access-Control-Allow-Origin", "https://52.41.36.82")
		w.Header().Set("Access-Control-Allow-Origin", "https://54.191.253.12")
		w.Header().Set("Access-Control-Allow-Origin", "https://44.226.122.3")
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Content-Type", "application/json")

		if r.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Allow-Methods", "POST")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			w.WriteHeader(http.StatusOK)
			return
		}

		if r.Method != "POST" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		var req struct {
			Expression string `json:"expression"`
		}
		json.NewDecoder(r.Body).Decode(&req)
		result := calculate(req.Expression)

		json.NewEncoder(w).Encode(struct {
			Result float64 `json:"result"`
		}{result})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "5001"
	}
	fmt.Printf("Listening on port %s\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println(err)
	}
}
