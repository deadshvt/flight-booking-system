package redis

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/deadshvt/flight-booking-system/bonus-service/internal/entity"
	"github.com/deadshvt/flight-booking-system/bonus-service/pkg/errs"
	"github.com/deadshvt/flight-booking-system/config"

	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
)

type Redis struct {
	Conn   *redis.Client
	Logger zerolog.Logger
}

func (cache *Redis) Connect(ctx context.Context) error {
	cache.Logger.Info().Msg("Connecting to Redis...")

	config.Load(".env") //TODO

	cache.Conn = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0, //TODO
	})

	return cache.Conn.Ping(ctx).Err()
}

func (cache *Redis) Disconnect(ctx context.Context) error {
	cache.Logger.Info().Msg("Disconnecting from Redis...")

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		if cache.Conn == nil {
			return nil
		}
	}

	return cache.Conn.Close()
}

func (cache *Redis) GetPrivilege(ctx context.Context, username string) (*entity.Privilege, error) {
	cache.Logger.Info().Msgf("Getting privilege...")

	key := GetPrivilegeKey(username)

	data, err := cache.Conn.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		return nil, errs.ErrPrivilegeNotFound
	}
	if err != nil {
		return nil, err
	}

	var privilege entity.Privilege
	err = json.Unmarshal([]byte(data), &privilege)
	if err != nil {
		return nil, err
	}

	return &privilege, nil
}

func (cache *Redis) SetPrivilege(ctx context.Context, privilege *entity.Privilege) error {
	cache.Logger.Info().Msgf("Setting privilege...")

	data, err := json.Marshal(privilege)
	if err != nil {
		return err
	}

	key := GetPrivilegeKey(privilege.Username)

	err = cache.Conn.Set(ctx, key, data, 0).Err()
	if err != nil {
		return err
	}

	err = cache.Conn.RPush(ctx, "privilege_keys", key).Err()
	if err != nil {
		return err
	}

	return nil
}

func (cache *Redis) SetPrivileges(ctx context.Context, privileges []*entity.Privilege) error {
	cache.Logger.Info().Msg("Setting privileges...")

	pipe := cache.Conn.TxPipeline()

	for _, privilege := range privileges {
		data, err := json.Marshal(privilege)
		if err != nil {
			return err
		}

		key := GetPrivilegeKey(privilege.Username)
		pipe.Set(ctx, key, data, 0)
		pipe.RPush(ctx, "privilege_keys", key)
	}

	_, err := pipe.Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

func GetPrivilegeKey(username string) string {
	return fmt.Sprintf("privilege:%s", username)
}
