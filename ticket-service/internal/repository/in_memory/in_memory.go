package in_memory

import (
	"context"
	"sync"

	"github.com/deadshvt/flight-booking-system/ticket-service/internal/entity"
	"github.com/deadshvt/flight-booking-system/ticket-service/internal/repository"
	"github.com/deadshvt/flight-booking-system/ticket-service/pkg/errs"
)

type InMemoryTicketRepository struct {
	tickets  map[string]*entity.Ticket
	ticketID int32
	mu       *sync.RWMutex
}

func NewInMemoryTicketRepository() repository.TicketRepository {
	return &InMemoryTicketRepository{
		tickets:  make(map[string]*entity.Ticket),
		ticketID: 1,
		mu:       &sync.RWMutex{},
	}
}

func (r *InMemoryTicketRepository) GetTicket(ctx context.Context, ticketUid string) (*entity.Ticket, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		if _, ok := r.tickets[ticketUid]; !ok {
			return nil, errs.ErrTicketNotFound
		}
	}

	return r.tickets[ticketUid], nil
}

func (r *InMemoryTicketRepository) CreateTicket(ctx context.Context, ticket *entity.Ticket) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		if _, ok := r.tickets[ticket.TicketUid]; ok {
			return errs.ErrTicketAlreadyExists
		}
	}

	ticket.ID = r.ticketID
	r.ticketID++

	r.tickets[ticket.TicketUid] = ticket

	return nil
}

func (r *InMemoryTicketRepository) GetTickets(ctx context.Context, username string) ([]*entity.Ticket, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var tickets []*entity.Ticket

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		for _, ticket := range r.tickets {
			if ticket.Username == username {
				tickets = append(tickets, ticket)
			}
		}
	}

	return tickets, nil
}

func (r *InMemoryTicketRepository) UpdateTicket(ctx context.Context, ticket *entity.Ticket) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		if _, ok := r.tickets[ticket.TicketUid]; !ok {
			return errs.ErrTicketNotFound
		}
	}

	r.tickets[ticket.TicketUid] = ticket
	
	return nil
}
