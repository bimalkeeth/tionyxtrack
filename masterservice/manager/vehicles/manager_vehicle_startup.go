package vehicles

import (
	"log"
	bs "tionyxtrack/masterservice/business"
	bu "tionyxtrack/masterservice/businesscontracts"
	ma "tionyxtrack/masterservice/manager"
)

var vehicleFac *bs.RuleFactory

func init() {
	conn, err := ma.Conn()
	if err != nil {
		log.Fatal("error in master manager initialisation")
	}
	vehicleFac = &bs.RuleFactory{Conn: conn}
}

type IVehicleManager interface {
	CreateVehicle(vehicle bu.VehicleBO) (uint, error)
	UpdateVehicle(vehicle bu.VehicleBO) (bool, error)
	DeleteVehicle(vehicleId uint) (bool, error)
	GetVehicleById(vehicleId uint) (bu.VehicleBO, error)
	GetVehicleByRegistration(registration string) (bu.VehicleBO, error)
	GetVehiclesByFleetId(fleetId uint) ([]bu.VehicleBO, error)
	CreateVehicleHistory(history bu.VehicleHistoryBO) (uint, error)
	UpdateVehicleHistory(history bu.VehicleHistoryBO) (bool, error)
	DeleteVehicleHistory(id uint) (bool, error)
	GetVehicleHistoryByVehicleId(vehicleId uint) ([]bu.VehicleHistoryBO, error)
	CreateVehicleLocation(ad bu.VehicleAddressBO) (uint, error)
	UpdateVehicleLocation(ad bu.VehicleAddressBO) (bool, error)
	DeleteVehicleLocation(id uint) (bool, error)
	GetVehicleLocationByVehicle(vehicleId uint) ([]bu.VehicleAddressBO, error)
	CreateVehicleMake(bo bu.VehicleMakeBO) (uint, error)
	UpdateVehicleMake(bo bu.VehicleMakeBO) (bool, error)
	DeleteVehicleMake(id uint) (bool, error)
	GetAllVehicleMake() ([]bu.VehicleMakeBO, error)
	GetVehicleMakeById(id uint) (bu.VehicleMakeBO, error)
	CreateVehicleModel(bo bu.VehicleModelBO) (uint, error)
	UpdateVehicleModel(bo bu.VehicleModelBO) (bool, error)
	DeleteVehicleModel(id uint) (bool, error)
	GetAllModelByMake(id uint) ([]bu.VehicleModelBO, error)
	GetModelById(id uint) (bu.VehicleModelBO, error)
	CreateVehicleReg(bo bu.VehicleTrackRegBO) (uint, error)
	UpdateVehicleReg(bo bu.VehicleTrackRegBO) (bool, error)
	DeleteVehicleReg(id uint) (bool, error)
	GetAllRegistrationsByVehicleId(id uint) ([]bu.VehicleTrackRegBO, error)
	GetActiveRegistrationsByVehicleId(id uint) (bu.VehicleTrackRegBO, error)
	CreateVehicleStatus(bo bu.VehicleStatusBO) (uint, error)
	UpdateVehicleStatus(bo bu.VehicleStatusBO) (bool, error)
	DeleteVehicleStatus(id uint) (bool, error)
	GetAllVehicleStatus() ([]bu.VehicleStatusBO, error)
	CreateVehicleOpBound(op bu.VehicleOperatorBoundBO) (uint, error)
	UpdateVehicleOpBound(op bu.VehicleOperatorBoundBO) (bool, error)
	DeleteVehicleOpBound(id uint) (bool, error)
}

type VehicleManager struct{}

func New() *VehicleManager {
	return &VehicleManager{}
}
