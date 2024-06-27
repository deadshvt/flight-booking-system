package grpc

import (
	"context"
	
	"github.com/deadshvt/flight-booking-system/bonus-service/internal/converter"
	"github.com/deadshvt/flight-booking-system/bonus-service/internal/repository"
	"github.com/deadshvt/flight-booking-system/bonus-service/internal/usecase"
	pb "github.com/deadshvt/flight-booking-system/bonus-service/proto"
	"github.com/deadshvt/flight-booking-system/logger"

	"github.com/rs/zerolog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type BonusServiceServer struct {
	pb.UnimplementedBonusServiceServer
	Repo    repository.BonusRepository
	Usecase *usecase.BonusUsecase
	Logger  zerolog.Logger
}

func NewBonusServiceServer(repo repository.BonusRepository,
	usecase *usecase.BonusUsecase, logger zerolog.Logger) *BonusServiceServer {
	return &BonusServiceServer{
		Repo:    repo,
		Usecase: usecase,
		Logger:  logger,
	}
}

func (s *BonusServiceServer) GetPrivilegeWithHistory(ctx context.Context,
	req *pb.GetPrivilegeWithHistoryRequest) (*pb.GetPrivilegeWithHistoryResponse, error) {
	logger.LogWithParams(s.Logger, "Getting privilege...", struct {
		Username string
	}{req.Username})

	privilege, err := s.Repo.GetPrivilegeWithHistory(ctx, req.Username)
	if err != nil {
		s.Logger.Err(err).Msg("Failed to get privilege with history")
		return nil, status.Errorf(codes.Internal, "failed to get privilege with history: %v", err)
	}

	protoHistory := make([]*pb.History, len(privilege.History))
	for i, history := range privilege.History {
		protoHistory[i] = converter.HistoryFromEntityToProto(history)
	}

	entityPrivilege := privilege.Privilege

	logger.LogWithParams(s.Logger, "Got privilege with history", struct {
		Username string
		Status   string
		Balance  int32
		Count    int
	}{entityPrivilege.Username, entityPrivilege.Status, entityPrivilege.Balance,
		len(privilege.History)})

	return &pb.GetPrivilegeWithHistoryResponse{
		Privilege: converter.PrivilegeFromEntityToProto(&entityPrivilege),
		History:   protoHistory,
	}, nil
}

func (s *BonusServiceServer) GetPrivilege(ctx context.Context,
	req *pb.GetPrivilegeRequest) (*pb.GetPrivilegeResponse, error) {
	logger.LogWithParams(s.Logger, "Getting privilege...", struct {
		Username string
	}{req.Username})

	privilege, err := s.Repo.GetPrivilege(ctx, req.Username)
	if err != nil {
		s.Logger.Err(err).Msg("Failed to get privilege")
		return nil, status.Errorf(codes.Internal, "failed to get privilege: %v", err)
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
	privilege := req.Privilege

	logger.LogWithParams(s.Logger, "Creating privilege...", struct {
		Username string
		Balance  int32
		Status   string
	}{privilege.Username, privilege.Balance, privilege.Status})

	err := s.Repo.CreatePrivilege(ctx, converter.PrivilegeFromProtoToEntity(privilege))
	if err != nil {
		s.Logger.Err(err).Msg("Failed to create privilege")
		return nil, status.Errorf(codes.Internal, "failed to create privilege: %v", err)
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
	privilege := req.Privilege

	logger.LogWithParams(s.Logger, "Updating privilege...", struct {
		Username string
		Balance  int32
		Status   string
	}{privilege.Username, privilege.Balance, privilege.Status})

	err := s.Repo.UpdatePrivilege(ctx, converter.PrivilegeFromProtoToEntity(privilege))
	if err != nil {
		s.Logger.Err(err).Msg("Failed to update privilege")
		return nil, status.Errorf(codes.Internal, "failed to update privilege: %v", err)
	}

	logger.LogWithParams(s.Logger, "Updated privilege", struct {
		Username string
		Status   string
		Balance  int32
	}{privilege.Username, privilege.Status, privilege.Balance})

	return &pb.UpdatePrivilegeResponse{}, nil
}

func (s *BonusServiceServer) CreateHistory(ctx context.Context,
	req *pb.CreateHistoryRequest) (*pb.CreateHistoryResponse, error) {
	history := req.History

	logger.LogWithParams(s.Logger, "Creating history...", struct {
		PrivilegeID   int32
		TicketUid     string
		Date          string
		BalanceDiff   int32
		OperationType string
	}{history.PrivilegeID, history.TicketUid, history.Date,
		history.BalanceDiff, history.OperationType})

	err := s.Repo.CreateHistory(ctx, converter.HistoryFromProtoToEntity(req.History))
	if err != nil {
		s.Logger.Err(err).Msg("Failed to create history")
		return nil, status.Errorf(codes.Internal, "failed to create history: %v", err)
	}

	logger.LogWithParams(s.Logger, "Created history", struct {
		PrivilegeID   int32
		TicketUid     string
		Date          string
		BalanceDiff   int32
		OperationType string
	}{history.PrivilegeID, history.TicketUid, history.Date,
		history.BalanceDiff, history.OperationType})

	return &pb.CreateHistoryResponse{}, nil
}
