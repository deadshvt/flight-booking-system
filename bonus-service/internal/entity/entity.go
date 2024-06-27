package entity

type Privilege struct {
	ID       int32
	Username string
	Status   string
	Balance  int32
}

type History struct {
	ID            int32
	PrivilegeID   int32
	TicketUid     string
	Date          string
	BalanceDiff   int32
	OperationType string
}

type PrivilegeWithHistory struct {
	Privilege Privilege
	History   []*History
}
