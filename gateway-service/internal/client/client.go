package client

import (
	"fmt"
	"os"
	"strings"

	bonuspb "github.com/deadshvt/flight-booking-system/bonus-service/proto"
	"github.com/deadshvt/flight-booking-system/config"
	flightpb "github.com/deadshvt/flight-booking-system/flight-service/proto"
	"github.com/deadshvt/flight-booking-system/gateway-service/internal/errs"
	ticketpb "github.com/deadshvt/flight-booking-system/ticket-service/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	Prefix = "dns:///"

	TicketConn = "ticketConn: "
	FlightConn = "flightConn: "
	BonusConn  = "bonusConn: "
)

type Gateway struct {
	TicketClient ticketpb.TicketServiceClient
	FlightClient flightpb.FlightServiceClient
	BonusClient  bonuspb.BonusServiceClient

	TicketConn *grpc.ClientConn
	FlightConn *grpc.ClientConn
	BonusConn  *grpc.ClientConn
}

func NewGateway() (*Gateway, error) {
	config.Load(".env") //TODO

	ticketConn, err := grpc.NewClient(Prefix+os.Getenv("TICKET_HOST")+":"+os.Getenv("TICKET_PORT"),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, errs.WrapError(errs.ErrCreateConnection, err)
	}

	flightConn, err := grpc.NewClient(Prefix+os.Getenv("FLIGHT_HOST")+":"+os.Getenv("FLIGHT_PORT"),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		err = errs.WrapError(errs.ErrCreateConnection, err)

		var errStrings []string

		errTicket := ticketConn.Close()
		if errTicket != nil {
			errStrings = append(errStrings, "ticketConn: "+errTicket.Error())
		}

		if len(errStrings) > 0 {
			err = fmt.Errorf("%v & %v", err,
				errs.WrapError(errs.ErrCloseConnection, fmt.Errorf("%s", strings.Join(errStrings, " & "))))
		}

		return nil, err
	}

	bonusConn, err := grpc.NewClient(Prefix+os.Getenv("BONUS_HOST")+":"+os.Getenv("BONUS_PORT"),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		err = errs.WrapError(errs.ErrCreateConnection, err)

		var errStrings []string

		errTicket := ticketConn.Close()
		if errTicket != nil {
			errStrings = append(errStrings, TicketConn+errTicket.Error())
		}
		errFlight := flightConn.Close()
		if errFlight != nil {
			errStrings = append(errStrings, FlightConn+errFlight.Error())
		}

		if len(errStrings) > 0 {
			err = fmt.Errorf("%v & %v", err,
				errs.WrapError(errs.ErrCloseConnection, fmt.Errorf("%s", strings.Join(errStrings, " & "))))
		}

		return nil, err
	}

	return &Gateway{
		TicketClient: ticketpb.NewTicketServiceClient(ticketConn),
		FlightClient: flightpb.NewFlightServiceClient(flightConn),
		BonusClient:  bonuspb.NewBonusServiceClient(bonusConn),
		TicketConn:   ticketConn,
		FlightConn:   flightConn,
		BonusConn:    bonusConn,
	}, nil
}

func (g *Gateway) Cleanup() error {
	var errStrings []string

	err := g.TicketConn.Close()
	if err != nil {
		errStrings = append(errStrings, TicketConn+err.Error())
	}
	err = g.FlightConn.Close()
	if err != nil {
		errStrings = append(errStrings, FlightConn+err.Error())
	}
	err = g.BonusConn.Close()
	if err != nil {
		errStrings = append(errStrings, BonusConn+err.Error())
	}

	if len(errStrings) > 0 {
		return errs.WrapError(errs.ErrCloseConnection, fmt.Errorf("%s", strings.Join(errStrings, " & ")))
	}

	return nil
}
