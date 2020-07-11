package entities

import "github.com/jinzhu/gorm"
import "errors"

type TableVehicleOperatorContacts struct {
	gorm.Model
	ContactId  uint                   `gorm:"column:contactid;not_null"`
	OperatorId uint                   `gorm:"column:operatorid;not_null"`
	Primary    bool                   `gorm:"column:primary;not_null"`
	Contact    *TableContact          `gorm:"foreignkey:contactid"`
	Operator   *TableVehicleOperators `gorm:"foreignkey:operatorid"`
}

func (t TableVehicleOperatorContacts) TableName() string {
	return "table_operatorcontacts"
}

func (t TableVehicleOperatorContacts) Validate(db *gorm.DB) {

	if t.ContactId == 0 {

		_ = db.AddError(errors.New("contact should contain value"))
	}
	if t.OperatorId == 0 {

		_ = db.AddError(errors.New("operator should contain value"))
	}
}
