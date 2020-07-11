package entities

import "github.com/jinzhu/gorm"
import "errors"

type TableVehicle struct {
	gorm.Model
	ModelId       uint                         `gorm:"column:modelid;not_null"`
	MakeId        uint                         `gorm:"column:makeid;not_null"`
	Registration  string                       `gorm:"column:registration;not_null"`
	FleetId       uint                         `gorm:"column:fleetid;not_null"`
	StatusId      uint                         `gorm:"column:statusid;not_null"`
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
	if t.ModelId == 0 {
		_ = db.AddError(errors.New("model should contain value"))
	}
	if t.MakeId == 0 {
		_ = db.AddError(errors.New("make should contain value"))
	}
	if len(t.Registration) == 0 {
		_ = db.AddError(errors.New("registration should contain value"))
	}
	if t.FleetId == 0 {
		_ = db.AddError(errors.New("fleet should contain value"))
	}
	if t.StatusId == 0 {
		_ = db.AddError(errors.New("status should contain value"))
	}
}
