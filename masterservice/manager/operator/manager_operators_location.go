package operator

import (
	uuid "github.com/satori/go.uuid"
	bs "tionyxtrack/masterservice/business"
	bu "tionyxtrack/masterservice/businesscontracts"
)

//----------------------------------------------------
//Create operation location
//----------------------------------------------------
func (o *OprManager) CreateOperatorLocation(bo bu.OperatorLocationBO) (uuid.UUID, error) {
	op := OpFac.New(bs.COperationLocation).(*bs.OperatorLocation)
	OpFac.Conn.Begin()
	res, err := op.CreateOperatorLocation(bo)
	if err != nil {
		OpFac.Conn.Rollback()
		return res, err
	}
	OpFac.Conn.Commit()
	return res, err
}

//----------------------------------------------------
//update operation location
//----------------------------------------------------
func (o *OprManager) UpdateOperatorLocation(bo bu.OperatorLocationBO) (bool, error) {
	op := OpFac.New(bs.COperationLocation).(*bs.OperatorLocation)
	OpFac.Conn.Begin()
	res, err := op.UpdateOperatorLocation(bo)
	if err != nil {
		OpFac.Conn.Rollback()
		return res, err
	}
	OpFac.Conn.Commit()
	return res, err
}

//----------------------------------------------------
//delete operation location
//----------------------------------------------------
func (o *OprManager) DeleteOperatorLocation(id uuid.UUID) (bool, error) {
	op := OpFac.New(bs.COperationLocation).(*bs.OperatorLocation)
	OpFac.Conn.Begin()
	res, err := op.DeleteOperatorLocation(id)
	if err != nil {
		OpFac.Conn.Rollback()
		return res, err
	}
	OpFac.Conn.Commit()
	return res, err
}

//----------------------------------------------------
//get operator location by operatorid
//----------------------------------------------------
func (o *OprManager) GetOperatorLocationByOperator(id uuid.UUID) ([]bu.OperatorLocationBO, error) {
	op := OpFac.New(bs.COperationLocation).(*bs.OperatorLocation)
	res, err := op.GetOperatorLocationByOperator(id)
	return res, err
}
