package business

import (
	"errors"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	bu "tionyxtrack/masterservice/businesscontracts"
	ent "tionyxtrack/masterservice/entities"
)

type IOperatorContact interface {
	CreateOperatorContact(contactId uuid.UUID, operatorId uuid.UUID, primary bool) (uuid.UUID, error)
	UpdateOperatorContact(id uuid.UUID, contactId uuid.UUID, operatorId uuid.UUID, primary bool) (bool, error)
	DeleteOperatorContact(id uuid.UUID) (bool, error)
	GetAllContactsByOperator(operatorId uuid.UUID) ([]bu.OperatorContactsBO, error)
}

type OperatorContact struct {
	Db *gorm.DB
}

func NewOperatorContact(db *gorm.DB) OperatorContact {
	return OperatorContact{Db: db}
}

//-----------------------------------------------------
//Create operator contact
//-----------------------------------------------------
func (o *OperatorContact) CreateOperatorContact(contactId uuid.UUID, operatorId uuid.UUID, primary bool) (uuid.UUID, error) {
	opContact := ent.TableVehicleOperatorContacts{ContactId: contactId, OperatorId: operatorId, Primary: primary}
	o.Db.Create(&opContact)
	return opContact.ID, nil
}

//-----------------------------------------------------
//Update operator contact
//-----------------------------------------------------
func (o *OperatorContact) UpdateOperatorContact(id uuid.UUID, contactId uuid.UUID, operatorId uuid.UUID, primary bool) (bool, error) {

	if primary {
		setOCPrimaryOff(o)
	}

	opContact := ent.TableVehicleOperatorContacts{}
	o.Db.First(&opContact, id)
	if opContact.ID == uuid.Nil {
		return false, errors.New("operator contact not found")
	}
	opContact.OperatorId = operatorId
	opContact.ContactId = contactId
	opContact.Primary = primary
	o.Db.Save(&opContact)
	return true, nil
}

func setOCPrimaryOff(f *OperatorContact) {
	oprCon := &ent.TableVehicleOperatorContacts{}
	f.Db.Where("primary = ?", true).First(&oprCon)
	if oprCon.ID != uuid.Nil {
		oprCon.Primary = false
		f.Db.Save(&oprCon)
	}
}

//-----------------------------------------------------
//Delete operator contact
//-----------------------------------------------------
func (o *OperatorContact) DeleteOperatorContact(id uuid.UUID) (bool, error) {
	opContact := ent.TableVehicleOperatorContacts{}
	o.Db.First(&opContact, id)
	if opContact.ID == uuid.Nil {
		return false, errors.New("operator contact not found")
	}
	o.Db.Delete(&opContact)
	return true, nil
}

//-----------------------------------------------------
//get contacts for operator
//-----------------------------------------------------
func (o *OperatorContact) GetAllContactsByOperator(operatorId uuid.UUID) ([]bu.OperatorContactsBO, error) {

	var operators []ent.TableVehicleOperatorContacts
	var oprResults []bu.OperatorContactsBO

	o.Db.Preload("Contact").Preload("Operator").
		Where(&ent.TableVehicleOperatorContacts{OperatorId: operatorId}).Find(&operators)

	for _, item := range operators {

		oprResults = append(oprResults, bu.OperatorContactsBO{
			Id:         item.ID,
			ContactId:  item.ContactId,
			OperatorId: item.OperatorId,
			Primary:    item.Primary,
			Contact: bu.ContactBO{
				Id:            item.ContactId,
				ContactTypeId: item.Contact.ContactTypeId,
				UpdatedAt:     item.Contact.UpdatedAt,
				Contact:       item.Contact.Contact,
			},
		})
	}
	return oprResults, nil
}
