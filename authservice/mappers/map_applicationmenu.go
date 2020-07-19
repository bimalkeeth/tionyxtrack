package mappers

import (
	"github.com/jinzhu/gorm"
	ent "tionyxtrack/authservice/entities"
)

func MapApplicationMenuTable(db *gorm.DB) {
	if !db.HasTable(&ent.TableApplicationMenu{}) {

	}
}
