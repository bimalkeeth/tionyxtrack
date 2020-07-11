package business

import (
	"errors"
	"github.com/jinzhu/gorm"
	bu "tionyxtrack/masterservice/businesscontracts"
	ent "tionyxtrack/masterservice/entities"
)

type IFleetLocation interface {
	CreateFleetLocation(fleetId uint, addressId uint, primary bool) (uint, error)
	UpdateFleetLocation(id uint, fleetId uint, addressId uint, primary bool) (bool, error)
	DeleteFleetLocation(id uint) (bool, error)
	GetLocationByFleetId(fleetId uint) ([]bu.FleetAddressBO, error)
}

type FleetLocation struct {
	Db *gorm.DB
}

func NewFleetLocation(db *gorm.DB) FleetLocation {
	return FleetLocation{Db: db}
}

//-----------------------------------------------------------
//Create Fleet location
//-----------------------------------------------------------
func (f *FleetLocation) CreateFleetLocation(fleetId uint, addressId uint, primary bool) (uint, error) {
	flContact := ent.TableFleetLocation{FleetId: fleetId, AddressId: addressId, Primary: primary}
	f.Db.Create(&flContact)
	return flContact.ID, nil
}

//----------------------------------------------------------
//Update fleet location
//----------------------------------------------------------
func (f *FleetLocation) UpdateFleetLocation(id uint, fleetId uint, addressId uint, primary bool) (bool, error) {

	if primary {
		setFLPrimaryOff(f)
	}
	fleetLoc := ent.TableFleetLocation{}
	f.Db.First(&fleetLoc, id)
	if fleetLoc.ID == 0 {
		return false, errors.New("fleet location can not be found")
	}
	fleetLoc.AddressId = addressId
	fleetLoc.FleetId = addressId
	fleetLoc.Primary = primary
	f.Db.Save(&fleetLoc)
	return true, nil
}

func setFLPrimaryOff(f *FleetLocation) {
	fleetLocation := &ent.TableFleetLocation{}
	f.Db.Where("primary = ?", true).First(&fleetLocation)
	if fleetLocation.ID > 0 {
		fleetLocation.Primary = false
		f.Db.Save(&fleetLocation)
	}
}

//----------------------------------------------------------
//Delete fleet location
//----------------------------------------------------------
func (f *FleetLocation) DeleteFleetLocation(id uint) (bool, error) {
	fleetLoc := ent.TableFleetLocation{}
	f.Db.First(&fleetLoc, id)
	if fleetLoc.ID == 0 {
		return false, errors.New("fleet location can not be found")
	}
	f.Db.Delete(&fleetLoc)
	return true, nil
}

//----------------------------------------------------------
//Get Fleet location by Fleet Id
//----------------------------------------------------------

func (f *FleetLocation) GetLocationByFleetId(fleetId uint) ([]bu.FleetAddressBO, error) {
	var fleetLocation []ent.TableFleetLocation
	var fleetLocationResult []bu.FleetAddressBO
	f.Db.Preload("Address").Preload("Fleet").Where(&ent.TableFleetLocation{FleetId: fleetId}).Find(&fleetLocation)

	for _, item := range fleetLocation {

		fleetLocationResult = append(fleetLocationResult, bu.FleetAddressBO{
			Id:        item.ID,
			FleetId:   item.FleetId,
			AddressId: item.AddressId,
			Primary:   item.Primary,
			Fleet: bu.FleetBO{
				Id:                   item.Fleet.ID,
				UpdatedAt:            item.Fleet.UpdatedAt,
				FleetCode:            item.Fleet.FleetCode,
				Name:                 item.Fleet.Name,
				SurName:              item.Fleet.SurName,
				OtherName:            item.Fleet.OtherName,
				DateRegistered:       item.Fleet.DateRegistered,
				RegistrationDuration: item.Fleet.RegistrationDuration,
			},
			Address: bu.AddressBO{
				Id:            item.Address.ID,
				Address:       item.Address.Address,
				UpdatedAt:     item.Address.UpdatedAt,
				AddressTypeId: item.Address.AddressTypeId,
				Location:      item.Address.Location,
				CountryId:     item.Address.CountryId,
				StateId:       item.Address.StateId,
				Suburb:        item.Address.Suburb,
				Street:        item.Address.Street,
			},
		})
	}
	return fleetLocationResult, nil
}
