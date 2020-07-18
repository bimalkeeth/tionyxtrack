package validation

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"regexp"
	pro "tionyxtrack/masterservice/proto"
)

func ValidateVehicle(method int, op *pro.VehicleProto) error {

	switch method {
	case CreateVehicle:
		return validation.ValidateStruct(op,
			validation.Field(op.FleetId, validation.Required, validation.Match(regexp.MustCompile(regex))),
			validation.Field(op.ModelId, validation.Required, validation.Match(regexp.MustCompile(regex))),
			validation.Field(op.MakeId, validation.Required, validation.Match(regexp.MustCompile(regex))),
			validation.Field(op.StatusId, validation.Required, validation.Match(regexp.MustCompile(regex))),
			validation.Field(op.OfficeName, validation.Required, validation.Nil.When(len(op.OfficeName) > 0)),
		)
	case UpdateVehicle:
		return validation.ValidateStruct(op,
			validation.Field(op.Id, validation.Required, validation.Match(regexp.MustCompile(regex))),
			validation.Field(op.FleetId, validation.Required, validation.Match(regexp.MustCompile(regex))),
			validation.Field(op.ModelId, validation.Required, validation.Match(regexp.MustCompile(regex))),
			validation.Field(op.MakeId, validation.Required, validation.Match(regexp.MustCompile(regex))),
			validation.Field(op.StatusId, validation.Required, validation.Match(regexp.MustCompile(regex))),
			validation.Field(op.OfficeName, validation.Required, validation.Nil.When(len(op.OfficeName) > 0)),
		)
	}
	return nil
}

func ValidateVehicleHistory(method int, op *pro.VehicleHistoryProto) error {

	switch method {
	case CreateVehicleHistory:
		return validation.ValidateStruct(op,

			validation.Field(op.VehicleId, validation.Required, validation.Match(regexp.MustCompile(regex))),
			validation.Field(op.ChangeDate, validation.Required, validation.By(IsDate)),
			validation.Field(op.FromStatusId, validation.Required, validation.Match(regexp.MustCompile(regex))),
			validation.Field(op.ToStatusId, validation.Required, validation.Match(regexp.MustCompile(regex))),
			validation.Field(op.OfficerName, validation.Required, validation.Nil.When(len(op.OfficerName) > 0)),
		)
	case UpdateVehicleHistory:
		return validation.ValidateStruct(op,
			validation.Field(op.Id, validation.Required, validation.Match(regexp.MustCompile(regex))),
			validation.Field(op.VehicleId, validation.Required, validation.Match(regexp.MustCompile(regex))),
			validation.Field(op.ChangeDate, validation.Required, validation.By(IsDate)),
			validation.Field(op.FromStatusId, validation.Required, validation.Match(regexp.MustCompile(regex))),
			validation.Field(op.ToStatusId, validation.Required, validation.Match(regexp.MustCompile(regex))),
			validation.Field(op.OfficerName, validation.Required, validation.Nil.When(len(op.OfficerName) > 0)),
		)
	}
	return nil
}

func ValidateVehicleLocation(method int, op *pro.VehicleAddressProto) error {

	switch method {
	case CreateVehicleHistory:
		return validation.ValidateStruct(op,
			validation.Field(op.VehicleId, validation.Required, validation.Match(regexp.MustCompile(regex))),
			validation.Field(op.AddressId, validation.Required, validation.Match(regexp.MustCompile(regex))),
		)
	case UpdateVehicleHistory:
		return validation.ValidateStruct(op,
			validation.Field(op.Id, validation.Required, validation.Match(regexp.MustCompile(regex))),
			validation.Field(op.VehicleId, validation.Required, validation.Match(regexp.MustCompile(regex))),
			validation.Field(op.AddressId, validation.Required, validation.Match(regexp.MustCompile(regex))),
		)
	}
	return nil
}

