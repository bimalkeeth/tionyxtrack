package entities

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"tionyxtrack/common"
)
import "errors"

type TableVehicleOperatorBound struct {
	common.Base
	OperatorId uuid.UUID              `gorm:"column:operatorid;not_null"`
	VehicleId  uuid.UUID              `gorm:"column:vehicleid;not_null"`
	Active     bool                   `gorm:"column:active;not_null"`
	Operator   *TableVehicleOperators `gorm:"foreignkey:operatorid"`
	Vehicle    *TableVehicle          `gorm:"foreignkey:vehicleid"`
}

func (t TableVehicleOperatorBound) TableName() string {
	return "table_vehicleoperatorbound"
}

func (t TableVehicleOperatorBound) Validate(db *gorm.DB) {

	if t.OperatorId == uuid.Nil {
		_ = db.AddError(errors.New("operator should contain value"))
	}
	if t.VehicleId == uuid.Nil {
		_ = db.AddError(errors.New("vehicle should contain value"))
	}
}
