package entities

import (
	"errors"
	"github.com/jinzhu/gorm"
	"time"
)

type TableVehicleHistory struct {
	gorm.Model
	VehicleId    uint                `gorm:"column:vehicleid;not_null"`
	ChangeDate   time.Time           `gorm:"column:changedate;not_null"`
	Description  string              `gorm:"column:description"`
	FromStatusId uint                `gorm:"column:fromstatusid;not_null"`
	ToStatusId   uint                `gorm:"column:tostatusid;not_null"`
	OfficerName  string              `gorm:"column:officername;not_null"`
	FromStatus   *TableVehicleStatus `gorm:"foreignkey:fromstatusid"`
	ToStatus     *TableVehicleStatus `gorm:"foreignkey:tostatusid"`
	Vehicle      *TableVehicle       `gorm:"foreignkey:vehicleid"`
}

func (t TableVehicleHistory) TableName() string {
	return "table_vehiclehistory"
}

func (t TableVehicleHistory) Validate(db *gorm.DB) {

	if len(t.OfficerName) == 0 {
		_ = db.AddError(errors.New("officer name should contain value"))
	}
	if t.VehicleId == 0 {
		_ = db.AddError(errors.New("vehicle should contain value"))
	}
	if t.ChangeDate.IsZero() {
		_ = db.AddError(errors.New("change date should contain value"))
	}
	if t.FromStatusId == 0 {
		_ = db.AddError(errors.New("from status should contain value"))
	}
	if t.ToStatusId == 0 {
		_ = db.AddError(errors.New("to status should contain value"))
	}
}
