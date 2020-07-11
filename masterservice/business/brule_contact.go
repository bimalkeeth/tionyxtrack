package business

import (
	"errors"
	"github.com/jinzhu/gorm"
	bu "tionyxtrack/masterservice/businesscontracts"
	en "tionyxtrack/masterservice/entities"
)

type IContact interface {
	CreateContact(con bu.ContactBO) (uint, error)
	UpdateContact(con bu.ContactBO) (bool, error)
	DeleteContact(Id uint) (bool, error)
	ContactById(Id uint) (bu.ContactBO, error)
}

type Contact struct{ Db *gorm.DB }

func NewContact(db *gorm.DB) *Contact { return &Contact{Db: db} }

//--------------------------------------------
//Create Contact
//---------------------------------------------
func (c *Contact) CreateContact(con bu.ContactBO) (uint, error) {

	cont := en.TableContact{Contact: con.Contact, ContactTypeId: con.ContactTypeId}
	c.Db.Create(&cont)
	return cont.ID, nil
}

//--------------------------------------------
//Update Contact
//---------------------------------------------
func (c *Contact) UpdateContact(con bu.ContactBO) (bool, error) {
	contact := en.TableContact{}
	c.Db.First(&contact, con.Id)
	if contact.ID == 0 {
		return false, errors.New("contact id cannot be found")
	}
	contact.ContactTypeId = con.ContactTypeId
	contact.Contact = con.Contact
	c.Db.Save(&contact)
	return true, nil
}

//--------------------------------------------
//Delete Contact
//---------------------------------------------
func (c *Contact) DeleteContact(Id uint) (bool, error) {
	contact := en.TableContact{}
	c.Db.First(&contact, Id)
	if contact.ID == 0 {
		return false, errors.New("contact id cannot be found")
	}
	c.Db.Delete(&contact)
	return true, nil
}

//--------------------------------------------
//Get Contact By Id
//---------------------------------------------
func (c *Contact) ContactById(Id uint) (bu.ContactBO, error) {

	contact := en.TableContact{}
	c.Db.First(&contact, Id)
	return bu.ContactBO{Id: contact.ID, ContactTypeId: contact.ContactTypeId, Contact: contact.Contact, UpdatedAt: contact.UpdatedAt}, nil
}
