package business

import (
	"errors"
	"github.com/jinzhu/gorm"
	bu "tionyxtrack/masterservice/businesscontracts"
	ent "tionyxtrack/masterservice/entities"
)

type IVehicleMake interface {
	CreateVehicleMake(bo bu.VehicleMakeBO) (uint, error)
	UpdateVehicleMake(bo bu.VehicleMakeBO) (bool, error)
	DeleteVehicleMake(id uint) (bool, error)
	GetAllVehicleMake() ([]bu.VehicleMakeBO, error)
	GetVehicleMakeById(id uint) (bu.VehicleMakeBO, error)
}

type VehicleMake struct {
	Db *gorm.DB
}

func NewVehicleMake(db *gorm.DB) VehicleMake {
	return VehicleMake{Db: db}
}

//--------------------------------------------
//Create Vehicle Make
//--------------------------------------------
func (v *VehicleMake) CreateVehicleMake(bo bu.VehicleMakeBO) (uint, error) {

	vehicleMake := ent.TableVehicleMake{
		CountryId: bo.CountryId,
		Make:      bo.Make,
	}
	v.Db.Create(&vehicleMake)
	return vehicleMake.ID, nil
}

//--------------------------------------------
//Update Vehicle Make
//--------------------------------------------
func (v *VehicleMake) UpdateVehicleMake(bo bu.VehicleMakeBO) (bool, error) {

	vehicleMake := ent.TableVehicleMake{}
	v.Db.First(&vehicleMake, bo.Id)
	if vehicleMake.ID == 0 {
		return false, errors.New("vehicle make could not be found")
	}
	vehicleMake.CountryId = bo.CountryId
	vehicleMake.Make = bo.Make
	v.Db.Save(&vehicleMake)
	return true, nil
}

//--------------------------------------------
//Delete Vehicle Make
//--------------------------------------------
func (v *VehicleMake) DeleteVehicleMake(id uint) (bool, error) {
	vehicleMake := ent.TableVehicleMake{}
	v.Db.First(&vehicleMake, id)
	if vehicleMake.ID == 0 {
		return false, errors.New("vehicle make could not be found")
	}
	v.Db.Delete(&vehicleMake)
	return true, nil
}

//--------------------------------------------
//Get All Vehicle Make
//--------------------------------------------
func (v *VehicleMake) GetAllVehicleMake() ([]bu.VehicleMakeBO, error) {
	var vehicleMakes []ent.TableVehicleMake
	var vehicleResult []bu.VehicleMakeBO
	v.Db.Preload("Country").Find(&vehicleMakes)
	for _, item := range vehicleMakes {

		vehicleResult = append(vehicleResult, bu.VehicleMakeBO{
			Id:        item.ID,
			Make:      item.Make,
			CountryId: item.CountryId,
			UpdateAt:  item.UpdatedAt,
			Country: bu.CountryBO{
				Id:          item.Country.ID,
				RegionId:    item.Country.RegionId,
				CountryName: item.Country.CountryName,
			},
		})
	}
	return vehicleResult, nil
}

//--------------------------------------------
//Get Vehicle Make By Id
//--------------------------------------------
func (v *VehicleMake) GetVehicleMakeById(id uint) (bu.VehicleMakeBO, error) {
	vehicleMake := ent.TableVehicleMake{}
	v.Db.Preload("Country").First(&vehicleMake, id)
	if vehicleMake.ID == 0 {
		return bu.VehicleMakeBO{}, errors.New("vehicle make could not be found")
	}
	result := bu.VehicleMakeBO{
		Id:        vehicleMake.ID,
		Make:      vehicleMake.Make,
		CountryId: vehicleMake.CountryId,
		UpdateAt:  vehicleMake.UpdatedAt,
		Country: bu.CountryBO{
			Id:          vehicleMake.Country.ID,
			RegionId:    vehicleMake.Country.RegionId,
			CountryName: vehicleMake.Country.CountryName,
		},
	}
	return result, nil
}
