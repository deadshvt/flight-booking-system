package entity

type Flight struct {
	ID            int32
	FlightNumber  string
	FromAirportID int32
	ToAirportID   int32
	Date          string
	Price         int32
}

type Airport struct {
	ID      int32
	Name    string
	City    string
	Country string
}

type FlightWithAirports struct {
	Flight      Flight
	FromAirport Airport
	ToAirport   Airport
}
