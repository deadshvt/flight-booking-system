package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"

	"github.com/deadshvt/flight-booking-system/bonus-service/internal/entity"
	"github.com/deadshvt/flight-booking-system/bonus-service/pkg/errs"
	"github.com/deadshvt/flight-booking-system/config"

	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
)

type TransactionFunc func(tx *sql.Tx) error

type Postgres struct {
	Conn   *sql.DB
	Logger zerolog.Logger
}

func (db *Postgres) Connect(ctx context.Context) error {
	db.Logger.Info().Msg("Connecting to Postgres...")

	config.Load(".env") //TODO

	dsn := os.Getenv("DB_DSN")
	var err error

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		db.Conn, err = sql.Open("postgres", dsn)
		if err != nil {
			return err
		}
	}

	return db.Conn.Ping()
}

func (db *Postgres) Disconnect(ctx context.Context) error {
	db.Logger.Info().Msg("Disconnecting from Postgres...")

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		if db.Conn == nil {
			return nil
		}
	}

	return db.Conn.Close()
}

func (db *Postgres) GetPrivileges(ctx context.Context, limit int32, offset int32) ([]*entity.Privilege, error) {
	db.Logger.Info().Msg("Getting privileges...")

	var privileges []*entity.Privilege

	err := db.WithTransaction(ctx, func(tx *sql.Tx) error {
		query := `
		SELECT
			p.id,
			p.username,
			p.status,
			p.balance
		FROM
			privilege p
		OFFSET $1
		LIMIT $2
		`

		stmt, err := tx.PrepareContext(ctx, query)
		if err != nil {
			return err
		}
		defer stmt.Close()

		rows, err := stmt.QueryContext(ctx, offset, limit)
		if errors.Is(err, sql.ErrNoRows) {
			return errs.ErrPrivilegesNotFound
		}
		if err != nil {
			return err
		}
		defer rows.Close()

		for rows.Next() {
			var privilege entity.Privilege
			err = rows.Scan(
				&privilege.ID,
				&privilege.Username,
				&privilege.Status,
				&privilege.Balance,
			)
			if err != nil {
				return err
			}

			privileges = append(privileges, &privilege)
		}

		return rows.Err()
	})

	if err != nil {
		return nil, err
	}

	return privileges, nil
}

func (db *Postgres) GetPrivilege(ctx context.Context, username string) (*entity.Privilege, error) {
	db.Logger.Info().Msg("Getting privilege...")

	var privilege *entity.Privilege

	err := db.WithTransaction(ctx, func(tx *sql.Tx) error {
		query := `
		SELECT
			p.id,
			p.username,
			p.status,
			p.balance
		FROM
			privilege p
		WHERE p.username = $1
		`

		stmt, err := tx.PrepareContext(ctx, query)
		if err != nil {
			return err
		}
		defer stmt.Close()

		err = stmt.QueryRowContext(ctx, username).
			Scan(&privilege.ID, &privilege.Username, &privilege.Status, &privilege.Balance)
		if errors.Is(err, sql.ErrNoRows) {
			return errs.ErrPrivilegeNotFound
		}
		return err
	})

	if err != nil {
		return nil, err
	}

	return privilege, nil
}

func (db *Postgres) CreatePrivilege(ctx context.Context, privilege *entity.Privilege) error {
	db.Logger.Info().Msg("Creating privilege...")

	return db.WithTransaction(ctx, func(tx *sql.Tx) error {
		query := `
		INSERT INTO flight (username, status, balance) 
		VALUES ($1, $2, $3)
		ON CONFLICT (username) DO NOTHING
		`

		stmt, err := tx.PrepareContext(ctx, query)
		if err != nil {
			return err
		}
		defer stmt.Close()

		res, err := stmt.ExecContext(ctx,
			privilege.Username, privilege.Status, privilege.Balance)
		if err != nil {
			return err
		}

		rowsAffected, err := res.RowsAffected()
		if err != nil {
			return err
		}
		if rowsAffected == 0 {
			return errs.ErrPrivilegeAlreadyExists
		}

		return nil
	})
}

func (db *Postgres) UpdatePrivilege(ctx context.Context, privilege *entity.Privilege) error {
	db.Logger.Info().Msg("Updating privilege...")

	return db.WithTransaction(ctx, func(tx *sql.Tx) error {
		query := `
		UPDATE privilege
		SET status = $2, balance = $3
		WHERE username = $1
		`

		stmt, err := tx.PrepareContext(ctx, query)
		if err != nil {
			return err
		}
		defer stmt.Close()

		res, err := stmt.ExecContext(ctx,
			privilege.Username, privilege.Status, privilege.Balance)
		if err != nil {
			return err
		}

		rowsAffected, err := res.RowsAffected()
		if err != nil {
			return err
		}
		if rowsAffected == 0 {
			return errs.ErrPrivilegeNotFound
		}

		return nil
	})
}

func (db *Postgres) GetHistory(ctx context.Context, privilegeID int32) ([]*entity.Operation, error) {
	db.Logger.Info().Msg("Getting history...")

	var history []*entity.Operation

	err := db.WithTransaction(ctx, func(tx *sql.Tx) error {
		query := `
		SELECT
			ph.id,
			ph.privilege_id,
			ph.ticket_uid,
			ph.datetime,
			ph.balance_diff,
			ph.operation_type
		FROM
			privilege_history ph
		WHERE ph.privilege_id = $1
		`

		stmt, err := tx.PrepareContext(ctx, query)
		if err != nil {
			return err
		}
		defer stmt.Close()

		rows, err := stmt.QueryContext(ctx, privilegeID)
		if errors.Is(err, sql.ErrNoRows) {
			return errs.ErrHistoryNotFound
		}
		if err != nil {
			return err
		}
		defer rows.Close()

		for rows.Next() {
			var operation entity.Operation
			err = rows.Scan(
				&operation.ID,
				&operation.PrivilegeID,
				&operation.TicketUid,
				&operation.Date,
				&operation.BalanceDiff,
				&operation.OperationType,
			)
			if err != nil {
				return err
			}

			history = append(history, &operation)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return history, nil
}

func (db *Postgres) CreateOperation(ctx context.Context, operation *entity.Operation) error {
	db.Logger.Info().Msg("Creating operation...")

	return db.WithTransaction(ctx, func(tx *sql.Tx) error {
		query := `
		INSERT INTO privilege_history (privilege_id, ticket_uid, datetime, balance_diff, operation_type) 
		VALUES ($1, $2, $3, $4, $5) 
		ON CONFLICT (privilege_id, ticket_uid, datetime, balance_diff, operation_type) DO NOTHING
		`

		stmt, err := tx.PrepareContext(ctx, query)
		if err != nil {
			return err
		}
		defer stmt.Close()

		res, err := stmt.ExecContext(ctx,
			operation.PrivilegeID, operation.TicketUid, operation.Date, operation.BalanceDiff, operation.OperationType)
		if err != nil {
			return err
		}

		rowsAffected, err := res.RowsAffected()
		if err != nil {
			return err
		}
		if rowsAffected == 0 {
			return errs.ErrOperationAlreadyExists
		}

		return nil
	})
}

func (db *Postgres) WithTransaction(ctx context.Context, fn TransactionFunc) error {
	tx, err := db.Conn.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelSerializable,
	})
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				db.Logger.Error().Msgf("Failed to rollback transaction: %v", rbErr)
				err = fmt.Errorf("%v & %v", err, rbErr)
			}
		}
	}()

	err = fn(tx)
	if err != nil {
		return err
	}

	return tx.Commit()
}
