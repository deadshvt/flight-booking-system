package repository

import (
	"context"

	"github.com/deadshvt/flight-booking-system/ticket-service/internal/entity"
)

type TicketRepository interface {
	GetTicket(ctx context.Context, ticketUid string) (*entity.Ticket, error)
	CreateTicket(ctx context.Context, ticket *entity.Ticket) error
	GetTickets(ctx context.Context, username string) ([]*entity.Ticket, error)
	UpdateTicket(ctx context.Context, ticket *entity.Ticket) error
}
