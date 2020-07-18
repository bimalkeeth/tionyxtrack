package vehicles

import (
	uuid "github.com/satori/go.uuid"
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
	CreateVehicle(vehicle bu.VehicleBO) (uuid.UUID, error)
	UpdateVehicle(vehicle bu.VehicleBO) (bool, error)
	DeleteVehicle(vehicleId uuid.UUID) (bool, error)
	GetVehicleById(vehicleId uuid.UUID) (bu.VehicleBO, error)
	GetVehicleByRegistration(registration string) (bu.VehicleBO, error)
	GetVehiclesByFleetId(fleetId uuid.UUID) ([]bu.VehicleBO, error)
	CreateVehicleHistory(history bu.VehicleHistoryBO) (uuid.UUID, error)
	UpdateVehicleHistory(history bu.VehicleHistoryBO) (bool, error)
	DeleteVehicleHistory(id uuid.UUID) (bool, error)
	GetVehicleHistoryByVehicleId(vehicleId uuid.UUID) ([]bu.VehicleHistoryBO, error)
	CreateVehicleLocation(ad bu.VehicleAddressBO) (uuid.UUID, error)
	UpdateVehicleLocation(ad bu.VehicleAddressBO) (bool, error)
	DeleteVehicleLocation(id uuid.UUID) (bool, error)
	GetVehicleLocationByVehicle(vehicleId uuid.UUID) ([]bu.VehicleAddressBO, error)
	CreateVehicleMake(bo bu.VehicleMakeBO) (uuid.UUID, error)
	UpdateVehicleMake(bo bu.VehicleMakeBO) (bool, error)
	DeleteVehicleMake(id uuid.UUID) (bool, error)
	GetAllVehicleMake() ([]bu.VehicleMakeBO, error)
	GetVehicleMakeById(id uuid.UUID) (bu.VehicleMakeBO, error)
	CreateVehicleModel(bo bu.VehicleModelBO) (uuid.UUID, error)
	UpdateVehicleModel(bo bu.VehicleModelBO) (bool, error)
	DeleteVehicleModel(id uuid.UUID) (bool, error)
	GetAllModelByMake(id uuid.UUID) ([]bu.VehicleModelBO, error)
	GetModelById(id uuid.UUID) (bu.VehicleModelBO, error)
	CreateVehicleReg(bo bu.VehicleTrackRegBO) (uuid.UUID, error)
	UpdateVehicleReg(bo bu.VehicleTrackRegBO) (bool, error)
	DeleteVehicleReg(id uuid.UUID) (bool, error)
	GetAllRegistrationsByVehicleId(id uuid.UUID) ([]bu.VehicleTrackRegBO, error)
	GetActiveRegistrationsByVehicleId(id uuid.UUID) (bu.VehicleTrackRegBO, error)
	CreateVehicleStatus(bo bu.VehicleStatusBO) (uuid.UUID, error)
	UpdateVehicleStatus(bo bu.VehicleStatusBO) (bool, error)
	DeleteVehicleStatus(id uuid.UUID) (bool, error)
	GetAllVehicleStatus() ([]bu.VehicleStatusBO, error)
	CreateVehicleOpBound(op bu.VehicleOperatorBoundBO) (uuid.UUID, error)
	UpdateVehicleOpBound(op bu.VehicleOperatorBoundBO) (bool, error)
	DeleteVehicleOpBound(id uuid.UUID) (bool, error)
}

type VehicleManager struct{}

func New() IVehicleManager {
	return &VehicleManager{}
}
