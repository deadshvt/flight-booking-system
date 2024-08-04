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

	bonusGrpc "github.com/deadshvt/flight-booking-system/bonus-service/internal/delivery/grpc"
	"github.com/deadshvt/flight-booking-system/bonus-service/internal/repository"
	"github.com/deadshvt/flight-booking-system/bonus-service/internal/repository/cache"
	"github.com/deadshvt/flight-booking-system/bonus-service/internal/repository/database"
	"github.com/deadshvt/flight-booking-system/bonus-service/internal/usecase"
	pb "github.com/deadshvt/flight-booking-system/bonus-service/proto"
	"github.com/deadshvt/flight-booking-system/config"
	"github.com/deadshvt/flight-booking-system/logger"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	baseLogger, err := logger.Init("bonus-service")
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}

	lg := logger.NewLogger(baseLogger, "bonus-service")

	lg.Info().Msg("Initialized logger")

	config.Load(".env")

	lg.Info().Msg("Loaded .env file")

	lis, err := net.Listen("tcp", ":"+os.Getenv("BONUS_PORT"))
	if err != nil {
		lg.Fatal().Msgf("Failed to listen: %v", err)
	}

	ctx := context.Background()

	// DB
	bonusDB, err := database.NewBonusDB(ctx, os.Getenv("BONUS_DB_TYPE"))
	if err != nil {
		lg.Fatal().Msgf("Failed to connect to database: %v", err)
	}

	// Cache
	bonusCache, err := cache.NewBonusCache(ctx, os.Getenv("BONUS_CACHE_TYPE"))
	if err != nil {
		lg.Fatal().Msgf("Failed to connect to cache: %v", err)
	}

	// Repository
	bonusRepoLogger := logger.NewLogger(baseLogger, "bonus-service/repository")
	bonusRepo := repository.NewBonusRepository(bonusDB, bonusCache, bonusRepoLogger)

	// Usecase
	bonusUsecaseLogger := logger.NewLogger(baseLogger, "ticket-service/usecase")
	bonusUsecase := usecase.NewBonusUsecase(bonusRepo, bonusUsecaseLogger)

	// Server
	bonusServerLogger := logger.NewLogger(baseLogger, "ticket-service/server")
	bonusServer := bonusGrpc.NewBonusServiceServer(bonusRepo, bonusUsecase, bonusServerLogger)

	s := grpc.NewServer()
	pb.RegisterBonusServiceServer(s, bonusServer)
	reflection.Register(s)

	lg.Info().Msgf("Starting bonus service on port :%s...", os.Getenv("BONUS_PORT"))

	go func() {
		err = s.Serve(lis)
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			lg.Error().Msgf("Failed to start bonus service: %v", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop

	lg.Info().Msg("Shutting down bonus service...")

	s.GracefulStop()

	lg.Info().Msg("Bonus service gracefully stopped")
}
