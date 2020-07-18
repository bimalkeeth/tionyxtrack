package operator

import (
	uuid "github.com/satori/go.uuid"
	bs "tionyxtrack/masterservice/business"
	bu "tionyxtrack/masterservice/businesscontracts"
)

//----------------------------------------------------
//Create operation contact
//----------------------------------------------------
func (o *OprManager) CreateOperatorContact(contactId uuid.UUID, operatorId uuid.UUID, primary bool) (uuid.UUID, error) {
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
func (o *OprManager) UpdateOperatorContact(id uuid.UUID, contactId uuid.UUID, operatorId uuid.UUID, primary bool) (bool, error) {
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
func (o *OprManager) DeleteOperatorContact(id uuid.UUID) (bool, error) {
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
func (o *OprManager) GetAllContactsByOperator(operatorId uuid.UUID) ([]bu.OperatorContactsBO, error) {
	op := OpFac.New(bs.COperationContact).(*bs.OperatorContact)
	res, err := op.GetAllContactsByOperator(operatorId)
	return res, err
}
