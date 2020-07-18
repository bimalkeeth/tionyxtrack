package protoapi

import (
	"context"
	timestamp "github.com/golang/protobuf/ptypes"
	uuid "github.com/satori/go.uuid"
	bu "tionyxtrack/masterservice/businesscontracts"
	val "tionyxtrack/masterservice/common/validation"
	flt "tionyxtrack/masterservice/manager/fleets"
	pro "tionyxtrack/masterservice/proto"
)

func (m *MasterService) CreateFleet(ctx context.Context, in *pro.RequestFleet, out *pro.ResponseCreateSuccess) error {

	defer RecoverError()

	if err := val.ValidateFleet(val.CreateFleet, in.Fleet); err != nil {
		out.Errors = ErrorResponse.GetCreateErrorJson(err)
		return nil
	}

	fltManager := flt.New()
	dateRegistered, err := timestamp.Timestamp(in.Fleet.DateRegistered)
	if err != nil {
		out.Errors = ErrorResponse.GetCreateErrorJson(err)
		return nil
	}
	countryId := uuid.FromStringOrNil(in.Fleet.CountryId)

	result, errs := fltManager.CreateFleet(bu.FleetBO{
		Name:                 in.Fleet.Name,
		CountryId:            countryId,
		FleetCode:            in.Fleet.FleetCode,
		SurName:              in.Fleet.SurName,
		OtherName:            in.Fleet.OtherName,
		RegistrationDuration: float64(in.Fleet.RegistrationDuration),
		DateRegistered:       dateRegistered,
	})
	out.Errors = ErrorResponse.GetCreateErrorJson(errs)
	out.Id = result.Id.String()
	return nil
}

func (m *MasterService) UpdateFleet(ctx context.Context, in *pro.RequestFleet, out *pro.ResponseSuccess) error {

	defer RecoverError()

	if err := val.ValidateFleet(val.UpdateFleet, in.Fleet); err != nil {
		out.Errors = ErrorResponse.GetCreateErrorJson(err)
		return nil
	}

	fltManager := flt.New()
	dateRegistered, err := timestamp.Timestamp(in.Fleet.DateRegistered)
	if err != nil {
		out.Errors = ErrorResponse.GetCreateErrorJson(err)
		return nil
	}

	result, errs := fltManager.UpdateFleet(bu.FleetBO{
		Id:                   uuid.FromStringOrNil(in.Fleet.Id),
		Name:                 in.Fleet.Name,
		CountryId:            uuid.FromStringOrNil(in.Fleet.CountryId),
		FleetCode:            in.Fleet.FleetCode,
		SurName:              in.Fleet.SurName,
		OtherName:            in.Fleet.OtherName,
		RegistrationDuration: float64(in.Fleet.RegistrationDuration),
		DateRegistered:       dateRegistered,
	})
	out.Errors = ErrorResponse.GetCreateErrorJson(errs)
	out.Success = result
	return nil
}

func (m *MasterService) DeleteFleet(ctx context.Context, in *pro.RequestDelete, out *pro.ResponseSuccess) error {

	defer RecoverError()

	if err := val.ValidateDelete(in); err != nil {
		out.Errors = ErrorResponse.GetCreateErrorJson(err)
		return nil
	}

	fltManager := flt.New()

	result, err := fltManager.DeleteFleet(uuid.FromStringOrNil(in.Id))
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Success = result
	return nil
}

func (m *MasterService) GetFleetById(ctx context.Context, in *pro.RequestKey, out *pro.ResponseFleet) error {

	defer RecoverError()

	fltManager := flt.New()

	result, err := fltManager.GetFleetById(uuid.FromStringOrNil(in.Id))
	if err != nil {
		out.Errors = ErrorResponse.GetCreateErrorJson(err)
		return nil
	}

	updatedat, _ := timestamp.TimestampProto(result.UpdatedAt)
	dateRegistered, _ := timestamp.TimestampProto(result.DateRegistered)

	out.Fleet = append(out.Fleet, &pro.FleetProto{
		Id:                   result.Id.String(),
		UpdatedAt:            updatedat,
		FleetCode:            result.FleetCode,
		Name:                 result.Name,
		SurName:              result.SurName,
		OtherName:            result.OtherName,
		DateRegistered:       dateRegistered,
		RegistrationDuration: float32(result.RegistrationDuration),
		FleetContactId:       result.FleetContactId.String(),
		FleetLocationId:      result.FleetLocationId.String(),
		CountryId:            result.CountryId.String(),
	})
	return nil
}

func (m *MasterService) CreateFleetContact(ctx context.Context, in *pro.RequestFleetContact, out *pro.ResponseCreateSuccess) error {

	defer RecoverError()

	if err := val.ValidateFleetContract(val.CreateFleetContract, in); err != nil {
		out.Errors = ErrorResponse.GetCreateErrorJson(err)
		return nil
	}

	fltManager := flt.New()

	result, err := fltManager.CreateFleetContact(uuid.FromStringOrNil(in.FleetId), uuid.FromStringOrNil(in.ContactId), in.Primary)
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Id = result.String()
	return nil
}

func (m *MasterService) UpdateFleetContact(ctx context.Context, in *pro.RequestFleetContact, out *pro.ResponseSuccess) error {

	if err := val.ValidateFleetContract(val.UpdateFleetContract, in); err != nil {
		out.Errors = ErrorResponse.GetCreateErrorJson(err)
		return nil
	}

	defer RecoverError()

	fltManager := flt.New()

	result, err := fltManager.UpdateFleetContact(uuid.FromStringOrNil(in.Id), uuid.FromStringOrNil(in.FleetId), uuid.FromStringOrNil(in.ContactId), in.Primary)
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Success = result
	return nil
}

