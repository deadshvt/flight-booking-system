package usecase

import (
	"github.com/deadshvt/flight-booking-system/bonus-service/internal/repository"

	"github.com/rs/zerolog"
)

type BonusUsecase struct {
	Repo   *repository.BonusRepository
	Logger zerolog.Logger
}

func NewBonusUsecase(repo *repository.BonusRepository, logger zerolog.Logger) *BonusUsecase {
	return &BonusUsecase{
		Repo:   repo,
		Logger: logger,
	}
}
