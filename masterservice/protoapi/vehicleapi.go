package protoapi

import (
	"context"
	timestamp "github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/empty"
	uuid "github.com/satori/go.uuid"
	bu "tionyxtrack/masterservice/businesscontracts"
	val "tionyxtrack/masterservice/common/validation"
	vh "tionyxtrack/masterservice/manager/vehicles"
	pro "tionyxtrack/masterservice/proto"
)

func (m *MasterService) CreateVehicle(ctx context.Context, in *pro.RequestVehicle, res *pro.ResponseCreateSuccess) error {

	defer RecoverError()

	if err := val.ValidateVehicle(val.CreateVehicle, in.Vehicle); err != nil {
		res.Errors = ErrorResponse.GetCreateErrorJson(err)
		return nil
	}

	vehManager := vh.New()
	result, err := vehManager.CreateVehicle(bu.VehicleBO{
		ModelId:      uuid.FromStringOrNil(in.Vehicle.ModelId),
		MakeId:       uuid.FromStringOrNil(in.Vehicle.MakeId),
		Registration: in.Vehicle.Registration,
		FleetId:      uuid.FromStringOrNil(in.Vehicle.FleetId),
		StatusId:     uuid.FromStringOrNil(in.Vehicle.StatusId),
		OfficeName:   in.Vehicle.OfficeName,
	})
	res.Errors = ErrorResponse.GetCreateErrorJson(err)
	res.Id = result.String()
	return nil
}

func (m *MasterService) UpdateVehicle(ctx context.Context, in *pro.RequestVehicle, res *pro.ResponseSuccess) error {

	defer RecoverError()

	if err := val.ValidateVehicle(val.UpdateVehicle, in.Vehicle); err != nil {
		res.Errors = ErrorResponse.GetCreateErrorJson(err)
		return nil
	}

	vehManager := vh.New()
	result, err := vehManager.UpdateVehicle(bu.VehicleBO{
		Id:           uuid.FromStringOrNil(in.Vehicle.Id),
		ModelId:      uuid.FromStringOrNil(in.Vehicle.ModelId),
		MakeId:       uuid.FromStringOrNil(in.Vehicle.MakeId),
		Registration: in.Vehicle.Registration,
		FleetId:      uuid.FromStringOrNil(in.Vehicle.FleetId),
		StatusId:     uuid.FromStringOrNil(in.Vehicle.StatusId),
		OfficeName:   in.Vehicle.OfficeName,
	})
	res.Success = result
	res.Errors = ErrorResponse.GetCreateErrorJson(err)
	return nil
}

func (m *MasterService) DeleteVehicle(ctx context.Context, in *pro.RequestDelete, res *pro.ResponseSuccess) error {

	defer RecoverError()

	if err := val.ValidateDelete(in); err != nil {
		res.Errors = ErrorResponse.GetCreateErrorJson(err)
		return nil
	}

	vehManager := vh.New()
	result, err := vehManager.DeleteVehicle(uuid.FromStringOrNil(in.Id))
	res.Errors = ErrorResponse.GetCreateErrorJson(err)
	res.Success = result
	return nil
}

