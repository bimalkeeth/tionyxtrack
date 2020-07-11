package mappers

import "github.com/jinzhu/gorm"
import ent "tionyxtrack/masterservice/entities"
//--------------------------------------
// Create vehicle status table
//--------------------------------------
func MapVehicleStatusTable(db *gorm.DB) {
	if !db.HasTable(&ent.TableVehicleStatus{}) {
		db.CreateTable(&ent.TableVehicleStatus{})
		db.Model(&ent.TableVehicleStatus{}).AddUniqueIndex("ux_vehiclestatus_statustype", "statustype")
	}
}
