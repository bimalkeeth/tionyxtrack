package protoapi

import (
	"context"
	timestamp "github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/empty"
	bu "tionyxtrack/masterservice/businesscontracts"
	vh "tionyxtrack/masterservice/manager/vehicles"
	pro "tionyxtrack/masterservice/proto"
)


func (m *MasterService) CreateVehicle(ctx context.Context, in *pro.RequestVehicle, res *pro.ResponseCreateSuccess) error {
	vehManager := vh.New()
	result, err := vehManager.CreateVehicle(bu.VehicleBO{
		ModelId:      uint(in.Vehicle.ModelId),
		MakeId:       uint(in.Vehicle.MakeId),
		Registration: in.Vehicle.Registration,
		FleetId:      uint(in.Vehicle.FleetId),
		StatusId:     uint(in.Vehicle.StatusId),
		OfficeName:   in.Vehicle.OfficeName,
	})
	res.Errors = ErrorResponse.GetCreateErrorJson(err)
	res.Id = uint64(result)
	return nil
}

func (m *MasterService) UpdateVehicle(ctx context.Context, in *pro.RequestVehicle, res *pro.ResponseSuccess) error {
	vehManager := vh.New()
	result, err := vehManager.UpdateVehicle(bu.VehicleBO{
		Id:           uint(in.Vehicle.Id),
		ModelId:      uint(in.Vehicle.ModelId),
		MakeId:       uint(in.Vehicle.MakeId),
		Registration: in.Vehicle.Registration,
		FleetId:      uint(in.Vehicle.FleetId),
		StatusId:     uint(in.Vehicle.StatusId),
		OfficeName:   in.Vehicle.OfficeName,
	})
	res.Success = result
	res.Errors = ErrorResponse.GetCreateErrorJson(err)
	return nil
}

func (m *MasterService) DeleteVehicle(ctx context.Context, in *pro.RequestDelete, res *pro.ResponseSuccess) error {
	vehManager := vh.New()
	result, err := vehManager.DeleteVehicle(uint(in.Id))
	res.Errors = ErrorResponse.GetCreateErrorJson(err)
	res.Success = result
	return nil
}

