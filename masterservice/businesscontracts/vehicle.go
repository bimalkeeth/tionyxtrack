package bucontracts

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type VehicleBO struct {
	Id            uuid.UUID
	ModelId       uuid.UUID
	MakeId        uuid.UUID
	Registration  string
	FleetId       uuid.UUID
	StatusId      uuid.UUID
	UpdatedAt     time.Time
	OfficeName    string
	VehicleModel  VehicleModelBO
	VehicleMake   VehicleMakeBO
	Fleet         FleetBO
	Status        VehicleStatusBO
	Locations     []VehicleAddressBO
	Operators     []VehicleOperatorBoundBO
	Registrations []VehicleTrackRegBO
}
