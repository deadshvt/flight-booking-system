package grpc

import (
	"context"
	"time"

	bonuspb "github.com/deadshvt/flight-booking-system/bonus-service/proto"
	flightpb "github.com/deadshvt/flight-booking-system/flight-service/proto"
	"github.com/deadshvt/flight-booking-system/gateway-service/internal/usecase"
	gatewaypb "github.com/deadshvt/flight-booking-system/gateway-service/proto"
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

	protoFlights, err := s.Usecase.GetFlightsWithAirports(ctx, req)
	if err != nil {
		s.Logger.Err(err).Msg("Failed to get flights with airports")
		return nil, err
	}

	logger.LogWithParams(s.Logger, "Got flights with airports", struct {
		Count int32
	}{protoFlights.GetTotalElements()})

	return protoFlights, nil
}

func (s *GatewayServer) GetTicketsWithAirports(ctx context.Context,
	req *gatewaypb.GetTicketsWithAirportsRequest) (*gatewaypb.GetTicketsWithAirportsResponse, error) {
	username, err := GetUsernameFromCtx(ctx)
	if err != nil {
		s.Logger.Err(err).Msg("Failed to get username from context")
		return nil, err
	}

	req.Username = username

	logger.LogWithParams(s.Logger, "Getting tickets with airports...", struct {
		Username string
	}{username})

	protoTickets, err := s.Usecase.GetTicketsWithAirports(ctx, req)
	if err != nil {
		s.Logger.Err(err).Msg("Failed to get tickets with airports")
		return nil, err
	}

	tickets := protoTickets.GetTickets()

	logger.LogWithParams(s.Logger, "Got tickets with airports", struct {
		Username string
		Count    int
	}{username, len(tickets)})

	return &gatewaypb.GetTicketsWithAirportsResponse{
		Tickets: tickets,
	}, nil
}

func (s *GatewayServer) GetTicketWithAirports(ctx context.Context,
	req *gatewaypb.GetTicketWithAirportsRequest) (*gatewaypb.GetTicketWithAirportsResponse, error) {
	username, err := GetUsernameFromCtx(ctx)
	if err != nil {
		s.Logger.Err(err).Msg("Failed to get username from context")
		return nil, err
	}

	req.Username = username

	ticketUid := req.GetTicketUid()

	logger.LogWithParams(s.Logger, "Getting ticket with airports...", struct {
		Username  string
		TicketUid string
	}{username, ticketUid})

	protoTicket, err := s.Usecase.GetTicketWithAirports(ctx, req)
	if err != nil {
		s.Logger.Err(err).Msg("Failed to get ticket with airports")
		return nil, err
	}

	ticket := protoTicket.GetTicket()

	logger.LogWithParams(s.Logger, "Got ticket with airports", struct {
		Username     string
		TicketUid    string
		FlightNumber string
		FromAirport  string
		ToAirport    string
		Date         time.Time
		Price        int32
		Status       string
	}{username, ticketUid, ticket.GetFlightNumber(),
		ticket.GetFromAirport(), ticket.GetToAirport(), ticket.GetDate().AsTime(),
		ticket.GetPrice(), ticket.GetStatus()})

	return protoTicket, nil
}

func (s *GatewayServer) PurchaseTicket(ctx context.Context,
	req *gatewaypb.PurchaseTicketRequest) (*gatewaypb.PurchaseTicketResponse, error) {
	username, err := GetUsernameFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	req.Username = username

	flightNumber := req.GetFlightNumber()
	price := req.GetPrice()
	paidFromBalance := req.GetPaidFromBalance()

	logger.LogWithParams(s.Logger, "Purchasing ticket...", struct {
		Username        string
		FlightNumber    string
		Price           int32
		PaidFromBalance bool
	}{username, flightNumber, price, paidFromBalance})

	protoTicket, err := s.Usecase.PurchaseTicket(ctx, req)
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
		Date          time.Time
		Price         int32
		PaidByMoney   int32
		PaidByBonuses int32
		Status        string
	}{username, protoTicket.GetTicketUid(), flightNumber, protoTicket.GetFromAirport(),
		protoTicket.GetToAirport(), protoTicket.GetDate().AsTime(), price,
		protoTicket.GetPaidByMoney(), protoTicket.GetPaidByBonuses(), protoTicket.GetStatus()})

	return protoTicket, nil
}

func (s *GatewayServer) ReturnTicket(ctx context.Context,
	req *gatewaypb.ReturnTicketRequest) (*gatewaypb.ReturnTicketResponse, error) {
	username, err := GetUsernameFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	req.Username = username

	ticketUid := req.GetTicketUid()

	logger.LogWithParams(s.Logger, "Returning ticket...", struct {
		Username  string
		TicketUid string
	}{username, ticketUid})

	_, err = s.Usecase.ReturnTicket(ctx, req)
	if err != nil {
		s.Logger.Err(err).Msg("Failed to return ticket")
		return nil, err
	}

	logger.LogWithParams(s.Logger, "Returned ticket", struct {
		Username  string
		TicketUid string
	}{username, ticketUid})

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
	}{username})

	protoPrivilege, err := s.Usecase.GetPrivilegeWithHistory(ctx, req)
	if err != nil {
		s.Logger.Err(err).Msg("Failed to get privilege with history")
		return nil, err
	}

	privilege := protoPrivilege.GetPrivilege()

	logger.LogWithParams(s.Logger, "Got privilege with history", struct {
		Username string
		Balance  int32
		Status   string
	}{username, privilege.GetBalance(), privilege.GetStatus()})

	return protoPrivilege, nil
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
	}{username})

	protoUser, err := s.Usecase.GetMe(ctx, req)
	if err != nil {
		s.Logger.Err(err).Msg("Failed to get user info")
		return nil, err
	}

	tickets := protoUser.GetTickets()
	privilege := protoUser.GetPrivilege()

	logger.LogWithParams(s.Logger, "Got user info", struct {
		Username string
		Count    int
		Balance  int32
		Status   string
	}{username, len(tickets), privilege.GetBalance(), privilege.GetStatus()})

	return protoUser, nil
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
