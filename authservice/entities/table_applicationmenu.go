package entities

import (
	"errors"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"tionyxtrack/common"
)

type TableApplicationMenu struct {
	common.Base
	RoleId uuid.UUID `gorm:"column:roleid;not_null"`
	Menu   string    `gorm:"column:menu;not_null"`
}

func (t TableApplicationMenu) TableName() string {
	return "table_applicationmenu"
}
func (t TableApplicationMenu) Validate(db *gorm.DB) {
	if t.RoleId == uuid.Nil {
		_ = db.AddError(errors.New("role name should be defined"))
	}
	if len(t.Menu) == 0 {
		_ = db.AddError(errors.New("menu should be defined"))
	}
}
