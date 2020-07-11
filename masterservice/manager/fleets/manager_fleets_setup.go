package fleets

import (
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
	DeleteFleet(id uint) (bool, error)
	GetFleetById(id uint) (bu.FleetBO, error)
	CreateFleetContact(fleetId uint, contactId uint, primary bool) (uint, error)
	UpdateFleetContact(id uint, fleetId uint, contactId uint, primary bool) (bool, error)
	DeleteFleetContact(id uint) (bool, error)
	GetContactByFleetId(fleetId uint) ([]bu.FleetContactBO, error)
	CreateFleetLocation(fleetId uint, addressId uint, primary bool) (uint, error)
	UpdateFleetLocation(id uint, fleetId uint, addressId uint, primary bool) (bool, error)
	DeleteFleetLocation(id uint) (bool, error)
	GetLocationByFleetId(fleetId uint) ([]bu.FleetAddressBO, error)
}

type FleetManager struct{}

func New() *FleetManager {
	return &FleetManager{}
}
