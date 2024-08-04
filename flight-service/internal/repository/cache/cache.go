package cache

import (
	"context"

	"github.com/deadshvt/flight-booking-system/flight-service/internal/entity"
	"github.com/deadshvt/flight-booking-system/flight-service/internal/repository/cache/redis"
	"github.com/deadshvt/flight-booking-system/flight-service/pkg/errs"
)

const (
	Redis = "redis"
)

type FlightCache interface {
	Connect(ctx context.Context) error
	Disconnect(ctx context.Context) error

	GetFlightWithAirports(ctx context.Context, flightNumber string) (*entity.FlightWithAirports, error)
	SetFlightWithAirports(ctx context.Context, flight *entity.FlightWithAirports) error

	GetFlightsWithAirports(ctx context.Context, limit int32, offset int32) ([]*entity.FlightWithAirports, error)
	SetFlightsWithAirports(ctx context.Context, flights []*entity.FlightWithAirports) error
}

func NewFlightCache(ctx context.Context, cacheType string) (FlightCache, error) {
	var cache FlightCache

	switch cacheType {
	case Redis:
		cache = &redis.Redis{}
	default:
		return nil, errs.ErrUnsupportedCacheType
	}

	err := cache.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return cache, nil
}
