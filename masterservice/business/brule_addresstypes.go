package business

import (
	"errors"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	bu "tionyxtrack/masterservice/businesscontracts"
	ent "tionyxtrack/masterservice/entities"
)

type IAddressTypes interface {
	CreateAddressType(addressType bu.AddressTypeBO) (uuid.UUID, error)
	UpdateAddressType(addressType bu.AddressTypeBO) (bool, error)
	DeleteAddressType(id uuid.UUID) (bool, error)
	GetAddressTypeById(id uuid.UUID) (bu.AddressTypeBO, error)
	GetAddressTypeByName(name string) (bu.AddressTypeBO, error)
	GetAll() ([]bu.AddressTypeBO, error)
	GetAllNames(namePart string) ([]bu.AddressTypeBO, error)
}
type AddressType struct{ Db *gorm.DB }

func NewAddressType(db *gorm.DB) *AddressType { return &AddressType{Db: db} }

//-----------------------------------------
// Create Address type
//-----------------------------------------
func (at *AddressType) CreateAddressType(addressType bu.AddressTypeBO) (uuid.UUID, error) {

	addType := ent.TableAddressType{AddressType: addressType.Name}
	at.Db.Create(&addType)
	return addType.ID, nil
}

//----------------------------------------
//Update Address type
//----------------------------------------
func (at *AddressType) UpdateAddressType(addressType bu.AddressTypeBO) (bool, error) {

	addressTypes := &ent.TableAddressType{}
	at.Db.First(&addressTypes, addressType.Id)
	if addressTypes.ID == uuid.Nil {
		return false, errors.New("address type cannot be found")
	}
	addressTypes.Name = addressType.Name
	at.Db.Save(&addressTypes)
	return true, nil
}

//-----------------------------------------
// Delete Address type
//-----------------------------------------
func (at *AddressType) DeleteAddressType(id uuid.UUID) (bool, error) {

	addressTypes := &ent.TableAddressType{}
	at.Db.First(&addressTypes, id)

	if addressTypes.ID == uuid.Nil {
		return false, errors.New("the record not exists in the storage")
	}
	at.Db.Delete(&addressTypes)
	return true, nil
}

//------------------------------------------
//Get Address type by Address Id
//------------------------------------------
func (at *AddressType) GetAddressTypeById(id uuid.UUID) (bu.AddressTypeBO, error) {

	addressTypes := &ent.TableAddressType{}
	at.Db.First(&addressTypes, id)

	result := bu.AddressTypeBO{}
	if addressTypes.ID == uuid.Nil {
		return result, errors.New("record not found")
	}
	return bu.AddressTypeBO{Name: addressTypes.Name, Id: addressTypes.ID}, nil
}

//------------------------------------------
//Get Address by Address name
//------------------------------------------
func (at *AddressType) GetAddressTypeByName(name string) (bu.AddressTypeBO, error) {

	addressType := ent.TableAddressType{}
	at.Db.Where(&ent.TableAddressType{Name: name}).First(&addressType)
	if addressType.ID == uuid.Nil {
		return bu.AddressTypeBO{}, errors.New("record not found")
	}
	return bu.AddressTypeBO{Name: addressType.Name, Id: addressType.ID}, nil
}

//------------------------------------------
//Get All Address
//------------------------------------------
func (at *AddressType) GetAll() ([]bu.AddressTypeBO, error) {

	var addressTypes []ent.TableAddressType
	var result []bu.AddressTypeBO

	at.Db.Find(&addressTypes)
	for _, item := range addressTypes {
		result = append(result, bu.AddressTypeBO{Name: item.Name, Id: item.ID})
	}
	return result, nil
}

//------------------------------------------
//Get all address by name part
//------------------------------------------
func (at *AddressType) GetAllNames(namePart string) ([]bu.AddressTypeBO, error) {
	var addressTypes []ent.TableAddressType
	at.Db.Where("name LIKE ?", "%"+namePart+"%").Find(&addressTypes)
	var result []bu.AddressTypeBO
	for _, item := range addressTypes {
		result = append(result, bu.AddressTypeBO{Name: item.Name, Id: item.ID})
	}
	return result, nil

}
