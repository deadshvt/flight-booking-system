package main

import (
	"errors"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	bonusGrpc "github.com/deadshvt/flight-booking-system/bonus-service/internal/delivery/grpc"
	repo "github.com/deadshvt/flight-booking-system/bonus-service/internal/repository/in_memory"
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

	lis, err := net.Listen("tcp", ":"+os.Getenv("BONUS_PORT"))
	if err != nil {
		lg.Fatal().Msgf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	bonusRepo := repo.NewInMemoryBonusRepository()
	bonusUsecaseLogger := logger.NewLogger(baseLogger, "ticket-service/usecase")
	bonusUsecase := usecase.NewBonusUsecase(bonusRepo, bonusUsecaseLogger)
	bonusServerLogger := logger.NewLogger(baseLogger, "ticket-service/server")
	bonusServer := bonusGrpc.NewBonusServiceServer(bonusRepo, bonusUsecase, bonusServerLogger)

	pb.RegisterBonusServiceServer(s, bonusServer)
	reflection.Register(s)

	lg.Info().Msgf("Starting bonus service on port :%s", os.Getenv("BONUS_PORT"))

	go func() {
		if err = s.Serve(lis); err != nil && !errors.Is(err, http.ErrServerClosed) {
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
