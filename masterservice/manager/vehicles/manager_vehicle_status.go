package vehicles

import (
	bs "tionyxtrack/masterservice/business"
	bu "tionyxtrack/masterservice/businesscontracts"
)

//------------------------------------------------
//Create vehicle status
//------------------------------------------------
func (v *VehicleManager) CreateVehicleStatus(bo bu.VehicleStatusBO) (uint, error) {
	vr := vehicleFac.New(bs.CVhStatus).(bs.VehicleStatus)
	vehicleFac.Conn.Begin()
	res, err := vr.CreateVehicleStatus(bo)
	if err != nil {
		vehicleFac.Conn.Rollback()
		return res, err
	}
	vehicleFac.Conn.Commit()
	return res, nil
}

//-------------------------------------------------
//Update vehicle status
//-------------------------------------------------
func (v *VehicleManager) UpdateVehicleStatus(bo bu.VehicleStatusBO) (bool, error) {

	vr := vehicleFac.New(bs.CVhStatus).(bs.VehicleStatus)
	vehicleFac.Conn.Begin()
	res, err := vr.UpdateVehicleStatus(bo)
	if err != nil {
		vehicleFac.Conn.Rollback()
		return res, err
	}
	vehicleFac.Conn.Commit()
	return res, nil
}

//-------------------------------------------------
//Delete vehicle status
//-------------------------------------------------
func (v *VehicleManager) DeleteVehicleStatus(id uint) (bool, error) {
	vr := vehicleFac.New(bs.CVhStatus).(bs.VehicleStatus)
	vehicleFac.Conn.Begin()
	res, err := vr.DeleteVehicleStatus(id)
	if err != nil {
		vehicleFac.Conn.Rollback()
		return res, err
	}
	vehicleFac.Conn.Commit()
	return res, nil
}

//------------------------------------------------
//Get all vehicle status
//------------------------------------------------
func (v *VehicleManager) GetAllVehicleStatus() ([]bu.VehicleStatusBO, error) {
	vr := vehicleFac.New(bs.CVhStatus).(bs.VehicleStatus)
	res, err := vr.GetAllVehicleStatus()
	return res, err
}
