package business

import (
	"errors"
	"github.com/jinzhu/gorm"
	bu "tionyxtrack/masterservice/businesscontracts"
	ent "tionyxtrack/masterservice/entities"
)

type IVehicleReg interface {
	CreateVehicleReg(bo bu.VehicleTrackRegBO) (uint, error)
	UpdateVehicleReg(bo bu.VehicleTrackRegBO) (bool, error)
	DeleteVehicleReg(id uint) (bool, error)
	GetAllRegistrationsByVehicleId(id uint) ([]bu.VehicleTrackRegBO, error)
	GetActiveRegistrationsByVehicleId(id uint) (bu.VehicleTrackRegBO, error)
}

type VehicleReg struct {
	Db *gorm.DB
}

func NewVhReg(db *gorm.DB) *VehicleReg {
	return &VehicleReg{Db: db}
}

//---------------------------------------------
//Create vehicle registration
//---------------------------------------------
func (r *VehicleReg) CreateVehicleReg(bo bu.VehicleTrackRegBO) (uint, error) {
	vhReg := ent.TableVehicleTrackReg{}
	vhReg.VehicleId = bo.VehicleId
	vhReg.Active = bo.Active
	vhReg.Duration = bo.Duration
	vhReg.ExpiredDate = bo.ExpiredDate
	vhReg.RegisterDate = bo.RegisterDate
	r.Db.Create(&vhReg)
	return vhReg.ID, nil

}

//---------------------------------------------
//Update vehicle registration
//---------------------------------------------
func (r *VehicleReg) UpdateVehicleReg(bo bu.VehicleTrackRegBO) (bool, error) {
	vhReg := ent.TableVehicleTrackReg{}
	r.Db.First(&vhReg, bo.Id)
	if vhReg.ID == 0 {
		return false, errors.New("registration not found")
	}
	vhReg.RegisterDate = bo.RegisterDate
	vhReg.ExpiredDate = bo.ExpiredDate
	vhReg.Duration = bo.Duration
	vhReg.Active = bo.Active
	vhReg.VehicleId = bo.Id
	r.Db.Save(&vhReg)
	return true, nil
}

//---------------------------------------------
//Delete vehicle registration
//---------------------------------------------
func (r *VehicleReg) DeleteVehicleReg(id uint) (bool, error) {
	vhReg := ent.TableVehicleTrackReg{}
	r.Db.First(&vhReg, id)
	if vhReg.ID == 0 {
		return false, errors.New("registration not found")
	}
	r.Db.Delete(&vhReg)
	return true, nil
}

//---------------------------------------------
//Get vehicle registrations by VehicleId
//---------------------------------------------
func (r *VehicleReg) GetAllRegistrationsByVehicleId(id uint) ([]bu.VehicleTrackRegBO, error) {
	var results []bu.VehicleTrackRegBO
	var vhResult []ent.TableVehicleTrackReg
	r.Db.Preload("Vehicle").Where("vehicleid = ?", id).Find(&vhResult)
	for _, vh := range vhResult {
		results = append(results, bu.VehicleTrackRegBO{
			Id:           vh.ID,
			RegisterDate: vh.RegisterDate,
			Duration:     vh.Duration,
			ExpiredDate:  vh.ExpiredDate,
			Active:       vh.Active,
			VehicleId:    vh.VehicleId,
			UpdatedAt:    vh.UpdatedAt,
			Vehicle: bu.VehicleBO{
				Id:           vh.Vehicle.ID,
				UpdatedAt:    vh.Vehicle.UpdatedAt,
				Registration: vh.Vehicle.Registration,
				StatusId:     vh.Vehicle.StatusId,
				MakeId:       vh.Vehicle.MakeId,
				ModelId:      vh.Vehicle.ModelId,
			},
		})
	}
	return results, nil
}

//---------------------------------------------
//Get active vehicle registrations by VehicleId
//---------------------------------------------
func (r *VehicleReg) GetActiveRegistrationsByVehicleId(id uint) (bu.VehicleTrackRegBO, error) {

	vh := ent.TableVehicleTrackReg{}
	r.Db.Preload("Vehicle").Where("vehicleid = ? and active= ?", id, true).First(&vh)

	return bu.VehicleTrackRegBO{
		Id:           vh.ID,
		RegisterDate: vh.RegisterDate,
		Duration:     vh.Duration,
		ExpiredDate:  vh.ExpiredDate,
		Active:       vh.Active,
		VehicleId:    vh.VehicleId,
		UpdatedAt:    vh.UpdatedAt,
		Vehicle: bu.VehicleBO{
			Id:           vh.Vehicle.ID,
			UpdatedAt:    vh.Vehicle.UpdatedAt,
			Registration: vh.Vehicle.Registration,
			StatusId:     vh.Vehicle.StatusId,
			MakeId:       vh.Vehicle.MakeId,
			ModelId:      vh.Vehicle.ModelId,
		},
	}, nil
}
