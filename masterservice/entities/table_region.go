package entities

import (
	"github.com/jinzhu/gorm"
	"tionyxtrack/common"
)
import "errors"

type TableRegion struct {
	common.Base
	Region     string          `gorm:"column:region;not_null"`
	RegionName string          `gorm:"column:regionname;not_null"`
	Countries  []*TableCountry `gorm:"foreignkey:regionid"`
}

func (t TableRegion) TableName() string {
	return "table_region"
}

func (t TableRegion) Validate(db *gorm.DB) {

	if len(t.Region) == 0 {

		_ = db.AddError(errors.New("region should contain value"))
	}
	if len(t.RegionName) == 0 {

		_ = db.AddError(errors.New("region name duration should contain value"))
	}
}