func composeVehicleRequest(vehicles []bu.VehicleBO) ([]*pro.VehicleProto, error) {

	var vehicleList = make([]*pro.VehicleProto, len(vehicles))
	for _, result := range vehicles {

		updateVehicleAt, _ := timestamp.TimestampProto(result.UpdatedAt)

		data := &pro.VehicleProto{
			Id:           result.Id.String(),
			OfficeName:   result.OfficeName,
			StatusId:     result.StatusId.String(),
			FleetId:      result.FleetId.String(),
			Registration: result.Registration,
			MakeId:       result.MakeId.String(),
			ModelId:      result.ModelId.String(),
			UpdatedAt:    updateVehicleAt,
			Status: &pro.VehicleStatusProto{
				Id:         result.Status.Id.String(),
				StatusType: result.Status.StatusType,
				StatusName: result.Status.StatusName,
			},
		}

		for _, loc := range result.Locations {

			updateLocAt, _ := timestamp.TimestampProto(loc.UpdateAt)
			location := &pro.VehicleAddressProto{
				AddressId: loc.AddressId.String(),
				VehicleId: loc.VehicleId.String(),
				UpdateAt:  updateLocAt,
				Address: &pro.AddressProto{
					Id:            loc.Address.Id.String(),
					Address:       loc.Address.Address,
					Street:        loc.Address.Street,
					Suburb:        loc.Address.Suburb,
					StateId:       loc.Address.StateId.String(),
					CountryId:     loc.Address.CountryId.String(),
					AddressTypeId: loc.Address.AddressTypeId.String(),
					Location:      loc.Address.Location,
					State: &pro.StateProto{
						Id:        loc.Address.State.Id.String(),
						Name:      loc.Address.State.Name,
						CountryId: loc.Address.State.CountryId.String(),
					},
				},
			}
			data.Locations = append(data.Locations, location)
		}
		for _, op := range result.Operators {
			opr := &pro.VehicleOperatorBoundProto{
				Id:         op.Id.String(),
				OperatorId: op.OperatorId.String(),
				VehicleId:  op.VehicleId.String(),
				Active:     op.Active,
				Operator: &pro.OperatorProto{
					Id:         op.Operator.Id.String(),
					Name:       op.Operator.Name,
					SurName:    op.Operator.SurName,
					Active:     op.Operator.Active,
					DrivingLic: op.Operator.DrivingLic,
				},
				Vehicle: nil,
			}
			data.Operators = append(data.Operators, opr)
		}

		for _, trc := range result.Registrations {
			expiryDate, _ := timestamp.TimestampProto(trc.ExpiredDate)
			updateDate, _ := timestamp.TimestampProto(trc.UpdatedAt)
			registrationDate, _ := timestamp.TimestampProto(trc.RegisterDate)
			track := &pro.VehicleTrackRegProto{
				Id:           trc.Id.String(),
				RegisterDate: registrationDate,
				Duration:     int32(trc.Duration),
				ExpiredDate:  expiryDate,
				Active:       trc.Active,
				VehicleId:    trc.VehicleId.String(),
				UpdatedAt:    updateDate,
			}
			data.Registrations = append(data.Registrations, track)
		}
		vehicleList = append(vehicleList, data)
	}
	return vehicleList, nil
}

func (m *MasterService) GetVehicleById(ctx context.Context, in *pro.RequestKey, res *pro.ResponseVehicle) error {

	defer RecoverError()

	if err := val.ValidateRequest(in); err != nil {
		res.Errors = ErrorResponse.GetCreateErrorJson(err)
		return nil
	}

	vehManager := vh.New()
	result, err := vehManager.GetVehicleById(uuid.FromStringOrNil(in.Id))
	res.Errors = ErrorResponse.GetCreateErrorJson(err)
	res.Vehicles, _ = composeVehicleRequest([]bu.VehicleBO{result})

	return nil

}

func (m *MasterService) GetVehicleByRegistration(ctx context.Context, in *pro.RequestByName, res *pro.ResponseVehicle) error {

	defer RecoverError()

	vehManager := vh.New()
	result, err := vehManager.GetVehicleByRegistration(in.Name)
	res.Errors = ErrorResponse.GetCreateErrorJson(err)

	res.Vehicles, _ = composeVehicleRequest([]bu.VehicleBO{result})

	return nil
}
func (m *MasterService) GetVehiclesByFleetId(ctx context.Context, in *pro.RequestKey, res *pro.ResponseVehicle) error {

	defer RecoverError()

	vehManager := vh.New()
	result, err := vehManager.GetVehiclesByFleetId(uuid.FromStringOrNil(in.Id))
	res.Errors = ErrorResponse.GetCreateErrorJson(err)
	res.Vehicles, _ = composeVehicleRequest(result)
	return nil
}

func (m *MasterService) CreateVehicleHistory(ctx context.Context, in *pro.RequestVehicleHistory, res *pro.ResponseCreateSuccess) error {

	defer RecoverError()

	if err := val.ValidateVehicleHistory(val.CreateVehicleHistory, in.VehicleHistory); err != nil {
		res.Errors = ErrorResponse.GetCreateErrorJson(err)
		return nil
	}

	vehManager := vh.New()
	history := in.VehicleHistory
	changeDate, _ := timestamp.Timestamp(history.ChangeDate)
	result, err := vehManager.CreateVehicleHistory(bu.VehicleHistoryBO{
		VehicleId:    uuid.FromStringOrNil(history.VehicleId),
		ChangeDate:   changeDate,
		Description:  history.Description,
		FromStatusId: uuid.FromStringOrNil(history.FromStatusId),
		OfficerName:  history.OfficerName,
		ToStatusId:   uuid.FromStringOrNil(history.ToStatusId)})
	res.Errors = ErrorResponse.GetCreateErrorJson(err)
	res.Id = result.String()
	return nil
}

