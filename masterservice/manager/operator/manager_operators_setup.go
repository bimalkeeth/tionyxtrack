package operator

import (
	"log"
	bs "tionyxtrack/masterservice/business"
	bu "tionyxtrack/masterservice/businesscontracts"
	ma "tionyxtrack/masterservice/manager"
)


var OpFac *bs.RuleFactory

func init() {
	conn, err := ma.Conn()
	if err != nil {
		log.Fatal("error in master manager initialisation")
	}
	OpFac = &bs.RuleFactory{Conn: conn}
}

type IOprManager interface {
	CreateOperator(bo bu.OperatorBO) (uint, error)
	UpdateOperator(bo bu.OperatorBO) (bool, error)
	DeleteOperator(id uint) (bool, error)
	GetOperatorById(id uint) (bu.OperatorBO, error)
	GetOperatorsByVehicleId(id uint) ([]bu.OperatorBO, error)
	CreateOperatorContact(contactId uint, operatorId uint, primary bool) (uint, error)
	UpdateOperatorContact(id uint, contactId uint, operatorId uint, primary bool) (bool, error)
	DeleteOperatorContact(id uint) (bool, error)
	GetAllContactsByOperator(operatorId uint) ([]bu.OperatorContactsBO, error)
	CreateOperatorLocation(bo bu.OperatorLocationBO) (uint, error)
	UpdateOperatorLocation(bo bu.OperatorLocationBO) (bool, error)
	DeleteOperatorLocation(id uint) (bool, error)
	GetOperatorLocationByOperator(id uint) ([]bu.OperatorLocationBO, error)
}

type OprManager struct{}

func New() *OprManager {
	return &OprManager{}
}
