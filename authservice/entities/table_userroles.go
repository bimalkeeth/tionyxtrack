package entities

import (
	"errors"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"tionyxtrack/common"
)

type TableUserRoles struct {
	common.Base
	UserId uuid.UUID `gorm:"column:userid;not_null"`
	RoleId uuid.UUID `gorm:"column:roleid;not_null"`
}

func (t TableUserRoles) TableName() string {
	return "table_userroles"
}

func (t TableUserRoles) Validate(db *gorm.DB) {
	if t.UserId == uuid.Nil {
		_ = db.AddError(errors.New("user id should be defined"))
	}
	if t.RoleId == uuid.Nil {
		_ = db.AddError(errors.New("role name should be defined"))
	}
}
