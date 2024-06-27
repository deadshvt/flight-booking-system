package usecase

import (
	"context"
	"math"
	"strings"
	"time"

	bonusErrs "github.com/deadshvt/flight-booking-system/bonus-service/pkg/errs"
	bonuspb "github.com/deadshvt/flight-booking-system/bonus-service/proto"
	flightErrs "github.com/deadshvt/flight-booking-system/flight-service/pkg/errs"
	flightpb "github.com/deadshvt/flight-booking-system/flight-service/proto"
	gatewaypb "github.com/deadshvt/flight-booking-system/gateway/proto"
	"github.com/deadshvt/flight-booking-system/logger"
	ticketErrs "github.com/deadshvt/flight-booking-system/ticket-service/pkg/errs"
	ticketpb "github.com/deadshvt/flight-booking-system/ticket-service/proto"

	"github.com/rs/zerolog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	StatusPurchase = "PAID"

	StatusBronze = "BRONZE"
	StatusSilver = "SILVER"
	StatusGold   = "GOLD"

	OperationFill  = "FILL_IN_BALANCE"
	OperationDebit = "DEBIT_THE_ACCOUNT"
)

type GatewayUsecase struct {
	TicketClient ticketpb.TicketServiceClient
	FlightClient flightpb.FlightServiceClient
	BonusClient  bonuspb.BonusServiceClient
	Logger       zerolog.Logger
}

func NewGatewayUsecase(ticketClient ticketpb.TicketServiceClient,
	flightClient flightpb.FlightServiceClient, bonusClient bonuspb.BonusServiceClient,
	logger zerolog.Logger) *GatewayUsecase {
	return &GatewayUsecase{
		TicketClient: ticketClient,
		FlightClient: flightClient,
		BonusClient:  bonusClient,
		Logger:       logger,
	}
}

func (u *GatewayUsecase) GetFlightsWithAirports(ctx context.Context,
	req *flightpb.GetFlightsWithAirportsRequest) (*flightpb.GetFlightsWithAirportsResponse, error) {
	u.Logger.Info().Msg("Getting flights with airports...")

	return u.FlightClient.GetFlightsWithAirports(ctx, req)
}

func (u *GatewayUsecase) GetTicketsWithAirports(ctx context.Context,
	req *gatewaypb.GetTicketsWithAirportsRequest) (*gatewaypb.GetTicketsWithAirportsResponse, error) {
	logger.LogWithParams(u.Logger, "Getting tickets with airports...", struct {
		Username string
	}{req.Username})

	tickets, err := u.TicketClient.GetTickets(ctx, &ticketpb.GetTicketsRequest{
		Username: req.Username,
	})
	if err != nil {
		u.Logger.Err(err).Msg("Failed to get tickets")
		return nil, status.Errorf(codes.Internal, "failed to get tickets: %v", err)
	}

	var protoTickets []*gatewaypb.TicketWithAirports
	for _, ticket := range tickets.Tickets {
		protoFlight, err := u.FlightClient.GetFlightWithAirports(ctx, &flightpb.GetFlightWithAirportsRequest{
			FlightNumber: ticket.FlightNumber,
		})
		if err != nil {
			u.Logger.Err(err).Msg("Failed to get flight with airports")
			return nil, status.Errorf(codes.Internal, "failed to get flight with airports: %v", err)
		}

		flight := protoFlight.Flight

		protoTickets = append(protoTickets, &gatewaypb.TicketWithAirports{
			TicketUid:    ticket.TicketUid,
			FlightNumber: ticket.FlightNumber,
			FromAirport:  flight.FromAirport,
			ToAirport:    flight.ToAirport,
			Date:         flight.Date,
			Price:        ticket.Price,
			Status:       ticket.Status,
		})
	}

	return &gatewaypb.GetTicketsWithAirportsResponse{
		Tickets: protoTickets,
	}, nil
}

func (u *GatewayUsecase) GetTicketWithAirports(ctx context.Context,
	req *gatewaypb.GetTicketWithAirportsRequest) (*gatewaypb.GetTicketWithAirportsResponse, error) {
	logger.LogWithParams(u.Logger, "Getting ticket with airports...", struct {
		Username  string
		TicketUid string
	}{req.Username, req.TicketUid})

	protoTicket, err := u.TicketClient.GetTicket(ctx, &ticketpb.GetTicketRequest{
		Username:  req.Username,
		TicketUid: req.TicketUid,
	})
	if err != nil && strings.Contains(err.Error(), ticketErrs.ErrTicketNotFound.Error()) {
		u.Logger.Err(err).Msg("Failed to get ticket")
		return nil, status.Errorf(codes.NotFound, ticketErrs.ErrTicketNotFound.Error())
	}
	if err != nil {
		u.Logger.Err(err).Msg("Failed to get ticket")
		return nil, status.Errorf(codes.Internal, "failed to get ticket: %v", err)
	}

	ticket := protoTicket.Ticket

	if ticket.Username != req.Username {
		u.Logger.Err(err).Msg("Wrong username")
		return nil, status.Errorf(codes.NotFound, ticketErrs.ErrTicketNotFound.Error())
	}

	protoFlight, err := u.FlightClient.GetFlightWithAirports(ctx, &flightpb.GetFlightWithAirportsRequest{
		FlightNumber: ticket.FlightNumber,
	})
	if err != nil {
		u.Logger.Err(err).Msg("Failed to get flight with airports")
		return nil, status.Errorf(codes.Internal, "failed to get flight with airports: %v", err)
	}

	flight := protoFlight.Flight

	return &gatewaypb.GetTicketWithAirportsResponse{
		Ticket: &gatewaypb.TicketWithAirports{
			TicketUid:    ticket.TicketUid,
			FlightNumber: ticket.FlightNumber,
			FromAirport:  flight.FromAirport,
			ToAirport:    flight.ToAirport,
			Date:         flight.Date,
			Price:        ticket.Price,
			Status:       ticket.Status,
		},
	}, nil
}

