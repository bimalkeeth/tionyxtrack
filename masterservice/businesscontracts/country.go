package bucontracts

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type CountryBO struct {
	Id          uuid.UUID
	CountryName string
	RegionId    uuid.UUID
	Region      RegionBO
	States      []StateBO
	UpdatedAt   time.Time
}
