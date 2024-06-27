package grpc

import (
	"context"

	bonuspb "github.com/deadshvt/flight-booking-system/bonus-service/proto"
	flightpb "github.com/deadshvt/flight-booking-system/flight-service/proto"
	"github.com/deadshvt/flight-booking-system/gateway/internal/usecase"
	gatewaypb "github.com/deadshvt/flight-booking-system/gateway/proto"
	"github.com/deadshvt/flight-booking-system/logger"

	"github.com/rs/zerolog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type GatewayServer struct {
	gatewaypb.UnimplementedGatewayServer
	Usecase *usecase.GatewayUsecase
	Logger  zerolog.Logger
}

func NewGatewayServer(usecase *usecase.GatewayUsecase, logger zerolog.Logger) *GatewayServer {
	return &GatewayServer{
		Usecase: usecase,
		Logger:  logger,
	}
}

func (s *GatewayServer) GetFlightsWithAirports(ctx context.Context,
	req *flightpb.GetFlightsWithAirportsRequest) (*flightpb.GetFlightsWithAirportsResponse, error) {
	s.Logger.Info().Msg("Getting flights with airports...")

	flights, err := s.Usecase.GetFlightsWithAirports(ctx, req)
	if err != nil {
		s.Logger.Err(err).Msg("Failed to get flights with airports")
		return nil, err
	}

	logger.LogWithParams(s.Logger, "Got flights with airports", struct {
		Count int32
	}{flights.TotalElements})

	return flights, nil
}

func (s *GatewayServer) GetTicketsWithAirports(ctx context.Context,
	req *gatewaypb.GetTicketsWithAirportsRequest) (*gatewaypb.GetTicketsWithAirportsResponse, error) {
	username, err := GetUsernameFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	req.Username = username

	logger.LogWithParams(s.Logger, "Getting tickets with airports...", struct {
		Username string
	}{req.Username})

	tickets, err := s.Usecase.GetTicketsWithAirports(ctx, req)
	if err != nil {
		s.Logger.Err(err).Msg("Failed to get tickets with airports")
		return nil, err
	}

	logger.LogWithParams(s.Logger, "Got tickets with airports", struct {
		Username string
		Count    int
	}{req.Username, len(tickets.Tickets)})

	return tickets, nil
}

func (s *GatewayServer) GetTicketWithAirports(ctx context.Context,
	req *gatewaypb.GetTicketWithAirportsRequest) (*gatewaypb.GetTicketWithAirportsResponse, error) {
	username, err := GetUsernameFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	req.Username = username

	logger.LogWithParams(s.Logger, "Getting ticket with airports...", struct {
		Username  string
		TicketUid string
	}{req.Username, req.TicketUid})

	protoTicket, err := s.Usecase.GetTicketWithAirports(ctx, req)
	if err != nil {
		s.Logger.Err(err).Msg("Failed to get ticket with airports")
		return nil, err
	}

	ticket := protoTicket.Ticket

	logger.LogWithParams(s.Logger, "Got ticket with airports", struct {
		Username     string
		TicketUid    string
		FlightNumber string
		FromAirport  string
		ToAirport    string
		Date         string
		Price        int32
		Status       string
	}{req.Username, req.TicketUid, ticket.FlightNumber,
		ticket.FromAirport, ticket.ToAirport, ticket.Date,
		ticket.Price, ticket.Status})

	return protoTicket, nil
}

func (s *GatewayServer) PurchaseTicket(ctx context.Context,
	req *gatewaypb.PurchaseTicketRequest) (*gatewaypb.PurchaseTicketResponse, error) {
	username, err := GetUsernameFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	req.Username = username

	logger.LogWithParams(s.Logger, "Purchasing ticket...", struct {
		Username        string
		FlightNumber    string
		Price           int32
		PaidFromBalance bool
	}{req.Username, req.FlightNumber, req.Price, req.PaidFromBalance})

	ticket, err := s.Usecase.PurchaseTicket(ctx, req)
	if err != nil {
		s.Logger.Err(err).Msg("Failed to purchase ticket")
		return nil, err
	}

	logger.LogWithParams(s.Logger, "Purchased ticket", struct {
		Username      string
		TicketUid     string
		FlightNumber  string
		FromAirport   string
		ToAirport     string
		Date          string
		Price         int32
		PaidByMoney   int32
		PaidByBonuses int32
		Status        string
		Balance       int32
	}{req.Username, ticket.TicketUid, ticket.FlightNumber,
		ticket.FromAirport, ticket.ToAirport, ticket.Date,
		ticket.Price, ticket.PaidByMoney, ticket.PaidByBonuses,
		ticket.Status, ticket.Privilege.Balance})

	return ticket, nil
}

func (s *GatewayServer) ReturnTicket(ctx context.Context,
	req *gatewaypb.ReturnTicketRequest) (*gatewaypb.ReturnTicketResponse, error) {
	username, err := GetUsernameFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	req.Username = username

	logger.LogWithParams(s.Logger, "Returning ticket...", struct {
		Username  string
		TicketUid string
	}{req.Username, req.TicketUid})

	_, err = s.Usecase.ReturnTicket(ctx, req)
	if err != nil {
		s.Logger.Err(err).Msg("Failed to return ticket")
		return nil, err
	}

	logger.LogWithParams(s.Logger, "Returned ticket", struct {
		Username  string
		TicketUid string
	}{req.Username, req.TicketUid})

	return &gatewaypb.ReturnTicketResponse{}, nil
}

func (s *GatewayServer) GetPrivilegeWithHistory(ctx context.Context,
	req *bonuspb.GetPrivilegeWithHistoryRequest) (*bonuspb.GetPrivilegeWithHistoryResponse, error) {
	username, err := GetUsernameFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	req.Username = username

	logger.LogWithParams(s.Logger, "Getting privilege with history...", struct {
		Username string
	}{req.Username})

	privilege, err := s.Usecase.GetPrivilegeWithHistory(ctx, req)
	if err != nil {
		s.Logger.Err(err).Msg("Failed to get privilege with history")
		return nil, err
	}

	logger.LogWithParams(s.Logger, "Got privilege with history", struct {
		Username string
		Balance  int32
		Status   string
	}{req.Username, privilege.Privilege.Balance, privilege.Privilege.Status})

	return privilege, nil
}

func (s *GatewayServer) GetMe(ctx context.Context,
	req *gatewaypb.GetMeRequest) (*gatewaypb.GetMeResponse, error) {
	username, err := GetUsernameFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	req.Username = username

	logger.LogWithParams(s.Logger, "Getting user info...", struct {
		Username string
	}{req.Username})

	user, err := s.Usecase.GetMe(ctx, req)
	if err != nil {
		s.Logger.Err(err).Msg("Failed to get user info")
		return nil, err
	}

	logger.LogWithParams(s.Logger, "Got user info", struct {
		Username string
		Count    int
		Balance  int32
		Status   string
	}{req.Username, len(user.Tickets), user.Privilege.Balance, user.Privilege.Status})

	return user, nil
}

func GetUsernameFromCtx(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", status.Errorf(codes.InvalidArgument, "missing metadata")
	}

	userNames := md.Get("x-user-name")
	if len(userNames) == 0 {
		return "", status.Errorf(codes.InvalidArgument, "missing X-User-Name")
	}

	if userNames[0] == "" {
		return "", status.Errorf(codes.InvalidArgument, "empty X-User-Name")
	}

	return userNames[0], nil
}
