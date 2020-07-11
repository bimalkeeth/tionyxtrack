package business

import (
	"errors"
	"github.com/jinzhu/gorm"
	bu "tionyxtrack/masterservice/businesscontracts"
	ent "tionyxtrack/masterservice/entities"
)

type IContactTypes interface {
	CreateContactType(contactType bu.ContactTypeBO) (uint, error)
	UpdateContactType(contactType bu.ContactTypeBO) (bool, error)
	DeleteContactType(id uint) (bool, error)
	GetContactTypeById(id uint) (bu.ContactTypeBO, error)
	GetContactTypeByName(name string) (bu.ContactTypeBO, error)
	GetAll() ([]bu.ContactTypeBO, error)
	GetAllNames(namePart string) ([]bu.ContactTypeBO, error)
}

type ContactType struct{ Db *gorm.DB }

//-------------------------------------------
//Create instance to through above interface
//-------------------------------------------
func NewContactType(db *gorm.DB) *ContactType {
	return &ContactType{Db: db}
}

//-------------------------------------------
// Create Address type
//-------------------------------------------
func (c *ContactType) CreateContactType(contactType bu.ContactTypeBO) (uint, error) {

	contType := ent.TableContactType{ContactType: contactType.ContactType}
	c.Db.Create(&contType)
	return contType.ID, nil
}

//-------------------------------------------
//Update Contact Type
//-------------------------------------------
func (c *ContactType) UpdateContactType(contactType bu.ContactTypeBO) (bool, error) {

	contacttype := ent.TableContactType{}
	c.Db.First(&contacttype, contactType.Id)
	if contacttype.ID == 0 {
		return false, errors.New("contact type not found")
	}
	contacttype.ContactType = contactType.ContactType
	c.Db.Save(&contactType)
	return true, nil
}

//-------------------------------------------
// Delete Contact Type
//-------------------------------------------
func (c *ContactType) DeleteContactType(id uint) (bool, error) {
	found := ent.TableContactType{}
	c.Db.First(&found, id)
	if found.ID == 0 {
		return false, errors.New("contact type not found")
	}
	c.Db.Delete(&found)
	return true, nil

}

//-------------------------------------------
// Get Contact type by Id
//-------------------------------------------
func (c *ContactType) GetContactTypeById(id uint) (bu.ContactTypeBO, error) {

	contactTypes := &ent.TableContactType{}
	c.Db.First(&contactTypes, id)

	result := bu.ContactTypeBO{}
	if contactTypes.ID == 0 {
		return result, errors.New("record not found")
	}
	return bu.ContactTypeBO{ContactType: contactTypes.ContactType, Id: contactTypes.ID}, nil
}

//-------------------------------------------
// Get Contact Type by Name
//-------------------------------------------
func (c *ContactType) GetContactTypeByName(name string) (bu.ContactTypeBO, error) {

	contactType := ent.TableContactType{}
	c.Db.Where(&ent.TableContactType{ContactType: name}).First(&contactType)
	if contactType.ID == 0 {
		return bu.ContactTypeBO{}, errors.New("record not found")
	}
	return bu.ContactTypeBO{ContactType: contactType.ContactType, Id: contactType.ID}, nil
}

//-------------------------------------------
// Get All Contact type Name Value
//-------------------------------------------
func (c *ContactType) GetAll() ([]bu.ContactTypeBO, error) {

	var contactTypes []ent.TableContactType
	var result []bu.ContactTypeBO

	c.Db.Find(&contactTypes)
	for _, item := range contactTypes {
		result = append(result, bu.ContactTypeBO{ContactType: item.ContactType, Id: item.ID})
	}
	return result, nil

}

//-------------------------------------------
//Get all Names by name like
//-------------------------------------------
func (c *ContactType) GetAllNames(namePart string) ([]bu.ContactTypeBO, error) {

	var contactTypes []ent.TableContactType
	c.Db.Where("contacttype LIKE ?", "%"+namePart+"%").Find(&contactTypes)
	var result []bu.ContactTypeBO
	for _, item := range contactTypes {
		result = append(result, bu.ContactTypeBO{ContactType: item.ContactType, Id: item.ID})
	}
	return result, nil

}
