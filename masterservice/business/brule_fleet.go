package business

import (
	"errors"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	bu "tionyxtrack/masterservice/businesscontracts"
	en "tionyxtrack/masterservice/entities"
)

type IFleet interface {
	CreateFleet(bo bu.FleetBO) (bu.FleetBO, error)
	UpdateFleet(bo bu.FleetBO) (bool, error)
	DeleteFleet(id uuid.UUID) (bool, error)
	GetFleetById(id uuid.UUID) (bu.FleetBO, error)
}

type Fleet struct{ Db *gorm.DB }

func NewFleet(db *gorm.DB) Fleet {
	return Fleet{Db: db}
}

//--------------------------------------------------------
//Create Fleet
//--------------------------------------------------------
func (f *Fleet) CreateFleet(bo bu.FleetBO) (bu.FleetBO, error) {

	fleet := &en.TableFleet{
		FleetCode:            bo.FleetCode,
		Name:                 bo.Name,
		SurName:              bo.SurName,
		OtherName:            bo.OtherName,
		DateRegistered:       bo.DateRegistered,
		RegistrationDuration: bo.RegistrationDuration,
		CountryId:            bo.CountryId,
	}
	f.Db.Create(fleet)
	bo.Id = fleet.ID
	return bo, nil
}

//----------------------------------------------------------
//Update Fleet
//----------------------------------------------------------
func (f *Fleet) UpdateFleet(bo bu.FleetBO) (bool, error) {

	fleet := &en.TableFleet{}
	f.Db.First(fleet, bo.Id)
	if fleet.ID == uuid.Nil {
		return false, errors.New("fleet can not be found")
	}

	fleet.SurName = bo.SurName
	fleet.Name = bo.Name
	fleet.FleetCode = bo.FleetCode
	fleet.OtherName = bo.OtherName
	fleet.DateRegistered = bo.DateRegistered
	fleet.RegistrationDuration = bo.RegistrationDuration
	fleet.CountryId = bo.CountryId
	f.Db.Save(fleet)

	return true, nil
}

//----------------------------------------------------------
//Update Fleet
//----------------------------------------------------------
func (f *Fleet) DeleteFleet(id uuid.UUID) (bool, error) {
	fleet := &en.TableFleet{}
	f.Db.First(fleet, id)
	if fleet.ID == uuid.Nil {
		return false, errors.New("fleet can not be found")
	}
	f.Db.Delete(fleet)
	return true, nil
}

//----------------------------------------------------------
//Get Fleet Id
//----------------------------------------------------------
func (f *Fleet) GetFleetById(id uuid.UUID) (bu.FleetBO, error) {
	fleet := &en.TableFleet{}
	f.Db.Preload("FleetContacts").
		Preload("FleetContacts.Contact").
		Preload("FleetLocations").
		Preload("FleetLocations.Address").
		First(fleet, id)
	if fleet.ID == uuid.Nil {
		return bu.FleetBO{}, errors.New("fleet can not be found")
	}

	result := bu.FleetBO{}
	result.Id = fleet.ID
	result.RegistrationDuration = fleet.RegistrationDuration
	result.DateRegistered = fleet.DateRegistered
	result.OtherName = fleet.OtherName
	result.FleetCode = fleet.FleetCode
	result.SurName = fleet.SurName
	result.Name = fleet.Name
	result.CountryId = fleet.CountryId

	result.FleetContacts = []bu.ContactBO{}
	for _, item := range fleet.FleetContacts {
		result.FleetContacts = append(result.FleetContacts, bu.ContactBO{
			Id:            item.Contact.ID,
			Contact:       item.Contact.Contact,
			ContactTypeId: item.Contact.ContactTypeId,
		})
	}
	result.Address = []bu.AddressBO{}
	for _, item := range fleet.FleetLocations {
		result.Address = append(result.Address, bu.AddressBO{
			Id:            item.Address.ID,
			Address:       item.Address.Address,
			Street:        item.Address.Street,
			Suburb:        item.Address.Suburb,
			StateId:       item.Address.StateId,
			CountryId:     item.Address.CountryId,
			AddressTypeId: item.Address.AddressTypeId,
			Location:      item.Address.Location,
		})
	}
	return result, nil
}
