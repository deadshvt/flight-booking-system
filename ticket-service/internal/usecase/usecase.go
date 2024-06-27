package usecase

import (
	"context"

	"github.com/deadshvt/flight-booking-system/ticket-service/internal/converter"
	"github.com/deadshvt/flight-booking-system/ticket-service/internal/repository"
	ticketpb "github.com/deadshvt/flight-booking-system/ticket-service/proto"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

const (
	StatusPaid     = "PAID"
	StatusCanceled = "CANCELED"
)

type TicketUsecase struct {
	Repo   repository.TicketRepository
	Logger zerolog.Logger
}

func NewTicketUsecase(repo repository.TicketRepository, logger zerolog.Logger) *TicketUsecase {
	return &TicketUsecase{
		Repo:   repo,
		Logger: logger,
	}
}

func (u *TicketUsecase) PurchaseTicket(ctx context.Context,
	req *ticketpb.PurchaseTicketRequest) (*ticketpb.PurchaseTicketResponse, error) {
	u.Logger.Info().Msg("Purchasing ticket...")

	ticketUid := uuid.New().String()

	protoTicket := &ticketpb.Ticket{
		Username:     req.Username,
		TicketUid:    ticketUid,
		FlightNumber: req.FlightNumber,
		Price:        req.Price,
		Status:       StatusPaid,
	}

	err := u.Repo.CreateTicket(ctx, converter.TicketFromProtoToEntity(protoTicket))
	if err != nil {
		u.Logger.Err(err).Msg("Failed to create ticket")
		return nil, err
	}

	return &ticketpb.PurchaseTicketResponse{TicketUid: ticketUid}, nil
}

func (u *TicketUsecase) ReturnTicket(ctx context.Context,
	req *ticketpb.ReturnTicketRequest) (*ticketpb.ReturnTicketResponse, error) {
	u.Logger.Info().Msg("Returning ticket...")

	ticket, err := u.Repo.GetTicket(ctx, req.TicketUid)
	if err != nil {
		u.Logger.Err(err).Msg("Failed to get ticket")
		return nil, err
	}

	ticket.Status = StatusCanceled

	err = u.Repo.UpdateTicket(ctx, ticket)
	if err != nil {
		u.Logger.Err(err).Msg("Failed to update ticket")
		return nil, err
	}

	return &ticketpb.ReturnTicketResponse{}, nil
}
