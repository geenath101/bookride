package core

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"ride-sharing/shared/types"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DefaultTripService struct {
	repo TripRepository
}

func NewService(repo TripRepository) *DefaultTripService {
	return &DefaultTripService{
		repo: repo,
	}
}

func (t DefaultTripService) CreateTrip(ctx *context.Context, fare *RideFareModel) (*TripModel, error) {
	m := &TripModel{
		ID:       primitive.NewObjectID(),
		Status:   "InProgress",
		RideFare: fare,
	}
	return t.repo.CreateTrip(ctx, m)
}

func (s DefaultTripService) GetRoute(ctx *context.Context, pickup, destination types.Coordinate) (*types.Route, error) {
	url := fmt.Sprintf("http://router.project-osrm.org/route/v1/driving/%f,%f;%f,%f?overview=full&geometrices=geojson",
		pickup.Longitude, pickup.Latitude, destination.Longitude, destination.Latitude)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch route from OSRM server", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read the response: %v", err)

	}

	var routeResp types.OsrmApiResponse
	if err := json.Unmarshal(body, &routeResp); err != nil {
		return nil, fmt.Errorf("failed to parse the response: %v", err)
	}

	return &routeResp, nil
}
