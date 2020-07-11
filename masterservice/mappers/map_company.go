package mappers

import "github.com/jinzhu/gorm"
import ent "tionyxtrack/masterservice/entities"

//-------------------------------------------
//Create Company Table
//-------------------------------------------

func MapCompanyTable(db *gorm.DB) {
	if !db.HasTable(&ent.TableCompany{}) {
		db.CreateTable(&ent.TableCompany{})
		db.Model(&ent.TableCompany{}).AddUniqueIndex("ux_table_company_name_uindex", "name")
		db.Model(&ent.TableCompany{}).AddForeignKey("addressid", "table_address(id)", "RESTRICT", "RESTRICT")
		db.Model(&ent.TableCompany{}).AddForeignKey("contactid", "table_contacts(id)", "RESTRICT", "RESTRICT")
	}
}
