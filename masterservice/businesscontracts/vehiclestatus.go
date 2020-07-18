package bucontracts

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type VehicleStatusBO struct {
	Id         uuid.UUID
	StatusType string
	StatusName string
	UpdatedAt  time.Time
}
