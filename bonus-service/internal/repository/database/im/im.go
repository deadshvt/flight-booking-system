package im

import (
	"context"
	"sync"

	"github.com/deadshvt/flight-booking-system/bonus-service/internal/entity"
	"github.com/deadshvt/flight-booking-system/bonus-service/internal/repository/database"
	"github.com/deadshvt/flight-booking-system/bonus-service/pkg/errs"
	"github.com/deadshvt/flight-booking-system/pkg/omap"
)

type InMemoryBonusDB struct {
	Privileges  *omap.OrderedMap[string, *entity.Privilege]
	Histories   map[int32][]*entity.Operation
	PrivilegeID int32
	OperationID int32
	Mu          *sync.RWMutex
}

func NewInMemoryBonusDB() database.BonusDB {
	return &InMemoryBonusDB{
		Privileges:  omap.NewOrderedMap[string, *entity.Privilege](),
		Histories:   make(map[int32][]*entity.Operation),
		PrivilegeID: 1,
		OperationID: 1,
		Mu:          &sync.RWMutex{},
	}
}

func (db *InMemoryBonusDB) Connect(_ context.Context) error {
	return nil
}

func (db *InMemoryBonusDB) Disconnect(_ context.Context) error {
	return nil
}

func (db *InMemoryBonusDB) GetPrivileges(ctx context.Context, limit int32, offset int32) ([]*entity.Privilege, error) {
	db.Mu.RLock()
	defer db.Mu.RUnlock()

	var privileges []*entity.Privilege

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		keys := db.Privileges.Keys()

		start := offset
		end := offset + limit
		length := int32(len(keys))

		if start >= length {
			return []*entity.Privilege{}, nil
		}

		if end > length {
			end = length
		}

		for i := start; i < end; i++ {
			privilege, _ := db.Privileges.Get(keys[i])
			privileges = append(privileges, privilege)
		}
	}

	return privileges, nil
}

func (db *InMemoryBonusDB) GetPrivilege(ctx context.Context, username string) (*entity.Privilege, error) {
	db.Mu.RLock()
	defer db.Mu.RUnlock()

	var privilege *entity.Privilege
	var ok bool

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		privilege, ok = db.Privileges.Get(username)
		if !ok {
			return nil, errs.ErrPrivilegeNotFound
		}
	}

	return privilege, nil
}

func (db *InMemoryBonusDB) CreatePrivilege(ctx context.Context, privilege *entity.Privilege) error {
	db.Mu.Lock()
	defer db.Mu.Unlock()

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		privilege.ID = db.PrivilegeID
		db.PrivilegeID += 1

		_, ok := db.Privileges.Get(privilege.Username)
		if ok {
			return errs.ErrPrivilegeAlreadyExists
		}
	}

	db.Privileges.Set(privilege.Username, privilege)

	return nil
}

func (db *InMemoryBonusDB) UpdatePrivilege(ctx context.Context, privilege *entity.Privilege) error {
	db.Mu.Lock()
	defer db.Mu.Unlock()

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		_, ok := db.Privileges.Get(privilege.Username)
		if !ok {
			return errs.ErrPrivilegeNotFound
		}
	}

	db.Privileges.Set(privilege.Username, privilege)

	return nil
}

func (db *InMemoryBonusDB) GetHistory(ctx context.Context, privilegeID int32) ([]*entity.Operation, error) {
	db.Mu.RLock()
	defer db.Mu.RUnlock()

	var operations []*entity.Operation
	var ok bool

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		operations, ok = db.Histories[privilegeID]
		if !ok {
			return nil, errs.ErrHistoryNotFound
		}
	}

	return operations, nil
}

func (db *InMemoryBonusDB) CreateOperation(ctx context.Context, operation *entity.Operation) error {
	db.Mu.Lock()
	defer db.Mu.Unlock()

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		operation.ID = db.OperationID
		db.OperationID++

		_, ok := db.Histories[operation.PrivilegeID]
		if ok {
			return errs.ErrOperationAlreadyExists
		}
	}

	db.Histories[operation.PrivilegeID] = append(db.Histories[operation.PrivilegeID], operation)

	return nil
}
