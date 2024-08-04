package repository

import (
	"context"

	"github.com/deadshvt/flight-booking-system/ticket-service/internal/entity"
	"github.com/deadshvt/flight-booking-system/ticket-service/internal/repository/cache"
	"github.com/deadshvt/flight-booking-system/ticket-service/internal/repository/database"

	"github.com/rs/zerolog"
)

type TicketRepository struct {
	DB     database.TicketDB
	Cache  cache.TicketCache
	Logger zerolog.Logger
}

func NewTicketRepository(db database.TicketDB, cache cache.TicketCache, logger zerolog.Logger) *TicketRepository {
	return &TicketRepository{
		DB:     db,
		Cache:  cache,
		Logger: logger,
	}
}

func (r *TicketRepository) LoadCacheFromDB(ctx context.Context) error {
	r.Logger.Info().Msg("Loading cache from db...")

	ticketsByUsername := make(map[string][]*entity.Ticket)

	const limit = 100
	offset := int32(0)

	for {
		tickets, err := r.DB.GetTickets(ctx, limit, offset)
		if err != nil {
			return err
		}

		if len(tickets) == 0 {
			break
		}

		for i, ticket := range tickets {
			ticketsByUsername[ticket.Username] = append(ticketsByUsername[ticket.Username], tickets[i])
		}

		offset += limit
	}

	for username, tickets := range ticketsByUsername {
		err := r.Cache.SetTickets(ctx, username, tickets)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *TicketRepository) GetTickets(ctx context.Context, username string) ([]*entity.Ticket, error) {
	r.Logger.Info().Msg("Getting tickets...")

	tickets, err := r.Cache.GetTickets(ctx, username)
	if err != nil {
		return r.DB.GetTicketsByUsername(ctx, username)
	}

	return tickets, nil
}

func (r *TicketRepository) GetTicket(ctx context.Context, username string, ticketUid string) (*entity.Ticket, error) {
	r.Logger.Info().Msg("Getting ticket...")

	ticket, err := r.Cache.GetTicket(ctx, username, ticketUid)
	if err != nil {
		return r.DB.GetTicket(ctx, ticketUid)
	}

	return ticket, nil
}

func (r *TicketRepository) CreateTicket(ctx context.Context, ticket *entity.Ticket) error {
	r.Logger.Info().Msg("Creating ticket...")

	err := r.DB.CreateTicket(ctx, ticket)
	if err != nil {
		return err
	}

	return r.Cache.SetTicket(ctx, ticket.Username, ticket)
}

func (r *TicketRepository) UpdateTicket(ctx context.Context, ticket *entity.Ticket) error {
	r.Logger.Info().Msg("Updating ticket...")

	err := r.DB.UpdateTicket(ctx, ticket)
	if err != nil {
		return err
	}

	return r.Cache.SetTicket(ctx, ticket.Username, ticket)
}
