package fleets

import (
	uuid "github.com/satori/go.uuid"
	bs "tionyxtrack/masterservice/business"
	bu "tionyxtrack/masterservice/businesscontracts"
)

//---------------------------------------------
//Create Fleet
//---------------------------------------------
func (f *FleetManager) CreateFleet(bo bu.FleetBO) (bu.FleetBO, error) {
	op := flFac.New(bs.CFleet).(*bs.Fleet)
	flFac.Conn.Begin()
	res, err := op.CreateFleet(bo)
	if err != nil {
		flFac.Conn.Rollback()
		return res, err
	}
	flFac.Conn.Commit()
	return res, err
}

//---------------------------------------------
//Update Fleet
//---------------------------------------------
func (f *FleetManager) UpdateFleet(bo bu.FleetBO) (bool, error) {
	op := flFac.New(bs.CFleet).(*bs.Fleet)
	flFac.Conn.Begin()
	res, err := op.UpdateFleet(bo)
	if err != nil {
		flFac.Conn.Rollback()
		return res, err
	}
	flFac.Conn.Commit()
	return res, err
}

//---------------------------------------------
//Delete Fleet
//---------------------------------------------
func (f *FleetManager) DeleteFleet(id uuid.UUID) (bool, error) {
	op := flFac.New(bs.CFleet).(*bs.Fleet)
	flFac.Conn.Begin()
	res, err := op.DeleteFleet(id)
	if err != nil {
		flFac.Conn.Rollback()
		return res, err
	}
	flFac.Conn.Commit()
	return res, err
}

//---------------------------------------------
//Get Fleet By Id
//---------------------------------------------
func (f *FleetManager) GetFleetById(id uuid.UUID) (bu.FleetBO, error) {
	op := flFac.New(bs.CFleet).(*bs.Fleet)
	res, err := op.GetFleetById(id)
	return res, err
}

//----------------------------------------------
//Create fleet Contact
//----------------------------------------------
func (f *FleetManager) CreateFleetContact(fleetId uuid.UUID, contactId uuid.UUID, primary bool) (uuid.UUID, error) {
	op := flFac.New(bs.CFleetContact).(*bs.FleetContact)
	flFac.Conn.Begin()
	res, err := op.CreateFleetContact(fleetId, contactId, primary)
	if err != nil {
		flFac.Conn.Rollback()
		return res, err
	}
	flFac.Conn.Commit()
	return res, err
}

//----------------------------------------------
//Update fleet Contact
//----------------------------------------------
func (f *FleetManager) UpdateFleetContact(id uuid.UUID, fleetId uuid.UUID, contactId uuid.UUID, primary bool) (bool, error) {
	op := flFac.New(bs.CFleetContact).(*bs.FleetContact)
	flFac.Conn.Begin()
	res, err := op.UpdateFleetContact(id, fleetId, contactId, primary)
	if err != nil {
		flFac.Conn.Rollback()
		return res, err
	}
	flFac.Conn.Commit()
	return res, err
}

//----------------------------------------------
//Delete fleet Contact
//----------------------------------------------
func (f *FleetManager) DeleteFleetContact(id uuid.UUID) (bool, error) {
	op := flFac.New(bs.CFleetContact).(*bs.FleetContact)
	flFac.Conn.Begin()
	res, err := op.DeleteFleetContact(id)
	if err != nil {
		flFac.Conn.Rollback()
		return res, err
	}
	flFac.Conn.Commit()
	return res, err
}

//----------------------------------------------
//Get fleet Contact by fleetId
//----------------------------------------------
func (f *FleetManager) GetContactByFleetId(fleetId uuid.UUID) ([]bu.FleetContactBO, error) {
	op := flFac.New(bs.CFleetContact).(*bs.FleetContact)
	res, err := op.GetContactByFleetId(fleetId)
	return res, err
}

//----------------------------------------------
//Create fleet location
//----------------------------------------------
func (f *FleetManager) CreateFleetLocation(fleetId uuid.UUID, addressId uuid.UUID, primary bool) (uuid.UUID, error) {
	op := flFac.New(bs.CFleetLocation).(*bs.FleetLocation)
	flFac.Conn.Begin()
	res, err := op.CreateFleetLocation(fleetId, addressId, primary)
	if err != nil {
		flFac.Conn.Rollback()
		return res, err
	}
	flFac.Conn.Commit()
	return res, err
}

//----------------------------------------------
//Update fleet location
//----------------------------------------------
func (f *FleetManager) UpdateFleetLocation(id uuid.UUID, fleetId uuid.UUID, addressId uuid.UUID, primary bool) (bool, error) {
	op := flFac.New(bs.CFleetLocation).(*bs.FleetLocation)
	flFac.Conn.Begin()
	res, err := op.UpdateFleetLocation(id, fleetId, addressId, primary)
	if err != nil {
		flFac.Conn.Rollback()
		return res, err
	}
	flFac.Conn.Commit()
	return res, err
}

//----------------------------------------------
//Delete fleet location
//----------------------------------------------
func (f *FleetManager) DeleteFleetLocation(id uuid.UUID) (bool, error) {
	op := flFac.New(bs.CFleetLocation).(*bs.FleetLocation)
	flFac.Conn.Begin()
	res, err := op.DeleteFleetLocation(id)
	if err != nil {
		flFac.Conn.Rollback()
		return res, err
	}
	flFac.Conn.Commit()
	return res, err
}

//----------------------------------------------
//Get fleet location by fleetId
//----------------------------------------------
func (f *FleetManager) GetLocationByFleetId(fleetId uuid.UUID) ([]bu.FleetAddressBO, error) {
	op := flFac.New(bs.CFleetLocation).(*bs.FleetLocation)
	res, err := op.GetLocationByFleetId(fleetId)
	return res, err
}
