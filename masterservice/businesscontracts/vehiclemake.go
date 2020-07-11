package bucontracts

import "time"

type VehicleMakeBO struct {
	Id        uint
	Make      string
	CountryId uint
	UpdateAt  time.Time
	Country   CountryBO
}
