package entities

import (
	"errors"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"tionyxtrack/common"
)

type TableFleetConfig struct {
	common.Base
	FleetId    uuid.UUID   `gorm:"column:fleetid;not_null"`
	ConfigId   uint        `gorm:"column:configid;not_null"`
	ConfigName string      `gorm:"column:configname;not_null"`
	IsHidden   bool        `gorm:"column:ishidden;not_null;default:false"`
	IsReadOnly bool        `gorm:"column:isreadonly;not_null;default:false"`
	Fleet      *TableFleet `gorm:"foreignkey:fleetid"`
}

func (t TableFleetConfig) TableName() string {
	return "table_fleetconfig"
}

func (t TableFleetConfig) Validate(db *gorm.DB) {

	if t.FleetId == uuid.Nil {
		_ = db.AddError(errors.New("fleet id should be defined"))
	}
	if t.ConfigId == 0 {
		_ = db.AddError(errors.New("config id should be defined"))
	}
	if len(t.ConfigName) == 0 {
		_ = db.AddError(errors.New("config id should be defined"))
	}
}
