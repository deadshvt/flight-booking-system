package repository

import (
	"context"

	"github.com/deadshvt/flight-booking-system/bonus-service/internal/entity"
	"github.com/deadshvt/flight-booking-system/bonus-service/internal/repository/cache"
	"github.com/deadshvt/flight-booking-system/bonus-service/internal/repository/database"

	"github.com/rs/zerolog"
)

type BonusRepository struct {
	DB     database.BonusDB
	Cache  cache.BonusCache
	Logger zerolog.Logger
}

func NewBonusRepository(db database.BonusDB, cache cache.BonusCache, logger zerolog.Logger) *BonusRepository {
	return &BonusRepository{
		DB:     db,
		Cache:  cache,
		Logger: logger,
	}
}

func (r *BonusRepository) LoadCacheFromDB(ctx context.Context) error {
	r.Logger.Info().Msg("Loading cache from db...")

	const limit = 100
	offset := int32(0)

	for {
		privileges, err := r.DB.GetPrivileges(ctx, limit, offset)
		if err != nil {
			return err
		}

		if len(privileges) == 0 {
			break
		}

		err = r.Cache.SetPrivileges(ctx, privileges)
		if err != nil {
			return err
		}

		offset += limit
	}

	return nil
}

func (r *BonusRepository) GetPrivilegeWithHistory(ctx context.Context,
	username string) (*entity.PrivilegeWithHistory, error) {
	r.Logger.Info().Msgf("Getting privilege with history with username: %s...", username)

	privilege, err := r.DB.GetPrivilege(ctx, username)
	if err != nil {
		return nil, err
	}

	history, err := r.DB.GetHistory(ctx, privilege.ID)
	if err != nil {
		return nil, err
	}

	return &entity.PrivilegeWithHistory{
		Privilege: privilege,
		History:   history,
	}, nil
}

func (r *BonusRepository) GetPrivilege(ctx context.Context, username string) (*entity.Privilege, error) {
	r.Logger.Info().Msgf("Getting privilege with username: %s...", username)

	privilege, err := r.Cache.GetPrivilege(ctx, username)
	if err != nil {
		return r.DB.GetPrivilege(ctx, username)
	}

	return privilege, nil
}

func (r *BonusRepository) CreatePrivilege(ctx context.Context, privilege *entity.Privilege) error {
	r.Logger.Info().Msgf("Creating privilege with username: %s...", privilege.Username)

	err := r.DB.CreatePrivilege(ctx, privilege)
	if err != nil {
		return err
	}

	return r.Cache.SetPrivilege(ctx, privilege)
}

func (r *BonusRepository) UpdatePrivilege(ctx context.Context, privilege *entity.Privilege) error {
	r.Logger.Info().Msgf("Updating privilege with username: %s...", privilege.Username)

	err := r.DB.UpdatePrivilege(ctx, privilege)
	if err != nil {
		return err
	}

	return r.Cache.SetPrivilege(ctx, privilege)
}

func (r *BonusRepository) GetHistory(ctx context.Context, privilegeID int32) ([]*entity.Operation, error) {
	r.Logger.Info().Msgf("Getting history with privilegeID: %d...", privilegeID)

	return r.DB.GetHistory(ctx, privilegeID)
}

func (r *BonusRepository) CreateOperation(ctx context.Context, operation *entity.Operation) error {
	r.Logger.Info().Msgf("Creating history with privilegeID: %d...", operation.PrivilegeID)

	return r.DB.CreateOperation(ctx, operation)
}
