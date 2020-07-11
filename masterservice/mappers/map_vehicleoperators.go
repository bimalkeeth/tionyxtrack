package mappers

import "github.com/jinzhu/gorm"
import ent "tionyxtrack/masterservice/entities"

//--------------------------------------
// Create vehicle operator table
//--------------------------------------
func MapVehicleOperatorsTable(db *gorm.DB) {

	if !db.HasTable(&ent.TableVehicleOperators{}) {

		db.CreateTable(&ent.TableVehicleOperators{})
		db.Model(&ent.TableVehicleOperators{}).AddUniqueIndex("ux_vehicleoperators_drivinglic", "drivinglic")
		db.Model(&ent.TableVehicleOperators{}).Association("Bounds")
		db.Model(&ent.TableVehicleOperators{}).Association("Locations")

	}
}
