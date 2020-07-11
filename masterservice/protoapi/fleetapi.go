package protoapi

import (
	"context"
	timestamp "github.com/golang/protobuf/ptypes"
	bu "tionyxtrack/masterservice/businesscontracts"
	flt "tionyxtrack/masterservice/manager/fleets"
	pro "tionyxtrack/masterservice/proto"
)


func (m *MasterService) CreateFleet(ctx context.Context, in *pro.RequestFleet, out *pro.ResponseCreateSuccess) error {
	fltManager := flt.New()

	dateRegistered, err := timestamp.Timestamp(in.Fleet.DateRegistered)
	if err != nil {
		out.Errors = ErrorResponse.GetCreateErrorJson(err)
		return nil
	}
	result, errs := fltManager.CreateFleet(bu.FleetBO{
		Name:                 in.Fleet.Name,
		CountryId:            uint(in.Fleet.CountryId),
		FleetCode:            in.Fleet.FleetCode,
		SurName:              in.Fleet.SurName,
		OtherName:            in.Fleet.OtherName,
		RegistrationDuration: float64(in.Fleet.RegistrationDuration),
		DateRegistered:       dateRegistered,
	})
	out.Errors = ErrorResponse.GetCreateErrorJson(errs)
	out.Id = uint64(result.Id)
	return nil
}

