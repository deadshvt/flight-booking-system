package entity

type Ticket struct {
	ID           int32
	TicketUid    string
	Username     string
	FlightNumber string
	Price        int32
	Status       string
}
