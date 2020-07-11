package protoapi

import (
	"context"
	timestamp "github.com/golang/protobuf/ptypes"
	bu "tionyxtrack/masterservice/businesscontracts"
	opr "tionyxtrack/masterservice/manager/operator"
	pro "tionyxtrack/masterservice/proto"
)


func (m *MasterService) CreateOperator(ctx context.Context, req *pro.RequestOperator, out *pro.ResponseCreateSuccess) error {
	oprManager := opr.New()
	result, err := oprManager.CreateOperator(bu.OperatorBO{
		Name:       req.Operator.Name,
		SurName:    req.Operator.SurName,
		DrivingLic: req.Operator.DrivingLic,
	})
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Id = uint64(result)
	return nil
}

func (m *MasterService) UpdateOperator(ctx context.Context, req *pro.RequestOperator, out *pro.ResponseSuccess) error {
	oprManager := opr.New()
	result, err := oprManager.UpdateOperator(bu.OperatorBO{
		Id:         uint(req.Operator.Id),
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
	oprManager := opr.New()
	result, err := oprManager.DeleteOperator(uint(req.Id))
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Success = result
	return nil
}

func (m *MasterService) GetOperatorById(ctx context.Context, req *pro.RequestKey, out *pro.ResponseOperator) error {
	oprManager := opr.New()
	result, err := oprManager.GetOperatorById(uint(req.Id))
	out.Errors = ErrorResponse.GetCreateErrorJson(err)

	data := &pro.OperatorProto{}
	data.Id = uint64(result.Id)
	data.Active = result.Active
	data.DrivingLic = result.DrivingLic
	data.SurName = result.SurName
	data.Name = result.Name

	for _, con := range result.Contacts {
		updateCon, _ := timestamp.TimestampProto(con.Contact.UpdatedAt)
		contact := &pro.OperatorContactsProto{
			Id:        uint64(con.Id),
			Primary:   con.Primary,
			ContactId: uint64(con.ContactId),
			Contact: &pro.ContactProto{
				Id:            uint64(con.Contact.Id),
				Contact:       con.Contact.Contact,
				ContactTypeId: uint64(con.Contact.ContactTypeId),
				UpdatedAt:     updateCon,
			},
		}
		data.Contacts = append(data.Contacts, contact)
	}
	for _, add := range result.Locations {
		updateLoc, _ := timestamp.TimestampProto(add.UpdateAt)
		address := &pro.OperatorLocationProto{
			Id:         uint64(add.Id),
			AddressId:  uint64(add.AddressId),
			OperatorId: uint64(add.OperatorId),
			Primary:    add.Primary,
			UpdateAt:   updateLoc,
		}
		updateAdd, _ := timestamp.TimestampProto(add.Address.UpdatedAt)
		addr := &pro.AddressProto{}
		addr.Address = add.Address.Address
		addr.Id = uint64(add.Address.Id)
		addr.CountryId = uint64(add.Address.CountryId)
		addr.AddressTypeId = uint64(add.Address.AddressTypeId)
		addr.StateId = uint64(add.Address.StateId)
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
			Id:           uint64(vh.Id),
			UpdatedAt:    updateVh,
			MakeId:       uint64(vh.MakeId),
			ModelId:      uint64(vh.ModelId),
			StatusId:     uint64(vh.StatusId),
			Registration: vh.Registration,
			FleetId:      uint64(vh.FleetId),
		}
		data.Vehicles = append(data.Vehicles, vehicle)
	}
	out.Operator = append(out.Operator, data)
	return nil

}

func (m *MasterService) GetOperatorsByVehicleId(ctx context.Context, req *pro.RequestKey, out *pro.ResponseOperator) error {
	oprManager := opr.New()
	result, err := oprManager.GetOperatorsByVehicleId(uint(req.Id))
	out.Errors = ErrorResponse.GetCreateErrorJson(err)

	for _, opr := range result {

		data := &pro.OperatorProto{}
		data.Id = uint64(opr.Id)
		data.Active = opr.Active
		data.DrivingLic = opr.DrivingLic
		data.SurName = opr.SurName
		data.Name = opr.Name

		for _, con := range opr.Contacts {
			updateCon, _ := timestamp.TimestampProto(con.Contact.UpdatedAt)
			contact := &pro.OperatorContactsProto{
				Id:        uint64(con.Id),
				Primary:   con.Primary,
				ContactId: uint64(con.ContactId),
				Contact: &pro.ContactProto{
					Id:            uint64(con.Contact.Id),
					Contact:       con.Contact.Contact,
					ContactTypeId: uint64(con.Contact.ContactTypeId),
					UpdatedAt:     updateCon,
				},
			}
			data.Contacts = append(data.Contacts, contact)
		}
		for _, add := range opr.Locations {
			updateLoc, _ := timestamp.TimestampProto(add.UpdateAt)
			address := &pro.OperatorLocationProto{
				Id:         uint64(add.Id),
				AddressId:  uint64(add.AddressId),
				OperatorId: uint64(add.OperatorId),
				Primary:    add.Primary,
				UpdateAt:   updateLoc,
			}
			updateAdd, _ := timestamp.TimestampProto(add.Address.UpdatedAt)
			addr := &pro.AddressProto{}
			addr.Address = add.Address.Address
			addr.Id = uint64(add.Address.Id)
			addr.CountryId = uint64(add.Address.CountryId)
			addr.AddressTypeId = uint64(add.Address.AddressTypeId)
			addr.StateId = uint64(add.Address.StateId)
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
				Id:           uint64(vh.Id),
				UpdatedAt:    updateVh,
				MakeId:       uint64(vh.MakeId),
				ModelId:      uint64(vh.ModelId),
				StatusId:     uint64(vh.StatusId),
				Registration: vh.Registration,
				FleetId:      uint64(vh.FleetId),
			}
			data.Vehicles = append(data.Vehicles, vehicle)
		}
		out.Operator = append(out.Operator, data)
	}
	return nil
}

