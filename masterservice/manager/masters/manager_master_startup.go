package masters

import (
	uuid "github.com/satori/go.uuid"
	"log"
	bs "tionyxtrack/masterservice/business"
	bu "tionyxtrack/masterservice/businesscontracts"
	ma "tionyxtrack/masterservice/manager"
)

var masFac *bs.RuleFactory

func init() {
	conn, err := ma.Conn()
	if err != nil {
		log.Fatal("error in master manager initialisation")
	}
	masFac = &bs.RuleFactory{Conn: conn}
}

type IMasterManager interface {
	CreateCompany(company bu.CompanyBO) (uuid.UUID, error)
	UpdateCompany(company bu.CompanyBO) (bool, error)
	DeleteCompany(id uuid.UUID) (bool, error)
	CreateAddressType(addressType bu.AddressTypeBO) (uuid.UUID, error)
	UpdateAddressType(addressType bu.AddressTypeBO) (bool, error)
	DeleteAddressType(id uuid.UUID) (bool, error)
	GetAddressTypeById(id uuid.UUID) (bu.AddressTypeBO, error)
	GetAddressTypeByName(name string) (bu.AddressTypeBO, error)
	GetAllAddressTypes() ([]bu.AddressTypeBO, error)
	GetAllAddressTypeNames(namePart string) ([]bu.AddressTypeBO, error)
	CreateRegion(bo bu.RegionBO) (uuid.UUID, error)
	UpdateRegion(bo bu.RegionBO) (bool, error)
	DeleteRegion(id uuid.UUID) (bool, error)
	GetAllRegion() ([]bu.RegionBO, error)
	GetRegionById(id uuid.UUID) (bu.RegionBO, error)
	GetRegionByName(name string) (bu.RegionBO, error)
	CreateState(bo bu.StateBO) (uuid.UUID, error)
	UpdateState(bo bu.StateBO) (bool, error)
	DeleteState(id uuid.UUID) (bool, error)
	GetStateById(id uuid.UUID) (bu.StateBO, error)
	GetStateByCountryId(id uuid.UUID) ([]bu.StateBO, error)
	GetStateByName(name string) (bu.StateBO, error)
	GetAllStates() ([]bu.StateBO, error)
}

type MasterManager struct{}

func New() IMasterManager {
	return &MasterManager{}
}
