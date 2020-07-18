package vehicles

import (
	uuid "github.com/satori/go.uuid"
	bs "tionyxtrack/masterservice/business"
	bu "tionyxtrack/masterservice/businesscontracts"
)

//-------------------------------------------
//Create Vehicle location
//-------------------------------------------
func (v *VehicleManager) CreateVehicleLocation(ad bu.VehicleAddressBO) (uuid.UUID, error) {
	vh := vehicleFac.New(bs.CVehicleLocation).(*bs.VehicleLocation)
	vehicleFac.Conn.Begin()
	res, err := vh.CreateVehicleLocation(ad)
	if err != nil {
		vehicleFac.Conn.Rollback()
		return res, err
	}
	vehicleFac.Conn.Commit()
	return res, nil
}

//-------------------------------------------
//Update vehicle location
//-------------------------------------------
func (v *VehicleManager) UpdateVehicleLocation(ad bu.VehicleAddressBO) (bool, error) {
	vh := vehicleFac.New(bs.CVehicleLocation).(*bs.VehicleLocation)
	vehicleFac.Conn.Begin()
	res, err := vh.UpdateVehicleLocation(ad)
	if err != nil {
		vehicleFac.Conn.Rollback()
		return res, err
	}
	vehicleFac.Conn.Commit()
	return res, nil
}

//-------------------------------------------
//Delete vehicle location
//-------------------------------------------
func (v *VehicleManager) DeleteVehicleLocation(id uuid.UUID) (bool, error) {
	vh := vehicleFac.New(bs.CVehicleLocation).(*bs.VehicleLocation)
	vehicleFac.Conn.Begin()
	res, err := vh.DeleteVehicleLocation(id)
	if err != nil {
		vehicleFac.Conn.Rollback()
		return res, err
	}
	vehicleFac.Conn.Commit()
	return res, nil
}

//------------------------------------------
//Get location by vehicle id
//------------------------------------------
func (v *VehicleManager) GetVehicleLocationByVehicle(vehicleId uuid.UUID) ([]bu.VehicleAddressBO, error) {
	vh := vehicleFac.New(bs.CVehicleLocation).(*bs.VehicleLocation)
	res, err := vh.GetVehicleLocationByVehicle(vehicleId)
	return res, err
}
