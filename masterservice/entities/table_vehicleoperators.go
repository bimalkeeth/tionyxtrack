package entities

import (
	"github.com/jinzhu/gorm"
	"tionyxtrack/common"
)
import "errors"

type TableVehicleOperators struct {
	common.Base
	Name       string                          `gorm:"column:name;not_null"`
	SurName    string                          `gorm:"column:surname;not_null"`
	DrivingLic string                          `gorm:"column:drivinglic;not_null"`
	Active     bool                            `gorm:"column:active;not_null;default:false"`
	Bounds     []*TableVehicleOperatorBound    `gorm:"foreignkey:operatorid"`
	Locations  []*TableVehicleOperatorLocation `gorm:"foreignkey:operatorid"`
	Contacts   []*TableVehicleOperatorContacts `gorm:"foreignkey:operatorid"`
}

func (t TableVehicleOperators) TableName() string {
	return "table_vehicleoperators"
}

func (t TableVehicleOperators) Validate(db *gorm.DB) {

	if len(t.Name) == 0 {
		_ = db.AddError(errors.New("name should contain value"))
	}
	if len(t.SurName) == 0 {
		_ = db.AddError(errors.New("surname should contain value"))
	}
	if len(t.DrivingLic) == 0 {
		_ = db.AddError(errors.New("driving license should contain value"))
	}
}
