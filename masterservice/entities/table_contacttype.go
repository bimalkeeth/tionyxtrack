package entities
import "github.com/jinzhu/gorm"
import "errors"

type TableContactType struct {
	gorm.Model
	ContactType string          `gorm:"column:contacttype;not_null;"`
	Contacts    []*TableContact `gorm:"foreignkey:contacttypeid"`
}

func (t TableContactType) TableName() string {
	return "table_contacttype"
}
func (t TableContactType) Validate(db *gorm.DB) {
	if len(t.ContactType) == 0 {
		_ = db.AddError(errors.New("contact type should contain value"))
	}
}
