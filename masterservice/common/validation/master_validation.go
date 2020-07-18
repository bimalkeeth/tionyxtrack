package validation

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"regexp"
	pro "tionyxtrack/masterservice/proto"
)

func ValidateCompany(method int, op *pro.CompanyProto) error {

	switch method {
	case CreateCompany:
		return validation.ValidateStruct(op,
			validation.Field(op.Name, validation.Required, validation.Nil.When(len(op.Name) > 0)),
			validation.Field(op.AddressId, validation.Required, validation.Match(regexp.MustCompile(regex))),
			validation.Field(op.ContactId, validation.Required, validation.Match(regexp.MustCompile(regex))),
		)
	case UpdateCompany:
		return validation.ValidateStruct(op,
			validation.Field(op.Id, validation.Required, validation.Match(regexp.MustCompile(regex))),
			validation.Field(op.Name, validation.Required, validation.Nil.When(len(op.Name) > 0)),
			validation.Field(op.AddressId, validation.Required, validation.Match(regexp.MustCompile(regex))),
			validation.Field(op.ContactId, validation.Required, validation.Match(regexp.MustCompile(regex))),
		)
	}
	return nil
}

func ValidateAddressType(method int, op *pro.AddressTypeProto) error {

	switch method {
	case CreateAddressType:
		return validation.ValidateStruct(op,
			validation.Field(op.Name, validation.Required, validation.Nil.When(len(op.Name) > 0)),
		)
	case UpdateAddressType:
		return validation.ValidateStruct(op,
			validation.Field(op.Id, validation.Required, validation.Match(regexp.MustCompile(regex))),
			validation.Field(op.Name, validation.Required, validation.Nil.When(len(op.Name) > 0)),
		)
	}
	return nil
}

func ValidateRegion(method int, op *pro.RegionProto) error {

	switch method {
	case CreateRegion:
		return validation.ValidateStruct(op,
			validation.Field(op.RegionName, validation.Required, validation.Nil.When(len(op.RegionName) > 0)),
			validation.Field(op.Region, validation.Required, validation.Nil.When(len(op.Region) > 0)),
		)
	case UpdateRegion:
		return validation.ValidateStruct(op,
			validation.Field(op.Id, validation.Required, validation.Match(regexp.MustCompile(regex))),
			validation.Field(op.RegionName, validation.Required, validation.Nil.When(len(op.RegionName) > 0)),
			validation.Field(op.Region, validation.Required, validation.Nil.When(len(op.Region) > 0)),
		)
	}
	return nil
}

func ValidateState(method int, op *pro.StateProto) error {

	switch method {
	case CreateState:
		return validation.ValidateStruct(op,
			validation.Field(op.Name, validation.Required, validation.Nil.When(len(op.Name) > 0)),
			validation.Field(op.CountryId, validation.Required, validation.Match(regexp.MustCompile(regex))),
		)
	case UpdateRegion:
		return validation.ValidateStruct(op,
			validation.Field(op.Id, validation.Required, validation.Match(regexp.MustCompile(regex))),
			validation.Field(op.Name, validation.Required, validation.Nil.When(len(op.Name) > 0)),
			validation.Field(op.CountryId, validation.Required, validation.Match(regexp.MustCompile(regex))),
		)
	}
	return nil
}

func ValidateContact(method int, op *pro.ContactProto) error {

	switch method {
	case CreateContact:
		return validation.ValidateStruct(op,
			validation.Field(op.Contact, validation.Required, validation.Nil.When(len(op.Contact) > 0)),
			validation.Field(op.ContactTypeId, validation.Required, validation.Match(regexp.MustCompile(regex))),
		)
	case UpdateContact:
		return validation.ValidateStruct(op,
			validation.Field(op.Id, validation.Required, validation.Match(regexp.MustCompile(regex))),
			validation.Field(op.Contact, validation.Required, validation.Nil.When(len(op.Contact) > 0)),
			validation.Field(op.ContactTypeId, validation.Required, validation.Match(regexp.MustCompile(regex))),
		)
	}
	return nil
}

func ValidateAddress(method int, op *pro.AddressProto) error {

	switch method {
	case CreateAddress:
		return validation.ValidateStruct(op,
			validation.Field(op.Address, validation.Required, validation.Nil.When(len(op.Address) > 0)),
			validation.Field(op.CountryId, validation.Required, validation.Match(regexp.MustCompile(regex))),
			validation.Field(op.StateId, validation.Required, validation.Match(regexp.MustCompile(regex))),
			validation.Field(op.AddressTypeId, validation.Required, validation.Match(regexp.MustCompile(regex))),
			validation.Field(op.Location, validation.Required),
			validation.Field(op.Street, validation.Required),
			validation.Field(op.Suburb, validation.Required),
		)
	case UpdateAddress:
		return validation.ValidateStruct(op,
			validation.Field(op.Id, validation.Required, validation.Match(regexp.MustCompile(regex))),
			validation.Field(op.Address, validation.Required, validation.Nil.When(len(op.Address) > 0)),
			validation.Field(op.CountryId, validation.Required, validation.Match(regexp.MustCompile(regex))),
			validation.Field(op.StateId, validation.Required, validation.Match(regexp.MustCompile(regex))),
			validation.Field(op.AddressTypeId, validation.Required, validation.Match(regexp.MustCompile(regex))),
			validation.Field(op.Location, validation.Required),
			validation.Field(op.Street, validation.Required),
			validation.Field(op.Suburb, validation.Required),
		)
	}
	return nil
}
