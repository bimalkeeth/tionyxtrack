package vehicles

import (
	bs "tionyxtrack/masterservice/business"
	bu "tionyxtrack/masterservice/businesscontracts"
)

//---------------------------------------------
//Create vehicle registration
//---------------------------------------------
func (v *VehicleManager) CreateVehicleReg(bo bu.VehicleTrackRegBO) (uint, error) {
	vr := vehicleFac.New(bs.CVhRegistration).(bs.VehicleReg)
	vehicleFac.Conn.Begin()
	result, err := vr.CreateVehicleReg(bo)
	if err != nil {
		vehicleFac.Conn.Rollback()
		return result, err
	}
	vehicleFac.Conn.Commit()
	return result, err
}

//---------------------------------------------
//Update vehicle registration
//---------------------------------------------
func (v *VehicleManager) UpdateVehicleReg(bo bu.VehicleTrackRegBO) (bool, error) {

	vr := vehicleFac.New(bs.CVhRegistration).(bs.VehicleReg)
	vehicleFac.Conn.Begin()
	result, err := vr.UpdateVehicleReg(bo)
	if err != nil {
		vehicleFac.Conn.Rollback()
		return result, err
	}
	vehicleFac.Conn.Commit()
	return result, err
}

//--------------------------------------------
//Delete vehicle registration
//--------------------------------------------
func (v *VehicleManager) DeleteVehicleReg(id uint) (bool, error) {
	vr := vehicleFac.New(bs.CVhRegistration).(bs.VehicleReg)
	vehicleFac.Conn.Begin()
	result, err := vr.DeleteVehicleReg(id)
	if err != nil {
		vehicleFac.Conn.Rollback()
		return result, err
	}
	vehicleFac.Conn.Commit()
	return result, err
}

//--------------------------------------------
//Get all vehicle registration by vehicle id
//--------------------------------------------
func (v *VehicleManager) GetAllRegistrationsByVehicleId(id uint) ([]bu.VehicleTrackRegBO, error) {
	vr := vehicleFac.New(bs.CVhRegistration).(bs.VehicleReg)
	result, err := vr.GetAllRegistrationsByVehicleId(id)
	return result, err
}

//-------------------------------------------
//Get all vehicle registration
//-------------------------------------------
func (v *VehicleManager) GetActiveRegistrationsByVehicleId(id uint) (bu.VehicleTrackRegBO, error) {
	vr := vehicleFac.New(bs.CVhRegistration).(bs.VehicleReg)
	result, err := vr.GetActiveRegistrationsByVehicleId(id)
	return result, err
}
