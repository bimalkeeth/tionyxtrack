package fleets

import (
	uuid "github.com/satori/go.uuid"
	"log"
	bs "tionyxtrack/masterservice/business"
	bu "tionyxtrack/masterservice/businesscontracts"
	ma "tionyxtrack/masterservice/manager"
)

var flFac *bs.RuleFactory

func init() {
	conn, err := ma.Conn()
	if err != nil {
		log.Fatal("error in master manager initialisation")
	}
	flFac = &bs.RuleFactory{Conn: conn}
}

type IFleetManager interface {
	CreateFleet(bo bu.FleetBO) (bu.FleetBO, error)
	UpdateFleet(bo bu.FleetBO) (bool, error)
	DeleteFleet(id uuid.UUID) (bool, error)
	GetFleetById(id uuid.UUID) (bu.FleetBO, error)
	CreateFleetContact(fleetId uuid.UUID, contactId uuid.UUID, primary bool) (uuid.UUID, error)
	UpdateFleetContact(id uuid.UUID, fleetId uuid.UUID, contactId uuid.UUID, primary bool) (bool, error)
	DeleteFleetContact(id uuid.UUID) (bool, error)
	GetContactByFleetId(fleetId uuid.UUID) ([]bu.FleetContactBO, error)
	CreateFleetLocation(fleetId uuid.UUID, addressId uuid.UUID, primary bool) (uuid.UUID, error)
	UpdateFleetLocation(id uuid.UUID, fleetId uuid.UUID, addressId uuid.UUID, primary bool) (bool, error)
	DeleteFleetLocation(id uuid.UUID) (bool, error)
	GetLocationByFleetId(fleetId uuid.UUID) ([]bu.FleetAddressBO, error)
}

type FleetManager struct{}

func New() *FleetManager {
	return &FleetManager{}
}
