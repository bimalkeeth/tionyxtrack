package manager

import (
	"github.com/jinzhu/gorm"
	"tionyxtrack/masterservice/connection"
)

//---------------------------------------
//Set database connection
//---------------------------------------
func Conn() (*gorm.DB, error) {
	conn := connection.New()
	db, err := conn.Open()
	if err != nil {
		return db, err
	}
	return db, nil
}
