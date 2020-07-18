package bucontracts

import uuid "github.com/satori/go.uuid"

type RegionBO struct {
	Id         uuid.UUID
	Region     string
	RegionName string
}
