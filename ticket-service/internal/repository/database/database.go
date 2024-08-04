package database

import (
	"context"

	"github.com/deadshvt/flight-booking-system/ticket-service/internal/entity"
	"github.com/deadshvt/flight-booking-system/ticket-service/internal/repository/database/postgres"
	"github.com/deadshvt/flight-booking-system/ticket-service/pkg/errs"
)

const (
	Postgres = "postgres"
)

type TicketDB interface {
	Connect(ctx context.Context) error
	Disconnect(ctx context.Context) error

	GetTickets(ctx context.Context, limit int32, offset int32) ([]*entity.Ticket, error)
	GetTicketsByUsername(ctx context.Context, username string) ([]*entity.Ticket, error)
	GetTicket(ctx context.Context, ticketUid string) (*entity.Ticket, error)
	CreateTicket(ctx context.Context, ticket *entity.Ticket) error
	UpdateTicket(ctx context.Context, ticket *entity.Ticket) error
}

func NewTicketDB(ctx context.Context, dbType string) (TicketDB, error) {
	var db TicketDB

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
