package entities

import (
	"errors"
	"github.com/jinzhu/gorm"
	"time"
)

type TableVehicleTrackReg struct {
	gorm.Model
	RegisterDate time.Time     `gorm:"column:registrationdate;not_null"`
	Duration     int           `gorm:"column:duration;not_null"`
	ExpiredDate  time.Time     `gorm:"column:expiredate"`
	Active       bool          `gorm:"column:active;not_null"`
	VehicleId    uint          `gorm:"column:vehicleid;not_null"`
	Vehicle      *TableVehicle `gorm:"foreignkey:vehicleid"`
}

func (t TableVehicleTrackReg) TableName() string {
	return "table_vehicletrackreg"
}

func (t TableVehicleTrackReg) Validate(db *gorm.DB) {
	if t.RegisterDate.IsZero() {
		_ = db.AddError(errors.New("registration date type should contain value"))
	}
	if t.Duration == 0 {
		_ = db.AddError(errors.New("duration name should contain value"))
	}
	if t.ExpiredDate.IsZero() {
		_ = db.AddError(errors.New("expire date name should contain value"))
	}
	if t.VehicleId == 0 {
		_ = db.AddError(errors.New("vehicle should contain value"))
	}
}
