package vehicles

import (
	bs "tionyxtrack/masterservice/business"
	bu "tionyxtrack/masterservice/businesscontracts"
)

//-----------------------------------------
//Create vehicle operator bound
//-----------------------------------------
func (v *VehicleManager) CreateVehicleOpBound(op bu.VehicleOperatorBoundBO) (uint, error) {
	vh := vehicleFac.New(bs.CVhOperatorBound).(*bs.VehicleOprBound)
	vehicleFac.Conn.Begin()
	res, err := vh.CreateVehicleOpBound(op)
	if err != nil {
		vehicleFac.Conn.Rollback()
		return res, err
	}
	vehicleFac.Conn.Commit()
	return res, nil
}

//-----------------------------------------
//Update vehicle operator bound
//-----------------------------------------
func (v *VehicleManager) UpdateVehicleOpBound(op bu.VehicleOperatorBoundBO) (bool, error) {
	vh := vehicleFac.New(bs.CVhOperatorBound).(*bs.VehicleOprBound)
	vehicleFac.Conn.Begin()
	res, err := vh.UpdateVehicleOpBound(op)
	if err != nil {
		vehicleFac.Conn.Rollback()
		return res, err
	}
	vehicleFac.Conn.Commit()
	return res, nil
}

//-----------------------------------------
//Delete vehicle operator bound
//-----------------------------------------
func (v *VehicleManager) DeleteVehicleOpBound(id uint) (bool, error) {
	vh := vehicleFac.New(bs.CVhOperatorBound).(*bs.VehicleOprBound)
	vehicleFac.Conn.Begin()
	res, err := vh.DeleteVehicleOpBound(id)
	if err != nil {
		vehicleFac.Conn.Rollback()
		return res, err
	}
	vehicleFac.Conn.Commit()
	return res, nil
}
