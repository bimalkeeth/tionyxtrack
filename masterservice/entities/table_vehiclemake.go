package entities

import "github.com/jinzhu/gorm"
import "errors"

type TableVehicleMake struct {
	gorm.Model
	Make      string        `gorm:"column:make;not_null"`
	CountryId uint          `gorm:"column:countryid;not_null"`
	Country   *TableCountry `gorm:"foreignkey:countryid"`
}

func (t TableVehicleMake) TableName() string {
	return "table_vehiclemake"
}

func (t TableVehicleMake) Validate(db *gorm.DB) {
	if len(t.Make) == 0 {
		_ = db.AddError(errors.New("make should contain value"))
	}
	if t.CountryId == 0 {
		_ = db.AddError(errors.New("country should contain value"))
	}
}
