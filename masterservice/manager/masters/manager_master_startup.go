package masters

import (
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
	CreateCompany(company bu.CompanyBO) (uint, error)
	UpdateCompany(company bu.CompanyBO) (bool, error)
	DeleteCompany(id uint) (bool, error)
	CreateAddressType(addressType bu.AddressTypeBO) (uint, error)
	UpdateAddressType(addressType bu.AddressTypeBO) (bool, error)
	DeleteAddressType(id uint) (bool, error)
	GetAddressTypeById(id uint) (bu.AddressTypeBO, error)
	GetAddressTypeByName(name string) (bu.AddressTypeBO, error)
	GetAllAddressTypes() ([]bu.AddressTypeBO, error)
	GetAllAddressTypeNames(namePart string) ([]bu.AddressTypeBO, error)
	CreateRegion(bo bu.RegionBO) (uint, error)
	UpdateRegion(bo bu.RegionBO) (bool, error)
	DeleteRegion(id uint) (bool, error)
	GetAllRegion() ([]bu.RegionBO, error)
	GetRegionById(id uint) (bu.RegionBO, error)
	GetRegionByName(name string) (bu.RegionBO, error)
	CreateState(bo bu.StateBO) (uint, error)
	UpdateState(bo bu.StateBO) (bool, error)
	DeleteState(id uint) (bool, error)
	GetStateById(id uint) (bu.StateBO, error)
	GetStateByCountryId(id uint) ([]bu.StateBO, error)
	GetStateByName(name string) (bu.StateBO, error)
	GetAllStates() ([]bu.StateBO, error)
}

type MasterManager struct{}

func New() IMasterManager {
	return &MasterManager{}
}
