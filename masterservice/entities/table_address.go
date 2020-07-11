package entities

import (
	"errors"
	"github.com/jinzhu/gorm"
)

//-----------------------------------------
// Table for address storage
//-----------------------------------------
type TableAddress struct {
	gorm.Model
	Address          string            `gorm:"column:address;not_null"`
	Street           string            `gorm:"column:street:not_null"`
	Suburb           string            `gorm:"column:suburb"`
	StateId          uint              `gorm:"column:stateid"`
	CountryId        uint              `gorm:"column:countryid"`
	AddressTypeId    uint              `gorm:"column:addresstypeid;not_null"`
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
	if t.AddressTypeId == 0 {
		_ = db.AddError(errors.New("address type should not be empty"))
	}
	if len(t.Location) == 0 {
		_ = db.AddError(errors.New("location should not be empty"))
	}
}
