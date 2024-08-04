package usecase

import (
	"context"

	"github.com/deadshvt/flight-booking-system/ticket-service/internal/entity"
	"github.com/deadshvt/flight-booking-system/ticket-service/internal/repository"
	"github.com/deadshvt/flight-booking-system/ticket-service/pkg/errs"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

const (
	StatusPaid     = "PAID"
	StatusCanceled = "CANCELED"
)

type TicketUsecase struct {
	Repo   *repository.TicketRepository
	Logger zerolog.Logger
}

func NewTicketUsecase(repo *repository.TicketRepository, logger zerolog.Logger) *TicketUsecase {
	return &TicketUsecase{
		Repo:   repo,
		Logger: logger,
	}
}

func (u *TicketUsecase) GetTicket(ctx context.Context,
	username string, ticketUid string) (*entity.Ticket, error) {
	u.Logger.Info().Msg("Getting ticket...")

	ticket, err := u.Repo.GetTicket(ctx, username, ticketUid)
	if err != nil {
		u.Logger.Err(err).Msg("Failed to get ticket")
		return nil, err
	}

	if ticket.Username != username {
		u.Logger.Error().Msg("Ticket does not belong to user")
		return nil, errs.ErrTicketDoesNotBelongToUser
	}

	return ticket, nil
}

func (u *TicketUsecase) PurchaseTicket(ctx context.Context,
	username string, flightNumber string, price int32) (string, error) {
	u.Logger.Info().Msg("Purchasing ticket...")

	ticketUid := uuid.New().String()

	err := u.Repo.CreateTicket(ctx, &entity.Ticket{
		Username:     username,
		TicketUid:    ticketUid,
		FlightNumber: flightNumber,
		Price:        price,
		Status:       StatusPaid,
	})
	if err != nil {
		u.Logger.Err(err).Msg("Failed to create ticket")
		return "", err
	}

	return ticketUid, nil
}

func (u *TicketUsecase) ReturnTicket(ctx context.Context,
	username string, ticketUid string) error {
	u.Logger.Info().Msg("Returning ticket...")

	ticket, err := u.Repo.GetTicket(ctx, username, ticketUid)
	if err != nil {
		u.Logger.Err(err).Msg("Failed to get ticket")
		return err
	}

	ticket.Status = StatusCanceled

	err = u.Repo.UpdateTicket(ctx, ticket)
	if err != nil {
		u.Logger.Err(err).Msg("Failed to update ticket")
		return err
	}

	return nil
}
