package bucontracts

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type ContactBO struct {
	Id            uuid.UUID
	Contact       string
	ContactTypeId uuid.UUID
	UpdatedAt     time.Time
}
