package connection

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

type IDatabase interface {
	Open() (*gorm.DB, error)
}
type DB struct{}

func New() IDatabase {
	return &DB{}
}

func (db DB) Open() (*gorm.DB, error) {
	pdb, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=teonyx_auth password=root sslmode=disable")
	if err != nil {
		log.Fatal("error in connecting to the database->", err)
		return nil, err
	}
	return pdb, nil
}
