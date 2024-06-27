package converter

import (
	"github.com/deadshvt/flight-booking-system/bonus-service/internal/entity"
	pb "github.com/deadshvt/flight-booking-system/bonus-service/proto"
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

func HistoryFromEntityToProto(history *entity.History) *pb.History {
	return &pb.History{
		ID:            history.ID,
		PrivilegeID:   history.PrivilegeID,
		Date:          history.Date,
		TicketUid:     history.TicketUid,
		BalanceDiff:   history.BalanceDiff,
		OperationType: history.OperationType,
	}
}

func HistoryFromProtoToEntity(history *pb.History) *entity.History {
	return &entity.History{
		ID:            history.ID,
		PrivilegeID:   history.PrivilegeID,
		Date:          history.Date,
		TicketUid:     history.TicketUid,
		BalanceDiff:   history.BalanceDiff,
		OperationType: history.OperationType,
	}
}
