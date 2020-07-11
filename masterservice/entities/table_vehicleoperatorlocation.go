package entities

import "github.com/jinzhu/gorm"
import "errors"

type TableVehicleOperatorLocation struct {
	gorm.Model
	AddressId  uint                   `gorm:"column:addressid;not_null"`
	OperatorId uint                   `gorm:"column:operatorid;not_null"`
	Primary    bool                   `gorm:"column:primary;not_null"`
	Address    *TableAddress          `gorm:"foreignkey:addressid"`
	Operator   *TableVehicleOperators `gorm:"foreignkey:operatorid"`
}

func (t TableVehicleOperatorLocation) TableName() string {
	return "table_vehicleoptlocation"
}

func (t TableVehicleOperatorLocation) Validate(db *gorm.DB) {

	if t.AddressId == 0 {
		_ = db.AddError(errors.New("address should contain value"))
	}
	if t.OperatorId == 0 {
		_ = db.AddError(errors.New("operator should contain value"))
	}
}
