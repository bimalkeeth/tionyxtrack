package bucontracts

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type VehicleTrackRegBO struct {
	Id           uuid.UUID
	RegisterDate time.Time
	Duration     int
	ExpiredDate  time.Time
	Active       bool
	VehicleId    uuid.UUID
	UpdatedAt    time.Time
	Vehicle      VehicleBO
}
