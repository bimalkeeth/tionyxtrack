package entities

import "github.com/jinzhu/gorm"
import "errors"

//-------------------------------------
//Table for storing contact for system
//-------------------------------------
type TableContact struct {
	gorm.Model
	Contact       string            `gorm:"column:contact;not_null"`
	ContactTypeId uint              `gorm:"column:contacttypeid;not_null"`
	ContactType   *TableContactType `gorm:"foreignkey:contacttypeid"`
}
func (t TableContact) TableName() string {
	return "table_contacts"
}
func (t TableContact) Validate(db *gorm.DB) {
	if len(t.Contact) == 0 {
		_ = db.AddError(errors.New("contact should should contain value"))
	}
	if t.ContactTypeId == 0 {
		_ = db.AddError(errors.New("contact type id should should contain value"))
	}
}
