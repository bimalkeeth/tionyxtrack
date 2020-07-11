package business

import (
	"errors"
	"github.com/jinzhu/gorm"
	bu "tionyxtrack/masterservice/businesscontracts"
	ent "tionyxtrack/masterservice/entities"
)

type IFleetContact interface {
	CreateFleetContact(fleetId uint, contactId uint, primary bool) (uint, error)
	UpdateFleetContact(id uint, fleetId uint, contactId uint, primary bool) (bool, error)
	DeleteFleetContact(id uint) (bool, error)
	GetContactByFleetId(fleetId uint) ([]bu.FleetContactBO, error)
}
type FleetContact struct {
	Db *gorm.DB
}

func NewFleetContact(db *gorm.DB) FleetContact {
	return FleetContact{Db: db}
}

//-------------------------------------------------
// Create Fleet Contact
//-------------------------------------------------
func (f *FleetContact) CreateFleetContact(fleetId uint, contactId uint, primary bool) (uint, error) {
	fleetCon := ent.TableFleetContact{FleetId: fleetId, ContactId: contactId, Primary: primary}
	f.Db.Create(&fleetCon)
	return fleetCon.ID, nil
}

//-------------------------------------------------
// Update Fleet Contact
//-------------------------------------------------

func (f *FleetContact) UpdateFleetContact(id uint, fleetId uint, contactId uint, primary bool) (bool, error) {

	if primary {
		setFleetContactPrimaryOff(f)
	}

	fleetContact := &ent.TableFleetContact{}
	f.Db.First(fleetContact, id)
	if fleetContact.ID == 0 {
		return false, errors.New("could not find fleet contact")
	}
	fleetContact.ContactId = contactId
	fleetContact.FleetId = fleetId
	fleetContact.Primary = primary
	f.Db.Save(fleetContact)
	return true, nil
}

func setFleetContactPrimaryOff(f *FleetContact) {
	fleetContact := &ent.TableFleetContact{}
	f.Db.Where("primary = ?", true).First(&fleetContact)
	if fleetContact.ID > 0 {
		fleetContact.Primary = false
		f.Db.Save(&fleetContact)
	}
}

//-------------------------------------------------
// Delete fleet contact
//-------------------------------------------------

func (f *FleetContact) DeleteFleetContact(id uint) (bool, error) {
	fleetContact := &ent.TableFleetContact{}
	f.Db.First(fleetContact, id)
	if fleetContact.ID == 0 {
		return false, errors.New("could not find fleet contact")
	}
	f.Db.Delete(fleetContact)
	return true, nil
}

//-------------------------------------------------
// Get contact by fleet contact id
//-------------------------------------------------

func (f *FleetContact) GetContactByFleetId(fleetId uint) ([]bu.FleetContactBO, error) {

	var fleetContact []ent.TableFleetContact
	var fleetResult []bu.FleetContactBO
	f.Db.Preload("Contact").Preload("Fleet").Where(&ent.TableFleetContact{FleetId: fleetId}).Find(&fleetContact)

	for _, item := range fleetContact {

		fleetResult = append(fleetResult, bu.FleetContactBO{
			Id:        item.ID,
			FleetId:   item.FleetId,
			ContactId: item.ContactId,
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
			Contact: bu.ContactBO{
				Id:            item.Contact.ID,
				Contact:       item.Contact.Contact,
				UpdatedAt:     item.Contact.UpdatedAt,
				ContactTypeId: item.Contact.ContactTypeId,
			},
		})
	}
	return fleetResult, nil
}