func (m *MasterService) GetVehicleById(ctx context.Context, in *pro.RequestKey, res *pro.ResponseVehicle) error {
	vehManager := vh.New()
	result, err := vehManager.GetVehicleById(uint(in.Id))
	res.Errors = ErrorResponse.GetCreateErrorJson(err)

	updateVehicleAt, _ := timestamp.TimestampProto(result.UpdatedAt)
	data := &pro.VehicleProto{
		Id:           uint64(result.Id),
		OfficeName:   result.OfficeName,
		StatusId:     uint64(result.StatusId),
		FleetId:      uint64(result.FleetId),
		Registration: result.Registration,
		MakeId:       uint64(result.MakeId),
		ModelId:      uint64(result.ModelId),
		UpdatedAt:    updateVehicleAt,
		Status: &pro.VehicleStatusProto{
			Id:         uint64(result.Status.Id),
			StatusType: result.Status.StatusType,
			StatusName: result.Status.StatusName,
		},
	}

	for _, loc := range result.Locations {

		updateLocAt, _ := timestamp.TimestampProto(loc.UpdateAt)
		location := &pro.VehicleAddressProto{
			AddressId: uint64(loc.AddressId),
			VehicleId: uint64(loc.VehicleId),
			UpdateAt:  updateLocAt,
			Address: &pro.AddressProto{
				Id:            uint64(loc.Address.Id),
				Address:       loc.Address.Address,
				Street:        loc.Address.Street,
				Suburb:        loc.Address.Suburb,
				StateId:       uint64(loc.Address.StateId),
				CountryId:     uint64(loc.Address.CountryId),
				AddressTypeId: uint64(loc.Address.AddressTypeId),
				Location:      loc.Address.Location,
				State: &pro.StateProto{
					Id:        uint64(loc.Address.State.Id),
					Name:      loc.Address.State.Name,
					CountryId: uint64(loc.Address.State.CountryId),
				},
			},
		}
		data.Locations = append(data.Locations, location)
	}
	for _, op := range result.Operators {
		opr := &pro.VehicleOperatorBoundProto{
			Id:         uint64(op.Id),
			OperatorId: uint64(op.OperatorId),
			VehicleId:  uint64(op.VehicleId),
			Active:     op.Active,
			Operator: &pro.OperatorProto{
				Id:         uint64(op.Operator.Id),
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
			Id:           uint64(trc.Id),
			RegisterDate: registrationDate,
			Duration:     int32(trc.Duration),
			ExpiredDate:  expiryDate,
			Active:       trc.Active,
			VehicleId:    uint64(trc.VehicleId),
			UpdatedAt:    updateDate,
		}
		data.Registrations = append(data.Registrations, track)
	}
	res.Vehicles = append(res.Vehicles, data)
	return nil

}

func (m *MasterService) GetVehicleByRegistration(ctx context.Context, in *pro.RequestByName, res *pro.ResponseVehicle) error {

	vehManager := vh.New()
	result, err := vehManager.GetVehicleByRegistration(in.Name)
	res.Errors = ErrorResponse.GetCreateErrorJson(err)

	updateVehicleAt, _ := timestamp.TimestampProto(result.UpdatedAt)
	data := &pro.VehicleProto{
		Id:           uint64(result.Id),
		OfficeName:   result.OfficeName,
		StatusId:     uint64(result.StatusId),
		FleetId:      uint64(result.FleetId),
		Registration: result.Registration,
		MakeId:       uint64(result.MakeId),
		ModelId:      uint64(result.ModelId),
		UpdatedAt:    updateVehicleAt,
		Status: &pro.VehicleStatusProto{
			Id:         uint64(result.Status.Id),
			StatusType: result.Status.StatusType,
			StatusName: result.Status.StatusName,
		},
	}

	for _, loc := range result.Locations {

		updateLocAt, _ := timestamp.TimestampProto(loc.UpdateAt)
		location := &pro.VehicleAddressProto{
			AddressId: uint64(loc.AddressId),
			VehicleId: uint64(loc.VehicleId),
			UpdateAt:  updateLocAt,
			Address: &pro.AddressProto{
				Id:            uint64(loc.Address.Id),
				Address:       loc.Address.Address,
				Street:        loc.Address.Street,
				Suburb:        loc.Address.Suburb,
				StateId:       uint64(loc.Address.StateId),
				CountryId:     uint64(loc.Address.CountryId),
				AddressTypeId: uint64(loc.Address.AddressTypeId),
				Location:      loc.Address.Location,
				State: &pro.StateProto{
					Id:        uint64(loc.Address.State.Id),
					Name:      loc.Address.State.Name,
					CountryId: uint64(loc.Address.State.CountryId),
				},
			},
		}
		data.Locations = append(data.Locations, location)
	}
	for _, op := range result.Operators {
		opr := &pro.VehicleOperatorBoundProto{
			Id:         uint64(op.Id),
			OperatorId: uint64(op.OperatorId),
			VehicleId:  uint64(op.VehicleId),
			Active:     op.Active,
			Operator: &pro.OperatorProto{
				Id:         uint64(op.Operator.Id),
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
			Id:           uint64(trc.Id),
			RegisterDate: registrationDate,
			Duration:     int32(trc.Duration),
			ExpiredDate:  expiryDate,
			Active:       trc.Active,
			VehicleId:    uint64(trc.VehicleId),
			UpdatedAt:    updateDate,
		}
		data.Registrations = append(data.Registrations, track)
	}
	res.Vehicles = append(res.Vehicles, data)

	return nil
}
func (m *MasterService) GetVehiclesByFleetId(ctx context.Context, in *pro.RequestKey, res *pro.ResponseVehicle) error {
	vehManager := vh.New()
	result, err := vehManager.GetVehiclesByFleetId(uint(in.Id))
	res.Errors = ErrorResponse.GetCreateErrorJson(err)

	for _, item := range result {
		updateVehicleAt, _ := timestamp.TimestampProto(item.UpdatedAt)
		data := &pro.VehicleProto{
			Id:           uint64(item.Id),
			OfficeName:   item.OfficeName,
			StatusId:     uint64(item.StatusId),
			FleetId:      uint64(item.FleetId),
			Registration: item.Registration,
			MakeId:       uint64(item.MakeId),
			ModelId:      uint64(item.ModelId),
			UpdatedAt:    updateVehicleAt,
			Status: &pro.VehicleStatusProto{
				Id:         uint64(item.Status.Id),
				StatusType: item.Status.StatusType,
				StatusName: item.Status.StatusName,
			},
		}

		for _, loc := range item.Locations {

			updateLocAt, _ := timestamp.TimestampProto(loc.UpdateAt)
			location := &pro.VehicleAddressProto{
				AddressId: uint64(loc.AddressId),
				VehicleId: uint64(loc.VehicleId),
				UpdateAt:  updateLocAt,
				Address: &pro.AddressProto{
					Id:            uint64(loc.Address.Id),
					Address:       loc.Address.Address,
					Street:        loc.Address.Street,
					Suburb:        loc.Address.Suburb,
					StateId:       uint64(loc.Address.StateId),
					CountryId:     uint64(loc.Address.CountryId),
					AddressTypeId: uint64(loc.Address.AddressTypeId),
					Location:      loc.Address.Location,
					State: &pro.StateProto{
						Id:        uint64(loc.Address.State.Id),
						Name:      loc.Address.State.Name,
						CountryId: uint64(loc.Address.State.CountryId),
					},
				},
			}
			data.Locations = append(data.Locations, location)
		}

		for _, op := range item.Operators {
			opr := &pro.VehicleOperatorBoundProto{
				Id:         uint64(op.Id),
				OperatorId: uint64(op.OperatorId),
				VehicleId:  uint64(op.VehicleId),
				Active:     op.Active,
				Operator: &pro.OperatorProto{
					Id:         uint64(op.Operator.Id),
					Name:       op.Operator.Name,
					SurName:    op.Operator.SurName,
					Active:     op.Operator.Active,
					DrivingLic: op.Operator.DrivingLic,
				},
				Vehicle: nil,
			}
			data.Operators = append(data.Operators, opr)
		}

		for _, trc := range item.Registrations {
			expiryDate, _ := timestamp.TimestampProto(trc.ExpiredDate)
			updateDate, _ := timestamp.TimestampProto(trc.UpdatedAt)
			registrationDate, _ := timestamp.TimestampProto(trc.RegisterDate)
			track := &pro.VehicleTrackRegProto{
				Id:           uint64(trc.Id),
				RegisterDate: registrationDate,
				Duration:     int32(trc.Duration),
				ExpiredDate:  expiryDate,
				Active:       trc.Active,
				VehicleId:    uint64(trc.VehicleId),
				UpdatedAt:    updateDate,
			}
			data.Registrations = append(data.Registrations, track)
		}
		res.Vehicles = append(res.Vehicles, data)
	}
	return nil
}

func (m *MasterService) CreateVehicleHistory(ctx context.Context, in *pro.RequestVehicleHistory, res *pro.ResponseCreateSuccess) error {
	vehManager := vh.New()
	history := in.VehicleHistory
	changeDate, _ := timestamp.Timestamp(history.ChangeDate)
	result, err := vehManager.CreateVehicleHistory(bu.VehicleHistoryBO{
		VehicleId:    uint(history.VehicleId),
		ChangeDate:   changeDate,
		Description:  history.Description,
		FromStatusId: uint(history.FromStatusId),
		OfficerName:  history.OfficerName,
		ToStatusId:   uint(history.ToStatusId)})
	res.Errors = ErrorResponse.GetCreateErrorJson(err)
	res.Id = uint64(result)
	return nil
}

func (m *MasterService) UpdateVehicleHistory(ctx context.Context, in *pro.RequestVehicleHistory, res *pro.ResponseSuccess) error {
	vehManager := vh.New()
	history := in.VehicleHistory
	changeDate, _ := timestamp.Timestamp(history.ChangeDate)
	result, err := vehManager.UpdateVehicleHistory(bu.VehicleHistoryBO{
		Id:           uint(history.Id),
		ToStatusId:   uint(history.ToStatusId),
		OfficerName:  history.OfficerName,
		FromStatusId: uint(history.FromStatusId),
		Description:  history.Description,
		ChangeDate:   changeDate,
		VehicleId:    uint(history.VehicleId),
	})
	res.Errors = ErrorResponse.GetCreateErrorJson(err)
	res.Success = result
	return nil
}

func (m *MasterService) DeleteVehicleHistory(ctx context.Context, in *pro.RequestDelete, out *pro.ResponseSuccess) error {
	vehManager := vh.New()
	result, err := vehManager.DeleteVehicleHistory(uint(in.Id))
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Success = result
	return nil
}

func (m *MasterService) GetVehicleHistoryByVehicleId(ctx context.Context, in *pro.RequestKey, out *pro.ResponseVehicleHistory) error {
	vehManager := vh.New()
	result, err := vehManager.GetVehicleHistoryByVehicleId(uint(in.Id))

	for _, item := range result {
		changeDate, _ := timestamp.TimestampProto(item.ChangeDate)
		out.VehicleHistory = append(out.VehicleHistory, &pro.VehicleHistoryProto{
			ChangeDate:   changeDate,
			Description:  item.Description,
			FromStatusId: uint64(item.FromStatusId),
			OfficerName:  item.OfficerName,
			VehicleId:    uint64(item.VehicleId),
			ToStatusId:   uint64(item.ToStatusId),
			Id:           uint64(item.Id),
			FromStatus: &pro.VehicleStatusProto{
				Id:         uint64(item.FromStatus.Id),
				StatusType: item.FromStatus.StatusType,
				StatusName: item.FromStatus.StatusName,
			},
			ToStatus: &pro.VehicleStatusProto{
				Id:         uint64(item.ToStatus.Id),
				StatusType: item.ToStatus.StatusType,
				StatusName: item.ToStatus.StatusName,
			},
		})
	}
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	return nil
}

func (m *MasterService) CreateVehicleLocation(ctx context.Context, in *pro.RequestVehicleAddress, out *pro.ResponseCreateSuccess) error {
	vehManager := vh.New()
	result, err := vehManager.CreateVehicleLocation(bu.VehicleAddressBO{
		VehicleId: uint(in.VehicleAddress.VehicleId),
		AddressId: uint(in.VehicleAddress.AddressId),
		Primary:   in.VehicleAddress.Primary,
	})
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Id = uint64(result)
	return nil
}

func (m *MasterService) UpdateVehicleLocation(ctx context.Context, in *pro.RequestVehicleAddress, out *pro.ResponseSuccess) error {
	vehManager := vh.New()
	result, err := vehManager.UpdateVehicleLocation(bu.VehicleAddressBO{
		VehicleId: uint(in.VehicleAddress.VehicleId),
		AddressId: uint(in.VehicleAddress.AddressId),
		Primary:   in.VehicleAddress.Primary,
	})
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Success = result
	return nil
}

func (m *MasterService) DeleteVehicleLocation(ctx context.Context, in *pro.RequestDelete, out *pro.ResponseSuccess) error {
	vehManager := vh.New()
	result, err := vehManager.DeleteVehicleLocation(uint(in.Id))
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Success = result
	return nil

}

func (m *MasterService) GetVehicleLocationByVehicle(ctx context.Context, in *pro.RequestKey, out *pro.ResponseVehicleAddress) error {
	vehManager := vh.New()
	result, err := vehManager.GetVehicleLocationByVehicle(uint(in.Id))

	for _, item := range result {

		vehicleDate, _ := timestamp.TimestampProto(item.UpdateAt)
		addUpdateAt, _ := timestamp.TimestampProto(item.Address.UpdatedAt)
		vehicleAddress := &pro.VehicleAddressProto{
			VehicleId: uint64(item.VehicleId),
			AddressId: uint64(item.AddressId),
			Primary:   item.Primary,
			UpdateAt:  vehicleDate,
			Address: &pro.AddressProto{
				Id:            uint64(item.Address.Id),
				Address:       item.Address.Address,
				Street:        item.Address.Street,
				Suburb:        item.Address.Suburb,
				StateId:       uint64(item.Address.StateId),
				CountryId:     uint64(item.Address.CountryId),
				AddressTypeId: uint64(item.Address.AddressTypeId),
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
	vehManager := vh.New()
	result, err := vehManager.CreateVehicleMake(bu.VehicleMakeBO{
		CountryId: uint(in.VehicleMake.CountryId),
		Make:      in.VehicleMake.Make,
	})

	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Id = uint64(result)
	return nil
}

func (m *MasterService) UpdateVehicleMake(ctx context.Context, in *pro.RequestVehicleMake, out *pro.ResponseSuccess) error {
	vehManager := vh.New()
	result, err := vehManager.UpdateVehicleMake(bu.VehicleMakeBO{
		Id:        uint(in.VehicleMake.Id),
		CountryId: uint(in.VehicleMake.CountryId),
		Make:      in.VehicleMake.Make,
	})
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Success = result
	return nil
}

func (m *MasterService) DeleteVehicleMake(ctx context.Context, in *pro.RequestDelete, out *pro.ResponseSuccess) error {
	vehManager := vh.New()
	result, err := vehManager.DeleteVehicleMake(uint(in.Id))
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Success = result
	return nil

}

func (m *MasterService) GetAllVehicleMake(ctx context.Context, in *empty.Empty, out *pro.ResponseVehicleMake) error {
	vehManager := vh.New()
	result, err := vehManager.GetAllVehicleMake()
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	for _, item := range result {
		updateAt, _ := timestamp.TimestampProto(item.UpdateAt)
		out.VehicleMake = append(out.VehicleMake, &pro.VehicleMakeProto{
			Id:        uint64(item.Id),
			Make:      item.Make,
			CountryId: uint64(item.CountryId),
			UpdateAt:  updateAt,
			Country: &pro.CountryProto{
				Id:          uint64(item.Country.Id),
				CountryName: item.Country.CountryName,
				RegionId:    uint64(item.Country.RegionId),
			},
		})
	}
	return nil
}

func (m *MasterService) GetVehicleMakeById(ctx context.Context, in *pro.RequestKey, out *pro.ResponseVehicleMake) error {
	vehManager := vh.New()
	result, err := vehManager.GetVehicleMakeById(uint(in.Id))
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	makeDate, _ := timestamp.TimestampProto(result.UpdateAt)
	out.VehicleMake = append(out.VehicleMake, &pro.VehicleMakeProto{
		Id:        uint64(result.Id),
		Make:      result.Make,
		CountryId: uint64(result.CountryId),
		UpdateAt:  makeDate,
		Country: &pro.CountryProto{
			Id:          uint64(result.Country.Id),
			CountryName: result.Country.CountryName,
			RegionId:    uint64(result.Country.RegionId),
		}})

	return nil
}

func (m *MasterService) CreateVehicleModel(ctx context.Context, in *pro.RequestVehicleModel, out *pro.ResponseCreateSuccess) error {
	vehManager := vh.New()
	result, err := vehManager.CreateVehicleModel(bu.VehicleModelBO{
		Description: in.VehicleModel.Description,
		ModelName:   in.VehicleModel.ModelName,
		MakeId:      uint(in.VehicleModel.MakeId),
	})
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Id = uint64(result)
	return nil
}

func (m *MasterService) UpdateVehicleModel(ctx context.Context, in *pro.RequestVehicleModel, out *pro.ResponseSuccess) error {
	vehManager := vh.New()
	result, err := vehManager.UpdateVehicleModel(bu.VehicleModelBO{
		MakeId:      uint(in.VehicleModel.MakeId),
		ModelName:   in.VehicleModel.ModelName,
		Description: in.VehicleModel.Description,
	})
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Success = result
	return nil
}

func (m *MasterService) DeleteVehicleModel(ctx context.Context, in *pro.RequestDelete, out *pro.ResponseSuccess) error {
	vehManager := vh.New()
	result, err := vehManager.DeleteVehicleModel(uint(in.Id))
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Success = result
	return nil
}
func (m *MasterService) GetAllModelByMake(ctx context.Context, in *pro.RequestKey, out *pro.ResponseVehicleModel) error {
	vehManager := vh.New()
	result, err := vehManager.GetAllModelByMake(uint(in.Id))
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	for _, item := range result {
		modelDate, _ := timestamp.TimestampProto(item.UpdatedAt)
		makeDate, _ := timestamp.TimestampProto(item.Make.UpdateAt)
		countryDate, _ := timestamp.TimestampProto(item.Make.Country.UpdatedAt)
		out.VehicleModel = append(out.VehicleModel, &pro.VehicleModelProto{
			Id:          uint64(item.Id),
			ModelName:   item.ModelName,
			Description: item.Description,
			MakeId:      uint64(item.MakeId),
			UpdatedAt:   modelDate,
			Make: &pro.VehicleMakeProto{
				Id:        uint64(item.Make.Id),
				Make:      item.Make.Make,
				CountryId: uint64(item.Make.CountryId),
				UpdateAt:  makeDate,
				Country: &pro.CountryProto{
					Id:          uint64(item.Make.Country.Id),
					CountryName: item.Make.Country.CountryName,
					RegionId:    uint64(item.Make.Country.RegionId),
					UpdatedAt:   countryDate,
				},
			},
		})
	}
	return nil
}

func (m *MasterService) GetModelById(ctx context.Context, in *pro.RequestKey, out *pro.ResponseVehicleModel) error {
	vehManager := vh.New()
	result, err := vehManager.GetModelById(uint(in.Id))
	out.Errors = ErrorResponse.GetCreateErrorJson(err)

	modelDate, _ := timestamp.TimestampProto(result.UpdatedAt)
	makeDate, _ := timestamp.TimestampProto(result.Make.UpdateAt)
	countryDate, _ := timestamp.TimestampProto(result.Make.Country.UpdatedAt)

	out.VehicleModel = append(out.VehicleModel, &pro.VehicleModelProto{
		Id:          uint64(result.Id),
		ModelName:   result.ModelName,
		Description: result.Description,
		MakeId:      uint64(result.MakeId),
		UpdatedAt:   modelDate,
		Make: &pro.VehicleMakeProto{
			Id:        uint64(result.Make.Id),
			Make:      result.Make.Make,
			CountryId: uint64(result.Make.CountryId),
			UpdateAt:  makeDate,
			Country: &pro.CountryProto{
				Id:          uint64(result.Make.Country.Id),
				CountryName: result.Make.Country.CountryName,
				RegionId:    uint64(result.Make.Country.RegionId),
				UpdatedAt:   countryDate,
			},
		},
	})
	return nil
}

func (m *MasterService) CreateVehicleReg(ctx context.Context, in *pro.RequestVehicleReg, out *pro.ResponseCreateSuccess) error {
	vehManager := vh.New()
	expireDate, _ := timestamp.Timestamp(in.VehicleReg.ExpiredDate)
	regDate, _ := timestamp.Timestamp(in.VehicleReg.RegisterDate)

	result, err := vehManager.CreateVehicleReg(bu.VehicleTrackRegBO{
		Active:       in.VehicleReg.Active,
		Duration:     int(in.VehicleReg.Duration),
		ExpiredDate:  expireDate,
		VehicleId:    uint(in.VehicleReg.VehicleId),
		RegisterDate: regDate,
	})
	out.Id = uint64(result)
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	return nil
}
func (m *MasterService) UpdateVehicleReg(ctx context.Context, in *pro.RequestVehicleReg, out *pro.ResponseSuccess) error {
	vehManager := vh.New()
	expireDate, _ := timestamp.Timestamp(in.VehicleReg.ExpiredDate)
	regDate, _ := timestamp.Timestamp(in.VehicleReg.RegisterDate)

	result, err := vehManager.UpdateVehicleReg(bu.VehicleTrackRegBO{
		Id:           uint(in.VehicleReg.Id),
		VehicleId:    uint(in.VehicleReg.VehicleId),
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
	vehManager := vh.New()
	result, err := vehManager.DeleteVehicleReg(uint(in.Id))
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Success = result
	return nil
}

func (m *MasterService) GetAllRegistrationsByVehicleId(ctx context.Context, in *pro.RequestKey, out *pro.ResponseVehicleReg) error {
	vehManager := vh.New()
	result, err := vehManager.GetAllRegistrationsByVehicleId(uint(in.Id))
	out.Errors = ErrorResponse.GetCreateErrorJson(err)

	for _, vh := range result {

		expireDate, _ := timestamp.TimestampProto(vh.ExpiredDate)
		regDate, _ := timestamp.TimestampProto(vh.RegisterDate)
		updateAt, _ := timestamp.TimestampProto(vh.UpdatedAt)
		vhUpdateAt, _ := timestamp.TimestampProto(vh.Vehicle.UpdatedAt)

		out.VehicleReg = append(out.VehicleReg, &pro.VehicleTrackRegProto{
			Id:           uint64(vh.Id),
			RegisterDate: regDate,
			Duration:     int32(vh.Duration),
			ExpiredDate:  expireDate,
			Active:       vh.Active,
			VehicleId:    uint64(vh.VehicleId),
			UpdatedAt:    updateAt,
			Vehicle: &pro.VehicleProto{
				Id:           uint64(vh.Vehicle.Id),
				UpdatedAt:    vhUpdateAt,
				Registration: vh.Vehicle.Registration,
				StatusId:     uint64(vh.Vehicle.StatusId),
				MakeId:       uint64(vh.Vehicle.MakeId),
				ModelId:      uint64(vh.Vehicle.ModelId),
			},
		})
	}
	return nil

}

func (m *MasterService) GetActiveRegistrationsByVehicleId(ctx context.Context, in *pro.RequestKey, out *pro.ResponseVehicleReg) error {

	vehManager := vh.New()
	result, err := vehManager.GetActiveRegistrationsByVehicleId(uint(in.Id))
	out.Errors = ErrorResponse.GetCreateErrorJson(err)

	expireDate, _ := timestamp.TimestampProto(result.ExpiredDate)
	regDate, _ := timestamp.TimestampProto(result.RegisterDate)
	updateAt, _ := timestamp.TimestampProto(result.UpdatedAt)
	vhUpdateAt, _ := timestamp.TimestampProto(result.Vehicle.UpdatedAt)

	out.VehicleReg = append(out.VehicleReg, &pro.VehicleTrackRegProto{
		Id:           uint64(result.Id),
		RegisterDate: regDate,
		Duration:     int32(result.Duration),
		ExpiredDate:  expireDate,
		Active:       result.Active,
		VehicleId:    uint64(result.VehicleId),
		UpdatedAt:    updateAt,
		Vehicle: &pro.VehicleProto{
			Id:           uint64(result.Vehicle.Id),
			UpdatedAt:    vhUpdateAt,
			Registration: result.Vehicle.Registration,
			StatusId:     uint64(result.Vehicle.StatusId),
			MakeId:       uint64(result.Vehicle.MakeId),
			ModelId:      uint64(result.Vehicle.ModelId),
		},
	})
	return nil
}

func (m *MasterService) CreateVehicleStatus(ctx context.Context, in *pro.RequestVehicleStatus, out *pro.ResponseCreateSuccess) error {
	vehManager := vh.New()
	result, err := vehManager.CreateVehicleStatus(bu.VehicleStatusBO{
		Id:         uint(in.VehicleStatus.Id),
		StatusName: in.VehicleStatus.StatusName,
		StatusType: in.VehicleStatus.StatusType,
	})

	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Success = true
	out.Id = uint64(result)
	return nil
}

func (m *MasterService) UpdateVehicleStatus(ctx context.Context, in *pro.RequestVehicleStatus, out *pro.ResponseSuccess) error {
	vehManager := vh.New()
	result, err := vehManager.UpdateVehicleStatus(bu.VehicleStatusBO{
		Id:         uint(in.VehicleStatus.Id),
		StatusType: in.VehicleStatus.StatusType,
		StatusName: in.VehicleStatus.StatusName,
	})
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Success = result
	return nil
}

func (m *MasterService) DeleteVehicleStatus(ctx context.Context, in *pro.RequestDelete, out *pro.ResponseSuccess) error {
	vehManager := vh.New()
	result, err := vehManager.DeleteVehicleStatus(uint(in.Id))
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Success = result
	return nil
}

func (m *MasterService) GetAllVehicleStatus(ctx context.Context, in *empty.Empty, out *pro.ResponseVehicleStatus) error {
	vehManager := vh.New()
	result, err := vehManager.GetAllVehicleStatus()
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	for _, item := range result {

		updateAt, _ := timestamp.TimestampProto(item.UpdatedAt)
		out.VehicleStatus = append(out.VehicleStatus, &pro.VehicleStatusProto{
			Id:         uint64(item.Id),
			StatusName: item.StatusName,
			StatusType: item.StatusType,
			UpdatedAt:  updateAt,
		})
	}
	return nil
}

func (m *MasterService) CreateVehicleOpBound(ctx context.Context, in *pro.RequestVehicleOprBound, out *pro.ResponseCreateSuccess) error {
	vehManager := vh.New()
	result, err := vehManager.CreateVehicleOpBound(bu.VehicleOperatorBoundBO{
		Active:     in.VehicleOprBound.Active,
		VehicleId:  uint(in.VehicleOprBound.VehicleId),
		OperatorId: uint(in.VehicleOprBound.OperatorId),
	})
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Id = uint64(result)
	return nil
}

func (m *MasterService) UpdateVehicleOpBound(ctx context.Context, in *pro.RequestVehicleOprBound, out *pro.ResponseSuccess) error {
	vehManager := vh.New()
	result, err := vehManager.UpdateVehicleOpBound(bu.VehicleOperatorBoundBO{
		Id:         uint(in.VehicleOprBound.Id),
		Active:     in.VehicleOprBound.Active,
		VehicleId:  uint(in.VehicleOprBound.VehicleId),
		OperatorId: uint(in.VehicleOprBound.OperatorId),
	})
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Success = result
	return nil
}

func (m *MasterService) DeleteVehicleOpBound(ctx context.Context, in *pro.RequestDelete, out *pro.ResponseSuccess) error {
	vehManager := vh.New()
	result, err := vehManager.DeleteVehicleOpBound(uint(in.Id))
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Success = result
	return nil
}
