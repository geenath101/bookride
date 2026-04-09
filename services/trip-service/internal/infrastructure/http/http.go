package http

import (
	"encoding/json"
	"net/http"
	"ride-sharing/services/trip-service/internal/service"
	"ride-sharing/shared/types"
)

type HttpHandler struct {
	Service service.TripService
}

type previewTripRequest struct {
	UserID      string           `json:"userID"`
	PickUp      types.Coordinate `json:"pickup"`
	Destination types.Coordinate `json:"destination"`
}

func (s *HttpHandler) HandleTripPreview(w http.ResponseWriter, r *http.Request) {
	var reqBody previewTripRequest
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "failed to parse json data ", http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	t, err := s.Service.GetRoute(ctx, &reqBody.PickUp, &reqBody.Destination)
	if err != nil {
		http.Error(w, "failed to fetch route information for trip preview ", http.StatusInternalServerError)
	}

	writeJSON(w, http.StatusOK, t)
}

func writeJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Conent-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}
