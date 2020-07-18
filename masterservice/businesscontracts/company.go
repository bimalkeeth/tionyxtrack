package bucontracts

import uuid "github.com/satori/go.uuid"

type CompanyBO struct {
	Id        uuid.UUID
	Name      string
	AddressId uuid.UUID
	ContactId uuid.UUID
	Address   AddressBO
	Contact   ContactBO
}
