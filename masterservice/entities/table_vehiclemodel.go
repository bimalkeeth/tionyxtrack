package entities

import "github.com/jinzhu/gorm"
import "errors"

type TableVehicleModel struct {
	gorm.Model
	ModelName   string            `gorm:"column:modelname;not_null"`
	Description string            `gorm:"column:description"`
	MakeId      uint              `gorm:"column:makeid;not_null"`
	Make        *TableVehicleMake `gorm:"foreignkey:makeid"`
}

func (t TableVehicleModel) TableName() string {
	return "table_vehiclemodel"
}

func (t TableVehicleModel) Validate(db *gorm.DB) {
	if len(t.ModelName) == 0 {
		_ = db.AddError(errors.New("model name should contain value"))
	}
	if t.MakeId == 0 {
		_ = db.AddError(errors.New("make should contain value"))
	}
}
