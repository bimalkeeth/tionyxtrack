package bucontracts

import "time"

type OperatorBO struct {
	Id         uint
	Name       string
	SurName    string
	Active     bool
	DrivingLic string
	UpdateAt   time.Time
	Locations  []*OperatorLocationBO
	Contacts   []*OperatorContactsBO
	Vehicles   []*VehicleBO
}
