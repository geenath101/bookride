package main

import (
	"encoding/json"
	"log"
	"net/http"
	"ride-sharing/services/api-gateway/grpc_clients"
	"ride-sharing/shared/contracts"
)

func handleTripPreview(w http.ResponseWriter, r *http.Request) {
	// todo : call the trip service
	var requestBody previewTripRequest
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "failed to parse JSON data", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if requestBody.UserID == "" {
		http.Error(w, "user id must be present", http.StatusBadRequest)
		return
	}

	tripSerice, err := grpc_clients.NewTripServiceClient("trip-service:9093")
	if err != nil {
		log.Printf("unable to establish a connection : %v", err)
		return
	}
	defer tripSerice.Close()

	tripResponse, err := tripSerice.Client.PreviewTrip(r.Context(), requestBody.ToProto())

	if err != nil {
		log.Printf("Failed to preview a trip: %v", err)
		http.Error(w, "got a error response from the grpc service", http.StatusInternalServerError)
		return
	}

	response := contracts.APIResponse{Data: tripResponse}
	writeJSON(w, http.StatusAccepted, response)
}

func writeJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Conent-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}