func (m *MasterService) UpdateVehicleHistory(ctx context.Context, in *pro.RequestVehicleHistory, res *pro.ResponseSuccess) error {

	defer RecoverError()

	if err := val.ValidateVehicleHistory(val.UpdateVehicleHistory, in.VehicleHistory); err != nil {
		res.Errors = ErrorResponse.GetCreateErrorJson(err)
		return nil
	}

	vehManager := vh.New()
	history := in.VehicleHistory
	changeDate, _ := timestamp.Timestamp(history.ChangeDate)
	result, err := vehManager.UpdateVehicleHistory(bu.VehicleHistoryBO{
		Id:           uuid.FromStringOrNil(history.Id),
		ToStatusId:   uuid.FromStringOrNil(history.ToStatusId),
		OfficerName:  history.OfficerName,
		FromStatusId: uuid.FromStringOrNil(history.FromStatusId),
		Description:  history.Description,
		ChangeDate:   changeDate,
		VehicleId:    uuid.FromStringOrNil(history.VehicleId),
	})
	res.Errors = ErrorResponse.GetCreateErrorJson(err)
	res.Success = result
	return nil
}

func (m *MasterService) DeleteVehicleHistory(ctx context.Context, in *pro.RequestDelete, out *pro.ResponseSuccess) error {

	defer RecoverError()

	if err := val.ValidateDelete(in); err != nil {
		out.Errors = ErrorResponse.GetCreateErrorJson(err)
		return nil
	}

	vehManager := vh.New()
	result, err := vehManager.DeleteVehicleHistory(uuid.FromStringOrNil(in.Id))
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Success = result
	return nil
}

func (m *MasterService) GetVehicleHistoryByVehicleId(ctx context.Context, in *pro.RequestKey, out *pro.ResponseVehicleHistory) error {

	defer RecoverError()

	if err := val.ValidateRequest(in); err != nil {
		out.Errors = ErrorResponse.GetCreateErrorJson(err)
		return nil
	}

	vehManager := vh.New()
	result, err := vehManager.GetVehicleHistoryByVehicleId(uuid.FromStringOrNil(in.Id))

	for _, item := range result {
		changeDate, _ := timestamp.TimestampProto(item.ChangeDate)
		out.VehicleHistory = append(out.VehicleHistory, &pro.VehicleHistoryProto{
			ChangeDate:   changeDate,
			Description:  item.Description,
			FromStatusId: item.FromStatusId.String(),
			OfficerName:  item.OfficerName,
			VehicleId:    item.VehicleId.String(),
			ToStatusId:   item.ToStatusId.String(),
			Id:           item.Id.String(),
			FromStatus: &pro.VehicleStatusProto{
				Id:         item.FromStatus.Id.String(),
				StatusType: item.FromStatus.StatusType,
				StatusName: item.FromStatus.StatusName,
			},
			ToStatus: &pro.VehicleStatusProto{
				Id:         item.ToStatus.Id.String(),
				StatusType: item.ToStatus.StatusType,
				StatusName: item.ToStatus.StatusName,
			},
		})
	}
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	return nil
}

func (m *MasterService) CreateVehicleLocation(ctx context.Context, in *pro.RequestVehicleAddress, out *pro.ResponseCreateSuccess) error {

	defer RecoverError()

	if err := val.ValidateVehicleLocation(val.CreateVehicleLocation, in.VehicleAddress); err != nil {
		out.Errors = ErrorResponse.GetCreateErrorJson(err)
		return nil
	}

	vehManager := vh.New()
	result, err := vehManager.CreateVehicleLocation(bu.VehicleAddressBO{
		VehicleId: uuid.FromStringOrNil(in.VehicleAddress.VehicleId),
		AddressId: uuid.FromStringOrNil(in.VehicleAddress.AddressId),
		Primary:   in.VehicleAddress.Primary,
	})
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Id = result.String()
	return nil
}

