package mappers

import (
	"log"
con "tionyxtrack/masterservice/connection"
)

type IEntityMapper interface {
	GenerateSchema() error
}

type SchemaGenerator struct{}

//-------------------------------------
// Create instance
//-------------------------------------
func New() SchemaGenerator {
	return SchemaGenerator{}
}

//--------------------------------------
// Create mapping
//--------------------------------------
func (t SchemaGenerator) GenerateSchema() error {
	db := con.New()
	dbase, err := db.Open()
	if err != nil {
		log.Fatal("error in connection")
	}
	MapAddressTypeTable(dbase)
	MapContactTypeTable(dbase)
	MapRegionTable(dbase)
	MapVehicleStatusTable(dbase)
	MapCountryTable(dbase)
	MapStatesTable(dbase)
	MapAddressTable(dbase)
	MapContactsTable(dbase)
	MapVehicleMakeTable(dbase)
	MapVehicleModelTable(dbase)
	MapCompanyTable(dbase)
	MapFleetTable(dbase)
	MapFleetContactTable(dbase)
	MapFleetLocationTable(dbase)
	MapVehicleTable(dbase)
	MapVehicleLocationTable(dbase)
	MapVehicleTrackRegTable(dbase)
	MapVehicleHistoryTable(dbase)
	MapVehicleOperatorsTable(dbase)
	MapVehicleOperatorBoundTable(dbase)
	MapVehicleOperatorLocationTable(dbase)
	MapVehicleOperatorContactsTable(dbase)
	return nil
}
