package bucontracts

type VehicleOperatorBoundBO struct {
	Id         uint
	OperatorId uint
	VehicleId  uint
	Active     bool
	Operator   *OperatorBO
	Vehicle    *VehicleBO
}
