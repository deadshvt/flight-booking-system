package usecase

import (
	"context"

	"github.com/deadshvt/flight-booking-system/flight-service/internal/converter"
	"github.com/deadshvt/flight-booking-system/flight-service/internal/repository"
	pb "github.com/deadshvt/flight-booking-system/flight-service/proto"

	"github.com/rs/zerolog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type FlightUsecase struct {
	Repo   repository.FlightRepository
	Logger zerolog.Logger
}

func NewFlightUsecase(repo repository.FlightRepository, logger zerolog.Logger) *FlightUsecase {
	return &FlightUsecase{
		Repo:   repo,
		Logger: logger,
	}
}

func (u *FlightUsecase) GetFlightsWithAirports(ctx context.Context,
	req *pb.GetFlightsWithAirportsRequest) (*pb.GetFlightsWithAirportsResponse, error) {
	u.Logger.Info().Msg("Getting flights...")

	page := req.Page
	if page < 0 {
		u.Logger.Error().Msg("Invalid page")
		return nil, status.Errorf(codes.InvalidArgument, "invalid page")
	}

	size := req.Size
	if size < 1 || size > 100 {
		u.Logger.Error().Msg("Invalid size")
		return nil, status.Errorf(codes.InvalidArgument, "invalid size")
	}

	flights, err := u.Repo.GetFlightsWithAirports(ctx)
	if err != nil {
		u.Logger.Err(err).Msg("Failed to get flights")
		return nil, status.Errorf(codes.Internal, "failed to get flights: %v", err)
	}

	start := (page - 1) * size
	end := start + size
	length := int32(len(flights))

	if start > length {
		start = length
	}
	if end > length {
		end = length
	}

	cutFlights := flights[start:end]

	protoFlights := make([]*pb.FlightWithAirports, len(cutFlights))
	for i, flight := range cutFlights {
		protoFlights[i] = converter.FlightFromEntityToProto(flight)
	}

	return &pb.GetFlightsWithAirportsResponse{
		Page:          page,
		PageSize:      size,
		TotalElements: int32(len(protoFlights)),
		Items:         protoFlights,
	}, nil
}
