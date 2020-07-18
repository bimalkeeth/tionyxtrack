package operator

import (
	uuid "github.com/satori/go.uuid"
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
	CreateOperator(bo bu.OperatorBO) (uuid.UUID, error)
	UpdateOperator(bo bu.OperatorBO) (bool, error)
	DeleteOperator(id uuid.UUID) (bool, error)
	GetOperatorById(id uuid.UUID) (bu.OperatorBO, error)
	GetOperatorsByVehicleId(id uuid.UUID) ([]bu.OperatorBO, error)
	CreateOperatorContact(contactId uuid.UUID, operatorId uuid.UUID, primary bool) (uuid.UUID, error)
	UpdateOperatorContact(id uuid.UUID, contactId uuid.UUID, operatorId uuid.UUID, primary bool) (bool, error)
	DeleteOperatorContact(id uuid.UUID) (bool, error)
	GetAllContactsByOperator(operatorId uuid.UUID) ([]bu.OperatorContactsBO, error)
	CreateOperatorLocation(bo bu.OperatorLocationBO) (uuid.UUID, error)
	UpdateOperatorLocation(bo bu.OperatorLocationBO) (bool, error)
	DeleteOperatorLocation(id uuid.UUID) (bool, error)
	GetOperatorLocationByOperator(id uuid.UUID) ([]bu.OperatorLocationBO, error)
}

type OprManager struct{}

func New() IOprManager {
	return &OprManager{}
}
