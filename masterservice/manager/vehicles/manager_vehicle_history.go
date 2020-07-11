package vehicles

import (
	bs "tionyxtrack/masterservice/business"
	bu "tionyxtrack/masterservice/businesscontracts"
)

//-----------------------------------------------
//Create Vehicle History
//-----------------------------------------------
func (v *VehicleManager) CreateVehicleHistory(history bu.VehicleHistoryBO) (uint, error) {
	vh := vehicleFac.New(bs.CVehicleHistory).(*bs.VehicleHistory)
	vehicleFac.Conn.Begin()
	res, err := vh.CreateVehicleHistory(history)
	if err != nil {
		vehicleFac.Conn.Rollback()
		return res, err
	}
	vehicleFac.Conn.Commit()
	return res, nil
}

//-----------------------------------------------
//Update Vehicle History
//-----------------------------------------------
func (v *VehicleManager) UpdateVehicleHistory(history bu.VehicleHistoryBO) (bool, error) {
	vh := vehicleFac.New(bs.CVehicleHistory).(*bs.VehicleHistory)
	vehicleFac.Conn.Begin()
	res, err := vh.UpdateVehicleHistory(history)
	if err != nil {
		vehicleFac.Conn.Rollback()
		return res, err
	}
	vehicleFac.Conn.Commit()
	return res, nil
}

//---------------------------------------------
//Delete Vehicle History
//---------------------------------------------
func (v *VehicleManager) DeleteVehicleHistory(id uint) (bool, error) {
	vh := vehicleFac.New(bs.CVehicleHistory).(*bs.VehicleHistory)
	vehicleFac.Conn.Begin()
	res, err := vh.DeleteVehicleHistory(id)
	if err != nil {
		vehicleFac.Conn.Rollback()
		return res, err
	}
	vehicleFac.Conn.Commit()
	return res, nil
}

//--------------------------------------------
//Get Vehicle By Id
//--------------------------------------------
func (v *VehicleManager) GetVehicleHistoryByVehicleId(vehicleId uint) ([]bu.VehicleHistoryBO, error) {
	vh := vehicleFac.New(bs.CVehicleHistory).(*bs.VehicleHistory)
	res, err := vh.GetVehicleHistoryByVehicleId(vehicleId)
	return res, err
}
