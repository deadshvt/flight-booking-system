package cache

import (
	"context"

	"github.com/deadshvt/flight-booking-system/bonus-service/internal/entity"
	"github.com/deadshvt/flight-booking-system/bonus-service/internal/repository/cache/redis"
	"github.com/deadshvt/flight-booking-system/bonus-service/pkg/errs"
)

const (
	Redis = "redis"
)

type BonusCache interface {
	Connect(ctx context.Context) error
	Disconnect(ctx context.Context) error

	GetPrivilege(ctx context.Context, username string) (*entity.Privilege, error)
	SetPrivilege(ctx context.Context, privilege *entity.Privilege) error

	SetPrivileges(ctx context.Context, privileges []*entity.Privilege) error
}

func NewBonusCache(ctx context.Context, cacheType string) (BonusCache, error) {
	var cache BonusCache

	switch cacheType {
	case Redis:
		cache = &redis.Redis{}
	default:
		return nil, errs.ErrUnsupportedCacheType
	}

	err := cache.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return cache, nil
}
