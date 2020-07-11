package mappers

import "github.com/jinzhu/gorm"
import ent "tionyxtrack/masterservice/entities"

//--------------------------------------
// Create Contact table
//--------------------------------------
func MapContactsTable(db *gorm.DB) {

	if !db.HasTable(&ent.TableContact{}) {

		db.CreateTable(&ent.TableContact{})
		db.Model(&ent.TableContact{}).AddForeignKey("contacttypeid", "table_contacttype(id)", "RESTRICT", "RESTRICT")
		db.Model(&ent.TableContact{}).AddUniqueIndex("table_contacts_typecontact_uindex", "contacttypeid", "contact")

	}
}
