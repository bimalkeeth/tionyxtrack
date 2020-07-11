package mappers

import "github.com/jinzhu/gorm"
import ent "tionyxtrack/masterservice/entities"

//--------------------------------------
// Create Country table
//--------------------------------------
func MapCountryTable(db *gorm.DB) {

	if !db.HasTable(&ent.TableCountry{}) {

		db.CreateTable(&ent.TableCountry{})
		db.Model(&ent.TableCountry{}).AddUniqueIndex("ux_country_countryname", "countryname")
		db.Model(&ent.TableCountry{}).AddForeignKey("regionid", "table_region(id)", "RESTRICT", "RESTRICT")
	}
}