func (m *MasterService) UpdateVehicleLocation(ctx context.Context, in *pro.RequestVehicleAddress, out *pro.ResponseSuccess) error {

	defer RecoverError()

	if err := val.ValidateVehicleLocation(val.UpdateVehicleLocation, in.VehicleAddress); err != nil {
		out.Errors = ErrorResponse.GetCreateErrorJson(err)
		return nil
	}

	vehManager := vh.New()
	result, err := vehManager.UpdateVehicleLocation(bu.VehicleAddressBO{
		VehicleId: uuid.FromStringOrNil(in.VehicleAddress.VehicleId),
		AddressId: uuid.FromStringOrNil(in.VehicleAddress.AddressId),
		Primary:   in.VehicleAddress.Primary,
	})
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Success = result
	return nil
}

func (m *MasterService) DeleteVehicleLocation(ctx context.Context, in *pro.RequestDelete, out *pro.ResponseSuccess) error {

	defer RecoverError()

	if err := val.ValidateDelete(in); err != nil {
		out.Errors = ErrorResponse.GetCreateErrorJson(err)
		return nil
	}

	vehManager := vh.New()
	result, err := vehManager.DeleteVehicleLocation(uuid.FromStringOrNil(in.Id))
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Success = result
	return nil

}

func (m *MasterService) GetVehicleLocationByVehicle(ctx context.Context, in *pro.RequestKey, out *pro.ResponseVehicleAddress) error {

	defer RecoverError()

	if err := val.ValidateRequest(in); err != nil {
		out.Errors = ErrorResponse.GetCreateErrorJson(err)
		return nil
	}

	vehManager := vh.New()
	result, err := vehManager.GetVehicleLocationByVehicle(uuid.FromStringOrNil(in.Id))

	for _, item := range result {

		vehicleDate, _ := timestamp.TimestampProto(item.UpdateAt)
		addUpdateAt, _ := timestamp.TimestampProto(item.Address.UpdatedAt)
		vehicleAddress := &pro.VehicleAddressProto{
			VehicleId: item.VehicleId.String(),
			AddressId: item.AddressId.String(),
			Primary:   item.Primary,
			UpdateAt:  vehicleDate,
			Address: &pro.AddressProto{
				Id:            item.Address.Id.String(),
				Address:       item.Address.Address,
				Street:        item.Address.Street,
				Suburb:        item.Address.Suburb,
				StateId:       item.Address.StateId.String(),
				CountryId:     item.Address.CountryId.String(),
				AddressTypeId: item.Address.AddressTypeId.String(),
				Location:      item.Address.Location,
				UpdatedAt:     addUpdateAt,
			},
		}
		out.Address = append(out.Address, vehicleAddress)
	}
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	return nil
}

func (m *MasterService) CreateVehicleMake(ctx context.Context, in *pro.RequestVehicleMake, out *pro.ResponseCreateSuccess) error {

	defer RecoverError()

	if err := val.ValidateVehicleMake(val.CreateVehicleMake, in.VehicleMake); err != nil {
		out.Errors = ErrorResponse.GetCreateErrorJson(err)
		return nil
	}

	vehManager := vh.New()
	result, err := vehManager.CreateVehicleMake(bu.VehicleMakeBO{
		CountryId: uuid.FromStringOrNil(in.VehicleMake.CountryId),
		Make:      in.VehicleMake.Make,
	})

	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Id = result.String()
	return nil
}

func (m *MasterService) UpdateVehicleMake(ctx context.Context, in *pro.RequestVehicleMake, out *pro.ResponseSuccess) error {

	defer RecoverError()

	if err := val.ValidateVehicleMake(val.UpdateVehicleMake, in.VehicleMake); err != nil {
		out.Errors = ErrorResponse.GetCreateErrorJson(err)
		return nil
	}

	vehManager := vh.New()
	result, err := vehManager.UpdateVehicleMake(bu.VehicleMakeBO{
		Id:        uuid.FromStringOrNil(in.VehicleMake.Id),
		CountryId: uuid.FromStringOrNil(in.VehicleMake.CountryId),
		Make:      in.VehicleMake.Make,
	})
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Success = result
	return nil
}

func (m *MasterService) DeleteVehicleMake(ctx context.Context, in *pro.RequestDelete, out *pro.ResponseSuccess) error {

	defer RecoverError()

	if err := val.ValidateDelete(in); err != nil {
		out.Errors = ErrorResponse.GetCreateErrorJson(err)
		return nil
	}

	vehManager := vh.New()
	result, err := vehManager.DeleteVehicleMake(uuid.FromStringOrNil(in.Id))
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Success = result
	return nil

}

