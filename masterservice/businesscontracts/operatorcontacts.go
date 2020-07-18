package bucontracts

import uuid "github.com/satori/go.uuid"

type OperatorContactsBO struct {
	Id         uuid.UUID
	ContactId  uuid.UUID
	OperatorId uuid.UUID
	Primary    bool
	Contact    ContactBO
}
