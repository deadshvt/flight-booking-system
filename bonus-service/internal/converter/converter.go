package converter

import (
	"github.com/deadshvt/flight-booking-system/bonus-service/internal/entity"
	pb "github.com/deadshvt/flight-booking-system/bonus-service/proto"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func PrivilegeFromEntityToProto(privilege *entity.Privilege) *pb.Privilege {
	return &pb.Privilege{
		ID:       privilege.ID,
		Username: privilege.Username,
		Balance:  privilege.Balance,
		Status:   privilege.Status,
	}
}

func PrivilegeFromProtoToEntity(privilege *pb.Privilege) *entity.Privilege {
	return &entity.Privilege{
		ID:       privilege.ID,
		Username: privilege.Username,
		Balance:  privilege.Balance,
		Status:   privilege.Status,
	}
}

func OperationFromEntityToProto(operation *entity.Operation) *pb.Operation {
	return &pb.Operation{
		ID:            operation.ID,
		PrivilegeID:   operation.PrivilegeID,
		TicketUid:     operation.TicketUid,
		Date:          timestamppb.New(operation.Date),
		BalanceDiff:   operation.BalanceDiff,
		OperationType: operation.OperationType,
	}
}

func OperationFromProtoToEntity(operation *pb.Operation) *entity.Operation {
	return &entity.Operation{
		ID:            operation.ID,
		PrivilegeID:   operation.PrivilegeID,
		TicketUid:     operation.TicketUid,
		Date:          operation.Date.AsTime(),
		BalanceDiff:   operation.BalanceDiff,
		OperationType: operation.OperationType,
	}
}
