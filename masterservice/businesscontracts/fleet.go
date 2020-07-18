package bucontracts

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type FleetBO struct {
	Id                   uuid.UUID
	UpdatedAt            time.Time
	FleetCode            string
	Name                 string
	SurName              string
	OtherName            string
	DateRegistered       time.Time
	RegistrationDuration float64
	FleetContactId       uuid.UUID
	FleetLocationId      uuid.UUID
	CountryId            uuid.UUID
	FleetContacts        []ContactBO
	Address              []AddressBO
}
