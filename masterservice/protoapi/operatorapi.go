package protoapi

import (
	"context"
	timestamp "github.com/golang/protobuf/ptypes"
	uuid "github.com/satori/go.uuid"
	bu "tionyxtrack/masterservice/businesscontracts"
	val "tionyxtrack/masterservice/common/validation"
	opr "tionyxtrack/masterservice/manager/operator"
	pro "tionyxtrack/masterservice/proto"
)

func (m *MasterService) CreateOperator(ctx context.Context, req *pro.RequestOperator, out *pro.ResponseCreateSuccess) error {

	defer RecoverError()

	oprManager := opr.New()

	if err := val.ValidateOperator(val.CreateOperator, req.Operator); err != nil {
		out.Errors = ErrorResponse.GetCreateErrorJson(err)
		return nil
	}

	result, err := oprManager.CreateOperator(bu.OperatorBO{
		Name:       req.Operator.Name,
		SurName:    req.Operator.SurName,
		DrivingLic: req.Operator.DrivingLic,
	})
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Id = result.String()
	return nil
}

func (m *MasterService) UpdateOperator(ctx context.Context, req *pro.RequestOperator, out *pro.ResponseSuccess) error {

	defer RecoverError()

	if err := val.ValidateOperator(val.UpdateOperator, req.Operator); err != nil {
		out.Errors = ErrorResponse.GetCreateErrorJson(err)
		return nil
	}

	oprManager := opr.New()
	result, err := oprManager.UpdateOperator(bu.OperatorBO{
		Id:         uuid.FromStringOrNil(req.Operator.Id),
		Name:       req.Operator.Name,
		SurName:    req.Operator.SurName,
		Active:     req.Operator.Active,
		DrivingLic: req.Operator.DrivingLic,
	})
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Success = result
	return nil
}

func (m *MasterService) DeleteOperator(ctx context.Context, req *pro.RequestDelete, out *pro.ResponseSuccess) error {

	defer RecoverError()

	if err := val.ValidateDelete(req); err != nil {
		out.Errors = ErrorResponse.GetCreateErrorJson(err)
		return nil
	}

	oprManager := opr.New()
	result, err := oprManager.DeleteOperator(uuid.FromStringOrNil(req.Id))
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Success = result
	return nil
}

func (m *MasterService) GetOperatorById(ctx context.Context, req *pro.RequestKey, out *pro.ResponseOperator) error {

	defer RecoverError()

	if err := val.ValidateRequest(req); err != nil {
		out.Errors = ErrorResponse.GetCreateErrorJson(err)
		return nil
	}

	oprManager := opr.New()
	result, err := oprManager.GetOperatorById(uuid.FromStringOrNil(req.Id))
	out.Errors = ErrorResponse.GetCreateErrorJson(err)

	data := &pro.OperatorProto{}
	data.Id = result.Id.String()
	data.Active = result.Active
	data.DrivingLic = result.DrivingLic
	data.SurName = result.SurName
	data.Name = result.Name

	for _, con := range result.Contacts {
		updateCon, _ := timestamp.TimestampProto(con.Contact.UpdatedAt)
		contact := &pro.OperatorContactsProto{
			Id:        con.Id.String(),
			Primary:   con.Primary,
			ContactId: con.ContactId.String(),
			Contact: &pro.ContactProto{
				Id:            con.Contact.Id.String(),
				Contact:       con.Contact.Contact,
				ContactTypeId: con.Contact.ContactTypeId.String(),
				UpdatedAt:     updateCon,
			},
		}
		data.Contacts = append(data.Contacts, contact)
	}
	for _, add := range result.Locations {
		updateLoc, _ := timestamp.TimestampProto(add.UpdateAt)
		address := &pro.OperatorLocationProto{
			Id:         add.Id.String(),
			AddressId:  add.AddressId.String(),
			OperatorId: add.OperatorId.String(),
			Primary:    add.Primary,
			UpdateAt:   updateLoc,
		}
		updateAdd, _ := timestamp.TimestampProto(add.Address.UpdatedAt)
		addr := &pro.AddressProto{}
		addr.Address = add.Address.Address
		addr.Id = add.Address.Id.String()
		addr.CountryId = add.Address.CountryId.String()
		addr.AddressTypeId = add.Address.AddressTypeId.String()
		addr.StateId = add.Address.StateId.String()
		addr.UpdatedAt = updateAdd
		addr.Location = add.Address.Location
		addr.Suburb = add.Address.Suburb
		addr.Street = add.Address.Street
		address.Address = addr
		data.Locations = append(data.Locations, address)
	}

	for _, vh := range result.Vehicles {
		updateVh, _ := timestamp.TimestampProto(vh.UpdatedAt)
		vehicle := &pro.VehicleProto{
			Id:           vh.Id.String(),
			UpdatedAt:    updateVh,
			MakeId:       vh.MakeId.String(),
			ModelId:      vh.ModelId.String(),
			StatusId:     vh.StatusId.String(),
			Registration: vh.Registration,
			FleetId:      vh.FleetId.String(),
		}
		data.Vehicles = append(data.Vehicles, vehicle)
	}
	out.Operator = append(out.Operator, data)
	return nil

}

