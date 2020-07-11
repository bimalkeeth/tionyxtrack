package bucontracts

import "time"

type VehicleHistoryBO struct {
	Id           uint
	VehicleId    uint
	ChangeDate   time.Time
	Description  string
	FromStatusId uint
	ToStatusId   uint
	OfficerName  string
	FromStatus   VehicleStatusBO
	ToStatus     VehicleStatusBO
}
