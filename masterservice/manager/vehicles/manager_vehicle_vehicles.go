package vehicles

import (
	"fmt"
	"time"
	bs "tionyxtrack/masterservice/business"
	bu "tionyxtrack/masterservice/businesscontracts"
)

//----------------------------------------
//Create Vehicle
//----------------------------------------
func (v *VehicleManager) CreateVehicle(vehicle bu.VehicleBO) (uint, error) {
	veh := vehicleFac.New(bs.CVehicle).(*bs.Vehicle)
	vehicleFac.Conn.Begin()
	res, err := veh.CreateVehicle(vehicle)

	if err != nil {
		vehicleFac.Conn.Rollback()
		return res, err
	}
	vehicleFac.Conn.Commit()
	return res, err
}

//-----------------------------------------
//Update Vehicle
//-----------------------------------------
func (v *VehicleManager) UpdateVehicle(vehicle bu.VehicleBO) (bool, error) {
	veh := vehicleFac.New(bs.CVehicle).(*bs.Vehicle)
	vehicleFac.Conn.Begin()

	res, err := veh.UpdateVehicle(vehicle)
	if err != nil {
		vehicleFac.Conn.Rollback()
		return res, err
	}
	vhExist, err := veh.GetVehicleById(vehicle.Id)
	if err != nil {
		vehicleFac.Conn.Rollback()
		return false, err
	}
	if vhExist.StatusId != vehicle.StatusId {
		his := vehicleFac.New(bs.CVehicleHistory).(bs.VehicleHistory)
		_, err := his.CreateVehicleHistory(bu.VehicleHistoryBO{VehicleId: vehicle.Id,
			Description:  fmt.Sprintf("Status change from  %d to %d", vhExist.StatusId, vehicle.StatusId),
			FromStatusId: vhExist.StatusId,
			ToStatusId:   vehicle.StatusId,
			OfficerName:  vehicle.OfficeName,
			ChangeDate:   time.Time{},
		})
		if err != nil {
			vehicleFac.Conn.Rollback()
			return false, err
		}
	}

	vehicleFac.Conn.Commit()
	return res, nil
}

//-----------------------------------------
//Delete Vehicle
//-----------------------------------------
func (v *VehicleManager) DeleteVehicle(vehicleId uint) (bool, error) {
	veh := vehicleFac.New(bs.CVehicle).(*bs.Vehicle)
	vehicleFac.Conn.Begin()
	res, err := veh.DeleteVehicle(vehicleId)
	if err != nil {
		vehicleFac.Conn.Rollback()
		return res, err
	}
	vehicleFac.Conn.Commit()
	return res, nil
}

//-----------------------------------------
//Get Vehicle by Id
//-----------------------------------------
func (v *VehicleManager) GetVehicleById(vehicleId uint) (bu.VehicleBO, error) {
	veh := vehicleFac.New(bs.CVehicle).(*bs.Vehicle)
	res, err := veh.GetVehicleById(vehicleId)
	return res, err
}

//-----------------------------------------
//Get Vehicle by Registration
//-----------------------------------------
func (v *VehicleManager) GetVehicleByRegistration(registration string) (bu.VehicleBO, error) {
	veh := vehicleFac.New(bs.CVehicle).(*bs.Vehicle)
	res, err := veh.GetVehicleByRegistration(registration)
	return res, err
}

//-----------------------------------------
//Get Vehicle by Fleet Id
//-----------------------------------------
func (v *VehicleManager) GetVehiclesByFleetId(fleetId uint) ([]bu.VehicleBO, error) {
	veh := vehicleFac.New(bs.CVehicle).(*bs.Vehicle)
	res, err := veh.GetVehiclesByFleetId(fleetId)
	return res, err
}
