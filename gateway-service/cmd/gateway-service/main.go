package main

import (
	"context"
	"errors"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/deadshvt/flight-booking-system/config"
	"github.com/deadshvt/flight-booking-system/gateway-service/internal/client"
	gatewayGrpc "github.com/deadshvt/flight-booking-system/gateway-service/internal/delivery/grpc"
	"github.com/deadshvt/flight-booking-system/gateway-service/internal/usecase"
	gatewaypb "github.com/deadshvt/flight-booking-system/gateway-service/proto"
	"github.com/deadshvt/flight-booking-system/logger"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
)

const (
	Key = "x-user-name"
)

func main() {
	baseLogger, err := logger.Init("gateway-service")
	if err != nil {
		log.Fatal().Msgf("Failed to initialize logger: %v", err)
	}

	lg := logger.NewLogger(baseLogger, "gateway-service")

	lg.Info().Msg("Initialized logger")

	config.Load(".env")

	lg.Info().Msg("Loaded .env file")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	runtimeMux := runtime.NewServeMux([]runtime.ServeMuxOption{
		runtime.WithIncomingHeaderMatcher(func(key string) (string, bool) {
			if strings.ToLower(key) == Key {
				return key, true
			}
			return runtime.DefaultHeaderMatcher(key)
		}),
		runtime.WithMetadata(func(ctx context.Context, req *http.Request) metadata.MD {
			if username := req.Header.Get("X-User-Name"); username != "" {
				return metadata.Pairs(Key, username)
			}
			return nil
		})}...)
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	err = gatewaypb.RegisterGatewayHandlerFromEndpoint(ctx, runtimeMux, ":"+os.Getenv("GRPC_PORT"), opts)
	if err != nil {
		lg.Fatal().Msgf("Failed to register http: %v", err)
	}

	gateway, err := client.NewGateway()
	if err != nil {
		lg.Fatal().Msgf("Failed to create client: %v", err)
	}
	defer func() {
		err = gateway.Cleanup()
		if err != nil {
			lg.Error().Msgf("Failed to close clients: %v", err)
		}
	}()

	// Usecase
	gatewayUsecaseLogger := logger.NewLogger(baseLogger, "gateway-service/usecase")
	gatewayUsecase := usecase.NewGatewayUsecase(gateway, gatewayUsecaseLogger)

	// Server
	gatewayServerLogger := logger.NewLogger(baseLogger, "gateway-service/server")
	gatewayServer := gatewayGrpc.NewGatewayServer(gatewayUsecase, gatewayServerLogger)
	if err != nil {
		lg.Fatal().Msgf("Failed to initialize gateway-service server: %v", err)
	}

	grpcSrv := grpc.NewServer()
	gatewaypb.RegisterGatewayServer(grpcSrv, gatewayServer)
	reflection.Register(grpcSrv)

	httpSrv := &http.Server{
		Addr:    ":" + os.Getenv("HTTP_PORT"),
		Handler: runtimeMux,
	}

	lis, err := net.Listen("tcp", ":"+os.Getenv("GRPC_PORT"))
	if err != nil {
		lg.Fatal().Msgf("Failed to listen: %v", err)
	}

	lg.Info().Msgf("Starting gateway-service server on port :%s", os.Getenv("GRPC_PORT"))

	go func() {
		err = grpcSrv.Serve(lis)
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			lg.Error().Msgf("Failed to start grpc server: %v", err)
		}
	}()

	lg.Info().Msgf("Starting http server on port :%s...", os.Getenv("HTTP_PORT"))

	go func() {
		err = httpSrv.ListenAndServe()
		if err != nil && !errors.Is(http.ErrServerClosed, err) {
			lg.Error().Msgf("Failed to start http server: %v", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop

	lg.Info().Msg("Shutting down servers...")

	grpcSrv.GracefulStop()

	ctx, cancel = context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err = httpSrv.Shutdown(ctx); err != nil {
		lg.Error().Msgf("Failed to shutdown http server: %v", err)
	}

	lg.Info().Msg("Servers gracefully stopped")
}
