package bucontracts

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type VehicleHistoryBO struct {
	Id           uuid.UUID
	VehicleId    uuid.UUID
	ChangeDate   time.Time
	Description  string
	FromStatusId uuid.UUID
	ToStatusId   uuid.UUID
	OfficerName  string
	FromStatus   VehicleStatusBO
	ToStatus     VehicleStatusBO
}
