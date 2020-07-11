package bucontracts

type CompanyBO struct {
	Id        uint
	Name      string
	AddressId uint
	ContactId uint
	Address   AddressBO
	Contact   ContactBO
}
