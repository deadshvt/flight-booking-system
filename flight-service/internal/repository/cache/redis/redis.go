package redis

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/deadshvt/flight-booking-system/config"
	"github.com/deadshvt/flight-booking-system/flight-service/internal/entity"
	"github.com/deadshvt/flight-booking-system/flight-service/pkg/errs"

	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
)

type Redis struct {
	Conn   *redis.Client
	Logger zerolog.Logger
}

func (cache *Redis) Connect(ctx context.Context) error {
	cache.Logger.Info().Msg("Connecting to Redis...")

	config.Load(".env") //TODO

	cache.Conn = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0, //TODO
	})

	return cache.Conn.Ping(ctx).Err()
}

func (cache *Redis) Disconnect(ctx context.Context) error {
	cache.Logger.Info().Msg("Disconnecting from Redis...")

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		if cache.Conn == nil {
			return nil
		}
	}

	return cache.Conn.Close()
}

func (cache *Redis) GetFlightWithAirports(ctx context.Context,
	flightNumber string) (*entity.FlightWithAirports, error) {
	cache.Logger.Info().Msgf("Getting flight with airports...")

	key := GetFlightKey(flightNumber)

	data, err := cache.Conn.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		return nil, errs.ErrFlightNotFound
	}
	if err != nil {
		return nil, err
	}

	var flight entity.FlightWithAirports
	err = json.Unmarshal([]byte(data), &flight)
	if err != nil {
		return nil, err
	}

	return &flight, nil
}

func (cache *Redis) SetFlightWithAirports(ctx context.Context, flight *entity.FlightWithAirports) error {
	cache.Logger.Info().Msgf("Setting flight with airports...")

	data, err := json.Marshal(flight)
	if err != nil {
		return err
	}

	key := GetFlightKey(flight.FlightNumber)

	err = cache.Conn.Set(ctx, key, data, 0).Err()
	if err != nil {
		return err
	}

	err = cache.Conn.RPush(ctx, "flight_keys", key).Err()
	if err != nil {
		return err
	}

	return nil
}

func (cache *Redis) GetFlightsWithAirports(ctx context.Context,
	limit int32, offset int32) ([]*entity.FlightWithAirports, error) {
	cache.Logger.Info().Msgf("Getting flights with airports...")

	keys, err := cache.Conn.LRange(ctx, "flight_keys", int64(offset), int64(offset+limit-1)).Result()
	if err != nil {
		return nil, err
	}

	pipe := cache.Conn.TxPipeline()
	cmds := make([]*redis.StringCmd, len(keys))
	for i := range keys {
		cmds[i] = pipe.Get(ctx, keys[i])
	}

	_, err = pipe.Exec(ctx)
	if err != nil {
		return nil, err
	}

	var flights []*entity.FlightWithAirports
	for _, cmd := range cmds {
		data, err := cmd.Result()
		if err != nil {
			return nil, err
		}

		var flight entity.FlightWithAirports
		err = json.Unmarshal([]byte(data), &flight)
		if err != nil {
			return nil, err
		}

		flights = append(flights, &flight)
	}

	return flights, nil
}

func (cache *Redis) SetFlightsWithAirports(ctx context.Context, flights []*entity.FlightWithAirports) error {
	cache.Logger.Info().Msg("Setting flights with airports...")

	pipe := cache.Conn.TxPipeline()

	for _, flight := range flights {
		data, err := json.Marshal(flight)
		if err != nil {
			return err
		}

		key := GetFlightKey(flight.FlightNumber)
		pipe.Set(ctx, key, data, 0)
		pipe.RPush(ctx, "flight_keys", key)
	}

	_, err := pipe.Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

func GetFlightKey(flightNumber string) string {
	return fmt.Sprintf("flight:%s", flightNumber)
}
