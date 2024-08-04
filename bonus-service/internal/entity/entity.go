package entity

import "time"

type Privilege struct {
	ID       int32
	Username string
	Status   string
	Balance  int32
}

type Operation struct {
	ID            int32
	PrivilegeID   int32
	TicketUid     string
	Date          time.Time
	BalanceDiff   int32
	OperationType string
}

type PrivilegeWithHistory struct {
	Privilege *Privilege
	History   []*Operation
}
