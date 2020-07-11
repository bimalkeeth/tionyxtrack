package bucontracts

import "time"

type VehicleBO struct {
	Id            uint
	ModelId       uint
	MakeId        uint
	Registration  string
	FleetId       uint
	StatusId      uint
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
