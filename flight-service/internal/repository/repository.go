package repository

import (
	"context"

	"github.com/deadshvt/flight-booking-system/flight-service/internal/entity"
	"github.com/deadshvt/flight-booking-system/flight-service/internal/repository/cache"
	"github.com/deadshvt/flight-booking-system/flight-service/internal/repository/database"

	"github.com/rs/zerolog"
)

type FlightRepository struct {
	DB     database.FlightDB
	Cache  cache.FlightCache
	Logger zerolog.Logger
}

func NewFlightRepository(db database.FlightDB, cache cache.FlightCache, logger zerolog.Logger) *FlightRepository {
	return &FlightRepository{
		DB:     db,
		Cache:  cache,
		Logger: logger,
	}
}

func (r *FlightRepository) LoadCacheFromDB(ctx context.Context) error {
	r.Logger.Info().Msg("Loading cache from db...")

	const limit = 100
	offset := int32(0)

	for {
		flights, err := r.DB.GetFlightsWithAirports(ctx, limit, offset)
		if err != nil {
			return err
		}

		if len(flights) == 0 {
			break
		}

		err = r.Cache.SetFlightsWithAirports(ctx, flights)
		if err != nil {
			return err
		}

		offset += limit
	}

	return nil
}

func (r *FlightRepository) GetFlightsWithAirports(ctx context.Context,
	limit int32, offset int32) ([]*entity.FlightWithAirports, error) {
	r.Logger.Info().Msgf("Getting flights with airports...")

	flights, err := r.Cache.GetFlightsWithAirports(ctx, limit, offset)
	if err != nil {
		return r.DB.GetFlightsWithAirports(ctx, limit, offset)
	}

	return flights, nil
}

func (r *FlightRepository) GetFlightWithAirports(ctx context.Context,
	flightNumber string) (*entity.FlightWithAirports, error) {
	r.Logger.Info().Msgf("Getting flight with airports with number: %s...", flightNumber)

	flight, err := r.Cache.GetFlightWithAirports(ctx, flightNumber)
	if err != nil {
		return r.DB.GetFlightWithAirports(ctx, flightNumber)
	}

	return flight, nil
}

func (r *FlightRepository) GetFlight(ctx context.Context, flightNumber string) (*entity.Flight, error) {
	r.Logger.Info().Msgf("Getting flight with number: %s...", flightNumber)

	return r.DB.GetFlight(ctx, flightNumber)
}

func (r *FlightRepository) CreateFlight(ctx context.Context, flight *entity.Flight) error {
	r.Logger.Info().Msgf("Creating flight with number: %s...", flight.FlightNumber)

	fromAirport, err := r.DB.GetAirport(ctx, flight.FromAirportID)
	if err != nil {
		return err
	}

	toAirport, err := r.DB.GetAirport(ctx, flight.ToAirportID)
	if err != nil {
		return err
	}

	err = r.DB.CreateFlight(ctx, flight)
	if err != nil {
		return err
	}

	return r.Cache.SetFlightWithAirports(ctx, &entity.FlightWithAirports{
		FlightNumber: flight.FlightNumber,
		FromAirport:  fromAirport.FullName(),
		ToAirport:    toAirport.FullName(),
		Date:         flight.Date,
		Price:        flight.Price,
	})
}

func (r *FlightRepository) GetAirport(ctx context.Context, airportID int32) (*entity.Airport, error) {
	r.Logger.Info().Msgf("Getting airport with ID: %d...", airportID)

	return r.DB.GetAirport(ctx, airportID)
}

func (r *FlightRepository) CreateAirport(ctx context.Context, airport *entity.Airport) error {
	r.Logger.Info().Msgf("Creating airport with name: %s...", airport.Name)

	return r.DB.CreateAirport(ctx, airport)
}
