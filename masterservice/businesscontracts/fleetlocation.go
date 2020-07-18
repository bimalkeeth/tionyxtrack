package bucontracts

import uuid "github.com/satori/go.uuid"

type FleetAddressBO struct {
	Id        uuid.UUID
	FleetId   uuid.UUID
	AddressId uuid.UUID
	Primary   bool
	Fleet     FleetBO
	Address   AddressBO
}
