package business

import (
	"errors"
	"github.com/jinzhu/gorm"
	bu "tionyxtrack/masterservice/businesscontracts"
	ent "tionyxtrack/masterservice/entities"
)

type IAddress interface {
	CreateAddress(address bu.AddressBO) (uint, error)
	UpdateAddress(address bu.AddressBO) (bool, error)
	DeleteAddress(id uint) (bool, error)
	GetAddressById(id uint) (bu.AddressBO, error)
	GetAddressByName(name string) ([]bu.AddressBO, error)
}

type Address struct {
	Db *gorm.DB
}

func NewAddress(db *gorm.DB) *Address { return &Address{Db: db} }

//---------------------------------------------------
//Create address
//---------------------------------------------------
func (a *Address) CreateAddress(address bu.AddressBO) (uint, error) {

	addr := ent.TableAddress{CountryId: address.CountryId,
		AddressTypeId: address.AddressTypeId,
		StateId:       address.StateId,
		Location:      address.Location,
		Address:       address.Address,
		Street:        address.Street,
		Suburb:        address.Suburb}
	a.Db.Create(&addr)
	return addr.ID, nil
}

//---------------------------------------------------
//Update Address
//---------------------------------------------------
func (a *Address) UpdateAddress(address bu.AddressBO) (bool, error) {

	addr := &ent.TableAddress{}
	a.Db.First(addr, address.Id)
	if addr.ID == 0 {
		return false, errors.New("address not found")
	}
	addr.Suburb = address.Suburb
	addr.Street = address.Street
	addr.Address = address.Address
	addr.Location = address.Location
	addr.StateId = address.StateId
	addr.AddressTypeId = address.AddressTypeId
	addr.CountryId = address.CountryId
	a.Db.Save(&addr)
	return true, nil
}

//---------------------------------------------------
//Delete Address
//---------------------------------------------------
func (a *Address) DeleteAddress(id uint) (bool, error) {

	address := &ent.TableAddress{}
	a.Db.First(&address, id)

	if address.ID == 0 {
		return false, errors.New("the record not exists in the storage")
	}
	a.Db.Delete(&address)
	return true, nil
}

//----------------------------------------------------
//Get Address by Id
//----------------------------------------------------
func (a *Address) GetAddressById(id uint) (bu.AddressBO, error) {

	address := &ent.TableAddress{}
	a.Db.First(&address, id)

	result := bu.AddressBO{}
	if address.ID == 0 {
		return result, errors.New("record not found")
	}
	return bu.AddressBO{CountryId: address.CountryId,
		AddressTypeId: address.AddressTypeId,
		StateId:       address.StateId,
		Location:      address.Location,
		Address:       address.Address,
		Street:        address.Street,
		Suburb:        address.Suburb,
		Id:            address.ID,
	}, nil
}

//----------------------------------------------------
//Get Address by Name
//----------------------------------------------------
func (a *Address) GetAddressByName(name string) ([]bu.AddressBO, error) {

	var address []ent.TableAddress
	a.Db.Where("name LIKE ?", "%"+name+"%").Find(&address)
	var result []bu.AddressBO
	for _, item := range address {
		result = append(result, bu.AddressBO{CountryId: item.CountryId,
			AddressTypeId: item.AddressTypeId,
			StateId:       item.StateId,
			Location:      item.Location,
			Address:       item.Address,
			Street:        item.Street,
			Suburb:        item.Suburb,
			Id:            item.ID,
		})
	}
	return result, nil
}
