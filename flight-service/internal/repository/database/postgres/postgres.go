package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/deadshvt/flight-booking-system/config"
	"github.com/deadshvt/flight-booking-system/flight-service/internal/entity"
	"github.com/deadshvt/flight-booking-system/flight-service/pkg/errs"

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

func (db *Postgres) GetFlightsWithAirports(ctx context.Context,
	limit int32, offset int32) ([]*entity.FlightWithAirports, error) {
	db.Logger.Info().Msg("Getting flights with airports...")

	var flightsWithAirports []*entity.FlightWithAirports

	err := db.WithTransaction(ctx, func(tx *sql.Tx) error {
		query := `
		SELECT
			f.id,
			f.flight_number,
			a_from.city || ' ' || a_from.name AS from_airport,
			a_to.city || ' ' || a_to.name AS to_airport,
			f.datetime,
			f.price
		FROM
			flight f
		JOIN
			airport a_from ON f.from_airport_id = a_from.id
		JOIN
			airport a_to ON f.to_airport_id = a_to.id
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
			return errs.ErrFlightsNotFound
		}
		if err != nil {
			return err
		}
		defer rows.Close()

		for rows.Next() {
			var flightWithAirports entity.FlightWithAirports
			err = rows.Scan(
				&flightWithAirports.ID,
				&flightWithAirports.FlightNumber,
				&flightWithAirports.FromAirport,
				&flightWithAirports.ToAirport,
				&flightWithAirports.Date,
				&flightWithAirports.Price,
			)
			if err != nil {
				return err
			}

			flightsWithAirports = append(flightsWithAirports, &flightWithAirports)
		}

		return rows.Err()
	})

	if err != nil {
		return nil, err
	}

	return flightsWithAirports, nil
}

//nolint:dupl
func (db *Postgres) GetFlightWithAirports(ctx context.Context,
	flightNumber string) (*entity.FlightWithAirports, error) {
	db.Logger.Info().Msg("Getting flight with airports...")

	var flightWithAirports *entity.FlightWithAirports

	err := db.WithTransaction(ctx, func(tx *sql.Tx) error {
		query := `
		SELECT
			f.id,
			f.flight_number,
			a_from.city || ' ' || a_from.name AS from_airport,
			a_to.city || ' ' || a_to.name AS to_airport,
			f.datetime,
			f.price
		FROM
			flight f
		JOIN
			airport a_from ON f.from_airport_id = a_from.id
		JOIN
			airport a_to ON f.to_airport_id = a_to.id
		WHERE f.flight_number = $1
		`

		stmt, err := tx.PrepareContext(ctx, query)
		if err != nil {
			return err
		}
		defer stmt.Close()

		err = stmt.QueryRowContext(ctx, flightNumber).
			Scan(&flightWithAirports.ID, &flightWithAirports.FlightNumber,
				&flightWithAirports.FromAirport, &flightWithAirports.ToAirport,
				&flightWithAirports.Date, &flightWithAirports.Price)
		if errors.Is(err, sql.ErrNoRows) {
			return errs.ErrFlightNotFound
		}
		return err
	})

	if err != nil {
		return nil, err
	}

	return flightWithAirports, nil
}

//nolint:dupl
func (db *Postgres) GetFlight(ctx context.Context, flightNumber string) (*entity.Flight, error) {
	db.Logger.Info().Msg("Getting flight...")

	var flight *entity.Flight

	err := db.WithTransaction(ctx, func(tx *sql.Tx) error {
		query := `
		SELECT
			f.id,
			f.flight_number,
			f.from_airport_id,
			f.to_airport_id,
			f.datetime,
			f.price
		FROM
			flight f
		WHERE f.flight_number = $1
		`

		stmt, err := tx.PrepareContext(ctx, query)
		if err != nil {
			return err
		}
		defer stmt.Close()

		err = stmt.QueryRowContext(ctx, flightNumber).
			Scan(&flight.ID, &flight.FlightNumber, &flight.FromAirportID, &flight.ToAirportID, &flight.Date, &flight.Price)
		if errors.Is(err, sql.ErrNoRows) {
			return errs.ErrFlightNotFound
		}
		return err
	})

	if err != nil {
		return nil, err
	}

	return flight, nil
}

func (db *Postgres) CreateFlight(ctx context.Context, flight *entity.Flight) error {
	db.Logger.Info().Msg("Creating flight...")

	return db.WithTransaction(ctx, func(tx *sql.Tx) error {
		query := `
		INSERT INTO flight (flight_number, from_airport_id, to_airport_id, datetime, price) 
		VALUES ($1, $2, $3, $4, $5)
		ON CONFLICT (flight_number) DO NOTHING
		`

		stmt, err := tx.PrepareContext(ctx, query)
		if err != nil {
			return err
		}
		defer stmt.Close()

		res, err := stmt.ExecContext(ctx,
			flight.FlightNumber, flight.FromAirportID, flight.ToAirportID, flight.Date, flight.Price)
		if err != nil {
			return err
		}

		rowsAffected, err := res.RowsAffected()
		if err != nil {
			return err
		}
		if rowsAffected == 0 {
			return errs.ErrFlightAlreadyExists
		}

		return nil
	})
}

func (db *Postgres) GetAirport(ctx context.Context, airportID int32) (*entity.Airport, error) {
	db.Logger.Info().Msg("Getting airport...")

	var airport *entity.Airport

	err := db.WithTransaction(ctx, func(tx *sql.Tx) error {
		query := `
		SELECT
			a.id,
			a.name,
			a.city,
			a.country
		FROM
			airport a
		WHERE a.id = $1
		`

		stmt, err := tx.PrepareContext(ctx, query)
		if err != nil {
			return err
		}
		defer stmt.Close()

		err = stmt.QueryRowContext(ctx, airportID).
			Scan(&airport.ID, &airport.Name, &airport.City, &airport.Country)
		if errors.Is(err, sql.ErrNoRows) {
			return errs.ErrAirportNotFound
		}
		return err
	})

	if err != nil {
		return nil, err
	}

	return airport, nil
}

func (db *Postgres) CreateAirport(ctx context.Context, airport *entity.Airport) error {
	db.Logger.Info().Msg("Creating airport...")

	return db.WithTransaction(ctx, func(tx *sql.Tx) error {
		query := `
		INSERT INTO airport (name, city, country) 
		VALUES ($1, $2, $3) 
		ON CONFLICT (name, city, country) DO NOTHING
		`

		stmt, err := tx.PrepareContext(ctx, query)
		if err != nil {
			return err
		}
		defer stmt.Close()

		res, err := stmt.ExecContext(ctx,
			airport.Name, airport.City, airport.Country)
		if err != nil {
			return err
		}

		rowsAffected, err := res.RowsAffected()
		if err != nil {
			return err
		}
		if rowsAffected == 0 {
			return errs.ErrAirportAlreadyExists
		}

		return nil
	})
}

func (db *Postgres) CreateAirports(ctx context.Context, airports []*entity.Airport) error {
	db.Logger.Info().Msg("Creating airports...")

	airportsLen := len(airports)
	if airportsLen == 0 {
		return nil
	}

	return db.WithTransaction(ctx, func(tx *sql.Tx) error {
		valueStrings := make([]string, 0, airportsLen)
		valueArgs := make([]interface{}, 0, airportsLen*3)
		for i, airport := range airports {
			valueStrings = append(valueStrings, fmt.Sprintf("($%d, $%d, $%d)", i*3+1, i*3+2, i*3+3))
			valueArgs = append(valueArgs, airport.Name, airport.City, airport.Country)
		}

		query := `
		INSERT INTO airport (name, city, country) 
		VALUES
		`

		stmt, err := tx.PrepareContext(ctx, strings.Join([]string{
			query, strings.Join(valueStrings, ","), "ON CONFLICT (name, city, country) DO NOTHING"}, " "))
		if err != nil {
			return err
		}
		defer stmt.Close()

		res, err := stmt.ExecContext(ctx,
			valueArgs...)
		if err != nil {
			return err
		}

		rowsAffected, err := res.RowsAffected()
		if err != nil {
			return err
		}
		if rowsAffected != int64(airportsLen) {
			return errs.ErrAirportAlreadyExists
		}

		return nil
	})
}

func (db *Postgres) AirportTableExists(ctx context.Context) (bool, error) {
	db.Logger.Info().Msg("Checking if airport table exists...")

	var count int

	err := db.WithTransaction(ctx, func(tx *sql.Tx) error {
		query := `
		SELECT
			COUNT(*)
		FROM
			airport a
		`

		stmt, err := tx.PrepareContext(ctx, query)
		if err != nil {
			return err
		}
		defer stmt.Close()

		return stmt.QueryRowContext(ctx).Scan(&count)
	})

	if err != nil {
		return false, err
	}

	return count > 0, nil
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
