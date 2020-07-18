package bucontracts

import uuid "github.com/satori/go.uuid"

type StateBO struct {
	Id        uuid.UUID
	Name      string
	CountryId uuid.UUID
}
