package contacts

import (
	uuid "github.com/satori/go.uuid"
	bs "tionyxtrack/masterservice/business"
	bu "tionyxtrack/masterservice/businesscontracts"
)

//---------------------------------------
//Create contact
//---------------------------------------
func (c *ContactManager) CreateContact(con bu.ContactBO) (uuid.UUID, error) {

	conFactory.Conn.Begin()
	contact := conFactory.New(bs.CContact).(*bs.Contact)
	id, err := contact.CreateContact(con)
	if err != nil {
		conFactory.Conn.Rollback()
		return uuid.Nil, err
	}
	conFactory.Conn.Commit()
	return id, nil
}

//--------------------------------------
//Update Contact
//--------------------------------------
func (c *ContactManager) UpdateContact(con bu.ContactBO) (bool, error) {

	conFactory.Conn.Begin()
	contact := conFactory.New(bs.CContact).(*bs.Contact)
	result, err := contact.UpdateContact(con)
	if err != nil {
		conFactory.Conn.Rollback()
		return false, err
	}
	conFactory.Conn.Commit()
	return result, nil
}

//------------------------------------
//Delete Contact
//------------------------------------
func (c *ContactManager) DeleteContact(id uuid.UUID) (bool, error) {

	conFactory.Conn.Begin()
	contact := conFactory.New(bs.CContact).(*bs.Contact)
	result, err := contact.DeleteContact(id)
	if err != nil {
		conFactory.Conn.Rollback()
		return false, err
	}
	conFactory.Conn.Commit()
	return result, nil
}

//------------------------------------
//Get Contact by Id
//------------------------------------
func (c *ContactManager) ContactById(Id uuid.UUID) (bu.ContactBO, error) {

	contact := conFactory.New(bs.CContact).(*bs.Contact)
	co, err := contact.ContactById(Id)
	if err != nil {

		return bu.ContactBO{}, err
	}
	return co, nil
}

//------------------------------------
//Create Address
//------------------------------------
func (c *ContactManager) CreateAddress(add bu.AddressBO) (uuid.UUID, error) {

	conFactory.Conn.Begin()
	address := conFactory.New(bs.CAddress).(*bs.Address)
	res, err := address.CreateAddress(add)
	if err != nil {
		conFactory.Conn.Rollback()
		return uuid.Nil, err
	}
	conFactory.Conn.Commit()
	return res, nil
}

//------------------------------------
//Update Address
//------------------------------------
func (c *ContactManager) UpdateAddress(add bu.AddressBO) (bool, error) {

	conFactory.Conn.Begin()
	address := conFactory.New(bs.CAddress).(*bs.Address)
	res, err := address.UpdateAddress(add)
	if err != nil {
		conFactory.Conn.Rollback()
		return false, err
	}
	conFactory.Conn.Commit()
	return res, nil
}

//------------------------------------
//Delete Address
//------------------------------------
func (c *ContactManager) DeleteAddress(id uuid.UUID) (bool, error) {

	conFactory.Conn.Begin()
	address := conFactory.New(bs.CAddress).(*bs.Address)
	res, err := address.DeleteAddress(id)
	if err != nil {
		conFactory.Conn.Rollback()
		return false, err
	}
	conFactory.Conn.Commit()
	return res, nil
}

//------------------------------------
//Get Address by AddressId
//------------------------------------
func (c *ContactManager) GetAddressById(id uuid.UUID) (bu.AddressBO, error) {

	address := conFactory.New(bs.CAddress).(*bs.Address)
	result, err := address.GetAddressById(id)
	if err != nil {
		return bu.AddressBO{}, err
	}
	return result, nil
}

//-------------------------------------
//Get Address By Name
//-------------------------------------
func (c *ContactManager) GetAddressByName(name string) ([]bu.AddressBO, error) {
	address := conFactory.New(bs.CAddress).(*bs.Address)
	result, err := address.GetAddressByName(name)
	if err != nil {
		return result, err
	}
	return result, nil
}
