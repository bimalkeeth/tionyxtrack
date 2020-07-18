package entities

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"tionyxtrack/common"
)
import "errors"

type TableVehicleModel struct {
	common.Base
	ModelName   string            `gorm:"column:modelname;not_null"`
	Description string            `gorm:"column:description"`
	MakeId      uuid.UUID         `gorm:"column:makeid;not_null"`
	Make        *TableVehicleMake `gorm:"foreignkey:makeid"`
}

func (t TableVehicleModel) TableName() string {
	return "table_vehiclemodel"
}

func (t TableVehicleModel) Validate(db *gorm.DB) {
	if len(t.ModelName) == 0 {
		_ = db.AddError(errors.New("model name should contain value"))
	}
	if t.MakeId == uuid.Nil {
		_ = db.AddError(errors.New("make should contain value"))
	}
}
