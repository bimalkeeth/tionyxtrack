package entities

import (
	"errors"
	"github.com/jinzhu/gorm"
)
//-------------------------------------
//Address type of the table
//-------------------------------------
type TableAddressType struct {
	gorm.Model
	AddressType string          `gorm:"column:addresstype;not_null"`
	Name        string          `gorm:"column:name;not_null"`
	Address     []*TableAddress `gorm:"foreignkey:addresstypeid"`
}
func (t TableAddressType) TableName() string {
	return "table_addresstype"
}
func (t TableAddressType) Validate(db *gorm.DB) {
	if len(t.Name) > 200 {
		_ = db.AddError(errors.New("name should not exceed length of 200 characters"))
	}
	if len(t.AddressType) > 4 {
		_ = db.AddError(errors.New("address type should not exceed length of 4 characters"))
	}
}
