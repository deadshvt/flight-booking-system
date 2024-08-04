package im

import (
	"context"
	"sync"

	"github.com/deadshvt/flight-booking-system/pkg/omap"
	"github.com/deadshvt/flight-booking-system/ticket-service/internal/entity"
	"github.com/deadshvt/flight-booking-system/ticket-service/internal/repository/database"
	"github.com/deadshvt/flight-booking-system/ticket-service/pkg/errs"
)

type InMemoryTicketDB struct {
	Tickets  *omap.OrderedMap[string, *entity.Ticket]
	TicketID int32
	Mu       *sync.RWMutex
}

func NewInMemoryTicketDB() database.TicketDB {
	return &InMemoryTicketDB{
		Tickets:  omap.NewOrderedMap[string, *entity.Ticket](),
		TicketID: 1,
		Mu:       &sync.RWMutex{},
	}
}

func (db *InMemoryTicketDB) Connect(_ context.Context) error {
	return nil
}

func (db *InMemoryTicketDB) Disconnect(_ context.Context) error {
	return nil
}

func (db *InMemoryTicketDB) GetTickets(ctx context.Context, limit int32, offset int32) ([]*entity.Ticket, error) {
	db.Mu.RLock()
	defer db.Mu.RUnlock()

	var tickets []*entity.Ticket

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		keys := db.Tickets.Keys()

		start := offset
		end := offset + limit
		length := int32(len(keys))

		if start >= length {
			return []*entity.Ticket{}, nil
		}

		if end > length {
			end = length
		}

		for i := start; i < end; i++ {
			ticket, _ := db.Tickets.Get(keys[i])
			tickets = append(tickets, ticket)
		}
	}

	return tickets, nil
}

func (db *InMemoryTicketDB) GetTicketsByUsername(ctx context.Context, username string) ([]*entity.Ticket, error) {
	db.Mu.RLock()
	defer db.Mu.RUnlock()

	var tickets []*entity.Ticket

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		keys := db.Tickets.Keys()
		for _, key := range keys {
			ticket, _ := db.Tickets.Get(key)
			if ticket.Username == username {
				tickets = append(tickets, ticket)
			}
		}
	}

	return tickets, nil
}

func (db *InMemoryTicketDB) GetTicket(ctx context.Context, ticketUid string) (*entity.Ticket, error) {
	db.Mu.RLock()
	defer db.Mu.RUnlock()

	var ticket *entity.Ticket
	var ok bool

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		ticket, ok = db.Tickets.Get(ticketUid)
		if !ok {
			return nil, errs.ErrTicketNotFound
		}
	}

	return ticket, nil
}

func (db *InMemoryTicketDB) CreateTicket(ctx context.Context, ticket *entity.Ticket) error {
	db.Mu.Lock()
	defer db.Mu.Unlock()

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		_, ok := db.Tickets.Get(ticket.TicketUid)
		if ok {
			return errs.ErrTicketAlreadyExists
		}
	}

	ticket.ID = db.TicketID
	db.TicketID++

	db.Tickets.Set(ticket.TicketUid, ticket)

	return nil
}

func (db *InMemoryTicketDB) UpdateTicket(ctx context.Context, ticket *entity.Ticket) error {
	db.Mu.Lock()
	defer db.Mu.Unlock()

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		_, ok := db.Tickets.Get(ticket.TicketUid)
		if !ok {
			return errs.ErrTicketNotFound
		}
	}

	db.Tickets.Set(ticket.TicketUid, ticket)

	return nil
}
