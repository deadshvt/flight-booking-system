package database

import (
	"context"

	"github.com/deadshvt/flight-booking-system/flight-service/internal/entity"
	"github.com/deadshvt/flight-booking-system/flight-service/internal/repository/database/postgres"
	"github.com/deadshvt/flight-booking-system/flight-service/pkg/errs"
)

const (
	Postgres = "postgres"
)

type FlightDB interface {
	Connect(ctx context.Context) error
	Disconnect(ctx context.Context) error

	GetFlightsWithAirports(ctx context.Context, limit int32, offset int32) ([]*entity.FlightWithAirports, error)
	GetFlightWithAirports(ctx context.Context, flightNumber string) (*entity.FlightWithAirports, error)
	GetFlight(ctx context.Context, flightNumber string) (*entity.Flight, error)
	CreateFlight(ctx context.Context, flight *entity.Flight) error

	GetAirport(ctx context.Context, airportID int32) (*entity.Airport, error)
	CreateAirport(ctx context.Context, airport *entity.Airport) error
	CreateAirports(ctx context.Context, airports []*entity.Airport) error
	AirportTableExists(ctx context.Context) (bool, error)
}

func NewFlightDB(ctx context.Context, dbType string) (FlightDB, error) {
	var db FlightDB

	switch dbType {
	case Postgres:
		db = &postgres.Postgres{}
	default:
		return nil, errs.ErrUnsupportedDBType
	}

	err := db.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
