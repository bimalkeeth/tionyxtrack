package bucontracts

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type VehicleAddressBO struct {
	Id        uuid.UUID
	AddressId uuid.UUID
	VehicleId uuid.UUID
	Primary   bool
	UpdateAt  time.Time
	Address   AddressBO
}
