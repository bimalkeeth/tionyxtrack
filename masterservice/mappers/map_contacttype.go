package mappers

import "github.com/jinzhu/gorm"
import ent "tionyxtrack/masterservice/entities"

//--------------------------------------
// Create Contact Type table
//--------------------------------------
func MapContactTypeTable(db *gorm.DB) {

	if !db.HasTable(&ent.TableContactType{}) {
		db.CreateTable(&ent.TableContactType{})
		db.Model(&ent.TableContactType{}).AddUniqueIndex("ux_addresstype_contacttype", "contacttype")
	}
}
