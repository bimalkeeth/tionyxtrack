package mappers

import "github.com/jinzhu/gorm"
import ent "tionyxtrack/masterservice/entities"

//--------------------------------------
// Create vehicle location table
//--------------------------------------
func MapVehicleLocationTable(db *gorm.DB) {

	if !db.HasTable(&ent.TableVehicleLocation{}) {
		db.CreateTable(&ent.TableVehicleLocation{})
		db.Model(&ent.TableVehicleLocation{}).AddForeignKey("vehicleid", "table_vehicles(id)", "RESTRICT", "RESTRICT")
		db.Model(&ent.TableVehicleLocation{}).AddForeignKey("addressid", "table_address(id)", "RESTRICT", "RESTRICT")
	}
}
