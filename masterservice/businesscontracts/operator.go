package bucontracts

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type OperatorBO struct {
	Id         uuid.UUID
	Name       string
	SurName    string
	Active     bool
	DrivingLic string
	UpdateAt   time.Time
	Locations  []*OperatorLocationBO
	Contacts   []*OperatorContactsBO
	Vehicles   []*VehicleBO
}
