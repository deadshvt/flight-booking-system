package in_memory

import (
	"context"
	"sync"

	"github.com/deadshvt/flight-booking-system/flight-service/internal/entity"
	"github.com/deadshvt/flight-booking-system/flight-service/internal/repository"
	"github.com/deadshvt/flight-booking-system/flight-service/pkg/errs"
)

type InMemoryFlightRepository struct {
	flights   map[string]*entity.Flight
	airports  map[int32]*entity.Airport
	flightID  int32
	airportID int32
	mu        *sync.RWMutex
}

func NewInMemoryFlightRepository() repository.FlightRepository {
	return &InMemoryFlightRepository{
		flights:   make(map[string]*entity.Flight),
		airports:  make(map[int32]*entity.Airport),
		flightID:  1,
		airportID: 1,
		mu:        &sync.RWMutex{},
	}
}

func (r *InMemoryFlightRepository) GetFlightWithAirports(ctx context.Context,
	flightNumber string) (*entity.FlightWithAirports, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var flight *entity.Flight
	var ok bool

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		flight, ok = r.flights[flightNumber]
		if !ok {
			return nil, errs.ErrFlightNotFound
		}

		_, ok = r.airports[flight.FromAirportID]
		if !ok {
			return nil, errs.ErrAirportNotFound
		}

		_, ok = r.airports[flight.ToAirportID]
		if !ok {
			return nil, errs.ErrAirportNotFound
		}
	}

	return &entity.FlightWithAirports{
		Flight:      *flight,
		FromAirport: *r.airports[flight.FromAirportID],
		ToAirport:   *r.airports[flight.ToAirportID],
	}, nil
}

func (r *InMemoryFlightRepository) GetFlight(ctx context.Context, flightNumber string) (*entity.Flight, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		_, ok := r.flights[flightNumber]
		if !ok {
			return nil, errs.ErrFlightNotFound
		}
	}

	return r.flights[flightNumber], nil
}

func (r *InMemoryFlightRepository) CreateFlight(ctx context.Context, flight *entity.Flight) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	flight.ID = r.flightID
	r.flightID++

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		if _, ok := r.flights[flight.FlightNumber]; ok {
			return errs.ErrFlightAlreadyExists
		}
	}

	r.flights[flight.FlightNumber] = flight

	return nil
}

func (r *InMemoryFlightRepository) GetAirport(ctx context.Context, airportID int32) (*entity.Airport, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		if _, ok := r.airports[airportID]; !ok {
			return nil, errs.ErrAirportNotFound
		}
	}

	return r.airports[airportID], nil
}

func (r *InMemoryFlightRepository) CreateAirport(ctx context.Context, airport *entity.Airport) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	airport.ID = r.airportID
	r.airportID++

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		if _, ok := r.airports[airport.ID]; ok {
			return errs.ErrAirportAlreadyExists
		}
	}

	r.airports[airport.ID] = airport

	return nil
}

func (r *InMemoryFlightRepository) GetFlightsWithAirports(ctx context.Context) ([]*entity.FlightWithAirports, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var flights []*entity.FlightWithAirports

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		for _, flight := range r.flights {
			flights = append(flights, &entity.FlightWithAirports{
				Flight:      *flight,
				FromAirport: *r.airports[flight.FromAirportID],
				ToAirport:   *r.airports[flight.ToAirportID],
			})
		}
	}

	return flights, nil
}
