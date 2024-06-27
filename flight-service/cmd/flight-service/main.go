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
	flightGrpc "github.com/deadshvt/flight-booking-system/flight-service/internal/delivery/grpc"
	repo "github.com/deadshvt/flight-booking-system/flight-service/internal/repository/in_memory"
	"github.com/deadshvt/flight-booking-system/flight-service/internal/usecase"
	pb "github.com/deadshvt/flight-booking-system/flight-service/proto"
	"github.com/deadshvt/flight-booking-system/logger"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	baseLogger, err := logger.Init("flight-service")
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}

	lg := logger.NewLogger(baseLogger, "flight-service")

	lg.Info().Msg("Initialized logger")

	config.Load(".env")

	lis, err := net.Listen("tcp", ":"+os.Getenv("FLIGHT_PORT"))
	if err != nil {
		lg.Fatal().Msgf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	flightRepo := repo.NewInMemoryFlightRepository()
	flightUsecaseLogger := logger.NewLogger(baseLogger, "flight-service/usecase")
	flightUsecase := usecase.NewFlightUsecase(flightRepo, flightUsecaseLogger)
	flightServerLogger := logger.NewLogger(baseLogger, "flight-service/server")
	flightServer := flightGrpc.NewFlightServiceServer(flightRepo, flightUsecase, flightServerLogger)

	pb.RegisterFlightServiceServer(s, flightServer)
	reflection.Register(s)

	lg.Info().Msgf("Starting flight service on port :%s", os.Getenv("FLIGHT_PORT"))

	go func() {
		if err = s.Serve(lis); err != nil && !errors.Is(err, http.ErrServerClosed) {
			lg.Error().Msgf("Failed to start flight service: %v", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop

	lg.Info().Msg("Shutting down flight service...")

	s.GracefulStop()

	lg.Info().Msg("Flight service gracefully stopped")
}
