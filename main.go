package main

import (
	"fmt"
	"log"
	"net/http"
	"retail-pulse-service/api"
	"retail-pulse-service/utils"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize the logger
	utils.InitLogger()
	utils.Log("Starting Retail Pulse Service...")

	// Initialize the router
	r := mux.NewRouter()

	// Define the routes for the API
	r.HandleFunc("/api/submit/", api.SubmitJob).Methods("POST")
	r.HandleFunc("/api/status", api.GetJobInfo).Methods("GET")

	// Start the server
	port := ":8080"
	fmt.Printf("Server is running on port %s\n", port)
	utils.Log(fmt.Sprintf("Server started on port %s", port))

	err := http.ListenAndServe(port, r)
	if err != nil {
		utils.Log(fmt.Sprintf("Error starting server: %v", err))
		log.Fatalf("Error starting server: %v", err)
	}
}
