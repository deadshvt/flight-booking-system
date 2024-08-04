package main

import (
	"context"
	"errors"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/deadshvt/flight-booking-system/config"
	flightGrpc "github.com/deadshvt/flight-booking-system/flight-service/internal/delivery/grpc"
	"github.com/deadshvt/flight-booking-system/flight-service/internal/repository"
	"github.com/deadshvt/flight-booking-system/flight-service/internal/repository/cache"
	"github.com/deadshvt/flight-booking-system/flight-service/internal/repository/database"
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

	lg.Info().Msg("Loaded .env file")

	lis, err := net.Listen("tcp", ":"+os.Getenv("FLIGHT_PORT"))
	if err != nil {
		lg.Fatal().Msgf("Failed to listen: %v", err)
	}

	ctx := context.Background()

	// DB
	flightDB, err := database.NewFlightDB(ctx, os.Getenv("FLIGHT_DB_TYPE"))
	if err != nil {
		lg.Fatal().Msgf("Failed to connect to database: %v", err)
	}

	// Cache
	flightCache, err := cache.NewFlightCache(ctx, os.Getenv("FLIGHT_CACHE_TYPE"))
	if err != nil {
		lg.Fatal().Msgf("Failed to connect to cache: %v", err)
	}

	// Repository
	flightRepoLogger := logger.NewLogger(baseLogger, "flight-service/repository")
	flightRepo := repository.NewFlightRepository(flightDB, flightCache, flightRepoLogger)

	// Usecase
	flightUsecaseLogger := logger.NewLogger(baseLogger, "flight-service/usecase")
	flightUsecase := usecase.NewFlightUsecase(flightRepo, flightUsecaseLogger)

	// Server
	flightServerLogger := logger.NewLogger(baseLogger, "flight-service/server")
	flightServer := flightGrpc.NewFlightServiceServer(flightRepo, flightUsecase, flightServerLogger)

	s := grpc.NewServer()
	pb.RegisterFlightServiceServer(s, flightServer)
	reflection.Register(s)

	lg.Info().Msgf("Starting flight service on port :%s...", os.Getenv("FLIGHT_PORT"))

	go func() {
		err = s.Serve(lis)
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
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
