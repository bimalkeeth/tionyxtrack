package mappers

import "github.com/jinzhu/gorm"
import ent "tionyxtrack/masterservice/entities"

//--------------------------------------
// Create Fleet Location table
//--------------------------------------
func MapVehicleOperatorContactsTable(db *gorm.DB) {

	if !db.HasTable(&ent.TableVehicleOperatorContacts{}) {

		db.CreateTable(&ent.TableVehicleOperatorContacts{})
		db.Model(&ent.TableVehicleOperatorContacts{}).AddUniqueIndex("ux_operatorcontacts_operationidcontactid", "operatorid", "contactid")
		db.Model(&ent.TableVehicleOperatorContacts{}).AddForeignKey("contactid", "table_contacts(id)", "RESTRICT", "RESTRICT")
		db.Model(&ent.TableVehicleOperatorContacts{}).AddForeignKey("operatorid", "table_vehicleoperators(id)", "RESTRICT", "RESTRICT")
	}
}
