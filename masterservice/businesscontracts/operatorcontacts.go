package bucontracts

type OperatorContactsBO struct {
	Id         uint
	ContactId  uint
	OperatorId uint
	Primary    bool
	Contact    ContactBO
}
