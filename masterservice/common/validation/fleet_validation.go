package validation

import (
	"github.com/go-ozzo/ozzo-validation/v4"
	"regexp"
	pro "tionyxtrack/masterservice/proto"
)

func ValidateFleet(method int, op *pro.FleetProto) error {

	switch method {
	case CreateFleet:
		return validation.ValidateStruct(op,
			validation.Field(op.Name, validation.Required, validation.Nil.When(len(op.Name) > 0)),
			validation.Field(op.SurName, validation.Required, validation.Nil.When(len(op.SurName) > 0)),
			validation.Field(op.CountryId, validation.Required, validation.Match(regexp.MustCompile(regex))),
			validation.Field(op.FleetCode, validation.Required, validation.Nil.When(len(op.FleetCode) > 0)),
		)
	case UpdateFleet:
		return validation.ValidateStruct(op,
			validation.Field(op.Id, validation.Required, validation.Match(regexp.MustCompile(regex))),
			validation.Field(op.Name, validation.Required, validation.Nil.When(len(op.Name) > 0)),
			validation.Field(op.SurName, validation.Required, validation.Nil.When(len(op.SurName) > 0)),
			validation.Field(op.CountryId, validation.Required, validation.Match(regexp.MustCompile(regex))),
			validation.Field(op.FleetCode, validation.Required, validation.Nil.When(len(op.FleetCode) > 0)),
		)
	}
	return nil
}

func ValidateFleetContract(method int, op *pro.RequestFleetContact) error {
	switch method {
	case CreateFleetContract:
		return validation.ValidateStruct(op,
			validation.Field(op.ContactId, validation.Required, validation.Match(regexp.MustCompile(regex))),
			validation.Field(op.FleetId, validation.Required, validation.Match(regexp.MustCompile(regex))),
		)
	case UpdateFleetContract:
		return validation.ValidateStruct(op,
			validation.Field(op.Id, validation.Required, validation.Match(regexp.MustCompile(regex))),
			validation.Field(op.ContactId, validation.Required, validation.Match(regexp.MustCompile(regex))),
			validation.Field(op.FleetId, validation.Required, validation.Match(regexp.MustCompile(regex))),
		)
	}
	return nil
}

func ValidateFleetLocation(method int, op *pro.RequestFleetLocation) error {
	switch method {
	case CreateFleetLocation:
		return validation.ValidateStruct(op,
			validation.Field(op.AddressId, validation.Required, validation.Match(regexp.MustCompile(regex))),
			validation.Field(op.FleetId, validation.Required, validation.Match(regexp.MustCompile(regex))),
		)
	case UpdateFleetLocation:
		return validation.ValidateStruct(op,
			validation.Field(op.Id, validation.Required, validation.Match(regexp.MustCompile(regex))),
			validation.Field(op.AddressId, validation.Required, validation.Match(regexp.MustCompile(regex))),
			validation.Field(op.FleetId, validation.Required, validation.Match(regexp.MustCompile(regex))),
		)
	}
	return nil
}
