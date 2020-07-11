package bucontracts

import "time"

type OperatorLocationBO struct {
	Id         uint
	AddressId  uint
	OperatorId uint
	Primary    bool
	UpdateAt   time.Time
	Address    AddressBO
	Operator   *OperatorBO
}
