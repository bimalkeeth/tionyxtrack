package operator

import (
	uuid "github.com/satori/go.uuid"
	bs "tionyxtrack/masterservice/business"
	bu "tionyxtrack/masterservice/businesscontracts"
)

//----------------------------------------------
//Create operator
//----------------------------------------------
func (o *OprManager) CreateOperator(bo bu.OperatorBO) (uuid.UUID, error) {

	op := OpFac.New(bs.CVhOperator).(*bs.Operator)
	OpFac.Conn.Begin()
	res, err := op.CreateOperator(bo)
	if err != nil {
		OpFac.Conn.Rollback()
		return res, err
	}
	OpFac.Conn.Commit()
	return res, nil
}

//----------------------------------------------
//Update operator
//----------------------------------------------
func (o *OprManager) UpdateOperator(bo bu.OperatorBO) (bool, error) {
	op := OpFac.New(bs.CVhOperator).(*bs.Operator)
	OpFac.Conn.Begin()
	res, err := op.UpdateOperator(bo)
	if err != nil {
		OpFac.Conn.Rollback()
		return res, err
	}
	OpFac.Conn.Commit()
	return res, nil
}

//----------------------------------------------
//Delete operator
//----------------------------------------------
func (o *OprManager) DeleteOperator(id uuid.UUID) (bool, error) {
	op := OpFac.New(bs.CVhOperator).(*bs.Operator)
	OpFac.Conn.Begin()
	res, err := op.DeleteOperator(id)
	if err != nil {
		OpFac.Conn.Rollback()
		return res, err
	}
	OpFac.Conn.Commit()
	return res, nil
}

//---------------------------------------------
//Get operator by Id
//---------------------------------------------
func (o *OprManager) GetOperatorById(id uuid.UUID) (bu.OperatorBO, error) {
	op := OpFac.New(bs.CVhOperator).(*bs.Operator)
	res, err := op.GetOperatorById(id)
	return res, err
}

//---------------------------------------------
//Get operator by vehicleid
//---------------------------------------------
func (o *OprManager) GetOperatorsByVehicleId(id uuid.UUID) ([]bu.OperatorBO, error) {
	op := OpFac.New(bs.CVhOperator).(*bs.Operator)
	res, err := op.GetOperatorsByVehicleId(id)
	return res, err
}
