package validation

import (
	"github.com/go-ozzo/ozzo-validation/v4"
	"regexp"
	pro "tionyxtrack/masterservice/proto"
)

func ValidateOperator(method int, op *pro.OperatorProto) error {

	switch method {
	case CreateOperator:
		return validation.ValidateStruct(op,
			validation.Field(op.Name, validation.Required, validation.Nil.When(len(op.Name) > 0)),
			validation.Field(op.SurName, validation.Required, validation.Nil.When(len(op.Name) > 0)),
			validation.Field(op.DrivingLic, validation.Required, validation.Nil.When(len(op.Name) > 0)),
		)
	case UpdateOperator:
		return validation.ValidateStruct(op,
			validation.Field(op.Id, validation.Required, validation.Match(regexp.MustCompile(regex))),
			validation.Field(op.Name, validation.Required, validation.Nil.When(len(op.Name) > 0)),
			validation.Field(op.SurName, validation.Required, validation.Nil.When(len(op.Name) > 0)),
			validation.Field(op.DrivingLic, validation.Required, validation.Nil.When(len(op.Name) > 0)),
		)
	}
	return nil

}

func ValidateOperatorContract(method int, op *pro.RequestOperatorContact) error {

	switch method {
	case CreateOperatorContract:
		return validation.ValidateStruct(op,
			validation.Field(op.ContactId, validation.Required, validation.Match(regexp.MustCompile(regex))),
			validation.Field(op.OperatorId, validation.Required, validation.Match(regexp.MustCompile(regex))),
		)
	case UpdateOperatorContract:
		return validation.ValidateStruct(op,
			validation.Field(op.Id, validation.Required, validation.Match(regexp.MustCompile(regex))),
			validation.Field(op.ContactId, validation.Required, validation.Match(regexp.MustCompile(regex))),
			validation.Field(op.OperatorId, validation.Required, validation.Match(regexp.MustCompile(regex))),
		)
	}
	return nil
}

func ValidateOperatorLocation(method int, op *pro.OperatorLocationProto) error {

	switch method {
	case CreateOperatorLocation:
		return validation.ValidateStruct(op,
			validation.Field(op.AddressId, validation.Required, validation.Match(regexp.MustCompile(regex))),
			validation.Field(op.OperatorId, validation.Required, validation.Match(regexp.MustCompile(regex))),
		)
	case UpdateOperatorLocation:
		return validation.ValidateStruct(op,
			validation.Field(op.Id, validation.Required, validation.Match(regexp.MustCompile(regex))),
			validation.Field(op.AddressId, validation.Required, validation.Match(regexp.MustCompile(regex))),
			validation.Field(op.OperatorId, validation.Required, validation.Match(regexp.MustCompile(regex))),
		)
	}
	return nil
}