func (m *MasterService) DeleteFleetContact(ctx context.Context, in *pro.RequestDelete, out *pro.ResponseSuccess) error {

	defer RecoverError()

	if err := val.ValidateDelete(in); err != nil {
		out.Errors = ErrorResponse.GetCreateErrorJson(err)
		return nil
	}

	fltManager := flt.New()

	result, err := fltManager.DeleteFleetContact(uuid.FromStringOrNil(in.Id))
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Success = result
	return nil

}

func (m *MasterService) GetContactByFleetId(ctx context.Context, in *pro.RequestKey, out *pro.ResponseFleetContact) error {

	defer RecoverError()

	fltManager := flt.New()
	result, err := fltManager.GetContactByFleetId(uuid.FromStringOrNil(in.Id))
	out.Errors = ErrorResponse.GetCreateErrorJson(err)

	for _, con := range result {
		updatedat, _ := timestamp.TimestampProto(con.Fleet.UpdatedAt)
		dateRegistered, _ := timestamp.TimestampProto(con.Fleet.DateRegistered)
		contact := &pro.FleetContactProto{
			Id:        con.Id.String(),
			FleetId:   con.FleetId.String(),
			ContactId: con.ContactId.String(),
			Primary:   con.Primary,
		}
		contact.Fleet = &pro.FleetProto{
			Id:             con.Fleet.Id.String(),
			FleetCode:      con.Fleet.FleetCode,
			Name:           con.Fleet.Name,
			SurName:        con.Fleet.SurName,
			OtherName:      con.Fleet.OtherName,
			DateRegistered: dateRegistered,
			UpdatedAt:      updatedat,
		}
		contact.Contact = &pro.ContactProto{
			Id:            con.Contact.Id.String(),
			Contact:       con.Contact.Contact,
			ContactTypeId: con.Contact.ContactTypeId.String(),
		}
		out.FleetContact = append(out.FleetContact, contact)
	}
	return nil
}

//---------------------------------------------------------
//Create Fleet location
//---------------------------------------------------------
func (m *MasterService) CreateFleetLocation(ctx context.Context, in *pro.RequestFleetLocation, out *pro.ResponseCreateSuccess) error {

	defer RecoverError()

	if err := val.ValidateFleetLocation(val.CreateFleetLocation, in); err != nil {
		out.Errors = ErrorResponse.GetCreateErrorJson(err)
		return nil
	}

	fltManager := flt.New()

	result, err := fltManager.CreateFleetLocation(uuid.FromStringOrNil(in.FleetId), uuid.FromStringOrNil(in.AddressId), in.Primary)
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Id = result.String()
	return nil
}

func (m *MasterService) UpdateFleetLocation(ctx context.Context, in *pro.RequestFleetLocation, out *pro.ResponseSuccess) error {

	if err := val.ValidateFleetLocation(val.UpdateFleetLocation, in); err != nil {
		out.Errors = ErrorResponse.GetCreateErrorJson(err)
		return nil
	}

	defer RecoverError()

	fltManager := flt.New()

	result, err := fltManager.UpdateFleetLocation(uuid.FromStringOrNil(in.Id), uuid.FromStringOrNil(in.FleetId), uuid.FromStringOrNil(in.AddressId), in.Primary)
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Success = result
	return nil
}

func (m *MasterService) DeleteFleetLocation(ctx context.Context, in *pro.RequestDelete, out *pro.ResponseSuccess) error {

	defer RecoverError()

	if err := val.ValidateDelete(in); err != nil {
		out.Errors = ErrorResponse.GetCreateErrorJson(err)
		return nil
	}

	fltManager := flt.New()

	result, err := fltManager.DeleteFleetLocation(uuid.FromStringOrNil(in.Id))
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Success = result
	return nil
}

func (m *MasterService) GetLocationByFleetId(ctx context.Context, in *pro.RequestKey, out *pro.ResponseFleetLocation) error {

	defer RecoverError()

	fltManager := flt.New()

	result, err := fltManager.GetLocationByFleetId(uuid.FromStringOrNil(in.Id))
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	for _, loc := range result {
		item := &pro.FleetLocationProto{
			Id:        loc.Id.String(),
			FleetId:   loc.FleetId.String(),
			AddressId: loc.AddressId.String(),
			Primary:   loc.Primary,
		}
		updatedat, _ := timestamp.TimestampProto(loc.Address.UpdatedAt)
		item.Address = &pro.AddressProto{
			Id:            loc.Address.Id.String(),
			Address:       loc.Address.Address,
			Street:        loc.Address.Street,
			Suburb:        loc.Address.Suburb,
			StateId:       loc.Address.StateId.String(),
			CountryId:     loc.Address.CountryId.String(),
			AddressTypeId: loc.Address.AddressTypeId.String(),
			Location:      loc.Address.Location,
			UpdatedAt:     updatedat,
		}
		fltUpdatedAt, _ := timestamp.TimestampProto(loc.Fleet.UpdatedAt)
		dateRegisterAt, _ := timestamp.TimestampProto(loc.Fleet.DateRegistered)
		item.Fleet = &pro.FleetProto{
			Id:             loc.Fleet.Id.String(),
			CountryId:      loc.Fleet.CountryId.String(),
			DateRegistered: dateRegisterAt,
			OtherName:      loc.Fleet.OtherName,
			SurName:        loc.Fleet.SurName,
			Name:           loc.Fleet.Name,
			UpdatedAt:      fltUpdatedAt,
		}
		out.FleetLocation = append(out.FleetLocation, item)
	}
	return nil
}
