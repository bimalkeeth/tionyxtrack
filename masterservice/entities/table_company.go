package entities
import "github.com/jinzhu/gorm"
import "errors"
//--------------------------------------------------
// Table for teonyx information
//--------------------------------------------------
type TableCompany struct {
	gorm.Model
	Name       string        `gorm:"column:name;not_null"`
	AddressId  uint          `gorm:"column:addressid;not_null"`
	ContractId uint          `gorm:"column:contactid;not_null"`
	Address    *TableAddress `gorm:"foreignkey:AddressId"`
	Contract   *TableContact `gorm:"foreignkey:contactid"`
}
func (t TableCompany) TableName() string {
	return "table_company"
}
func (t TableCompany) Validate(db *gorm.DB) {
	if len(t.Name) == 0 {
		_ = db.AddError(errors.New("name should not be empty"))
	}
	if t.AddressId == 0 {
		_ = db.AddError(errors.New("address should not be empty"))
	}
	if t.ContractId == 0 {
		_ = db.AddError(errors.New("contact should not be empty"))
	}
}
