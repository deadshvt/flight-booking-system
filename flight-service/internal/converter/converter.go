package converter

import (
	"fmt"

	"github.com/deadshvt/flight-booking-system/flight-service/internal/entity"
	pb "github.com/deadshvt/flight-booking-system/flight-service/proto"
)

func FlightFromEntityToProto(flight *entity.FlightWithAirports) *pb.FlightWithAirports {
	return &pb.FlightWithAirports{
		FlightNumber: flight.Flight.FlightNumber,
		FromAirport:  fmt.Sprintf("%s %s", flight.FromAirport.City, flight.FromAirport.Name),
		ToAirport:    fmt.Sprintf("%s %s", flight.ToAirport.City, flight.ToAirport.Name),
		Date:         flight.Flight.Date,
		Price:        flight.Flight.Price,
	}
}
