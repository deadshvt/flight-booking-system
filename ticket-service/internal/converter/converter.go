package converter

import (
	"github.com/deadshvt/flight-booking-system/ticket-service/internal/entity"
	pb "github.com/deadshvt/flight-booking-system/ticket-service/proto"
)

func TicketFromEntityToProto(ticket *entity.Ticket) *pb.Ticket {
	return &pb.Ticket{
		ID:           ticket.ID,
		Username:     ticket.Username,
		TicketUid:    ticket.TicketUid,
		FlightNumber: ticket.FlightNumber,
		Price:        ticket.Price,
		Status:       ticket.Status,
	}
}

func TicketFromProtoToEntity(ticket *pb.Ticket) *entity.Ticket {
	return &entity.Ticket{
		ID:           ticket.ID,
		Username:     ticket.Username,
		TicketUid:    ticket.TicketUid,
		FlightNumber: ticket.FlightNumber,
		Price:        ticket.Price,
		Status:       ticket.Status,
	}
}
