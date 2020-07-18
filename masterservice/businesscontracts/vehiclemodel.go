package bucontracts

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type VehicleModelBO struct {
	Id          uuid.UUID
	ModelName   string
	Description string
	MakeId      uuid.UUID
	UpdatedAt   time.Time
	Make        VehicleMakeBO
}