func (m *MasterService) GetAllVehicleMake(ctx context.Context, in *empty.Empty, out *pro.ResponseVehicleMake) error {

	defer RecoverError()

	vehManager := vh.New()
	result, err := vehManager.GetAllVehicleMake()
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	for _, item := range result {
		updateAt, _ := timestamp.TimestampProto(item.UpdateAt)
		out.VehicleMake = append(out.VehicleMake, &pro.VehicleMakeProto{
			Id:        item.Id.String(),
			Make:      item.Make,
			CountryId: item.CountryId.String(),
			UpdateAt:  updateAt,
			Country: &pro.CountryProto{
				Id:          item.Country.Id.String(),
				CountryName: item.Country.CountryName,
				RegionId:    item.Country.RegionId.String(),
			},
		})
	}
	return nil
}

func (m *MasterService) GetVehicleMakeById(ctx context.Context, in *pro.RequestKey, out *pro.ResponseVehicleMake) error {

	defer RecoverError()

	if err := val.ValidateRequest(in); err != nil {
		out.Errors = ErrorResponse.GetCreateErrorJson(err)
		return nil
	}

	vehManager := vh.New()
	result, err := vehManager.GetVehicleMakeById(uuid.FromStringOrNil(in.Id))
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	makeDate, _ := timestamp.TimestampProto(result.UpdateAt)
	out.VehicleMake = append(out.VehicleMake, &pro.VehicleMakeProto{
		Id:        result.Id.String(),
		Make:      result.Make,
		CountryId: result.CountryId.String(),
		UpdateAt:  makeDate,
		Country: &pro.CountryProto{
			Id:          result.Country.Id.String(),
			CountryName: result.Country.CountryName,
			RegionId:    result.Country.RegionId.String(),
		}})

	return nil
}

func (m *MasterService) CreateVehicleModel(ctx context.Context, in *pro.RequestVehicleModel, out *pro.ResponseCreateSuccess) error {

	defer RecoverError()

	if err := val.ValidateVehicleModel(val.CreateVehicleModel, in.VehicleModel); err != nil {
		out.Errors = ErrorResponse.GetCreateErrorJson(err)
		return nil
	}

	vehManager := vh.New()
	result, err := vehManager.CreateVehicleModel(bu.VehicleModelBO{
		Description: in.VehicleModel.Description,
		ModelName:   in.VehicleModel.ModelName,
		MakeId:      uuid.FromStringOrNil(in.VehicleModel.MakeId),
	})
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Id = result.String()
	return nil
}

func (m *MasterService) UpdateVehicleModel(ctx context.Context, in *pro.RequestVehicleModel, out *pro.ResponseSuccess) error {

	defer RecoverError()

	if err := val.ValidateVehicleModel(val.UpdateVehicleModel, in.VehicleModel); err != nil {
		out.Errors = ErrorResponse.GetCreateErrorJson(err)
		return nil
	}

	vehManager := vh.New()
	result, err := vehManager.UpdateVehicleModel(bu.VehicleModelBO{
		MakeId:      uuid.FromStringOrNil(in.VehicleModel.MakeId),
		ModelName:   in.VehicleModel.ModelName,
		Description: in.VehicleModel.Description,
	})
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Success = result
	return nil
}

func (m *MasterService) DeleteVehicleModel(ctx context.Context, in *pro.RequestDelete, out *pro.ResponseSuccess) error {

	defer RecoverError()

	if err := val.ValidateDelete(in); err != nil {
		out.Errors = ErrorResponse.GetCreateErrorJson(err)
		return nil
	}

	vehManager := vh.New()
	result, err := vehManager.DeleteVehicleModel(uuid.FromStringOrNil(in.Id))
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Success = result
	return nil
}
func (m *MasterService) GetAllModelByMake(ctx context.Context, in *pro.RequestKey, out *pro.ResponseVehicleModel) error {

	defer RecoverError()

	if err := val.ValidateRequest(in); err != nil {
		out.Errors = ErrorResponse.GetCreateErrorJson(err)
		return nil
	}

	vehManager := vh.New()
	result, err := vehManager.GetAllModelByMake(uuid.FromStringOrNil(in.Id))
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	for _, item := range result {
		modelDate, _ := timestamp.TimestampProto(item.UpdatedAt)
		makeDate, _ := timestamp.TimestampProto(item.Make.UpdateAt)
		countryDate, _ := timestamp.TimestampProto(item.Make.Country.UpdatedAt)
		out.VehicleModel = append(out.VehicleModel, &pro.VehicleModelProto{
			Id:          item.Id.String(),
			ModelName:   item.ModelName,
			Description: item.Description,
			MakeId:      item.MakeId.String(),
			UpdatedAt:   modelDate,
			Make: &pro.VehicleMakeProto{
				Id:        item.Make.Id.String(),
				Make:      item.Make.Make,
				CountryId: item.Make.CountryId.String(),
				UpdateAt:  makeDate,
				Country: &pro.CountryProto{
					Id:          item.Make.Country.Id.String(),
					CountryName: item.Make.Country.CountryName,
					RegionId:    item.Make.Country.RegionId.String(),
					UpdatedAt:   countryDate,
				},
			},
		})
	}
	return nil
}

