package mappers

import (
	"log"
	con "tionyxtrack/authservice/connection"
)

type IEntityMapper interface {
	GenerateSchema() error
}

type SchemaGenerator struct{}

func New() IEntityMapper {
	return SchemaGenerator{}
}

func (t SchemaGenerator) GenerateSchema() error {
	connection := con.New()
	dbase, err := connection.Open()
	if err != nil {
		log.Fatal("error in connection")
	}
	MapApplicationMenuTable(dbase)
	return nil
}
