package repository

import (
	"context"
	"ride-sharing/services/trip-service/internal/core"
)

type InmemRepository struct {
	trips     map[string]*core.TripModel
	rideFares map[string]*core.RideFareModel
}

func NewInmemRepository() *InmemRepository {
	return &InmemRepository{
		trips:     make(map[string]*core.TripModel),
		rideFares: make(map[string]*core.RideFareModel),
	}
}

func (in *InmemRepository) CreateTrip(ctx context.Context, trip *core.TripModel) (*core.TripModel, error) {
	in.trips[trip.ID.Hex()] = trip
	return trip, nil
}

func (in *InmemRepository) Create(core.TripModel) {

}

func (in *InmemRepository) Update(core.TripModel) {

}

func (in *InmemRepository) Delete(core.TripModel) {

}
