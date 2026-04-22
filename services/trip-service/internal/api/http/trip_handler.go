package http

import (
	"encoding/json"
	"net/http"
	"ride-sharing/services/trip-service/internal/core"
	"ride-sharing/shared/contracts"
	"ride-sharing/shared/types"
)

type HttpHandler struct {
	Service core.TripService
}

type PreviewTripRequest struct {
	UserID      string           `json:"userID"`
	PickUp      types.Coordinate `json:"pickup"`
	Destination types.Coordinate `json:"destination"`
}

func HandleTripPreview(w http.ResponseWriter, r *http.Request) {
	// todo : call the trip service
	response := contracts.APIResponse{Data: "Ok"}
	writeJSON(w, http.StatusAccepted, response)
}

func writeJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Conent-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}
