package bucontracts

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type VehicleMakeBO struct {
	Id        uuid.UUID
	Make      string
	CountryId uuid.UUID
	UpdateAt  time.Time
	Country   CountryBO
}
