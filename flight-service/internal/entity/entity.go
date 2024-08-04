package entity

import (
	"fmt"
	"time"
)

type Flight struct {
	ID            int32
	FlightNumber  string
	FromAirportID int32
	ToAirportID   int32
	Date          time.Time
	Price         int32
}

type Airport struct {
	ID      int32
	Name    string
	City    string
	Country string
}

func (a *Airport) FullName() string {
	return fmt.Sprintf("%s %s", a.City, a.Name)
}

type FlightWithAirports struct {
	ID           int32
	FlightNumber string
	FromAirport  string
	ToAirport    string
	Date         time.Time
	Price        int32
}
