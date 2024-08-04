package cache

import (
	"context"

	"github.com/deadshvt/flight-booking-system/ticket-service/internal/entity"
	"github.com/deadshvt/flight-booking-system/ticket-service/internal/repository/cache/redis"
	"github.com/deadshvt/flight-booking-system/ticket-service/pkg/errs"
)

const (
	Redis = "redis"
)

type TicketCache interface {
	Connect(ctx context.Context) error
	Disconnect(ctx context.Context) error

	GetTicket(ctx context.Context, username string, ticketUid string) (*entity.Ticket, error)
	SetTicket(ctx context.Context, username string, ticket *entity.Ticket) error

	GetTickets(ctx context.Context, username string) ([]*entity.Ticket, error)
	SetTickets(ctx context.Context, username string, tickets []*entity.Ticket) error
}

func NewTicketCache(ctx context.Context, cacheType string) (TicketCache, error) {
	var cache TicketCache

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
