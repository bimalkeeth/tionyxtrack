package mappers

import "github.com/jinzhu/gorm"
import ent "tionyxtrack/masterservice/entities"

//--------------------------------------
// Create Fleet Location table
//--------------------------------------
func MapFleetLocationTable(db *gorm.DB) {

	if !db.HasTable(&ent.TableFleetLocation{}) {

		db.CreateTable(&ent.TableFleetLocation{})
		db.Model(&ent.TableFleetLocation{}).AddUniqueIndex("ux_fleetlocation_fleetidaddressid", "fleetid", "addressid")
		db.Model(&ent.TableFleetLocation{}).AddForeignKey("fleetid", "table_fleet(id)", "RESTRICT", "RESTRICT")
		db.Model(&ent.TableFleetLocation{}).AddForeignKey("addressid", "table_address(id)", "RESTRICT", "RESTRICT")
	}
}
