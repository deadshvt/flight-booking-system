package redis

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/deadshvt/flight-booking-system/config"
	"github.com/deadshvt/flight-booking-system/ticket-service/internal/entity"
	"github.com/deadshvt/flight-booking-system/ticket-service/pkg/errs"

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

func (cache *Redis) GetTicket(ctx context.Context, username string, ticketUid string) (*entity.Ticket, error) {
	cache.Logger.Info().Msgf("Getting ticket...")

	key := GetTicketKey(username)

	data, err := cache.Conn.HGet(ctx, key, ticketUid).Result()
	if errors.Is(err, redis.Nil) {
		return nil, errs.ErrTicketNotFound
	}
	if err != nil {
		return nil, err
	}

	var ticket entity.Ticket
	err = json.Unmarshal([]byte(data), &ticket)
	if err != nil {
		return nil, err
	}

	return &ticket, nil
}

func (cache *Redis) SetTicket(ctx context.Context, username string, ticket *entity.Ticket) error {
	cache.Logger.Info().Msgf("Setting ticket...")

	data, err := json.Marshal(ticket)
	if err != nil {
		return err
	}

	key := GetTicketKey(username)

	return cache.Conn.HSet(ctx, key, ticket.TicketUid, data).Err()
}

func (cache *Redis) GetTickets(ctx context.Context, username string) ([]*entity.Ticket, error) {
	cache.Logger.Info().Msgf("Getting tickets...")

	key := GetTicketKey(username)

	fields, err := cache.Conn.HGetAll(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	if len(fields) == 0 {
		return nil, errs.ErrTicketsNotFound
	}

	var tickets []*entity.Ticket
	for i := range fields {
		var ticket entity.Ticket
		err = json.Unmarshal([]byte(fields[i]), &ticket)
		if err != nil {
			return nil, err
		}

		tickets = append(tickets, &ticket)
	}

	return tickets, nil
}

func (cache *Redis) SetTickets(ctx context.Context, username string, tickets []*entity.Ticket) error {
	cache.Logger.Info().Msgf("Setting tickets...")

	key := GetTicketKey(username)

	pipe := cache.Conn.TxPipeline()

	for _, ticket := range tickets {
		data, err := json.Marshal(ticket)
		if err != nil {
			return err
		}

		pipe.HSet(ctx, key, ticket.TicketUid, data)
	}

	_, err := pipe.Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

func GetTicketKey(username string) string {
	return fmt.Sprintf("ticket:%s", username)
}
