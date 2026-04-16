package core

import (
	"context"
	"ride-sharing/shared/types"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// this contains the model classes and requried service definitions , which need to be implemented
// if you required to add mappers which used to map data between api and domain/service.
// add them here

type TripModel struct {
	ID primitive.ObjectID
	// user who book the trip
	UserID   string
	Status   string
	RideFare *RideFareModel
}

type TripRepository interface {
	CreateTrip(ctx *context.Context, trip *TripModel) (*TripModel, error)
}

type TripService interface {
	CreateTrip(ctx *context.Context, fare *RideFareModel) (*TripModel, error)
	GetRoute(ctx *context.Context, pickup, destination types.Coordinate) (*types.Route, error)
}