func (m *MasterService) GetModelById(ctx context.Context, in *pro.RequestKey, out *pro.ResponseVehicleModel) error {

	defer RecoverError()

	if err := val.ValidateRequest(in); err != nil {
		out.Errors = ErrorResponse.GetCreateErrorJson(err)
		return nil
	}

	vehManager := vh.New()
	result, err := vehManager.GetModelById(uuid.FromStringOrNil(in.Id))
	out.Errors = ErrorResponse.GetCreateErrorJson(err)

	modelDate, _ := timestamp.TimestampProto(result.UpdatedAt)
	makeDate, _ := timestamp.TimestampProto(result.Make.UpdateAt)
	countryDate, _ := timestamp.TimestampProto(result.Make.Country.UpdatedAt)

	out.VehicleModel = append(out.VehicleModel, &pro.VehicleModelProto{
		Id:          result.Id.String(),
		ModelName:   result.ModelName,
		Description: result.Description,
		MakeId:      result.MakeId.String(),
		UpdatedAt:   modelDate,
		Make: &pro.VehicleMakeProto{
			Id:        result.Make.Id.String(),
			Make:      result.Make.Make,
			CountryId: result.Make.CountryId.String(),
			UpdateAt:  makeDate,
			Country: &pro.CountryProto{
				Id:          result.Make.Country.Id.String(),
				CountryName: result.Make.Country.CountryName,
				RegionId:    result.Make.Country.RegionId.String(),
				UpdatedAt:   countryDate,
			},
		},
	})
	return nil
}

func (m *MasterService) CreateVehicleReg(ctx context.Context, in *pro.RequestVehicleReg, out *pro.ResponseCreateSuccess) error {

	defer RecoverError()

	if err := val.ValidateVehicleReg(val.CreateVehicleReg, in.VehicleReg); err != nil {
		out.Errors = ErrorResponse.GetCreateErrorJson(err)
		return nil
	}

	vehManager := vh.New()
	expireDate, _ := timestamp.Timestamp(in.VehicleReg.ExpiredDate)
	regDate, _ := timestamp.Timestamp(in.VehicleReg.RegisterDate)

	result, err := vehManager.CreateVehicleReg(bu.VehicleTrackRegBO{
		Active:       in.VehicleReg.Active,
		Duration:     int(in.VehicleReg.Duration),
		ExpiredDate:  expireDate,
		VehicleId:    uuid.FromStringOrNil(in.VehicleReg.VehicleId),
		RegisterDate: regDate,
	})
	out.Id = result.String()
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	return nil
}
func (m *MasterService) UpdateVehicleReg(ctx context.Context, in *pro.RequestVehicleReg, out *pro.ResponseSuccess) error {

	defer RecoverError()

	if err := val.ValidateVehicleReg(val.UpdateVehicleReg, in.VehicleReg); err != nil {
		out.Errors = ErrorResponse.GetCreateErrorJson(err)
		return nil
	}

	vehManager := vh.New()
	expireDate, _ := timestamp.Timestamp(in.VehicleReg.ExpiredDate)
	regDate, _ := timestamp.Timestamp(in.VehicleReg.RegisterDate)

	result, err := vehManager.UpdateVehicleReg(bu.VehicleTrackRegBO{
		Id:           uuid.FromStringOrNil(in.VehicleReg.Id),
		VehicleId:    uuid.FromStringOrNil(in.VehicleReg.VehicleId),
		ExpiredDate:  expireDate,
		Duration:     int(in.VehicleReg.Duration),
		Active:       in.VehicleReg.Active,
		RegisterDate: regDate,
	})
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Success = result

	return nil
}

