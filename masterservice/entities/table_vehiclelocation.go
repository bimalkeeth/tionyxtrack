package entities

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"tionyxtrack/common"
)
import "errors"

type TableVehicleLocation struct {
	common.Base
	AddressId uuid.UUID     `gorm:"column:addressid;not_null"`
	VehicleId uuid.UUID     `gorm:"column:vehicleid;not_null"`
	Primary   bool          `gorm:"column:primary;not_null"`
	Address   *TableAddress `gorm:"foreignkey:addressid"`
	Vehicle   *TableVehicle `gorm:"foreignkey:vehicleid"`
}

func (t TableVehicleLocation) TableName() string {
	return "table_vehiclelocation"
}

func (t TableVehicleLocation) Validate(db *gorm.DB) {
	if t.AddressId == uuid.Nil {
		_ = db.AddError(errors.New("address should contain value"))
	}
	if t.VehicleId == uuid.Nil {
		_ = db.AddError(errors.New("vehicle should contain value"))
	}
}
