package bucontracts

type FleetContactBO struct {
	Id        uint
	FleetId   uint
	ContactId uint
	Primary   bool
	Fleet     FleetBO
	Contact   ContactBO
}