func (m *MasterService) CreateOperatorContact(ctx context.Context, req *pro.RequestOperatorContact, out *pro.ResponseCreateSuccess) error {
	oprManager := opr.New()
	result, err := oprManager.CreateOperatorContact(uint(req.ContactId), uint(req.OperatorId), req.Primary)
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Id = uint64(result)
	return nil
}

func (m *MasterService) UpdateOperatorContact(ctx context.Context, in *pro.RequestOperatorContact, out *pro.ResponseSuccess) error {
	oprManager := opr.New()
	result, err := oprManager.UpdateOperatorContact(uint(in.Id), uint(in.ContactId), uint(in.OperatorId), in.Primary)
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Success = result
	return nil
}

func (m *MasterService) DeleteOperatorContact(ctx context.Context, in *pro.RequestDelete, out *pro.ResponseSuccess) error {
	oprManager := opr.New()
	result, err := oprManager.DeleteOperatorContact(uint(in.Id))
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Success = result
	return nil
}

func (m *MasterService) GetAllContactsByOperator(ctx context.Context, in *pro.RequestKey, out *pro.ResponseOperatorContacts) error {
	oprManager := opr.New()
	result, err := oprManager.GetAllContactsByOperator(uint(in.Id))
	out.Errors = ErrorResponse.GetCreateErrorJson(err)

	for _, op := range result {
		oprCon := pro.OperatorContactsProto{}
		oprCon.Id = uint64(op.Id)
		oprCon.Primary = op.Primary
		oprCon.OperatorId = uint64(op.OperatorId)
		oprCon.ContactId = uint64(op.ContactId)
		conUpdate, _ := timestamp.TimestampProto(op.Contact.UpdatedAt)
		oprCon.Contact = &pro.ContactProto{
			Contact:       op.Contact.Contact,
			Id:            uint64(op.Contact.Id),
			ContactTypeId: uint64(op.Contact.ContactTypeId),
			UpdatedAt:     conUpdate,
		}
		out.OperatorContacts = append(out.OperatorContacts, &oprCon)
	}
	return nil
}

func (m *MasterService) CreateOperatorLocation(ctx context.Context, in *pro.RequestOperatorLocation, out *pro.ResponseCreateSuccess) error {
	oprManager := opr.New()
	result, err := oprManager.CreateOperatorLocation(bu.OperatorLocationBO{
		OperatorId: uint(in.OperatorLocation.OperatorId),
		Primary:    in.OperatorLocation.Primary,
		AddressId:  uint(in.OperatorLocation.AddressId),
	})
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Id = uint64(result)
	return nil
}

func (m *MasterService) UpdateOperatorLocation(ctx context.Context, in *pro.RequestOperatorLocation, out *pro.ResponseSuccess) error {
	oprManager := opr.New()
	result, err := oprManager.UpdateOperatorLocation(bu.OperatorLocationBO{
		Id:         uint(in.OperatorLocation.Id),
		OperatorId: uint(in.OperatorLocation.OperatorId),
		Primary:    in.OperatorLocation.Primary,
		AddressId:  uint(in.OperatorLocation.AddressId),
	})
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Success = result
	return nil
}
func (m *MasterService) DeleteOperatorLocation(ctx context.Context, in *pro.RequestDelete, out *pro.ResponseSuccess) error {
	oprManager := opr.New()
	result, err := oprManager.DeleteOperatorLocation(uint(in.Id))
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Success = result
	return nil
}

func (m *MasterService) GetOperatorLocationByOperator(ctx context.Context, in *pro.RequestKey, out *pro.ResponseOperatorLocation) error {
	oprManager := opr.New()
	result, err := oprManager.GetOperatorLocationByOperator(uint(in.Id))
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	for _, item := range result {
		updateOpr, _ := timestamp.TimestampProto(item.UpdateAt)
		opLoc := &pro.OperatorLocationProto{
			Id:         uint64(item.Id),
			AddressId:  uint64(item.AddressId),
			OperatorId: uint64(item.OperatorId),
			Primary:    item.Primary,
			UpdateAt:   updateOpr,
		}
		updateAdd, _ := timestamp.TimestampProto(item.Address.UpdatedAt)
		opLoc.Address = &pro.AddressProto{
			Id:            uint64(item.Address.Id),
			Address:       item.Address.Address,
			Street:        item.Address.Street,
			Suburb:        item.Address.Suburb,
			StateId:       uint64(item.Address.StateId),
			CountryId:     uint64(item.Address.CountryId),
			AddressTypeId: uint64(item.Address.AddressTypeId),
			Location:      item.Address.Location,
			UpdatedAt:     updateAdd,
		}
		out.OperatorLocation = append(out.OperatorLocation, opLoc)
	}
	return nil
}
