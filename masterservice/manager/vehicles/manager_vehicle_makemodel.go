package vehicles

import (
	uuid "github.com/satori/go.uuid"
	bs "tionyxtrack/masterservice/business"
	bu "tionyxtrack/masterservice/businesscontracts"
)

//---------------------------------------------
//Create vehicle make
//---------------------------------------------
func (v *VehicleManager) CreateVehicleMake(bo bu.VehicleMakeBO) (uuid.UUID, error) {
	vh := vehicleFac.New(bs.CVehicleMake).(*bs.VehicleMake)
	vehicleFac.Conn.Begin()
	res, err := vh.CreateVehicleMake(bo)
	if err != nil {
		vehicleFac.Conn.Rollback()
		return res, err
	}
	vehicleFac.Conn.Commit()
	return res, err
}

//--------------------------------------------
//Update vehicle make
//--------------------------------------------
func (v *VehicleManager) UpdateVehicleMake(bo bu.VehicleMakeBO) (bool, error) {
	vh := vehicleFac.New(bs.CVehicleMake).(*bs.VehicleMake)
	vehicleFac.Conn.Begin()
	res, err := vh.UpdateVehicleMake(bo)
	if err != nil {
		vehicleFac.Conn.Rollback()
		return res, err
	}
	vehicleFac.Conn.Commit()
	return res, err
}

//--------------------------------------------
//Delete Vehicle make
//--------------------------------------------
func (v *VehicleManager) DeleteVehicleMake(id uuid.UUID) (bool, error) {
	vh := vehicleFac.New(bs.CVehicleMake).(*bs.VehicleMake)
	vehicleFac.Conn.Begin()
	res, err := vh.DeleteVehicleMake(id)
	if err != nil {
		vehicleFac.Conn.Rollback()
		return res, err
	}
	vehicleFac.Conn.Commit()
	return res, err
}

//------------------------------------------
//Get all vehicle make
//------------------------------------------
func (v *VehicleManager) GetAllVehicleMake() ([]bu.VehicleMakeBO, error) {
	vh := vehicleFac.New(bs.CVehicleMake).(*bs.VehicleMake)
	res, err := vh.GetAllVehicleMake()
	return res, err
}

//------------------------------------------
//Get  vehicle make by id
//------------------------------------------
func (v *VehicleManager) GetVehicleMakeById(id uuid.UUID) (bu.VehicleMakeBO, error) {
	vh := vehicleFac.New(bs.CVehicleMake).(*bs.VehicleMake)
	res, err := vh.GetVehicleMakeById(id)
	return res, err
}

//------------------------------------------
//Create Vehicle Model
//------------------------------------------
func (v *VehicleManager) CreateVehicleModel(bo bu.VehicleModelBO) (uuid.UUID, error) {
	vh := vehicleFac.New(bs.CVehicleModel).(*bs.VehicleModel)
	vehicleFac.Conn.Begin()
	res, err := vh.CreateVehicleModel(bo)
	if err != nil {
		vehicleFac.Conn.Rollback()
		return res, err
	}
	vehicleFac.Conn.Commit()
	return res, err
}

//------------------------------------------
//Update vehicle model
//------------------------------------------
func (v *VehicleManager) UpdateVehicleModel(bo bu.VehicleModelBO) (bool, error) {
	vh := vehicleFac.New(bs.CVehicleModel).(*bs.VehicleModel)
	vehicleFac.Conn.Begin()
	res, err := vh.UpdateVehicleModel(bo)
	if err != nil {
		vehicleFac.Conn.Rollback()
		return res, err
	}
	vehicleFac.Conn.Commit()
	return res, err
}

//------------------------------------------
//Delete vehicle model
//------------------------------------------
func (v *VehicleManager) DeleteVehicleModel(id uuid.UUID) (bool, error) {
	vh := vehicleFac.New(bs.CVehicleModel).(*bs.VehicleModel)
	vehicleFac.Conn.Begin()
	res, err := vh.DeleteVehicleModel(id)
	if err != nil {
		vehicleFac.Conn.Rollback()
		return res, err
	}
	vehicleFac.Conn.Commit()
	return res, err
}

//------------------------------------------
//Get all models by make
//------------------------------------------
func (v *VehicleManager) GetAllModelByMake(id uuid.UUID) ([]bu.VehicleModelBO, error) {
	vh := vehicleFac.New(bs.CVehicleModel).(*bs.VehicleModel)
	res, err := vh.GetAllModelByMake(id)
	return res, err

}

//-------------------------------------------
//Get Model by Id
//-------------------------------------------
func (v *VehicleManager) GetModelById(id uuid.UUID) (bu.VehicleModelBO, error) {
	vh := vehicleFac.New(bs.CVehicleModel).(*bs.VehicleModel)
	res, err := vh.GetModelById(id)
	return res, err
}
