package repository

import (
	"context"
	"ride-sharing/services/trip-service/internal/domain"
)

type InmemRepository struct {
	trips     map[string]*domain.TripModel
	rideFares map[string]*domain.RideFareModel
}

func NewInmemRepository() *InmemRepository {
	return &InmemRepository{
		trips:     make(map[string]*domain.TripModel),
		rideFares: make(map[string]*domain.RideFareModel),
	}
}

func (in *InmemRepository) CreateTrip(ctx context.Context, trip *domain.TripModel) (*domain.TripModel, error) {
	in.trips[trip.ID.Hex()] = trip
	return trip, nil
}

func (in *InmemRepository) Create(domain.TripModel) {

}

func (in *InmemRepository) Update(domain.TripModel) {

}

func (in *InmemRepository) Delete(domain.TripModel) {

}
