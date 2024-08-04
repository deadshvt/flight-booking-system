package usecase

import (
	"context"
	"errors"
	"math"
	"sync"
	"time"

	bonusErrs "github.com/deadshvt/flight-booking-system/bonus-service/pkg/errs"
	bonuspb "github.com/deadshvt/flight-booking-system/bonus-service/proto"
	flightpb "github.com/deadshvt/flight-booking-system/flight-service/proto"
	"github.com/deadshvt/flight-booking-system/gateway-service/internal/client"
	gatewaypb "github.com/deadshvt/flight-booking-system/gateway-service/proto"
	"github.com/deadshvt/flight-booking-system/logger"
	ticketpb "github.com/deadshvt/flight-booking-system/ticket-service/proto"

	"github.com/rs/zerolog"
	"google.golang.org/protobuf/types/known/timestamppb"
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
	Gateway *client.Gateway
	Logger  zerolog.Logger
}

func NewGatewayUsecase(gateway *client.Gateway, logger zerolog.Logger) *GatewayUsecase {
	return &GatewayUsecase{
		Gateway: gateway,
		Logger:  logger,
	}
}

func (u *GatewayUsecase) GetFlightsWithAirports(ctx context.Context,
	req *flightpb.GetFlightsWithAirportsRequest) (*flightpb.GetFlightsWithAirportsResponse, error) {
	u.Logger.Info().Msg("Getting flights with airports...")

	return u.Gateway.FlightClient.GetFlightsWithAirports(ctx, req)
}

