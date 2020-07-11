package mappers

import "github.com/jinzhu/gorm"
import ent "tionyxtrack/masterservice/entities"

//--------------------------------------
// Create vehicle operator location table
//--------------------------------------
func MapVehicleOperatorLocationTable(db *gorm.DB) {

	if !db.HasTable(&ent.TableVehicleLocation{}) {

		db.CreateTable(&ent.TableVehicleLocation{})
		db.Model(&ent.TableVehicleLocation{}).AddUniqueIndex("ux_vehicleoperators_operatoridaddressid", "operatorid", "addressid")
		db.Model(&ent.TableVehicleLocation{}).AddForeignKey("operatorid", "table_vehicleoperators(id)", "RESTRICT", "RESTRICT")
		db.Model(&ent.TableVehicleLocation{}).AddForeignKey("addressid", "table_address(id)", "RESTRICT", "RESTRICT")
	}
}
