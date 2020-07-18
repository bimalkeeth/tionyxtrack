package entities

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"tionyxtrack/common"
)
import "errors"

type TableVehicleOperatorLocation struct {
	common.Base
	AddressId  uuid.UUID              `gorm:"column:addressid;not_null"`
	OperatorId uuid.UUID              `gorm:"column:operatorid;not_null"`
	Primary    bool                   `gorm:"column:primary;not_null"`
	Address    *TableAddress          `gorm:"foreignkey:addressid"`
	Operator   *TableVehicleOperators `gorm:"foreignkey:operatorid"`
}

func (t TableVehicleOperatorLocation) TableName() string {
	return "table_vehicleoptlocation"
}

func (t TableVehicleOperatorLocation) Validate(db *gorm.DB) {

	if t.AddressId == uuid.Nil {
		_ = db.AddError(errors.New("address should contain value"))
	}
	if t.OperatorId == uuid.Nil {
		_ = db.AddError(errors.New("operator should contain value"))
	}
}
