package mappers

import "github.com/jinzhu/gorm"
import ent "tionyxtrack/masterservice/entities"

//--------------------------------------
// Create Fleet Contact table
//--------------------------------------
func MapFleetContactTable(db *gorm.DB) {

	if !db.HasTable(&ent.TableFleetContact{}) {

		db.CreateTable(&ent.TableFleetContact{})
		db.Model(&ent.TableFleetContact{}).AddUniqueIndex("ux_fleetcontact_fleetidcontactid", "fleetid", "contactid")
		db.Model(&ent.TableFleetContact{}).AddForeignKey("fleetid", "table_fleet(id)", "RESTRICT", "RESTRICT")
		db.Model(&ent.TableFleetContact{}).AddForeignKey("contactid", "table_contacts(id)", "RESTRICT", "RESTRICT")
	}
}
