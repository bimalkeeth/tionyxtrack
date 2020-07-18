package contacts

import (
	uuid "github.com/satori/go.uuid"
	"log"
	bs "tionyxtrack/masterservice/business"
	bu "tionyxtrack/masterservice/businesscontracts"
	ma "tionyxtrack/masterservice/manager"
)

type IContactManager interface {
	CreateContact(con bu.ContactBO) (uuid.UUID, error)
	UpdateContact(con bu.ContactBO) (bool, error)
	DeleteContact(id uuid.UUID) (bool, error)
	ContactById(Id uuid.UUID) (bu.ContactBO, error)
	CreateAddress(add bu.AddressBO) (uuid.UUID, error)
	UpdateAddress(add bu.AddressBO) (bool, error)
	DeleteAddress(id uuid.UUID) (bool, error)
	GetAddressById(id uuid.UUID) (bu.AddressBO, error)
	GetAddressByName(name string) ([]bu.AddressBO, error)
}

type ContactManager struct{}

func New() IContactManager {
	return &ContactManager{}
}

var conFactory *bs.RuleFactory

func init() {
	conn, err := ma.Conn()
	if err != nil {
		log.Fatal("error in contact manager initialisation")
	}
	conFactory = &bs.RuleFactory{Conn: conn}
}
