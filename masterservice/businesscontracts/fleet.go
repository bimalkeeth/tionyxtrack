package bucontracts

import "time"

type FleetBO struct {
	Id                   uint
	UpdatedAt            time.Time
	FleetCode            string
	Name                 string
	SurName              string
	OtherName            string
	DateRegistered       time.Time
	RegistrationDuration float64
	FleetContactId       uint
	FleetLocationId      uint
	CountryId            uint
	FleetContacts        []ContactBO
	Address              []AddressBO
}
