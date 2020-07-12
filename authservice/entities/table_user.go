package entities

import (
	"errors"
	"github.com/jinzhu/gorm"
	"time"
	"tionyxtrack/common"
)

type TableUser struct {
	common.Base
	EmailAddress    string    `gorm:"column:emailaddress;not_null"`
	Password        string    `gorm:"column:password;not_null"`
	Salt            string    `gorm:"column:salt;not_null"`
	FirstName       string    `gorm:"column:firstname;not_null"`
	LastName        string    `gorm:"column:lastname;not_null"`
	IsEmailVerified bool      `gorm:"column:isemailverified;not_null;default:false"`
	EmailVerifiedAt time.Time `gorm:"column:emailverifiedat"`
}

func (t TableUser) TableName() string {
	return "table_user"
}
func (t TableUser) Validate(db *gorm.DB) {
	if len(t.EmailAddress) == 0 {
		_ = db.AddError(errors.New("email address should be defined"))
	}
	if len(t.Password) == 0 {
		_ = db.AddError(errors.New("password should be defined"))
	}
	if len(t.FirstName) == 0 {
		_ = db.AddError(errors.New("first name should be defined"))
	}
	if len(t.LastName) == 0 {
		_ = db.AddError(errors.New("last name should be defined"))
	}
}
