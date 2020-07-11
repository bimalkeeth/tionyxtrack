package bucontracts

import "time"

type CountryBO struct {
	Id          uint
	CountryName string
	RegionId    uint
	Region      RegionBO
	States      []StateBO
	UpdatedAt   time.Time
}
