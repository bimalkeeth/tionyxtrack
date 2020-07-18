package bucontracts

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type OperatorLocationBO struct {
	Id         uuid.UUID
	AddressId  uuid.UUID
	OperatorId uuid.UUID
	Primary    bool
	UpdateAt   time.Time
	Address    AddressBO
	Operator   *OperatorBO
}
