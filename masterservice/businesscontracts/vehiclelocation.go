package bucontracts

import "time"

type VehicleAddressBO struct {
	Id        uint
	AddressId uint
	VehicleId uint
	Primary   bool
	UpdateAt  time.Time
	Address   AddressBO
}
