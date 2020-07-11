package operator

import (
	bs "tionyxtrack/masterservice/business"
	bu "tionyxtrack/masterservice/businesscontracts"
)

//----------------------------------------------------
//Create operation contact
//----------------------------------------------------
func (o *OprManager) CreateOperatorContact(contactId uint, operatorId uint, primary bool) (uint, error) {
	op := OpFac.New(bs.COperationContact).(*bs.OperatorContact)
	OpFac.Conn.Begin()
	res, err := op.CreateOperatorContact(contactId, operatorId, primary)
	if err != nil {
		OpFac.Conn.Rollback()
		return res, err
	}
	OpFac.Conn.Commit()
	return res, err
}

//----------------------------------------------------
//Update operation contact
//----------------------------------------------------
func (o *OprManager) UpdateOperatorContact(id uint, contactId uint, operatorId uint, primary bool) (bool, error) {
	op := OpFac.New(bs.COperationContact).(*bs.OperatorContact)
	OpFac.Conn.Begin()
	res, err := op.UpdateOperatorContact(id, contactId, operatorId, primary)
	if err != nil {
		OpFac.Conn.Rollback()
		return res, err
	}
	OpFac.Conn.Commit()
	return res, err
}

//----------------------------------------------------
//Delete operation contact
//----------------------------------------------------
func (o *OprManager) DeleteOperatorContact(id uint) (bool, error) {
	op := OpFac.New(bs.COperationContact).(*bs.OperatorContact)
	OpFac.Conn.Begin()
	res, err := op.DeleteOperatorContact(id)
	if err != nil {
		OpFac.Conn.Rollback()
		return res, err
	}
	OpFac.Conn.Commit()
	return res, err
}

//----------------------------------------------------
//Get all operation contact
//----------------------------------------------------
func (o *OprManager) GetAllContactsByOperator(operatorId uint) ([]bu.OperatorContactsBO, error) {
	op := OpFac.New(bs.COperationContact).(*bs.OperatorContact)
	res, err := op.GetAllContactsByOperator(operatorId)
	return res, err
}