func (m *MasterService) DeleteVehicleReg(ctx context.Context, in *pro.RequestDelete, out *pro.ResponseSuccess) error {

	defer RecoverError()

	if err := val.ValidateDelete(in); err != nil {
		out.Errors = ErrorResponse.GetCreateErrorJson(err)
		return nil
	}

	vehManager := vh.New()
	result, err := vehManager.DeleteVehicleReg(uuid.FromStringOrNil(in.Id))
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Success = result
	return nil
}

func (m *MasterService) GetAllRegistrationsByVehicleId(ctx context.Context, in *pro.RequestKey, out *pro.ResponseVehicleReg) error {

	defer RecoverError()

	if err := val.ValidateRequest(in); err != nil {
		out.Errors = ErrorResponse.GetCreateErrorJson(err)
		return nil
	}

	vehManager := vh.New()
	result, err := vehManager.GetAllRegistrationsByVehicleId(uuid.FromStringOrNil(in.Id))
	out.Errors = ErrorResponse.GetCreateErrorJson(err)

	for _, vh := range result {

		expireDate, _ := timestamp.TimestampProto(vh.ExpiredDate)
		regDate, _ := timestamp.TimestampProto(vh.RegisterDate)
		updateAt, _ := timestamp.TimestampProto(vh.UpdatedAt)
		vhUpdateAt, _ := timestamp.TimestampProto(vh.Vehicle.UpdatedAt)

		out.VehicleReg = append(out.VehicleReg, &pro.VehicleTrackRegProto{
			Id:           vh.Id.String(),
			RegisterDate: regDate,
			Duration:     int32(vh.Duration),
			ExpiredDate:  expireDate,
			Active:       vh.Active,
			VehicleId:    vh.VehicleId.String(),
			UpdatedAt:    updateAt,
			Vehicle: &pro.VehicleProto{
				Id:           vh.Vehicle.Id.String(),
				UpdatedAt:    vhUpdateAt,
				Registration: vh.Vehicle.Registration,
				StatusId:     vh.Vehicle.StatusId.String(),
				MakeId:       vh.Vehicle.MakeId.String(),
				ModelId:      vh.Vehicle.ModelId.String(),
			},
		})
	}
	return nil

}

func (m *MasterService) GetActiveRegistrationsByVehicleId(ctx context.Context, in *pro.RequestKey, out *pro.ResponseVehicleReg) error {

	defer RecoverError()

	if err := val.ValidateRequest(in); err != nil {
		out.Errors = ErrorResponse.GetCreateErrorJson(err)
		return nil
	}

	vehManager := vh.New()
	result, err := vehManager.GetActiveRegistrationsByVehicleId(uuid.FromStringOrNil(in.Id))
	out.Errors = ErrorResponse.GetCreateErrorJson(err)

	expireDate, _ := timestamp.TimestampProto(result.ExpiredDate)
	regDate, _ := timestamp.TimestampProto(result.RegisterDate)
	updateAt, _ := timestamp.TimestampProto(result.UpdatedAt)
	vhUpdateAt, _ := timestamp.TimestampProto(result.Vehicle.UpdatedAt)

	out.VehicleReg = append(out.VehicleReg, &pro.VehicleTrackRegProto{
		Id:           result.Id.String(),
		RegisterDate: regDate,
		Duration:     int32(result.Duration),
		ExpiredDate:  expireDate,
		Active:       result.Active,
		VehicleId:    result.VehicleId.String(),
		UpdatedAt:    updateAt,
		Vehicle: &pro.VehicleProto{
			Id:           result.Vehicle.Id.String(),
			UpdatedAt:    vhUpdateAt,
			Registration: result.Vehicle.Registration,
			StatusId:     result.Vehicle.StatusId.String(),
			MakeId:       result.Vehicle.MakeId.String(),
			ModelId:      result.Vehicle.ModelId.String(),
		},
	})
	return nil
}

func (m *MasterService) CreateVehicleStatus(ctx context.Context, in *pro.RequestVehicleStatus, out *pro.ResponseCreateSuccess) error {

	defer RecoverError()

	if err := val.ValidateVehicleStatus(val.CreateVehicleStatus, in.VehicleStatus); err != nil {
		out.Errors = ErrorResponse.GetCreateErrorJson(err)
		return nil
	}

	vehManager := vh.New()
	result, err := vehManager.CreateVehicleStatus(bu.VehicleStatusBO{
		Id:         uuid.FromStringOrNil(in.VehicleStatus.Id),
		StatusName: in.VehicleStatus.StatusName,
		StatusType: in.VehicleStatus.StatusType,
	})

	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Success = true
	out.Id = result.String()
	return nil
}