func (m *MasterService) GetOperatorsByVehicleId(ctx context.Context, req *pro.RequestKey, out *pro.ResponseOperator) error {

	defer RecoverError()

	if err := val.ValidateRequest(req); err != nil {
		out.Errors = ErrorResponse.GetCreateErrorJson(err)
		return nil
	}

	oprManager := opr.New()
	result, err := oprManager.GetOperatorsByVehicleId(uuid.FromStringOrNil(req.Id))
	out.Errors = ErrorResponse.GetCreateErrorJson(err)

	for _, opr := range result {

		data := &pro.OperatorProto{}
		data.Id = opr.Id.String()
		data.Active = opr.Active
		data.DrivingLic = opr.DrivingLic
		data.SurName = opr.SurName
		data.Name = opr.Name

		for _, con := range opr.Contacts {
			updateCon, _ := timestamp.TimestampProto(con.Contact.UpdatedAt)
			contact := &pro.OperatorContactsProto{
				Id:        con.Id.String(),
				Primary:   con.Primary,
				ContactId: con.ContactId.String(),
				Contact: &pro.ContactProto{
					Id:            con.Contact.Id.String(),
					Contact:       con.Contact.Contact,
					ContactTypeId: con.Contact.ContactTypeId.String(),
					UpdatedAt:     updateCon,
				},
			}
			data.Contacts = append(data.Contacts, contact)
		}
		for _, add := range opr.Locations {
			updateLoc, _ := timestamp.TimestampProto(add.UpdateAt)
			address := &pro.OperatorLocationProto{
				Id:         add.Id.String(),
				AddressId:  add.AddressId.String(),
				OperatorId: add.OperatorId.String(),
				Primary:    add.Primary,
				UpdateAt:   updateLoc,
			}
			updateAdd, _ := timestamp.TimestampProto(add.Address.UpdatedAt)
			addr := &pro.AddressProto{}
			addr.Address = add.Address.Address
			addr.Id = add.Address.Id.String()
			addr.CountryId = add.Address.CountryId.String()
			addr.AddressTypeId = add.Address.AddressTypeId.String()
			addr.StateId = add.Address.StateId.String()
			addr.UpdatedAt = updateAdd
			addr.Location = add.Address.Location
			addr.Suburb = add.Address.Suburb
			addr.Street = add.Address.Street
			address.Address = addr
			data.Locations = append(data.Locations, address)
		}

		for _, vh := range opr.Vehicles {
			updateVh, _ := timestamp.TimestampProto(vh.UpdatedAt)
			vehicle := &pro.VehicleProto{
				Id:           vh.Id.String(),
				UpdatedAt:    updateVh,
				MakeId:       vh.MakeId.String(),
				ModelId:      vh.ModelId.String(),
				StatusId:     vh.StatusId.String(),
				Registration: vh.Registration,
				FleetId:      vh.FleetId.String(),
			}
			data.Vehicles = append(data.Vehicles, vehicle)
		}
		out.Operator = append(out.Operator, data)
	}
	return nil
}

func (m *MasterService) CreateOperatorContact(ctx context.Context, req *pro.RequestOperatorContact, out *pro.ResponseCreateSuccess) error {

	defer RecoverError()

	if err := val.ValidateOperatorContract(val.CreateOperatorContract, req); err != nil {
		out.Errors = ErrorResponse.GetCreateErrorJson(err)
		return nil
	}

	oprManager := opr.New()
	result, err := oprManager.CreateOperatorContact(uuid.FromStringOrNil(req.ContactId), uuid.FromStringOrNil(req.OperatorId), req.Primary)
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Id = result.String()
	return nil
}

func (m *MasterService) UpdateOperatorContact(ctx context.Context, in *pro.RequestOperatorContact, out *pro.ResponseSuccess) error {

	defer RecoverError()

	if err := val.ValidateOperatorContract(val.UpdateOperatorContract, in); err != nil {
		out.Errors = ErrorResponse.GetCreateErrorJson(err)
		return nil
	}

	oprManager := opr.New()
	result, err := oprManager.UpdateOperatorContact(uuid.FromStringOrNil(in.Id), uuid.FromStringOrNil(in.ContactId), uuid.FromStringOrNil(in.OperatorId), in.Primary)
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Success = result
	return nil
}

func (m *MasterService) DeleteOperatorContact(ctx context.Context, in *pro.RequestDelete, out *pro.ResponseSuccess) error {

	defer RecoverError()

	if err := val.ValidateDelete(in); err != nil {
		out.Errors = ErrorResponse.GetCreateErrorJson(err)
		return nil
	}

	oprManager := opr.New()
	result, err := oprManager.DeleteOperatorContact(uuid.FromStringOrNil(in.Id))
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Success = result
	return nil
}

