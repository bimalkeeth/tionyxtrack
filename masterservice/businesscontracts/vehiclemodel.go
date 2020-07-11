package bucontracts

import "time"

type VehicleModelBO struct {
	Id          uint
	ModelName   string
	Description string
	MakeId      uint
	UpdatedAt   time.Time
	Make        VehicleMakeBO
}
