package im

import (
	"context"
	"sync"

	"github.com/deadshvt/flight-booking-system/flight-service/internal/entity"
	"github.com/deadshvt/flight-booking-system/flight-service/internal/repository/database"
	"github.com/deadshvt/flight-booking-system/flight-service/pkg/errs"
	"github.com/deadshvt/flight-booking-system/pkg/omap"
)

type InMemoryFlightDB struct {
	Flights   *omap.OrderedMap[string, *entity.Flight]
	Airports  map[int32]*entity.Airport
	FlightID  int32
	AirportID int32
	Mu        *sync.RWMutex
}

func NewInMemoryFlightDB() database.FlightDB {
	return &InMemoryFlightDB{
		Flights:   omap.NewOrderedMap[string, *entity.Flight](),
		Airports:  make(map[int32]*entity.Airport),
		FlightID:  1,
		AirportID: 1,
		Mu:        &sync.RWMutex{},
	}
}

func (db *InMemoryFlightDB) Connect(_ context.Context) error {
	return nil
}

func (db *InMemoryFlightDB) Disconnect(_ context.Context) error {
	return nil
}

func (db *InMemoryFlightDB) GetFlightsWithAirports(ctx context.Context,
	limit int32, offset int32) ([]*entity.FlightWithAirports, error) {
	db.Mu.RLock()
	defer db.Mu.RUnlock()

	var flightsWithAirports []*entity.FlightWithAirports

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		keys := db.Flights.Keys()

		start := offset
		end := offset + limit
		length := int32(len(keys))

		if start >= length {
			return []*entity.FlightWithAirports{}, nil
		}

		if end > length {
			end = length
		}

		for i := start; i < end; i++ {
			flightNumber := keys[i]
			flight, _ := db.Flights.Get(flightNumber)
			fromAirport := db.Airports[flight.FromAirportID]
			toAirport := db.Airports[flight.ToAirportID]

			flightsWithAirports = append(flightsWithAirports, &entity.FlightWithAirports{
				FlightNumber: flightNumber,
				FromAirport:  fromAirport.FullName(),
				ToAirport:    toAirport.FullName(),
				Date:         flight.Date,
				Price:        flight.Price,
			})
		}
	}

	return flightsWithAirports, nil
}

func (db *InMemoryFlightDB) GetFlightWithAirports(ctx context.Context,
	flightNumber string) (*entity.FlightWithAirports, error) {
	db.Mu.RLock()
	defer db.Mu.RUnlock()

	var flight *entity.Flight
	var fromAirport *entity.Airport
	var toAirport *entity.Airport
	var ok bool

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		flight, ok = db.Flights.Get(flightNumber)
		if !ok {
			return nil, errs.ErrFlightNotFound
		}

		fromAirport, ok = db.Airports[flight.FromAirportID]
		if !ok {
			return nil, errs.ErrAirportNotFound
		}

		toAirport, ok = db.Airports[flight.ToAirportID]
		if !ok {
			return nil, errs.ErrAirportNotFound
		}
	}

	return &entity.FlightWithAirports{
		FlightNumber: flightNumber,
		FromAirport:  fromAirport.FullName(),
		ToAirport:    toAirport.FullName(),
		Date:         flight.Date,
		Price:        flight.Price,
	}, nil
}

func (db *InMemoryFlightDB) GetFlight(ctx context.Context, flightNumber string) (*entity.Flight, error) {
	db.Mu.RLock()
	defer db.Mu.RUnlock()

	var flight *entity.Flight
	var ok bool

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		flight, ok = db.Flights.Get(flightNumber)
		if !ok {
			return nil, errs.ErrFlightNotFound
		}
	}

	return flight, nil
}

func (db *InMemoryFlightDB) CreateFlight(ctx context.Context, flight *entity.Flight) error {
	db.Mu.Lock()
	defer db.Mu.Unlock()

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		flight.ID = db.FlightID
		db.FlightID++

		_, ok := db.Flights.Get(flight.FlightNumber)
		if ok {
			return errs.ErrFlightAlreadyExists
		}
	}

	db.Flights.Set(flight.FlightNumber, flight)

	return nil
}

func (db *InMemoryFlightDB) GetAirport(ctx context.Context, airportID int32) (*entity.Airport, error) {
	db.Mu.RLock()
	defer db.Mu.RUnlock()

	var airport *entity.Airport
	var ok bool

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		airport, ok = db.Airports[airportID]
		if !ok {
			return nil, errs.ErrAirportNotFound
		}
	}

	return airport, nil
}

func (db *InMemoryFlightDB) CreateAirport(ctx context.Context, airport *entity.Airport) error {
	db.Mu.Lock()
	defer db.Mu.Unlock()

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		airport.ID = db.AirportID
		db.AirportID++

		_, ok := db.Airports[airport.ID]
		if ok {
			return errs.ErrAirportAlreadyExists
		}
	}

	db.Airports[airport.ID] = airport

	return nil
}

func (db *InMemoryFlightDB) CreateAirports(ctx context.Context, airports []*entity.Airport) error {
	db.Mu.Lock()
	defer db.Mu.Unlock()

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		for _, airport := range airports {
			airport.ID = db.AirportID
			db.AirportID++
		}

		for _, airport := range airports {
			_, ok := db.Airports[airport.ID]
			if ok {
				return errs.ErrAirportAlreadyExists
			}
		}
	}

	for _, airport := range airports {
		db.Airports[airport.ID] = airport
	}

	return nil
}

func (db *InMemoryFlightDB) AirportTableExists(_ context.Context) (bool, error) {
	return true, nil
}
