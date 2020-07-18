package entities

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"tionyxtrack/common"
)
import "errors"

type TableVehicle struct {
	common.Base
	ModelId       uuid.UUID                    `gorm:"column:modelid;not_null"`
	MakeId        uuid.UUID                    `gorm:"column:makeid;not_null"`
	Registration  string                       `gorm:"column:registration;not_null"`
	FleetId       uuid.UUID                    `gorm:"column:fleetid;not_null"`
	StatusId      uuid.UUID                    `gorm:"column:statusid;not_null"`
	VehicleModel  *TableVehicleModel           `gorm:"foreignkey:modelid"`
	VehicleMake   *TableVehicleMake            `gorm:"foreignkey:makeid"`
	Fleet         *TableFleet                  `gorm:"foreignkey:fleetid"`
	Status        *TableVehicleStatus          `gorm:"foreignkey:statusid"`
	Locations     []*TableVehicleLocation      `gorm:"foreignkey:vehicleid"`
	History       []*TableVehicleHistory       `gorm:"foreignkey:vehicleid"`
	Operators     []*TableVehicleOperatorBound `gorm:"foreignkey:vehicleid"`
	Registrations []*TableVehicleTrackReg      `gorm:"foreignkey:vehicleId"`
}

func (t TableVehicle) TableName() string {
	return "table_vehicles"
}

func (t TableVehicle) Validate(db *gorm.DB) {
	if t.ModelId == uuid.Nil {
		_ = db.AddError(errors.New("model should contain value"))
	}
	if t.MakeId == uuid.Nil {
		_ = db.AddError(errors.New("make should contain value"))
	}
	if len(t.Registration) == 0 {
		_ = db.AddError(errors.New("registration should contain value"))
	}
	if t.FleetId == uuid.Nil {
		_ = db.AddError(errors.New("fleet should contain value"))
	}
	if t.StatusId == uuid.Nil {
		_ = db.AddError(errors.New("status should contain value"))
	}
}