func (u *GatewayUsecase) GetTicketsWithAirports(ctx context.Context,
	req *gatewaypb.GetTicketsWithAirportsRequest) (*gatewaypb.GetTicketsWithAirportsResponse, error) {
	username := req.GetUsername()

	logger.LogWithParams(u.Logger, "Getting tickets with airports...", struct {
		Username string
	}{username})

	protoTickets, err := u.Gateway.TicketClient.GetTickets(ctx, &ticketpb.GetTicketsRequest{
		Username: username,
	})
	if err != nil {
		u.Logger.Err(err).Msg("Failed to get tickets")
		return nil, err
	}

	var (
		wg                       sync.WaitGroup
		mu                       sync.Mutex
		ch                       = make(chan error, 1)
		protoTicketsWithAirports = make([]*gatewaypb.TicketWithAirports, 0, len(protoTickets.GetTickets()))
	)

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	for _, ticket := range protoTickets.GetTickets() {
		wg.Add(1)

		go func(ticket *ticketpb.Ticket) {
			defer wg.Done()

			select {
			case <-ctx.Done():
				return
			default:
			}

			protoFlight, err := u.Gateway.FlightClient.GetFlightWithAirports(ctx,
				&flightpb.GetFlightWithAirportsRequest{
					FlightNumber: ticket.GetFlightNumber(),
				})
			if err != nil {
				u.Logger.Err(err).Msgf("Failed to get flight with airports for flight number %s",
					ticket.GetFlightNumber())
				ch <- err
				cancel()
				return
			}

			flight := protoFlight.GetFlight()

			mu.Lock()
			protoTicketsWithAirports = append(protoTicketsWithAirports, &gatewaypb.TicketWithAirports{
				TicketUid:    ticket.GetTicketUid(),
				FlightNumber: ticket.GetFlightNumber(),
				FromAirport:  flight.GetFromAirport(),
				ToAirport:    flight.GetToAirport(),
				Date:         flight.GetDate(),
				Price:        ticket.GetPrice(),
				Status:       ticket.GetStatus(),
			})
			mu.Unlock()
		}(ticket)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	err = <-ch
	if err != nil {
		return nil, err
	}

	return &gatewaypb.GetTicketsWithAirportsResponse{
		Tickets: protoTicketsWithAirports,
	}, nil
}

func (u *GatewayUsecase) GetTicketWithAirports(ctx context.Context,
	req *gatewaypb.GetTicketWithAirportsRequest) (*gatewaypb.GetTicketWithAirportsResponse, error) {
	username := req.GetUsername()
	ticketUid := req.GetTicketUid()

	logger.LogWithParams(u.Logger, "Getting ticket with airports...", struct {
		Username  string
		TicketUid string
	}{username, ticketUid})

	protoTicket, err := u.Gateway.TicketClient.GetTicket(ctx, &ticketpb.GetTicketRequest{
		Username:  username,
		TicketUid: ticketUid,
	})
	if err != nil {
		u.Logger.Err(err).Msg("Failed to get ticket")
		return nil, err
	}

	ticket := protoTicket.GetTicket()

	protoFlight, err := u.Gateway.FlightClient.GetFlightWithAirports(ctx, &flightpb.GetFlightWithAirportsRequest{
		FlightNumber: ticket.GetFlightNumber(),
	})
	if err != nil {
		u.Logger.Err(err).Msg("Failed to get flight with airports")
		return nil, err
	}

	flight := protoFlight.Flight

	return &gatewaypb.GetTicketWithAirportsResponse{
		Ticket: &gatewaypb.TicketWithAirports{
			TicketUid:    ticket.GetTicketUid(),
			FlightNumber: ticket.GetFlightNumber(),
			FromAirport:  flight.GetFromAirport(),
			ToAirport:    flight.GetToAirport(),
			Date:         flight.GetDate(),
			Price:        ticket.GetPrice(),
			Status:       ticket.GetStatus(),
		},
	}, nil
}

func (u *GatewayUsecase) PurchaseTicket(ctx context.Context,
	req *gatewaypb.PurchaseTicketRequest) (*gatewaypb.PurchaseTicketResponse, error) {
	username := req.GetUsername()
	flightNumber := req.GetFlightNumber()
	price := req.GetPrice()
	paidFromBalance := req.GetPaidFromBalance()

	logger.LogWithParams(u.Logger, "Purchasing ticket...", struct {
		Username     string
		FlightNumber string
		Price        int32
	}{username, flightNumber, price})

	protoFlight, err := u.Gateway.FlightClient.GetFlightWithAirports(ctx, &flightpb.GetFlightWithAirportsRequest{
		FlightNumber: flightNumber,
	})
	if err != nil {
		u.Logger.Err(err).Msg("Failed to get flight with airports")
		return nil, err
	}

	flight := protoFlight.GetFlight()

	protoTicket, err := u.Gateway.TicketClient.PurchaseTicket(ctx, &ticketpb.PurchaseTicketRequest{
		Username:     username,
		FlightNumber: flightNumber,
		Price:        price,
	})
	if err != nil {
		u.Logger.Err(err).Msg("Failed to purchase ticket")
		return nil, err
	}

	ticketUid := protoTicket.GetTicketUid()

	var privilege *bonuspb.Privilege

	protoPrivilege, err := u.Gateway.BonusClient.GetPrivilege(ctx, &bonuspb.GetPrivilegeRequest{
		Username: username,
	})
	if err != nil && !errors.Is(err, bonusErrs.ErrPrivilegeNotFound) {
		u.Logger.Err(err).Msg("Failed to get privilege")
		return nil, err
	}
	if err != nil && errors.Is(err, bonusErrs.ErrPrivilegeNotFound) {
		privilege = &bonuspb.Privilege{
			Username: username,
			Balance:  0,
			Status:   StatusBronze,
		}
		_, err = u.Gateway.BonusClient.CreatePrivilege(ctx, &bonuspb.CreatePrivilegeRequest{
			Privilege: privilege,
		})
		if err != nil {
			u.Logger.Err(err).Msg("Failed to create privilege")
			return nil, err
		}
	} else {
		privilege = protoPrivilege.GetPrivilege()
	}

	operationType := OperationFill
	balanceDiff := int32(math.Floor(float64(req.Price) * 0.1))
	paidByMoney := price
	paidByBonuses := int32(0)

	if paidFromBalance {
		operationType = OperationDebit
		balanceDiff = -privilege.Balance
		paidByBonuses = privilege.Balance
		paidByMoney = price - paidByMoney
	}

	privilege.Balance += balanceDiff
	UpdatePrivilegeStatus(privilege)

	_, err = u.Gateway.BonusClient.UpdatePrivilege(ctx, &bonuspb.UpdatePrivilegeRequest{
		Privilege: privilege,
	})
	if err != nil {
		u.Logger.Err(err).Msg("Failed to update privilege")
		return nil, err
	}

	_, err = u.Gateway.BonusClient.CreateOperation(ctx, &bonuspb.CreateOperationRequest{
		Operation: &bonuspb.Operation{
			PrivilegeID:   privilege.GetID(),
			TicketUid:     ticketUid,
			Date:          timestamppb.New(time.Now()),
			BalanceDiff:   balanceDiff,
			OperationType: operationType,
		},
	})
	if err != nil {
		u.Logger.Err(err).Msg("Failed to create operation")
		return nil, err
	}

	return &gatewaypb.PurchaseTicketResponse{
		TicketUid:     ticketUid,
		FlightNumber:  flightNumber,
		FromAirport:   flight.GetFromAirport(),
		ToAirport:     flight.GetToAirport(),
		Date:          flight.GetDate(),
		Price:         flight.GetPrice(),
		PaidByMoney:   paidByMoney,
		PaidByBonuses: paidByBonuses,
		Status:        StatusPurchase,
		Privilege: &gatewaypb.PrivilegeShortInfo{
			Balance: privilege.GetBalance(),
			Status:  privilege.GetStatus(),
		},
	}, nil
}

func (u *GatewayUsecase) GetPrivilegeWithHistory(ctx context.Context,
	req *bonuspb.GetPrivilegeWithHistoryRequest) (*bonuspb.GetPrivilegeWithHistoryResponse, error) {
	username := req.GetUsername()

	logger.LogWithParams(u.Logger, "Getting privilege with history...", struct {
		Username string
	}{username})

	return u.Gateway.BonusClient.GetPrivilegeWithHistory(ctx, req)
}

func (u *GatewayUsecase) ReturnTicket(ctx context.Context,
	req *gatewaypb.ReturnTicketRequest) (*gatewaypb.ReturnTicketResponse, error) {
	username := req.GetUsername()
	ticketUid := req.GetTicketUid()

	logger.LogWithParams(u.Logger, "Returning ticket...", struct {
		Username  string
		TicketUid string
	}{username, ticketUid})

	_, err := u.Gateway.TicketClient.ReturnTicket(ctx, &ticketpb.ReturnTicketRequest{
		Username:  username,
		TicketUid: ticketUid,
	})
	if err != nil {
		u.Logger.Err(err).Msg("Failed to return ticket")
		return nil, err
	}

	protoPrivilegeWithHistory, err := u.Gateway.BonusClient.GetPrivilegeWithHistory(ctx,
		&bonuspb.GetPrivilegeWithHistoryRequest{
			Username: username,
		})
	if err != nil {
		u.Logger.Err(err).Msg("Failed to get privilege with history")
		return nil, err
	}

	privilege := protoPrivilegeWithHistory.GetPrivilege()
	history := protoPrivilegeWithHistory.GetHistory()

	var balanceDiff int32
	for _, h := range history {
		if h.GetTicketUid() == ticketUid {
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

	_, err = u.Gateway.BonusClient.UpdatePrivilege(ctx, &bonuspb.UpdatePrivilegeRequest{
		Privilege: privilege,
	})
	if err != nil {
		u.Logger.Err(err).Msg("Failed to update privilege")
		return nil, err
	}

	_, err = u.Gateway.BonusClient.CreateOperation(ctx, &bonuspb.CreateOperationRequest{
		Operation: &bonuspb.Operation{
			PrivilegeID:   privilege.GetID(),
			TicketUid:     ticketUid,
			Date:          timestamppb.New(time.Now()),
			BalanceDiff:   balanceDiff,
			OperationType: operationType,
		},
	})
	if err != nil {
		u.Logger.Err(err).Msg("Failed to create history")
		return nil, err
	}

	return &gatewaypb.ReturnTicketResponse{}, nil
}

func (u *GatewayUsecase) GetMe(ctx context.Context,
	req *gatewaypb.GetMeRequest) (*gatewaypb.GetMeResponse, error) {
	username := req.GetUsername()

	logger.LogWithParams(u.Logger, "Getting user info...", struct {
		Username string
	}{username})

	protoTickets, err := u.GetTicketsWithAirports(ctx, &gatewaypb.GetTicketsWithAirportsRequest{
		Username: username,
	})
	if err != nil {
		u.Logger.Err(err).Msg("Failed to get tickets with airports")
		return nil, err
	}

	protoPrivilege, err := u.Gateway.BonusClient.GetPrivilege(ctx, &bonuspb.GetPrivilegeRequest{
		Username: username,
	})
	if err != nil {
		u.Logger.Err(err).Msg("Failed to get privilege with history")
		return nil, err
	}

	privilege := protoPrivilege.GetPrivilege()
	privilegeShortInfo := &gatewaypb.PrivilegeShortInfo{
		Balance: privilege.GetBalance(),
		Status:  privilege.GetStatus(),
	}

	return &gatewaypb.GetMeResponse{
		Tickets:   protoTickets.GetTickets(),
		Privilege: privilegeShortInfo,
	}, nil
}

func UpdatePrivilegeStatus(privilege *bonuspb.Privilege) {
	switch {
	case privilege.Balance < 0:
		privilege.Balance = 0
	case privilege.Balance < 1000:
		privilege.Status = StatusBronze
	case privilege.Balance < 5000:
		privilege.Status = StatusSilver
	default:
		privilege.Status = StatusGold
	}
}
