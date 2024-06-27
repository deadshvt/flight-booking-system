package repository

import (
	"context"

	"github.com/deadshvt/flight-booking-system/flight-service/internal/entity"
)

type FlightRepository interface {
	GetFlightWithAirports(ctx context.Context, flightNumber string) (*entity.FlightWithAirports, error)
	GetFlight(ctx context.Context, flightNumber string) (*entity.Flight, error)
	CreateFlight(ctx context.Context, flight *entity.Flight) error
	GetAirport(ctx context.Context, airportID int32) (*entity.Airport, error)
	CreateAirport(ctx context.Context, airport *entity.Airport) error
	GetFlightsWithAirports(ctx context.Context) ([]*entity.FlightWithAirports, error)
}
