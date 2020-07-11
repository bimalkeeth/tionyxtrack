package mappers

import "github.com/jinzhu/gorm"
import ent "tionyxtrack/masterservice/entities"

//--------------------------------------
// Create vehicle operator bound table
//--------------------------------------
func MapVehicleOperatorBoundTable(db *gorm.DB) {

	if !db.HasTable(&ent.TableVehicleOperatorBound{}) {

		db.CreateTable(&ent.TableVehicleOperatorBound{})
		db.Model(&ent.TableVehicleOperatorBound{}).AddUniqueIndex("ux_vehicleoperatorbound_operatoridvehicleid", "operatorid", "vehicleid")
		db.Model(&ent.TableVehicleOperatorBound{}).AddForeignKey("operatorid", "table_vehicleoperators(id)", "RESTRICT", "RESTRICT")
		db.Model(&ent.TableVehicleOperatorBound{}).AddForeignKey("vehicleid", "table_vehicles(id)", "RESTRICT", "RESTRICT")
	}
}
