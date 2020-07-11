package mappers

import "github.com/jinzhu/gorm"
import ent "tionyxtrack/masterservice/entities"

//--------------------------------------
// Create Operator Contact table
//--------------------------------------
func MapStatesTable(db *gorm.DB) {

	if !db.HasTable(&ent.TableState{}) {

		db.CreateTable(&ent.TableState{})
		db.Model(&ent.TableState{}).AddUniqueIndex("ux_states_name", "name")
		db.Model(&ent.TableState{}).AddForeignKey("countryid", "table_country(id)", "RESTRICT", "RESTRICT")
	}
}
