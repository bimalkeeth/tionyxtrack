package entities

import "github.com/jinzhu/gorm"
import "errors"

type TableVehicleOperatorBound struct {
	gorm.Model
	OperatorId uint                   `gorm:"column:operatorid;not_null"`
	VehicleId  uint                   `gorm:"column:vehicleid;not_null"`
	Active     bool                   `gorm:"column:active;not_null"`
	Operator   *TableVehicleOperators `gorm:"foreignkey:operatorid"`
	Vehicle    *TableVehicle          `gorm:"foreignkey:vehicleid"`
}

func (t TableVehicleOperatorBound) TableName() string {
	return "table_vehicleoperatorbound"
}

func (t TableVehicleOperatorBound) Validate(db *gorm.DB) {

	if t.OperatorId == 0 {
		_ = db.AddError(errors.New("operator should contain value"))
	}
	if t.VehicleId == 0 {
		_ = db.AddError(errors.New("vehicle should contain value"))
	}
}
