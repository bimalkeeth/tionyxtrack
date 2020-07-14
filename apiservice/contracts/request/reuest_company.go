package request

type CompanyRequest struct {
	Id        uint64 `json:"id"`
	Name      string `json:"name"`
	AddressId uint64 `json:"addressid"`
	ContactId uint64 `json:"contactid"`
}
