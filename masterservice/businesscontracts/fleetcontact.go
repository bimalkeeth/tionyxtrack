package bucontracts

import uuid "github.com/satori/go.uuid"

type FleetContactBO struct {
	Id        uuid.UUID
	FleetId   uuid.UUID
	ContactId uuid.UUID
	Primary   bool
	Fleet     FleetBO
	Contact   ContactBO
}
