package entities

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"tionyxtrack/common"
)
import "errors"

type TableCountry struct {
	common.Base
	CountryName string        `gorm:"column:countryname;not_null"`
	RegionId    uuid.UUID     `gorm:"column:regionid;not_null"`
	Region      *TableRegion  `gorm:"foreignkey:regionid"`
	States      []*TableState `gorm:"foreignkey:countryId"`
}

func (t TableCountry) TableName() string {
	return "table_country"
}
func (t TableCountry) Validate(db *gorm.DB) {
	if len(t.CountryName) == 0 {
		_ = db.AddError(errors.New("country name should contain value"))
	}
	if t.RegionId == uuid.Nil {
		_ = db.AddError(errors.New("region should contain value"))
	}
}