func (u *GatewayUsecase) PurchaseTicket(ctx context.Context,
	req *gatewaypb.PurchaseTicketRequest) (*gatewaypb.PurchaseTicketResponse, error) {
	logger.LogWithParams(u.Logger, "Purchasing ticket...", struct {
		Username     string
		FlightNumber string
	}{req.Username, req.FlightNumber})

	protoFlight, err := u.FlightClient.GetFlightWithAirports(ctx, &flightpb.GetFlightWithAirportsRequest{
		FlightNumber: req.FlightNumber,
	})
	if err != nil && strings.Contains(err.Error(), flightErrs.ErrFlightNotFound.Error()) {
		u.Logger.Err(err).Msg("Failed to get flight with airports")
		return nil, status.Errorf(codes.NotFound, "failed to get flight with airports: %v", err)
	}
	if err != nil {
		u.Logger.Err(err).Msg("Failed to get flight with airports")
		return nil, status.Errorf(codes.Internal, "failed to get flight with airports: %v", err)
	}

	flight := protoFlight.Flight

	protoTicket, err := u.TicketClient.PurchaseTicket(ctx, &ticketpb.PurchaseTicketRequest{
		Username:     req.Username,
		FlightNumber: req.FlightNumber,
		Price:        flight.Price,
	})
	if err != nil {
		u.Logger.Err(err).Msg("Failed to purchase ticket")
		return nil, status.Errorf(codes.Internal, "failed to purchase ticket: %v", err)
	}

	var privilege *bonuspb.Privilege

	protoPrivilege, err := u.BonusClient.GetPrivilege(ctx, &bonuspb.GetPrivilegeRequest{
		Username: req.Username,
	})
	if err != nil && !strings.Contains(err.Error(), bonusErrs.ErrPrivilegeNotFound.Error()) {
		u.Logger.Err(err).Msg("Failed to get privilege")
		return nil, status.Errorf(codes.Internal, "failed to get privilege: %v", err)
	}
	if err != nil && strings.Contains(err.Error(), bonusErrs.ErrPrivilegeNotFound.Error()) {
		privilege = &bonuspb.Privilege{
			Username: req.Username,
			Balance:  0,
			Status:   StatusBronze,
		}
		_, err = u.BonusClient.CreatePrivilege(ctx, &bonuspb.CreatePrivilegeRequest{
			Privilege: privilege,
		})
	} else {
		privilege = protoPrivilege.Privilege
	}

	operationType := OperationFill
	balanceDiff := int32(math.Floor(float64(req.Price) * 0.1))
	paidByMoney := req.Price
	paidByBonuses := int32(0)

	if req.PaidFromBalance {
		operationType = OperationDebit
		balanceDiff = -privilege.Balance
		paidByBonuses = privilege.Balance
		paidByMoney = req.Price - paidByMoney
	}

	privilege.Balance += balanceDiff
	UpdatePrivilegeStatus(privilege)

	_, err = u.BonusClient.UpdatePrivilege(ctx, &bonuspb.UpdatePrivilegeRequest{
		Privilege: privilege,
	})
	if err != nil {
		u.Logger.Err(err).Msg("Failed to update privilege")
		return nil, status.Errorf(codes.Internal, "failed to update privilege: %v", err)
	}

	_, err = u.BonusClient.CreateHistory(ctx, &bonuspb.CreateHistoryRequest{
		History: &bonuspb.History{
			PrivilegeID:   privilege.ID,
			TicketUid:     protoTicket.TicketUid,
			Date:          time.Now().Format(time.RFC3339),
			BalanceDiff:   balanceDiff,
			OperationType: operationType,
		},
	})
	if err != nil {
		u.Logger.Err(err).Msg("Failed to create history")
		return nil, status.Errorf(codes.Internal, "failed to create history: %v", err)
	}

	return &gatewaypb.PurchaseTicketResponse{
		TicketUid:     protoTicket.TicketUid,
		FlightNumber:  req.FlightNumber,
		FromAirport:   flight.FromAirport,
		ToAirport:     flight.ToAirport,
		Date:          flight.Date,
		Price:         flight.Price,
		PaidByMoney:   paidByMoney,
		PaidByBonuses: paidByBonuses,
		Status:        StatusPurchase,
		Privilege: &gatewaypb.PrivilegeShortInfo{
			Balance: privilege.Balance,
			Status:  privilege.Status,
		},
	}, nil
}

