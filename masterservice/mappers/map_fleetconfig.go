package mappers

import (
	"github.com/jinzhu/gorm"
	ent "tionyxtrack/masterservice/entities"
)

func MapFleetConfigTable(db *gorm.DB) {

	if !db.HasTable(&ent.TableFleetConfig{}) {
		db.CreateTable(&ent.TableFleetConfig{})
		db.Model(&ent.TableFleetConfig{}).AddUniqueIndex("ux_fleetcontact_fleetidcontactid", "fleetid", "id")
		db.Model(&ent.TableFleetContact{}).AddForeignKey("fleetid", "table_fleet(id)", "RESTRICT", "RESTRICT")
	}
}
