package business

import (
	"errors"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	bu "tionyxtrack/masterservice/businesscontracts"
	ent "tionyxtrack/masterservice/entities"
)

type IVehicleLocation interface {
	CreateVehicleLocation(ad bu.VehicleAddressBO) (uuid.UUID, error)
	UpdateVehicleLocation(ad bu.VehicleAddressBO) (bool, error)
	DeleteVehicleLocation(id uuid.UUID) (bool, error)
	GetVehicleLocationByVehicle(vehicleId uuid.UUID) ([]bu.VehicleAddressBO, error)
}

type VehicleLocation struct {
	Db *gorm.DB
}

func NewVehicleLocation(db *gorm.DB) VehicleLocation {
	return VehicleLocation{Db: db}
}

//----------------------------------------------------------
// Create Vehicle location
//----------------------------------------------------------
func (l *VehicleLocation) CreateVehicleLocation(ad bu.VehicleAddressBO) (uuid.UUID, error) {

	vehicleLocation := ent.TableVehicleLocation{VehicleId: ad.VehicleId, AddressId: ad.AddressId, Primary: ad.Primary}
	l.Db.Create(&vehicleLocation)
	return vehicleLocation.ID, nil
}

//----------------------------------------------------------
// update Vehicle location
//----------------------------------------------------------
func (l *VehicleLocation) UpdateVehicleLocation(ad bu.VehicleAddressBO) (bool, error) {

	if ad.Primary {
		setVHPrimaryOff(l)
	}
	vehicleLocation := ent.TableVehicleLocation{}
	l.Db.First(&vehicleLocation, ad.Id)
	if vehicleLocation.ID == uuid.Nil {
		return false, errors.New("vehicle location could not be found")
	}
	vehicleLocation.AddressId = ad.AddressId
	vehicleLocation.VehicleId = ad.VehicleId
	vehicleLocation.Primary = ad.Primary

	l.Db.Save(&vehicleLocation)
	return true, nil
}

func setVHPrimaryOff(f *VehicleLocation) {
	vhLoc := &ent.TableVehicleLocation{}
	f.Db.Where("primary = ?", true).First(&vhLoc)
	if vhLoc.ID != uuid.Nil {
		vhLoc.Primary = false
		f.Db.Save(&vhLoc)
	}
}

//----------------------------------------------------------
// delete Vehicle location
//----------------------------------------------------------
func (l *VehicleLocation) DeleteVehicleLocation(id uuid.UUID) (bool, error) {
	vehicleLocation := ent.TableVehicleLocation{}
	l.Db.First(&vehicleLocation, id)
	if vehicleLocation.ID == uuid.Nil {
		return false, errors.New("vehicle location could not be found")
	}
	l.Db.Delete(&vehicleLocation)
	return true, nil
}

//----------------------------------------------------------
// get Vehicle location by vehicleid
//----------------------------------------------------------
func (l *VehicleLocation) GetVehicleLocationByVehicle(vehicleId uuid.UUID) ([]bu.VehicleAddressBO, error) {
	var vehicleLocation []ent.TableVehicleLocation
	var locationResult []bu.VehicleAddressBO

	l.Db.Preload("Address").Preload("Vehicle").Where("vehicleid = ?", vehicleId).Find(&vehicleLocation)
	for _, item := range vehicleLocation {
		locationResult = append(locationResult, bu.VehicleAddressBO{
			VehicleId: item.VehicleId,
			AddressId: item.AddressId,
			Id:        item.ID,
			UpdateAt:  item.UpdatedAt,
			Address: bu.AddressBO{
				Id:            item.Address.ID,
				Address:       item.Address.Address,
				Street:        item.Address.Street,
				Suburb:        item.Address.Suburb,
				StateId:       item.Address.StateId,
				CountryId:     item.Address.CountryId,
				AddressTypeId: item.Address.AddressTypeId,
				Location:      item.Address.Location,
				UpdatedAt:     item.Address.UpdatedAt,
			},
		})
	}
	return locationResult, nil
}
