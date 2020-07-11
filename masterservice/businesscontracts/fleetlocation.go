package bucontracts

type FleetAddressBO struct {
	Id        uint
	FleetId   uint
	AddressId uint
	Primary   bool
	Fleet     FleetBO
	Address   AddressBO
}
