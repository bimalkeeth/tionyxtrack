package mappers

import "github.com/jinzhu/gorm"
import ent "tionyxtrack/masterservice/entities"

//--------------------------------------
// Create Address table
//--------------------------------------
func MapAddressTable(db *gorm.DB) {

	if !db.HasTable(&ent.TableAddress{}) {

		db.CreateTable(&ent.TableAddress{})
		db.Model(&ent.TableAddress{}).AddForeignKey("addresstypeid", "table_addresstype(id)", "RESTRICT", "RESTRICT")
		db.Model(&ent.TableAddress{}).AddForeignKey("countryid", "table_country(id)", "RESTRICT", "RESTRICT")
		db.Model(&ent.TableAddress{}).AddForeignKey("stateid", "table_states(id)", "RESTRICT", "RESTRICT")
	}
}
