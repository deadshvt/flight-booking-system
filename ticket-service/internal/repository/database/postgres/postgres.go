package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"

	"github.com/deadshvt/flight-booking-system/config"
	"github.com/deadshvt/flight-booking-system/ticket-service/internal/entity"
	"github.com/deadshvt/flight-booking-system/ticket-service/pkg/errs"

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

func (db *Postgres) GetTickets(ctx context.Context, limit int32, offset int32) ([]*entity.Ticket, error) {
	db.Logger.Info().Msg("Getting tickets...")

	var tickets []*entity.Ticket

	err := db.WithTransaction(ctx, func(tx *sql.Tx) error {
		query := `
		SELECT
			t.id,
			t.ticket_uid,
			t.flight_number,
			t.username,
			t.status,
			t.price
		FROM
			ticket t
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
			return errs.ErrTicketsNotFound
		}
		if err != nil {
			return err
		}
		defer rows.Close()

		for rows.Next() {
			var ticket entity.Ticket
			err = rows.Scan(
				&ticket.ID,
				&ticket.TicketUid,
				&ticket.FlightNumber,
				&ticket.Username,
				&ticket.Status,
				&ticket.Price,
			)
			if err != nil {
				return err
			}

			tickets = append(tickets, &ticket)
		}

		return rows.Err()
	})

	if err != nil {
		return nil, err
	}

	return tickets, nil
}

func (db *Postgres) GetTicketsByUsername(ctx context.Context, username string) ([]*entity.Ticket, error) {
	db.Logger.Info().Msg("Getting tickets by username...")

	var tickets []*entity.Ticket

	err := db.WithTransaction(ctx, func(tx *sql.Tx) error {
		query := `
		SELECT
			t.id,
			t.ticket_uid,
			t.flight_number,
			t.username,
			t.status,
			t.price
		FROM
			ticket t
		WHERE
			t.username = $1
		`

		stmt, err := tx.PrepareContext(ctx, query)
		if err != nil {
			return err
		}
		defer stmt.Close()

		rows, err := stmt.QueryContext(ctx, username)
		if errors.Is(err, sql.ErrNoRows) {
			return errs.ErrTicketsNotFound
		}
		if err != nil {
			return err
		}
		defer rows.Close()

		for rows.Next() {
			var ticket entity.Ticket
			err = rows.Scan(
				&ticket.ID,
				&ticket.TicketUid,
				&ticket.FlightNumber,
				&ticket.Username,
				&ticket.Status,
				&ticket.Price,
			)
			if err != nil {
				return err
			}

			tickets = append(tickets, &ticket)
		}

		return rows.Err()
	})

	if err != nil {
		return nil, err
	}

	return tickets, nil
}

func (db *Postgres) GetTicket(ctx context.Context, ticketUid string) (*entity.Ticket, error) {
	db.Logger.Info().Msg("Getting ticket...")

	var ticket *entity.Ticket

	err := db.WithTransaction(ctx, func(tx *sql.Tx) error {
		query := `
		SELECT
			t.id,
			t.ticket_uid,
			t.flight_number,
			t.username,
			t.status,
			t.price
		FROM
			ticket t
		WHERE
			t.ticket_uid = $1
		`

		stmt, err := tx.PrepareContext(ctx, query)
		if err != nil {
			return err
		}
		defer stmt.Close()

		err = stmt.QueryRowContext(ctx, ticketUid).
			Scan(&ticket.ID, &ticket.TicketUid, &ticket.FlightNumber, &ticket.Username, &ticket.Status, &ticket.Price)
		if errors.Is(err, sql.ErrNoRows) {
			return errs.ErrTicketNotFound
		}
		return err
	})

	if err != nil {
		return nil, err
	}

	return ticket, nil
}

func (db *Postgres) CreateTicket(ctx context.Context, ticket *entity.Ticket) error {
	db.Logger.Info().Msg("Creating ticket...")

	return db.WithTransaction(ctx, func(tx *sql.Tx) error {
		query := `
		INSERT INTO ticket (ticket_uid, flight_number, username, status, price) 
		VALUES ($1, $2, $3, $4, $5)
		ON CONFLICT (ticket_uid) DO NOTHING
		`

		stmt, err := tx.PrepareContext(ctx, query)
		if err != nil {
			return err
		}
		defer stmt.Close()

		res, err := stmt.ExecContext(ctx,
			ticket.TicketUid, ticket.FlightNumber, ticket.Username, ticket.Status, ticket.Price)
		if err != nil {
			return err
		}

		rowsAffected, err := res.RowsAffected()
		if err != nil {
			return err
		}
		if rowsAffected == 0 {
			return errs.ErrTicketAlreadyExists
		}

		return nil
	})
}

func (db *Postgres) UpdateTicket(ctx context.Context, ticket *entity.Ticket) error {
	db.Logger.Info().Msg("Updating ticket...")

	return db.WithTransaction(ctx, func(tx *sql.Tx) error {
		query := `
		UPDATE ticket
		SET flight_number = $2, username = $3, status = $4, price = $5
		WHERE ticket_uid = $1
		`

		stmt, err := tx.PrepareContext(ctx, query)
		if err != nil {
			return err
		}
		defer stmt.Close()

		res, err := stmt.ExecContext(ctx,
			ticket.TicketUid, ticket.FlightNumber, ticket.Username, ticket.Status, ticket.Price)
		if err != nil {
			return err
		}

		rowsAffected, err := res.RowsAffected()
		if err != nil {
			return err
		}
		if rowsAffected == 0 {
			return errs.ErrTicketNotFound
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