func (m *MasterService) UpdateFleet(ctx context.Context, in *pro.RequestFleet, out *pro.ResponseSuccess) error {
	fltManager := flt.New()
	dateRegistered, err := timestamp.Timestamp(in.Fleet.DateRegistered)
	if err != nil {
		out.Errors = ErrorResponse.GetCreateErrorJson(err)
		return nil
	}
	result, errs := fltManager.UpdateFleet(bu.FleetBO{
		Id:                   uint(in.Fleet.Id),
		Name:                 in.Fleet.Name,
		CountryId:            uint(in.Fleet.CountryId),
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
	fltManager := flt.New()
	result, err := fltManager.DeleteFleet(uint(in.Id))
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Success = result
	return nil
}

func (m *MasterService) GetFleetById(ctx context.Context, in *pro.RequestKey, out *pro.ResponseFleet) error {
	fltManager := flt.New()
	result, err := fltManager.GetFleetById(uint(in.Id))
	out.Errors = ErrorResponse.GetCreateErrorJson(err)

	updatedat, _ := timestamp.TimestampProto(result.UpdatedAt)
	dateRegistered, _ := timestamp.TimestampProto(result.DateRegistered)

	out.Fleet = append(out.Fleet, &pro.FleetProto{
		Id:                   uint64(result.Id),
		UpdatedAt:            updatedat,
		FleetCode:            result.FleetCode,
		Name:                 result.Name,
		SurName:              result.SurName,
		OtherName:            result.OtherName,
		DateRegistered:       dateRegistered,
		RegistrationDuration: float32(result.RegistrationDuration),
		FleetContactId:       uint64(result.FleetContactId),
		FleetLocationId:      uint64(result.FleetLocationId),
		CountryId:            uint64(result.CountryId),
	})
	return nil
}

func (m *MasterService) CreateFleetContact(ctx context.Context, in *pro.RequestFleetContact, out *pro.ResponseCreateSuccess) error {
	fltManager := flt.New()
	result, err := fltManager.CreateFleetContact(uint(in.FleetId), uint(in.ContactId), in.Primary)
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Id = uint64(result)
	return nil
}

func (m *MasterService) UpdateFleetContact(ctx context.Context, in *pro.RequestFleetContact, out *pro.ResponseSuccess) error {
	fltManager := flt.New()
	result, err := fltManager.UpdateFleetContact(uint(in.Id), uint(in.FleetId), uint(in.ContactId), in.Primary)
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Success = result
	return nil
}

func (m *MasterService) DeleteFleetContact(ctx context.Context, in *pro.RequestDelete, out *pro.ResponseSuccess) error {
	fltManager := flt.New()
	result, err := fltManager.DeleteFleetContact(uint(in.Id))
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Success = result
	return nil

}

func (m *MasterService) GetContactByFleetId(ctx context.Context, in *pro.RequestKey, out *pro.ResponseFleetContact) error {
	fltManager := flt.New()
	result, err := fltManager.GetContactByFleetId(uint(in.Id))
	out.Errors = ErrorResponse.GetCreateErrorJson(err)

	for _, con := range result {

		updatedat, _ := timestamp.TimestampProto(con.Fleet.UpdatedAt)
		dateRegistered, _ := timestamp.TimestampProto(con.Fleet.DateRegistered)
		contact := &pro.FleetContactProto{
			Id:        uint64(con.Id),
			FleetId:   uint64(con.FleetId),
			ContactId: uint64(con.ContactId),
			Primary:   con.Primary,
		}
		contact.Fleet = &pro.FleetProto{
			Id:             uint64(con.Fleet.Id),
			FleetCode:      con.Fleet.FleetCode,
			Name:           con.Fleet.Name,
			SurName:        con.Fleet.SurName,
			OtherName:      con.Fleet.OtherName,
			DateRegistered: dateRegistered,
			UpdatedAt:      updatedat,
		}
		contact.Contact = &pro.ContactProto{
			Id:            uint64(con.Contact.Id),
			Contact:       con.Contact.Contact,
			ContactTypeId: uint64(con.Contact.ContactTypeId),
		}
		out.FleetContact = append(out.FleetContact, contact)
	}
	return nil
}

func (m *MasterService) CreateFleetLocation(ctx context.Context, in *pro.RequestFleetLocation, out *pro.ResponseCreateSuccess) error {
	fltManager := flt.New()
	result, err := fltManager.CreateFleetLocation(uint(in.FleetId), uint(in.AddressId), in.Primary)
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Id = uint64(result)
	return nil
}

func (m *MasterService) UpdateFleetLocation(ctx context.Context, in *pro.RequestFleetLocation, out *pro.ResponseSuccess) error {
	fltManager := flt.New()
	result, err := fltManager.UpdateFleetLocation(uint(in.Id), uint(in.FleetId), uint(in.AddressId), in.Primary)
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Success = result
	return nil
}

func (m *MasterService) DeleteFleetLocation(ctx context.Context, in *pro.RequestDelete, out *pro.ResponseSuccess) error {
	fltManager := flt.New()
	result, err := fltManager.DeleteFleetLocation(uint(in.Id))
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Success = result
	return nil
}

func (m *MasterService) GetLocationByFleetId(ctx context.Context, in *pro.RequestKey, out *pro.ResponseFleetLocation) error {
	fltManager := flt.New()
	result, err := fltManager.GetLocationByFleetId(uint(in.Id))
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	for _, loc := range result {
		item := &pro.FleetLocationProto{
			Id:        uint64(loc.Id),
			FleetId:   uint64(loc.FleetId),
			AddressId: uint64(loc.AddressId),
			Primary:   loc.Primary,
		}
		updatedat, _ := timestamp.TimestampProto(loc.Address.UpdatedAt)
		item.Address = &pro.AddressProto{
			Id:            uint64(loc.Address.Id),
			Address:       loc.Address.Address,
			Street:        loc.Address.Street,
			Suburb:        loc.Address.Suburb,
			StateId:       uint64(loc.Address.StateId),
			CountryId:     uint64(loc.Address.CountryId),
			AddressTypeId: uint64(loc.Address.AddressTypeId),
			Location:      loc.Address.Location,
			UpdatedAt:     updatedat,
		}
		fltUpdatedAt, _ := timestamp.TimestampProto(loc.Fleet.UpdatedAt)
		dateRegisterAt, _ := timestamp.TimestampProto(loc.Fleet.DateRegistered)
		item.Fleet = &pro.FleetProto{
			Id:             uint64(loc.Fleet.Id),
			CountryId:      uint64(loc.Fleet.CountryId),
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
