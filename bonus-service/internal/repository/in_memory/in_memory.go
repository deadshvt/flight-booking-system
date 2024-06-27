package in_memory

import (
	"context"
	"sync"

	"github.com/deadshvt/flight-booking-system/bonus-service/internal/entity"
	"github.com/deadshvt/flight-booking-system/bonus-service/internal/repository"
	"github.com/deadshvt/flight-booking-system/bonus-service/pkg/errs"
)

type InMemoryBonusRepository struct {
	privileges  map[string]*entity.Privilege
	history     map[int32][]*entity.History
	privilegeID int32
	historyID   int32
	mu          *sync.RWMutex
}

func NewInMemoryBonusRepository() repository.BonusRepository {
	return &InMemoryBonusRepository{
		privileges:  make(map[string]*entity.Privilege),
		history:     make(map[int32][]*entity.History),
		privilegeID: 1,
		historyID:   1,
		mu:          &sync.RWMutex{},
	}
}

func (r *InMemoryBonusRepository) GetPrivilegeWithHistory(ctx context.Context,
	username string) (*entity.PrivilegeWithHistory, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var privilege *entity.Privilege
	var ok bool

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		privilege, ok = r.privileges[username]
		if !ok {
			return nil, errs.ErrPrivilegeNotFound
		}

		if _, ok = r.history[privilege.ID]; !ok {
			return nil, errs.ErrHistoryNotFound
		}
	}

	return &entity.PrivilegeWithHistory{
		Privilege: *privilege,
		History:   r.history[privilege.ID],
	}, nil
}

func (r *InMemoryBonusRepository) GetPrivilege(ctx context.Context, username string) (*entity.Privilege, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		if _, ok := r.privileges[username]; !ok {
			return nil, errs.ErrPrivilegeNotFound
		}
	}

	return r.privileges[username], nil
}

func (r *InMemoryBonusRepository) CreatePrivilege(ctx context.Context, privilege *entity.Privilege) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	privilege.ID = r.privilegeID
	r.privilegeID += 1

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		if _, ok := r.privileges[privilege.Username]; ok {
			return errs.ErrPrivilegeAlreadyExists
		}
	}

	r.privileges[privilege.Username] = privilege

	return nil
}

func (r *InMemoryBonusRepository) UpdatePrivilege(ctx context.Context, privilege *entity.Privilege) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		if _, ok := r.privileges[privilege.Username]; !ok {
			return errs.ErrPrivilegeNotFound
		}
	}

	r.privileges[privilege.Username] = privilege

	return nil
}

func (r *InMemoryBonusRepository) GetHistory(ctx context.Context, privilegeID int32) ([]*entity.History, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		if _, ok := r.history[privilegeID]; !ok {
			return nil, errs.ErrHistoryNotFound
		}
	}

	return r.history[privilegeID], nil
}

func (r *InMemoryBonusRepository) CreateHistory(ctx context.Context, history *entity.History) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	history.ID = r.historyID
	r.historyID++

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		if _, ok := r.history[history.PrivilegeID]; ok {
			return errs.ErrHistoryAlreadyExists
		}
	}

	r.history[history.PrivilegeID] = append(r.history[history.PrivilegeID], history)

	return nil
}