func (m *MasterService) UpdateVehicleStatus(ctx context.Context, in *pro.RequestVehicleStatus, out *pro.ResponseSuccess) error {

	defer RecoverError()

	if err := val.ValidateVehicleStatus(val.UpdateVehicleStatus, in.VehicleStatus); err != nil {
		out.Errors = ErrorResponse.GetCreateErrorJson(err)
		return nil
	}

	vehManager := vh.New()
	result, err := vehManager.UpdateVehicleStatus(bu.VehicleStatusBO{
		Id:         uuid.FromStringOrNil(in.VehicleStatus.Id),
		StatusType: in.VehicleStatus.StatusType,
		StatusName: in.VehicleStatus.StatusName,
	})
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Success = result
	return nil
}

func (m *MasterService) DeleteVehicleStatus(ctx context.Context, in *pro.RequestDelete, out *pro.ResponseSuccess) error {

	defer RecoverError()

	if err := val.ValidateDelete(in); err != nil {
		out.Errors = ErrorResponse.GetCreateErrorJson(err)
		return nil
	}

	vehManager := vh.New()
	result, err := vehManager.DeleteVehicleStatus(uuid.FromStringOrNil(in.Id))
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Success = result
	return nil
}

func (m *MasterService) GetAllVehicleStatus(ctx context.Context, in *empty.Empty, out *pro.ResponseVehicleStatus) error {

	defer RecoverError()

	vehManager := vh.New()
	result, err := vehManager.GetAllVehicleStatus()
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	for _, item := range result {

		updateAt, _ := timestamp.TimestampProto(item.UpdatedAt)
		out.VehicleStatus = append(out.VehicleStatus, &pro.VehicleStatusProto{
			Id:         item.Id.String(),
			StatusName: item.StatusName,
			StatusType: item.StatusType,
			UpdatedAt:  updateAt,
		})
	}
	return nil
}

func (m *MasterService) CreateVehicleOpBound(ctx context.Context, in *pro.RequestVehicleOprBound, out *pro.ResponseCreateSuccess) error {

	defer RecoverError()

	if err := val.ValidateVehicleOperatorBound(val.CreateVehicleOpBound, in.VehicleOprBound); err != nil {
		out.Errors = ErrorResponse.GetCreateErrorJson(err)
		return nil
	}

	vehManager := vh.New()
	result, err := vehManager.CreateVehicleOpBound(bu.VehicleOperatorBoundBO{
		Active:     in.VehicleOprBound.Active,
		VehicleId:  uuid.FromStringOrNil(in.VehicleOprBound.VehicleId),
		OperatorId: uuid.FromStringOrNil(in.VehicleOprBound.OperatorId),
	})
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Id = result.String()
	return nil
}

func (m *MasterService) UpdateVehicleOpBound(ctx context.Context, in *pro.RequestVehicleOprBound, out *pro.ResponseSuccess) error {

	defer RecoverError()

	if err := val.ValidateVehicleOperatorBound(val.UpdateVehicleOpBound, in.VehicleOprBound); err != nil {
		out.Errors = ErrorResponse.GetCreateErrorJson(err)
		return nil
	}

	vehManager := vh.New()
	result, err := vehManager.UpdateVehicleOpBound(bu.VehicleOperatorBoundBO{
		Id:         uuid.FromStringOrNil(in.VehicleOprBound.Id),
		Active:     in.VehicleOprBound.Active,
		VehicleId:  uuid.FromStringOrNil(in.VehicleOprBound.VehicleId),
		OperatorId: uuid.FromStringOrNil(in.VehicleOprBound.OperatorId),
	})
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Success = result
	return nil
}

func (m *MasterService) DeleteVehicleOpBound(ctx context.Context, in *pro.RequestDelete, out *pro.ResponseSuccess) error {

	defer RecoverError()

	if err := val.ValidateDelete(in); err != nil {
		out.Errors = ErrorResponse.GetCreateErrorJson(err)
		return nil
	}

	vehManager := vh.New()
	result, err := vehManager.DeleteVehicleOpBound(uuid.FromStringOrNil(in.Id))
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Success = result
	return nil
}
