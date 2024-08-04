package converter

import (
	"github.com/deadshvt/flight-booking-system/flight-service/internal/entity"
	pb "github.com/deadshvt/flight-booking-system/flight-service/proto"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func FlightFromEntityToProto(flight *entity.FlightWithAirports) *pb.FlightWithAirports {
	return &pb.FlightWithAirports{
		FlightNumber: flight.FlightNumber,
		FromAirport:  flight.FromAirport,
		ToAirport:    flight.ToAirport,
		Date:         timestamppb.New(flight.Date),
		Price:        flight.Price,
	}
}
