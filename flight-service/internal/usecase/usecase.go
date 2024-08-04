package usecase

import (
	"context"

	"github.com/deadshvt/flight-booking-system/flight-service/internal/entity"
	"github.com/deadshvt/flight-booking-system/flight-service/internal/repository"
	"github.com/deadshvt/flight-booking-system/flight-service/pkg/errs"

	"github.com/rs/zerolog"
)

type FlightUsecase struct {
	Repo   *repository.FlightRepository
	Logger zerolog.Logger
}

func NewFlightUsecase(repo *repository.FlightRepository, logger zerolog.Logger) *FlightUsecase {
	return &FlightUsecase{
		Repo:   repo,
		Logger: logger,
	}
}

func (u *FlightUsecase) GetFlightsWithAirports(ctx context.Context,
	page int32, size int32) ([]*entity.FlightWithAirports, error) {
	u.Logger.Info().Msg("Getting flights with airports...")

	if page < 1 {
		u.Logger.Error().Msg("Invalid page")
		return nil, errs.ErrInvalidPage
	}

	if size < 1 || size > 100 {
		u.Logger.Error().Msg("Invalid size")
		return nil, errs.ErrInvalidSize
	}

	flights, err := u.Repo.GetFlightsWithAirports(ctx, size, (page-1)*size)
	if err != nil {
		u.Logger.Err(err).Msg("Failed to get flights")
		return nil, err
	}

	return flights, nil
}
