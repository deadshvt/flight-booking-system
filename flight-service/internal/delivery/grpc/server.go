package grpc

import (
	"context"
	"errors"
	"time"

	"github.com/deadshvt/flight-booking-system/flight-service/internal/converter"
	"github.com/deadshvt/flight-booking-system/flight-service/internal/repository"
	"github.com/deadshvt/flight-booking-system/flight-service/internal/usecase"
	"github.com/deadshvt/flight-booking-system/flight-service/pkg/errs"
	pb "github.com/deadshvt/flight-booking-system/flight-service/proto"
	"github.com/deadshvt/flight-booking-system/logger"

	"github.com/rs/zerolog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type FlightServiceServer struct {
	pb.UnimplementedFlightServiceServer
	Repo    *repository.FlightRepository
	Usecase *usecase.FlightUsecase
	Logger  zerolog.Logger
}

func NewFlightServiceServer(repo *repository.FlightRepository,
	usecase *usecase.FlightUsecase, logger zerolog.Logger) *FlightServiceServer {
	return &FlightServiceServer{
		Repo:    repo,
		Usecase: usecase,
		Logger:  logger,
	}
}

func (s *FlightServiceServer) GetFlightsWithAirports(ctx context.Context,
	req *pb.GetFlightsWithAirportsRequest) (*pb.GetFlightsWithAirportsResponse, error) {
	s.Logger.Info().Msg("Getting flights with airports...")

	page := req.GetPage()
	size := req.GetSize()

	flights, err := s.Usecase.GetFlightsWithAirports(ctx, page, size)
	if err != nil {
		s.Logger.Err(err).Msg("Failed to get flights")
		return nil, s.HandleError(err)
	}

	flightLen := int32(len(flights))

	protoFlights := make([]*pb.FlightWithAirports, flightLen)
	for i, flight := range flights {
		protoFlights[i] = converter.FlightFromEntityToProto(flight)
	}

	logger.LogWithParams(s.Logger, "Got flights with airports", struct {
		Count int32
	}{flightLen})

	return &pb.GetFlightsWithAirportsResponse{
		Page:          page,
		PageSize:      size,
		TotalElements: flightLen,
		Items:         protoFlights,
	}, nil
}

func (s *FlightServiceServer) GetFlightWithAirports(ctx context.Context,
	req *pb.GetFlightWithAirportsRequest) (*pb.GetFlightWithAirportsResponse, error) {
	flightNumber := req.GetFlightNumber()

	logger.LogWithParams(s.Logger, "Getting flight with airports...", struct {
		FlightNumber string
	}{flightNumber})

	flight, err := s.Repo.GetFlightWithAirports(ctx, flightNumber)
	if err != nil {
		s.Logger.Err(err).Msg("Failed to get flight")
		return nil, s.HandleError(err)
	}

	logger.LogWithParams(s.Logger, "Got flight with airports", struct {
		FlightNumber string
		FromAirport  string
		ToAirport    string
		Date         time.Time
		Price        int32
	}{flight.FlightNumber, flight.FromAirport,
		flight.ToAirport, flight.Date, flight.Price})

	return &pb.GetFlightWithAirportsResponse{
		Flight: converter.FlightFromEntityToProto(flight),
	}, nil
}

func (s *FlightServiceServer) HandleError(err error) error {
	switch {
	case errors.Is(err, errs.ErrInvalidPage):
		return status.Error(codes.InvalidArgument, err.Error())
	case errors.Is(err, errs.ErrInvalidSize):
		return status.Error(codes.InvalidArgument, err.Error())
	case errors.Is(err, errs.ErrFlightNotFound):
		return status.Error(codes.NotFound, err.Error())
	default:
		return status.Error(codes.Internal, err.Error())
	}
}
