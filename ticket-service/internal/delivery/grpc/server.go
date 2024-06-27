package grpc

import (
	"context"

	"github.com/deadshvt/flight-booking-system/logger"
	"github.com/deadshvt/flight-booking-system/ticket-service/internal/converter"
	"github.com/deadshvt/flight-booking-system/ticket-service/internal/repository"
	"github.com/deadshvt/flight-booking-system/ticket-service/internal/usecase"
	pb "github.com/deadshvt/flight-booking-system/ticket-service/proto"

	"github.com/rs/zerolog"
)

type TicketServiceServer struct {
	pb.UnimplementedTicketServiceServer
	Repo    repository.TicketRepository
	Usecase *usecase.TicketUsecase
	Logger  zerolog.Logger
}

func NewTicketServiceServer(repo repository.TicketRepository,
	usecase *usecase.TicketUsecase, logger zerolog.Logger) *TicketServiceServer {
	return &TicketServiceServer{
		Repo:    repo,
		Usecase: usecase,
		Logger:  logger,
	}
}

func (s *TicketServiceServer) GetTickets(ctx context.Context,
	req *pb.GetTicketsRequest) (*pb.GetTicketsResponse, error) {
	logger.LogWithParams(s.Logger, "Getting tickets...", struct {
		Username string
	}{req.Username})

	tickets, err := s.Repo.GetTickets(ctx, req.Username)
	if err != nil {
		s.Logger.Err(err).Msg("Failed to get ticket")
		return nil, err
	}

	protoTickets := make([]*pb.Ticket, len(tickets))
	for i, ticket := range tickets {
		protoTickets[i] = converter.TicketFromEntityToProto(ticket)
	}

	logger.LogWithParams(s.Logger, "Got tickets", struct {
		Username string
		Count    int
	}{req.Username, len(tickets)})

	return &pb.GetTicketsResponse{Tickets: protoTickets}, nil
}

func (s *TicketServiceServer) GetTicket(ctx context.Context,
	req *pb.GetTicketRequest) (*pb.GetTicketResponse, error) {
	logger.LogWithParams(s.Logger, "Getting ticket...", struct {
		Username  string
		TicketUid string
	}{req.Username, req.TicketUid})

	ticket, err := s.Repo.GetTicket(ctx, req.TicketUid)
	if err != nil {
		s.Logger.Err(err).Msg("Failed to get ticket")
		return nil, err
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
	logger.LogWithParams(s.Logger, "Purchasing ticket...", struct {
		Username     string
		FlightNumber string
	}{req.Username, req.FlightNumber})

	protoTicket, err := s.Usecase.PurchaseTicket(ctx, req)
	if err != nil {
		s.Logger.Err(err).Msg("Failed to purchase ticket")
		return nil, err
	}

	logger.LogWithParams(s.Logger, "Purchased ticket", struct {
		Username     string
		TicketUid    string
		FlightNumber string
	}{req.Username, protoTicket.TicketUid, req.FlightNumber})

	return protoTicket, nil
}

func (s *TicketServiceServer) ReturnTicket(ctx context.Context,
	req *pb.ReturnTicketRequest) (*pb.ReturnTicketResponse, error) {
	logger.LogWithParams(s.Logger, "Returning ticket...", struct {
		Username  string
		TicketUid string
	}{req.Username, req.TicketUid})

	protoTicket, err := s.Usecase.ReturnTicket(ctx, req)
	if err != nil {
		s.Logger.Err(err).Msg("Failed to return ticket")
		return nil, err
	}

	logger.LogWithParams(s.Logger, "Returned ticket", struct {
		Username  string
		TicketUid string
	}{req.Username, req.TicketUid})

	return protoTicket, nil
}
