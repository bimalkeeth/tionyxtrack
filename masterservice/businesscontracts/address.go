package bucontracts

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type AddressBO struct {
	Id            uuid.UUID
	Address       string
	Street        string
	Suburb        string
	StateId       uuid.UUID
	CountryId     uuid.UUID
	AddressTypeId uuid.UUID
	Location      string
	AddressType   AddressTypeBO
	State         StateBO
	Country       CountryBO
	UpdatedAt     time.Time
}
