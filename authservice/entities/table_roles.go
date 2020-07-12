package entities

import (
	"errors"
	"github.com/jinzhu/gorm"
)

type TableRoles struct {
	gorm.Model
	Name        string `gorm:"column:name;not_null"`
	Description string `gorm:"column:description"`
}

func (t TableRoles) TableName() string {
	return "table_roles"
}

func (t TableRoles) Validate(db *gorm.DB) {
	if len(t.Name) == 0 {
		_ = db.AddError(errors.New("role name should be defined"))
	}
}