func ValidateVehicleMake(method int, op *pro.VehicleMakeProto) error {

	switch method {
	case CreateVehicleMake:
		return validation.ValidateStruct(op,
			validation.Field(op.CountryId, validation.Required, validation.Match(regexp.MustCompile(regex))),
			validation.Field(op.Make, validation.Required, validation.Nil.When(len(op.Make) > 0)),
		)
	case UpdateVehicleMake:
		return validation.ValidateStruct(op,
			validation.Field(op.Id, validation.Required, validation.Match(regexp.MustCompile(regex))),
			validation.Field(op.CountryId, validation.Required, validation.Match(regexp.MustCompile(regex))),
			validation.Field(op.Make, validation.Required, validation.Nil.When(len(op.Make) > 0)),
		)
	}
	return nil
}

func ValidateVehicleModel(method int, op *pro.VehicleModelProto) error {

	switch method {
	case CreateVehicleModel:
		return validation.ValidateStruct(op,
			validation.Field(op.MakeId, validation.Required, validation.Match(regexp.MustCompile(regex))),
			validation.Field(op.ModelName, validation.Required, validation.Nil.When(len(op.ModelName) > 0)),
		)
	case UpdateVehicleModel:
		return validation.ValidateStruct(op,
			validation.Field(op.Id, validation.Required, validation.Match(regexp.MustCompile(regex))),
			validation.Field(op.MakeId, validation.Required, validation.Match(regexp.MustCompile(regex))),
			validation.Field(op.ModelName, validation.Required, validation.Nil.When(len(op.ModelName) > 0)),
		)
	}
	return nil
}

func ValidateVehicleReg(method int, op *pro.VehicleTrackRegProto) error {

	switch method {
	case CreateVehicleReg:
		return validation.ValidateStruct(op,
			validation.Field(op.VehicleId, validation.Required, validation.Match(regexp.MustCompile(regex))),
			validation.Field(op.Duration, validation.Required, validation.By(GreaterThan)),
			validation.Field(op.RegisterDate, validation.Required, validation.By(IsDate)),
			validation.Field(op.ExpiredDate, validation.Required, validation.By(IsDate)),
		)
	case UpdateVehicleReg:
		return validation.ValidateStruct(op,
			validation.Field(op.Id, validation.Required, validation.Match(regexp.MustCompile(regex))),
			validation.Field(op.VehicleId, validation.Required, validation.Match(regexp.MustCompile(regex))),
			validation.Field(op.Duration, validation.Required, validation.By(GreaterThan)),
			validation.Field(op.RegisterDate, validation.Required, validation.By(IsDate)),
			validation.Field(op.ExpiredDate, validation.Required, validation.By(IsDate)),
		)
	}
	return nil
}

func ValidateVehicleStatus(method int, op *pro.VehicleStatusProto) error {

	switch method {
	case CreateVehicleStatus:
		return validation.ValidateStruct(op,
			validation.Field(op.StatusType, validation.Required, validation.Nil.When(len(op.StatusType) > 0)),
			validation.Field(op.StatusName, validation.Required, validation.Nil.When(len(op.StatusName) > 0)),
		)
	case UpdateVehicleStatus:
		return validation.ValidateStruct(op,
			validation.Field(op.Id, validation.Required, validation.Match(regexp.MustCompile(regex))),
			validation.Field(op.StatusType, validation.Required, validation.Nil.When(len(op.StatusType) > 0)),
			validation.Field(op.StatusName, validation.Required, validation.Nil.When(len(op.StatusName) > 0)),
		)
	}
	return nil
}

func ValidateVehicleOperatorBound(method int, op *pro.VehicleOperatorBoundProto) error {

	switch method {
	case CreateVehicleOpBound:
		return validation.ValidateStruct(op,
			validation.Field(op.VehicleId, validation.Required, validation.Match(regexp.MustCompile(regex))),
			validation.Field(op.OperatorId, validation.Required, validation.Match(regexp.MustCompile(regex))),
		)
	case UpdateVehicleOpBound:
		return validation.ValidateStruct(op,
			validation.Field(op.Id, validation.Required, validation.Match(regexp.MustCompile(regex))),
			validation.Field(op.VehicleId, validation.Required, validation.Match(regexp.MustCompile(regex))),
			validation.Field(op.OperatorId, validation.Required, validation.Match(regexp.MustCompile(regex))),
		)
	}
	return nil
}
