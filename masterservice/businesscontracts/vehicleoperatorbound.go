package bucontracts

import uuid "github.com/satori/go.uuid"

type VehicleOperatorBoundBO struct {
	Id         uuid.UUID
	OperatorId uuid.UUID
	VehicleId  uuid.UUID
	Active     bool
	Operator   *OperatorBO
	Vehicle    *VehicleBO
}
