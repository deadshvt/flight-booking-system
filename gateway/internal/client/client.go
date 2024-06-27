package client

import (
	"fmt"
	"os"
	"strings"

	bonuspb "github.com/deadshvt/flight-booking-system/bonus-service/proto"
	flightpb "github.com/deadshvt/flight-booking-system/flight-service/proto"
	ticketpb "github.com/deadshvt/flight-booking-system/ticket-service/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	Prefix  = "dns:///"
	Address = "localhost:"
)

func Init() (ticketpb.TicketServiceClient, flightpb.FlightServiceClient, bonuspb.BonusServiceClient, func() error, error) {
	ticketConn, err := grpc.NewClient(Prefix+Address+os.Getenv("TICKET_PORT"),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, nil, nil, nil, err
	}

	flightConn, err := grpc.NewClient(Prefix+Address+os.Getenv("FLIGHT_PORT"),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		var errStrings []string

		errTicket := ticketConn.Close()
		if errTicket != nil {
			errStrings = append(errStrings, errTicket.Error())
		}

		return nil, nil, nil, nil, fmt.Errorf("failed to close connections: %s",
			strings.Join(errStrings, " & "))
	}

	bonusConn, err := grpc.NewClient(Prefix+Address+os.Getenv("BONUS_PORT"),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		var errStrings []string

		errFlight := flightConn.Close()
		if errFlight != nil {
			errStrings = append(errStrings, errFlight.Error())
		}
		errTicket := ticketConn.Close()
		if errTicket != nil {
			errStrings = append(errStrings, errTicket.Error())
		}

		return nil, nil, nil, nil, fmt.Errorf("failed to close connections: %s",
			strings.Join(errStrings, " & "))
	}

	ticketClient := ticketpb.NewTicketServiceClient(ticketConn)
	flightClient := flightpb.NewFlightServiceClient(flightConn)
	bonusClient := bonuspb.NewBonusServiceClient(bonusConn)

	cleanup := func() error {
		var errStrings []string

		if err = ticketConn.Close(); err != nil {
			errStrings = append(errStrings, "ticketConn: "+err.Error())
		}
		if err = flightConn.Close(); err != nil {
			errStrings = append(errStrings, "flightConn: "+err.Error())
		}
		if err = bonusConn.Close(); err != nil {
			errStrings = append(errStrings, "bonusConn: "+err.Error())
		}

		if len(errStrings) > 0 {
			return fmt.Errorf("failed to close connections: %s", strings.Join(errStrings, " & "))
		}

		return nil
	}

	return ticketClient, flightClient, bonusClient, cleanup, nil
}
