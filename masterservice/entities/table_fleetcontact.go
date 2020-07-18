package entities

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"tionyxtrack/common"
)
import "errors"

type TableFleetContact struct {
	common.Base
	ContactId uuid.UUID     `gorm:"column:contactid;not_null"`
	FleetId   uuid.UUID     `gorm:"column:fleetid;not_null"`
	Primary   bool          `gorm:"column:primary;not_null"`
	Contact   *TableContact `gorm:"foreignkey:contactid"`
	Fleet     *TableFleet   `gorm:"foreignkey:fleetid"`
}

func (t TableFleetContact) TableName() string {
	return "table_fleetcontact"
}

func (t TableFleetContact) Validate(db *gorm.DB) {
	if t.ContactId == uuid.Nil {
		_ = db.AddError(errors.New("contact should contain value"))
	}
	if t.FleetId == uuid.Nil {
		_ = db.AddError(errors.New("fleet should contain value"))
	}
}
