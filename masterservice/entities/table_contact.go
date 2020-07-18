package entities

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"tionyxtrack/common"
)
import "errors"

//-------------------------------------
//Table for storing contact for system
//-------------------------------------
type TableContact struct {
	common.Base
	Contact       string            `gorm:"column:contact;not_null"`
	ContactTypeId uuid.UUID         `gorm:"column:contacttypeid;not_null"`
	ContactType   *TableContactType `gorm:"foreignkey:contacttypeid"`
}

func (t TableContact) TableName() string {
	return "table_contacts"
}
func (t TableContact) Validate(db *gorm.DB) {
	if len(t.Contact) == 0 {
		_ = db.AddError(errors.New("contact should should contain value"))
	}
	if t.ContactTypeId == uuid.Nil {
		_ = db.AddError(errors.New("contact type id should should contain value"))
	}
}
