package bucontracts

import "time"

type VehicleTrackRegBO struct {
	Id           uint
	RegisterDate time.Time
	Duration     int
	ExpiredDate  time.Time
	Active       bool
	VehicleId    uint
	UpdatedAt    time.Time
	Vehicle      VehicleBO
}
