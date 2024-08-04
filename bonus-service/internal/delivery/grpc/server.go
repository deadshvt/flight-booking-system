package grpc

import (
	"context"
	"errors"
	"time"

	"github.com/deadshvt/flight-booking-system/bonus-service/internal/converter"
	"github.com/deadshvt/flight-booking-system/bonus-service/internal/repository"
	"github.com/deadshvt/flight-booking-system/bonus-service/internal/usecase"
	"github.com/deadshvt/flight-booking-system/bonus-service/pkg/errs"
	pb "github.com/deadshvt/flight-booking-system/bonus-service/proto"
	"github.com/deadshvt/flight-booking-system/logger"

	"github.com/rs/zerolog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type BonusServiceServer struct {
	pb.UnimplementedBonusServiceServer
	Repo    *repository.BonusRepository
	Usecase *usecase.BonusUsecase
	Logger  zerolog.Logger
}

func NewBonusServiceServer(repo *repository.BonusRepository,
	usecase *usecase.BonusUsecase, logger zerolog.Logger) *BonusServiceServer {
	return &BonusServiceServer{
		Repo:    repo,
		Usecase: usecase,
		Logger:  logger,
	}
}

func (s *BonusServiceServer) GetPrivilegeWithHistory(ctx context.Context,
	req *pb.GetPrivilegeWithHistoryRequest) (*pb.GetPrivilegeWithHistoryResponse, error) {
	username := req.GetUsername()

	logger.LogWithParams(s.Logger, "Getting privilege with history...", struct {
		Username string
	}{username})

	privilegeWithHistory, err := s.Repo.GetPrivilegeWithHistory(ctx, username)
	if err != nil {
		s.Logger.Err(err).Msg("Failed to get privilege with history")
		return nil, s.HandleError(err)
	}

	protoHistory := make([]*pb.Operation, len(privilegeWithHistory.History))
	for i, operation := range privilegeWithHistory.History {
		protoHistory[i] = converter.OperationFromEntityToProto(operation)
	}

	privilege := privilegeWithHistory.Privilege

	logger.LogWithParams(s.Logger, "Got privilege with history", struct {
		Username string
		Status   string
		Balance  int32
		Count    int
	}{privilege.Username, privilege.Status, privilege.Balance,
		len(privilegeWithHistory.History)})

	return &pb.GetPrivilegeWithHistoryResponse{
		Privilege: converter.PrivilegeFromEntityToProto(privilege),
		History:   protoHistory,
	}, nil
}

func (s *BonusServiceServer) GetPrivilege(ctx context.Context,
	req *pb.GetPrivilegeRequest) (*pb.GetPrivilegeResponse, error) {
	username := req.GetUsername()

	logger.LogWithParams(s.Logger, "Getting privilege...", struct {
		Username string
	}{username})

	privilege, err := s.Repo.GetPrivilege(ctx, username)
	if err != nil {
		s.Logger.Err(err).Msg("Failed to get privilege")
		return nil, s.HandleError(err)
	}

	logger.LogWithParams(s.Logger, "Got privilege", struct {
		Username string
		Status   string
		Balance  int32
	}{privilege.Username, privilege.Status, privilege.Balance})

	return &pb.GetPrivilegeResponse{
		Privilege: converter.PrivilegeFromEntityToProto(privilege),
	}, nil
}

func (s *BonusServiceServer) CreatePrivilege(ctx context.Context,
	req *pb.CreatePrivilegeRequest) (*pb.CreatePrivilegeResponse, error) {
	privilege := converter.PrivilegeFromProtoToEntity(req.GetPrivilege())

	logger.LogWithParams(s.Logger, "Creating privilege...", struct {
		Username string
		Balance  int32
		Status   string
	}{privilege.Username, privilege.Balance, privilege.Status})

	err := s.Repo.CreatePrivilege(ctx, privilege)
	if err != nil {
		s.Logger.Err(err).Msg("Failed to create privilege")
		return nil, s.HandleError(err)
	}

	logger.LogWithParams(s.Logger, "Created privilege", struct {
		Username string
		Status   string
		Balance  int32
	}{privilege.Username, privilege.Status, privilege.Balance})

	return &pb.CreatePrivilegeResponse{}, nil
}

func (s *BonusServiceServer) UpdatePrivilege(ctx context.Context,
	req *pb.UpdatePrivilegeRequest) (*pb.UpdatePrivilegeResponse, error) {
	privilege := converter.PrivilegeFromProtoToEntity(req.GetPrivilege())

	logger.LogWithParams(s.Logger, "Updating privilege...", struct {
		Username string
		Balance  int32
		Status   string
	}{privilege.Username, privilege.Balance, privilege.Status})

	err := s.Repo.UpdatePrivilege(ctx, privilege)
	if err != nil {
		s.Logger.Err(err).Msg("Failed to update privilege")
		return nil, s.HandleError(err)
	}

	logger.LogWithParams(s.Logger, "Updated privilege", struct {
		Username string
		Status   string
		Balance  int32
	}{privilege.Username, privilege.Status, privilege.Balance})

	return &pb.UpdatePrivilegeResponse{}, nil
}

func (s *BonusServiceServer) CreateOperation(ctx context.Context,
	req *pb.CreateOperationRequest) (*pb.CreateOperationResponse, error) {
	operation := converter.OperationFromProtoToEntity(req.GetOperation())

	logger.LogWithParams(s.Logger, "Creating history...", struct {
		PrivilegeID   int32
		TicketUid     string
		Date          time.Time
		BalanceDiff   int32
		OperationType string
	}{operation.PrivilegeID, operation.TicketUid, operation.Date,
		operation.BalanceDiff, operation.OperationType})

	err := s.Repo.CreateOperation(ctx, operation)
	if err != nil {
		s.Logger.Err(err).Msg("Failed to create history")
		return nil, s.HandleError(err)
	}

	logger.LogWithParams(s.Logger, "Created history", struct {
		PrivilegeID   int32
		TicketUid     string
		Date          time.Time
		BalanceDiff   int32
		OperationType string
	}{operation.PrivilegeID, operation.TicketUid, operation.Date,
		operation.BalanceDiff, operation.OperationType})

	return &pb.CreateOperationResponse{}, nil
}

func (s *BonusServiceServer) HandleError(err error) error {
	switch {
	case errors.Is(err, errs.ErrPrivilegeAlreadyExists):
		return status.Error(codes.AlreadyExists, err.Error())
	case errors.Is(err, errs.ErrPrivilegeNotFound):
		return status.Error(codes.NotFound, err.Error())
	case errors.Is(err, errs.ErrHistoryNotFound):
		return status.Error(codes.NotFound, err.Error())
	default:
		return status.Error(codes.Internal, err.Error())
	}
}
