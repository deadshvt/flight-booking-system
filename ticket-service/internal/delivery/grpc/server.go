package grpc

import (
	"context"
	"errors"

	"github.com/deadshvt/flight-booking-system/logger"
	"github.com/deadshvt/flight-booking-system/ticket-service/internal/converter"
	"github.com/deadshvt/flight-booking-system/ticket-service/internal/repository"
	"github.com/deadshvt/flight-booking-system/ticket-service/internal/usecase"
	"github.com/deadshvt/flight-booking-system/ticket-service/pkg/errs"
	pb "github.com/deadshvt/flight-booking-system/ticket-service/proto"

	"github.com/rs/zerolog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TicketServiceServer struct {
	pb.UnimplementedTicketServiceServer
	Repo    *repository.TicketRepository
	Usecase *usecase.TicketUsecase
	Logger  zerolog.Logger
}

func NewTicketServiceServer(repo *repository.TicketRepository,
	usecase *usecase.TicketUsecase, logger zerolog.Logger) *TicketServiceServer {
	return &TicketServiceServer{
		Repo:    repo,
		Usecase: usecase,
		Logger:  logger,
	}
}

func (s *TicketServiceServer) GetTickets(ctx context.Context,
	req *pb.GetTicketsRequest) (*pb.GetTicketsResponse, error) {
	username := req.GetUsername()

	logger.LogWithParams(s.Logger, "Getting tickets...", struct {
		Username string
	}{username})

	tickets, err := s.Repo.GetTickets(ctx, username)
	if err != nil {
		s.Logger.Err(err).Msg("Failed to get tickets")
		return nil, s.HandleError(err)
	}

	protoTickets := make([]*pb.Ticket, len(tickets))
	for i, ticket := range tickets {
		protoTickets[i] = converter.TicketFromEntityToProto(ticket)
	}

	logger.LogWithParams(s.Logger, "Got tickets", struct {
		Username string
		Count    int
	}{username, len(tickets)})

	return &pb.GetTicketsResponse{Tickets: protoTickets}, nil
}

func (s *TicketServiceServer) GetTicket(ctx context.Context,
	req *pb.GetTicketRequest) (*pb.GetTicketResponse, error) {
	username := req.GetUsername()
	ticketUid := req.GetTicketUid()

	logger.LogWithParams(s.Logger, "Getting ticket...", struct {
		Username  string
		TicketUid string
	}{username, ticketUid})

	ticket, err := s.Usecase.GetTicket(ctx, username, ticketUid)
	if err != nil {
		s.Logger.Err(err).Msg("Failed to get ticket")
		return nil, s.HandleError(err)
	}

	logger.LogWithParams(s.Logger, "Got ticket", struct {
		Username     string
		TicketUid    string
		FlightNumber string
		Price        int32
		Status       string
	}{ticket.Username, ticket.TicketUid, ticket.FlightNumber,
		ticket.Price, ticket.Status})

	return &pb.GetTicketResponse{Ticket: converter.TicketFromEntityToProto(ticket)}, nil
}

func (s *TicketServiceServer) PurchaseTicket(ctx context.Context,
	req *pb.PurchaseTicketRequest) (*pb.PurchaseTicketResponse, error) {
	username := req.GetUsername()
	flightNumber := req.GetFlightNumber()
	price := req.GetPrice()

	logger.LogWithParams(s.Logger, "Purchasing ticket...", struct {
		Username     string
		FlightNumber string
	}{username, flightNumber})

	ticketUid, err := s.Usecase.PurchaseTicket(ctx, username, flightNumber, price)
	if err != nil {
		s.Logger.Err(err).Msg("Failed to purchase ticket")
		return nil, s.HandleError(err)
	}

	logger.LogWithParams(s.Logger, "Purchased ticket", struct {
		Username     string
		FlightNumber string
		TicketUid    string
	}{username, flightNumber, ticketUid})

	return &pb.PurchaseTicketResponse{}, nil
}

func (s *TicketServiceServer) ReturnTicket(ctx context.Context,
	req *pb.ReturnTicketRequest) (*pb.ReturnTicketResponse, error) {
	username := req.GetUsername()
	ticketUid := req.GetTicketUid()

	logger.LogWithParams(s.Logger, "Returning ticket...", struct {
		Username  string
		TicketUid string
	}{username, ticketUid})

	err := s.Usecase.ReturnTicket(ctx, username, ticketUid)
	if err != nil {
		s.Logger.Err(err).Msg("Failed to return ticket")
		return nil, s.HandleError(err)
	}

	logger.LogWithParams(s.Logger, "Returned ticket", struct {
		Username  string
		TicketUid string
	}{username, ticketUid})

	return &pb.ReturnTicketResponse{}, nil
}

func (s *TicketServiceServer) HandleError(err error) error {
	switch {
	case errors.Is(err, errs.ErrTicketsNotFound):
		return status.Error(codes.NotFound, err.Error())
	case errors.Is(err, errs.ErrTicketNotFound):
		return status.Error(codes.NotFound, err.Error())
	default:
		return status.Error(codes.Internal, err.Error())
	}
}
