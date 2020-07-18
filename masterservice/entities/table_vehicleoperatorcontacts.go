package entities

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"tionyxtrack/common"
)
import "errors"

type TableVehicleOperatorContacts struct {
	common.Base
	ContactId  uuid.UUID              `gorm:"column:contactid;not_null"`
	OperatorId uuid.UUID              `gorm:"column:operatorid;not_null"`
	Primary    bool                   `gorm:"column:primary;not_null"`
	Contact    *TableContact          `gorm:"foreignkey:contactid"`
	Operator   *TableVehicleOperators `gorm:"foreignkey:operatorid"`
}

func (t TableVehicleOperatorContacts) TableName() string {
	return "table_operatorcontacts"
}

func (t TableVehicleOperatorContacts) Validate(db *gorm.DB) {

	if t.ContactId == uuid.Nil {

		_ = db.AddError(errors.New("contact should contain value"))
	}
	if t.OperatorId == uuid.Nil {

		_ = db.AddError(errors.New("operator should contain value"))
	}
}
