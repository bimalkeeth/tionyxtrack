package mappers

import "github.com/jinzhu/gorm"
import ent "tionyxtrack/masterservice/entities"

//--------------------------------------
// Create vehicle model table
//--------------------------------------
func MapVehicleModelTable(db *gorm.DB) {

	if !db.HasTable(&ent.TableVehicleModel{}) {
		db.CreateTable(&ent.TableVehicleModel{})
		db.Model(&ent.TableVehicleModel{}).AddUniqueIndex("ux_vehiclemodel_modelname", "modelname")
		db.Model(&ent.TableVehicleModel{}).AddForeignKey("makeid", "table_vehiclemake(id)", "RESTRICT", "RESTRICT")

	}
}
