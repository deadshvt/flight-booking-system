package repository

import (
	"context"

	"github.com/deadshvt/flight-booking-system/bonus-service/internal/entity"
)

type BonusRepository interface {
	GetPrivilegeWithHistory(ctx context.Context, username string) (*entity.PrivilegeWithHistory, error)
	GetPrivilege(ctx context.Context, username string) (*entity.Privilege, error)
	CreatePrivilege(ctx context.Context, privilege *entity.Privilege) error
	UpdatePrivilege(ctx context.Context, privilege *entity.Privilege) error
	GetHistory(ctx context.Context, privilegeID int32) ([]*entity.History, error)
	CreateHistory(ctx context.Context, history *entity.History) error
}
