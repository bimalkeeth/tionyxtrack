package mappers

import "github.com/jinzhu/gorm"
import ent "tionyxtrack/masterservice/entities"

//--------------------------------------
// Create Operator Contact table
//--------------------------------------
func MapRegionTable(db *gorm.DB) {

	if !db.HasTable(&ent.TableRegion{}) {
		db.CreateTable(&ent.TableRegion{})
		db.Model(&ent.TableRegion{}).AddUniqueIndex("ux_regioname", "regionname")
	}
}
