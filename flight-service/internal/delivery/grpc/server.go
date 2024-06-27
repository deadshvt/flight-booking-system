package grpc

import (
	"context"

	"github.com/deadshvt/flight-booking-system/flight-service/internal/converter"
	"github.com/deadshvt/flight-booking-system/flight-service/internal/repository"
	"github.com/deadshvt/flight-booking-system/flight-service/internal/usecase"
	pb "github.com/deadshvt/flight-booking-system/flight-service/proto"
	"github.com/deadshvt/flight-booking-system/logger"

	"github.com/rs/zerolog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type FlightServiceServer struct {
	pb.UnimplementedFlightServiceServer
	Repo    repository.FlightRepository
	Usecase *usecase.FlightUsecase
	Logger  zerolog.Logger
}

func NewFlightServiceServer(repo repository.FlightRepository,
	usecase *usecase.FlightUsecase, logger zerolog.Logger) *FlightServiceServer {
	return &FlightServiceServer{
		Repo:    repo,
		Usecase: usecase,
		Logger:  logger,
	}
}

func (s *FlightServiceServer) GetFlightsWithAirports(ctx context.Context,
	req *pb.GetFlightsWithAirportsRequest) (*pb.GetFlightsWithAirportsResponse, error) {
	s.Logger.Info().Msg("Getting flights...")

	protoFlights, err := s.Usecase.GetFlightsWithAirports(ctx, req)
	if err != nil {
		s.Logger.Err(err).Msg("Failed to get flights")
		return nil, status.Errorf(codes.Internal, "failed to get flights: %v", err)
	}

	logger.LogWithParams(s.Logger, "Got flights", struct {
		Count int32
	}{protoFlights.TotalElements})

	return protoFlights, nil
}

func (s *FlightServiceServer) GetFlightWithAirports(ctx context.Context,
	req *pb.GetFlightWithAirportsRequest) (*pb.GetFlightWithAirportsResponse, error) {
	logger.LogWithParams(s.Logger, "Getting flight...", struct {
		FlightNumber string
	}{req.FlightNumber})

	flight, err := s.Repo.GetFlightWithAirports(ctx, req.FlightNumber)
	if err != nil {
		s.Logger.Err(err).Msg("Failed to get flight")
		return nil, status.Errorf(codes.Internal, "failed to get flight: %v", err)
	}

	logger.LogWithParams(s.Logger, "Got flight", struct {
		FlightNumber string
		FromAirport  string
		ToAirport    string
		Date         string
		Price        int32
	}{flight.Flight.FlightNumber, flight.FromAirport.Name,
		flight.ToAirport.Name, flight.Flight.Date, flight.Flight.Price})

	return &pb.GetFlightWithAirportsResponse{
		Flight: converter.FlightFromEntityToProto(flight),
	}, nil
}
