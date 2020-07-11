package bucontracts

import "time"

type VehicleStatusBO struct {
	Id         uint
	StatusType string
	StatusName string
	UpdatedAt  time.Time
}