func (m *MasterService) GetAllContactsByOperator(ctx context.Context, in *pro.RequestKey, out *pro.ResponseOperatorContacts) error {

	defer RecoverError()

	if err := val.ValidateRequest(in); err != nil {
		out.Errors = ErrorResponse.GetCreateErrorJson(err)
		return nil
	}

	oprManager := opr.New()
	result, err := oprManager.GetAllContactsByOperator(uuid.FromStringOrNil(in.Id))
	out.Errors = ErrorResponse.GetCreateErrorJson(err)

	for _, op := range result {
		oprCon := pro.OperatorContactsProto{}
		oprCon.Id = op.Id.String()
		oprCon.Primary = op.Primary
		oprCon.OperatorId = op.OperatorId.String()
		oprCon.ContactId = op.ContactId.String()
		conUpdate, _ := timestamp.TimestampProto(op.Contact.UpdatedAt)
		oprCon.Contact = &pro.ContactProto{
			Contact:       op.Contact.Contact,
			Id:            op.Contact.Id.String(),
			ContactTypeId: op.Contact.ContactTypeId.String(),
			UpdatedAt:     conUpdate,
		}
		out.OperatorContacts = append(out.OperatorContacts, &oprCon)
	}
	return nil
}

func (m *MasterService) CreateOperatorLocation(ctx context.Context, in *pro.RequestOperatorLocation, out *pro.ResponseCreateSuccess) error {

	defer RecoverError()

	if err := val.ValidateOperatorLocation(val.CreateOperatorLocation, in.OperatorLocation); err != nil {
		out.Errors = ErrorResponse.GetCreateErrorJson(err)
		return nil
	}

	oprManager := opr.New()
	result, err := oprManager.CreateOperatorLocation(bu.OperatorLocationBO{
		OperatorId: uuid.FromStringOrNil(in.OperatorLocation.OperatorId),
		Primary:    in.OperatorLocation.Primary,
		AddressId:  uuid.FromStringOrNil(in.OperatorLocation.AddressId),
	})
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Id = result.String()
	return nil
}

func (m *MasterService) UpdateOperatorLocation(ctx context.Context, in *pro.RequestOperatorLocation, out *pro.ResponseSuccess) error {

	defer RecoverError()

	if err := val.ValidateOperatorLocation(val.UpdateOperatorLocation, in.OperatorLocation); err != nil {
		out.Errors = ErrorResponse.GetCreateErrorJson(err)
		return nil
	}

	oprManager := opr.New()
	result, err := oprManager.UpdateOperatorLocation(bu.OperatorLocationBO{
		Id:         uuid.FromStringOrNil(in.OperatorLocation.Id),
		OperatorId: uuid.FromStringOrNil(in.OperatorLocation.OperatorId),
		Primary:    in.OperatorLocation.Primary,
		AddressId:  uuid.FromStringOrNil(in.OperatorLocation.AddressId),
	})
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Success = result
	return nil
}
func (m *MasterService) DeleteOperatorLocation(ctx context.Context, in *pro.RequestDelete, out *pro.ResponseSuccess) error {

	defer RecoverError()

	if err := val.ValidateDelete(in); err != nil {
		out.Errors = ErrorResponse.GetCreateErrorJson(err)
		return nil
	}

	oprManager := opr.New()
	result, err := oprManager.DeleteOperatorLocation(uuid.FromStringOrNil(in.Id))
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Success = result
	return nil
}

func (m *MasterService) GetOperatorLocationByOperator(ctx context.Context, in *pro.RequestKey, out *pro.ResponseOperatorLocation) error {

	defer RecoverError()

	if err := val.ValidateRequest(in); err != nil {
		out.Errors = ErrorResponse.GetCreateErrorJson(err)
		return nil
	}

	oprManager := opr.New()
	result, err := oprManager.GetOperatorLocationByOperator(uuid.FromStringOrNil(in.Id))
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	for _, item := range result {
		updateOpr, _ := timestamp.TimestampProto(item.UpdateAt)
		opLoc := &pro.OperatorLocationProto{
			Id:         item.Id.String(),
			AddressId:  item.AddressId.String(),
			OperatorId: item.OperatorId.String(),
			Primary:    item.Primary,
			UpdateAt:   updateOpr,
		}
		updateAdd, _ := timestamp.TimestampProto(item.Address.UpdatedAt)
		opLoc.Address = &pro.AddressProto{
			Id:            item.Address.Id.String(),
			Address:       item.Address.Address,
			Street:        item.Address.Street,
			Suburb:        item.Address.Suburb,
			StateId:       item.Address.StateId.String(),
			CountryId:     item.Address.CountryId.String(),
			AddressTypeId: item.Address.AddressTypeId.String(),
			Location:      item.Address.Location,
			UpdatedAt:     updateAdd,
		}
		out.OperatorLocation = append(out.OperatorLocation, opLoc)
	}
	return nil
}
