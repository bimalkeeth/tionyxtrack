package entities

import "github.com/jinzhu/gorm"
import "errors"

type TableState struct {
	gorm.Model
	Name      string        `gorm:"column:name;not_null"`
	CountryId uint          `gorm:"column:countryid;not_null"`
	Country   *TableCountry `gorm:"foreignkey:countryid"`
}

func (t TableState) TableName() string {
	return "table_states"
}

func (t TableState) Validate(db *gorm.DB) {

	if len(t.Name) == 0 {
		_ = db.AddError(errors.New("name should contain value"))
	}
	if t.CountryId == 0 {
		_ = db.AddError(errors.New("country should contain value"))
	}
}
