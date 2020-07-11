package entities

import "github.com/jinzhu/gorm"
import "errors"

type TableVehicleStatus struct {
	gorm.Model
	StatusType string `gorm:"column:statustype;not_null"`
	StatusName string `gorm:"column:statusname;not_null"`
}

func (t TableVehicleStatus) TableName() string {
	return "table_vehiclestatus"
}

func (t TableVehicleStatus) Validate(db *gorm.DB) {
	if len(t.StatusType) == 0 {
		_ = db.AddError(errors.New("status type should contain value"))
	}
	if len(t.StatusName) == 0 {
		_ = db.AddError(errors.New("status name should contain value"))
	}
}
