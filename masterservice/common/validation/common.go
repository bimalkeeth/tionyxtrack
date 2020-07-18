package validation

import (
	"errors"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/golang/protobuf/ptypes/timestamp"
	"regexp"
	pro "tionyxtrack/masterservice/proto"
)

const (
	CreateFleet            = 1
	UpdateFleet            = 2
	CreateFleetContract    = 4
	UpdateFleetContract    = 5
	CreateFleetLocation    = 6
	UpdateFleetLocation    = 7
	CreateCompany          = 8
	UpdateCompany          = 9
	CreateAddressType      = 10
	UpdateAddressType      = 11
	CreateRegion           = 12
	UpdateRegion           = 13
	CreateState            = 14
	UpdateState            = 15
	CreateContact          = 16
	UpdateContact          = 17
	CreateAddress          = 18
	UpdateAddress          = 19
	CreateOperator         = 20
	UpdateOperator         = 21
	CreateOperatorContract = 22
	UpdateOperatorContract = 23
	CreateOperatorLocation = 24
	UpdateOperatorLocation = 25
	CreateVehicle          = 26
	UpdateVehicle          = 27
	CreateVehicleHistory   = 28
	UpdateVehicleHistory   = 29
	CreateVehicleLocation  = 30
	UpdateVehicleLocation  = 31
	CreateVehicleMake      = 32
	UpdateVehicleMake      = 33
	CreateVehicleModel     = 34
	UpdateVehicleModel     = 35
	CreateVehicleReg       = 36
	UpdateVehicleReg       = 37
	CreateVehicleStatus    = 38
	UpdateVehicleStatus    = 39
	CreateVehicleOpBound   = 40
	UpdateVehicleOpBound   = 41
)

var regex = "/^[0-9A-F]{8}-[0-9A-F]{4}-[4][0-9A-F]{3}-[89AB][0-9A-F]{3}-[0-9A-F]{12}$/i"

func ValidateDelete(op *pro.RequestDelete) error {
	return validation.ValidateStruct(op,
		validation.Field(op.Id, validation.Required, validation.Match(regexp.MustCompile(regex))),
	)
}

func ValidateRequest(op *pro.RequestKey) error {
	return validation.ValidateStruct(op,
		validation.Field(op.Id, validation.Required, validation.Match(regexp.MustCompile(regex))),
	)
}

func IsDate(value interface{}) error {
	s, _ := value.(*timestamp.Timestamp)
	if !s.IsValid() {
		return errors.New("date is not valid")
	}
	return nil
}

func GreaterThan(value interface{}) error {
	s, _ := value.(int)
	if s == 0 {
		return errors.New("value should be  > 0 ")
	}
	return nil
}
