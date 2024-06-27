package main

import (
	"errors"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/deadshvt/flight-booking-system/config"
	"github.com/deadshvt/flight-booking-system/logger"
	ticketGrpc "github.com/deadshvt/flight-booking-system/ticket-service/internal/delivery/grpc"
	repo "github.com/deadshvt/flight-booking-system/ticket-service/internal/repository/in_memory"
	"github.com/deadshvt/flight-booking-system/ticket-service/internal/usecase"
	pb "github.com/deadshvt/flight-booking-system/ticket-service/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	baseLogger, err := logger.Init("ticket-service")
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}

	lg := logger.NewLogger(baseLogger, "ticket-service")

	lg.Info().Msg("Initialized logger")

	config.Load(".env")

	lis, err := net.Listen("tcp", ":"+os.Getenv("TICKET_PORT"))
	if err != nil {
		lg.Fatal().Msgf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	ticketRepo := repo.NewInMemoryTicketRepository()
	ticketUsecaseLogger := logger.NewLogger(baseLogger, "ticket-service/usecase")
	ticketUsecase := usecase.NewTicketUsecase(ticketRepo, ticketUsecaseLogger)
	ticketServerLogger := logger.NewLogger(baseLogger, "ticket-service/server")
	ticketServer := ticketGrpc.NewTicketServiceServer(ticketRepo, ticketUsecase, ticketServerLogger)

	pb.RegisterTicketServiceServer(s, ticketServer)
	reflection.Register(s)

	lg.Info().Msgf("Starting ticket service on port :%s", os.Getenv("TICKET_PORT"))

	go func() {
		if err = s.Serve(lis); err != nil && !errors.Is(err, http.ErrServerClosed) {
			lg.Error().Msgf("Failed to start ticket service: %v", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop

	lg.Info().Msg("Shutting down ticket service...")

	s.GracefulStop()

	lg.Info().Msg("Ticket service gracefully stopped")
}
