package database

import (
	"context"

	"github.com/deadshvt/flight-booking-system/bonus-service/internal/entity"
	"github.com/deadshvt/flight-booking-system/bonus-service/internal/repository/database/postgres"
	"github.com/deadshvt/flight-booking-system/bonus-service/pkg/errs"
)

const (
	Postgres = "postgres"
)

type BonusDB interface {
	Connect(ctx context.Context) error
	Disconnect(ctx context.Context) error

	GetPrivileges(ctx context.Context, limit int32, offset int32) ([]*entity.Privilege, error)
	GetPrivilege(ctx context.Context, username string) (*entity.Privilege, error)
	CreatePrivilege(ctx context.Context, privilege *entity.Privilege) error
	UpdatePrivilege(ctx context.Context, privilege *entity.Privilege) error

	GetHistory(ctx context.Context, privilegeID int32) ([]*entity.Operation, error)
	CreateOperation(ctx context.Context, operation *entity.Operation) error
}

func NewBonusDB(ctx context.Context, dbType string) (BonusDB, error) {
	var db BonusDB

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