func (u *GatewayUsecase) GetPrivilegeWithHistory(ctx context.Context,
	req *bonuspb.GetPrivilegeWithHistoryRequest) (*bonuspb.GetPrivilegeWithHistoryResponse, error) {
	u.Logger.Info().Msg("Getting privilege with history...")

	return u.BonusClient.GetPrivilegeWithHistory(ctx, req)
}

func (u *GatewayUsecase) ReturnTicket(ctx context.Context,
	req *gatewaypb.ReturnTicketRequest) (*gatewaypb.ReturnTicketResponse, error) {
	u.Logger.Info().Msg("Returning ticket...")

	_, err := u.TicketClient.ReturnTicket(ctx, &ticketpb.ReturnTicketRequest{
		Username:  req.Username,
		TicketUid: req.TicketUid,
	})
	if err != nil {
		u.Logger.Err(err).Msg("Failed to return ticket")
		return nil, status.Errorf(codes.Internal, "failed to return ticket: %v", err)
	}

	protoPrivilegeWithHistory, err := u.BonusClient.GetPrivilegeWithHistory(ctx, &bonuspb.GetPrivilegeWithHistoryRequest{
		Username: req.Username,
	})
	if err != nil {
		u.Logger.Err(err).Msg("Failed to get privilege with history")
		return nil, status.Errorf(codes.Internal, "failed to get privilege with history: %v", err)
	}

	privilege := protoPrivilegeWithHistory.Privilege
	history := protoPrivilegeWithHistory.History
	var balanceDiff int32

	for _, h := range history {
		if h.TicketUid == req.TicketUid {
			privilege.Balance -= h.BalanceDiff
			balanceDiff = -h.BalanceDiff
			break
		}
	}

	operationType := OperationFill
	if balanceDiff < 0 {
		operationType = OperationDebit
	}

	UpdatePrivilegeStatus(privilege)

	_, err = u.BonusClient.UpdatePrivilege(ctx, &bonuspb.UpdatePrivilegeRequest{
		Privilege: privilege,
	})
	if err != nil {
		u.Logger.Err(err).Msg("Failed to update privilege")
		return nil, status.Errorf(codes.Internal, "failed to update privilege: %v", err)
	}

	_, err = u.BonusClient.CreateHistory(ctx, &bonuspb.CreateHistoryRequest{
		History: &bonuspb.History{
			PrivilegeID:   privilege.ID,
			TicketUid:     req.TicketUid,
			Date:          time.Now().Format(time.RFC3339),
			BalanceDiff:   balanceDiff,
			OperationType: operationType,
		},
	})
	if err != nil {
		u.Logger.Err(err).Msg("Failed to create history")
		return nil, status.Errorf(codes.Internal, "failed to create history: %v", err)
	}

	return &gatewaypb.ReturnTicketResponse{}, nil
}

func (u *GatewayUsecase) GetMe(ctx context.Context,
	req *gatewaypb.GetMeRequest) (*gatewaypb.GetMeResponse, error) {
	u.Logger.Info().Msg("Getting user info...")

	tickets, err := u.GetTicketsWithAirports(ctx, &gatewaypb.GetTicketsWithAirportsRequest{
		Username: req.Username,
	})
	if err != nil {
		u.Logger.Err(err).Msg("Failed to get tickets with airports")
		return nil, status.Errorf(codes.Internal, "failed to get tickets with airports: %v", err)
	}

	privilege, err := u.BonusClient.GetPrivilege(ctx, &bonuspb.GetPrivilegeRequest{
		Username: req.Username,
	})
	if err != nil {
		u.Logger.Err(err).Msg("Failed to get privilege with history")
		return nil, status.Errorf(codes.Internal, "failed to get privilege with history: %v", err)
	}

	privilegeShortInfo := &gatewaypb.PrivilegeShortInfo{
		Balance: privilege.Privilege.Balance,
		Status:  privilege.Privilege.Status,
	}

	return &gatewaypb.GetMeResponse{
		Tickets:   tickets.Tickets,
		Privilege: privilegeShortInfo,
	}, nil
}

func UpdatePrivilegeStatus(privilege *bonuspb.Privilege) {
	if privilege.Balance < 0 {
		privilege.Balance = 0
	} else if privilege.Balance < 1000 {
		privilege.Status = StatusBronze
	} else if privilege.Balance < 5000 {
		privilege.Status = StatusSilver
	} else {
		privilege.Status = StatusGold
	}
}
