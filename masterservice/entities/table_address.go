package entities

import (
	"errors"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"tionyxtrack/common"
)

//-----------------------------------------
// Table for address storage
//-----------------------------------------
type TableAddress struct {
	common.Base
	Address          string            `gorm:"column:address;not_null"`
	Street           string            `gorm:"column:street:not_null"`
	Suburb           string            `gorm:"column:suburb:not_null"`
	StateId          uuid.UUID         `gorm:"column:stateid:not_null"`
	CountryId        uuid.UUID         `gorm:"column:countryid:not_null"`
	AddressTypeId    uuid.UUID         `gorm:"column:addresstypeid;not_null"`
	Location         string            `gorm:"column:location;not_null"`
	TableAddressType *TableAddressType `gorm:"foreignkey:addresstypeid"`
	State            *TableState       `gorm:"foreignkey:stateid"`
	Country          *TableCountry     `gorm:"foreignkey:countryid"`
}

func (t TableAddress) TableName() string {
	return "table_address"
}

func (t TableAddress) Validate(db *gorm.DB) {
	if len(t.Address) == 0 {
		_ = db.AddError(errors.New("address should not be empty"))
	}
	if len(t.Street) == 0 {
		_ = db.AddError(errors.New("address should not be empty"))
	}
	if t.AddressTypeId == uuid.Nil {
		_ = db.AddError(errors.New("address type should not be empty"))
	}
	if len(t.Location) == 0 {
		_ = db.AddError(errors.New("location should not be empty"))
	}
}
