package entities

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"tionyxtrack/common"
)
import "errors"

type TableFleetLocation struct {
	common.Base
	FleetId   uuid.UUID     `gorm:"column:fleetid;not_null"`
	AddressId uuid.UUID     `gorm:"column:addressid;not_null"`
	Primary   bool          `gorm:"column:primary;not_null"`
	Fleet     *TableFleet   `gorm:"foreignkey:fleetid"`
	Address   *TableAddress `gorm:"foreignkey:addressid"`
}

func (t TableFleetLocation) TableName() string {
	return "table_fleetlocation"
}

func (t TableFleetLocation) Validate(db *gorm.DB) {

	if t.FleetId == uuid.Nil {

		_ = db.AddError(errors.New("fleet should contain value"))
	}
	if t.AddressId == uuid.Nil {

		_ = db.AddError(errors.New("address should contain value"))
	}
}
