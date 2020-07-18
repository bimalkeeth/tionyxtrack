package protoapi

import (
	"context"
	"errors"
	timestamp "github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/empty"
	uuid "github.com/satori/go.uuid"
	"time"
	bu "tionyxtrack/masterservice/businesscontracts"
	msg "tionyxtrack/masterservice/common/messages"
	val "tionyxtrack/masterservice/common/validation"
	con "tionyxtrack/masterservice/manager/contacts"
	mst "tionyxtrack/masterservice/manager/masters"
	pro "tionyxtrack/masterservice/proto"
)

func (m *MasterService) CreateCompany(ctx context.Context, in *pro.RequestCompany, out *pro.ResponseCreateSuccess) error {

	if err := val.ValidateCompany(val.CreateCompany, in.Company); err != nil {
		out.Errors = ErrorResponse.GetCreateErrorJson(err)
		return nil
	}

	defer RecoverError()

	masterManager := mst.New()

	companyRequest := bu.CompanyBO{
		Name:      in.Company.Name,
		ContactId: uuid.FromStringOrNil(in.Company.ContactId),
		AddressId: uuid.FromStringOrNil(in.Company.AddressId),
	}
	id, err := masterManager.CreateCompany(companyRequest)
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Id = id.String()
	return nil
}

func (m *MasterService) UpdateCompany(ctx context.Context, in *pro.RequestCompany, out *pro.ResponseSuccess) error {

	defer RecoverError()

	if err := val.ValidateCompany(val.UpdateCompany, in.Company); err != nil {
		out.Errors = ErrorResponse.GetCreateErrorJson(err)
		return nil
	}

	masterManager := mst.New()

	companyRequest := bu.CompanyBO{
		Id:        uuid.FromStringOrNil(in.Company.Id),
		Name:      in.Company.Name,
		ContactId: uuid.FromStringOrNil(in.Company.ContactId),
		AddressId: uuid.FromStringOrNil(in.Company.AddressId),
	}
	success, err := masterManager.UpdateCompany(companyRequest)
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Success = success
	return nil
}

func (m *MasterService) DeleteCompany(ctx context.Context, in *pro.RequestDelete, out *pro.ResponseSuccess) error {

	defer RecoverError()

	if err := val.ValidateDelete(in); err != nil {
		out.Errors = ErrorResponse.GetCreateErrorJson(err)
		return nil
	}

	masterManager := mst.New()

	if uuid.FromStringOrNil(in.Id) == uuid.Nil {
		out.Errors = ErrorResponse.GetCreateErrorJson(errors.New(msg.ErrorNotDefinedCompanyId))
		return nil
	}

	result, err := masterManager.DeleteCompany(uuid.FromStringOrNil(in.Id))
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Success = result
	return nil
}

func (m *MasterService) CreateAddressType(ctx context.Context, in *pro.RequestAddressType, out *pro.ResponseCreateSuccess) error {

	defer RecoverError()

	if err := val.ValidateAddressType(val.CreateAddressType, in.AddressType); err != nil {
		out.Errors = ErrorResponse.GetCreateErrorJson(err)
		return nil
	}

	masterManager := mst.New()

	result, err := masterManager.CreateAddressType(bu.AddressTypeBO{
		Name: in.AddressType.Name,
	})
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Id = result.String()
	return nil
}

func (m *MasterService) UpdateAddressType(ctx context.Context, in *pro.RequestAddressType, out *pro.ResponseSuccess) error {

	defer RecoverError()

	if err := val.ValidateAddressType(val.UpdateAddressType, in.AddressType); err != nil {
		out.Errors = ErrorResponse.GetCreateErrorJson(err)
		return nil
	}

	masterManager := mst.New()

	result, err := masterManager.UpdateAddressType(bu.AddressTypeBO{
		Name: in.AddressType.Name,
	})

	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Success = result
	return nil
}

func (m *MasterService) DeleteAddressType(ctx context.Context, in *pro.RequestDelete, out *pro.ResponseSuccess) error {

	defer RecoverError()

	if err := val.ValidateDelete(in); err != nil {
		out.Errors = ErrorResponse.GetCreateErrorJson(err)
		return nil
	}

	masterManager := mst.New()

	result, err := masterManager.DeleteAddressType(uuid.FromStringOrNil(in.Id))

	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Success = result
	return nil
}

func (m *MasterService) GetAddressTypeById(ctx context.Context, in *pro.RequestKey, out *pro.ResponseAddressType) error {

	defer RecoverError()

	masterManager := mst.New()
	result, err := masterManager.GetAddressTypeById(uuid.FromStringOrNil(in.Id))

	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	var addressType []*pro.AddressTypeProto
	addressType = append(addressType, &pro.AddressTypeProto{
		Id:   result.Id.String(),
		Name: result.Name,
	})
	out.AddressType = addressType
	return nil
}

func (m *MasterService) GetAddressTypeByName(ctx context.Context, in *pro.RequestByName, out *pro.ResponseAddressType) error {

	defer RecoverError()

	masterManager := mst.New()
	result, err := masterManager.GetAddressTypeByName(in.Name)
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	var addressType []*pro.AddressTypeProto
	addressType = append(addressType, &pro.AddressTypeProto{
		Id:   result.Id.String(),
		Name: result.Name,
	})
	out.AddressType = addressType
	return nil
}

func (m *MasterService) GetAllAddressTypes(ctx context.Context, in *empty.Empty, out *pro.ResponseAddressType) error {

	defer RecoverError()

	masterManager := mst.New()
	result, err := masterManager.GetAllAddressTypes()
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	var addressType []*pro.AddressTypeProto
	for _, item := range result {
		addressType = append(addressType, &pro.AddressTypeProto{
			Id:   item.Id.String(),
			Name: item.Name,
		})
	}
	out.AddressType = addressType
	return nil
}

func (m *MasterService) GetAllAddressTypeNames(ctx context.Context, in *pro.RequestByName, out *pro.ResponseAddressType) error {

	defer RecoverError()

	masterManager := mst.New()
	result, err := masterManager.GetAllAddressTypeNames(in.Name)

	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	var addressType []*pro.AddressTypeProto
	for _, item := range result {
		addressType = append(addressType, &pro.AddressTypeProto{
			Id:   item.Id.String(),
			Name: item.Name,
		})
	}
	out.AddressType = addressType
	return nil
}

func (m *MasterService) CreateRegion(ctx context.Context, in *pro.RequestRegion, out *pro.ResponseCreateSuccess) error {

	defer RecoverError()

	if err := val.ValidateRegion(val.CreateRegion, in.Region); err != nil {
		out.Errors = ErrorResponse.GetCreateErrorJson(err)
		return nil
	}

	masterManager := mst.New()

	result, err := masterManager.CreateRegion(bu.RegionBO{Region: in.Region.Region})
	response := &pro.ResponseCreateSuccess{}
	response.Errors = ErrorResponse.GetCreateErrorJson(err)
	response.Id = result.String()
	return nil
}

func (m *MasterService) UpdateRegion(ctx context.Context, in *pro.RequestRegion, out *pro.ResponseSuccess) error {

	defer RecoverError()

	if err := val.ValidateRegion(val.UpdateRegion, in.Region); err != nil {
		out.Errors = ErrorResponse.GetCreateErrorJson(err)
		return nil
	}

	masterManager := mst.New()

	result, err := masterManager.UpdateRegion(bu.RegionBO{Id: uuid.FromStringOrNil(in.Region.Id), Region: in.Region.Region})
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Success = result
	return nil
}

func (m *MasterService) DeleteRegion(ctx context.Context, in *pro.RequestDelete, out *pro.ResponseSuccess) error {

	defer RecoverError()

	if err := val.ValidateDelete(in); err != nil {
		out.Errors = ErrorResponse.GetCreateErrorJson(err)
		return nil
	}

	masterManager := mst.New()

	result, err := masterManager.DeleteRegion(uuid.FromStringOrNil(in.Id))
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Success = result
	return nil
}

func (m *MasterService) GetAllRegion(ctx context.Context, in *empty.Empty, out *pro.ResponseRegion) error {

	defer RecoverError()

	masterManager := mst.New()

	result, err := masterManager.GetAllRegion()

	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	for _, item := range result {
		out.Region = append(out.Region, &pro.RegionProto{Region: item.Region,
			RegionName: item.RegionName,
			Id:         item.Id.String(),
		})
	}
	return nil
}

func (m *MasterService) GetRegionById(ctx context.Context, in *pro.RequestKey, out *pro.ResponseRegion) error {

	defer RecoverError()

	masterManager := mst.New()

	result, err := masterManager.GetRegionById(uuid.FromStringOrNil(in.Id))

	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Region = append(out.Region, &pro.RegionProto{Region: result.Region, RegionName: result.RegionName, Id: result.Id.String()})
	return nil
}

func (m *MasterService) GetRegionByName(ctx context.Context, in *pro.RequestByName, out *pro.ResponseRegion) error {

	defer RecoverError()

	masterManager := mst.New()
	result, err := masterManager.GetRegionByName(in.Name)
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Region = append(out.Region, &pro.RegionProto{Region: result.Region, RegionName: result.RegionName, Id: result.Id.String()})
	return nil

}

func (m *MasterService) CreateState(ctx context.Context, in *pro.RequestState, out *pro.ResponseCreateSuccess) error {

	defer RecoverError()

	if err := val.ValidateState(val.CreateState, in.State); err != nil {
		out.Errors = ErrorResponse.GetCreateErrorJson(err)
		return nil
	}

	masterManager := mst.New()

	result, err := masterManager.CreateState(bu.StateBO{Name: in.State.Name,
		CountryId: uuid.FromStringOrNil(in.State.CountryId)})

	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Id = result.String()
	return nil
}

func (m *MasterService) UpdateState(ctx context.Context, in *pro.RequestState, out *pro.ResponseSuccess) error {

	defer RecoverError()

	if err := val.ValidateState(val.UpdateState, in.State); err != nil {
		out.Errors = ErrorResponse.GetCreateErrorJson(err)
		return nil
	}

	masterManager := mst.New()

	result, err := masterManager.UpdateState(bu.StateBO{Name: in.State.Name,
		CountryId: uuid.FromStringOrNil(in.State.CountryId), Id: uuid.FromStringOrNil(in.State.Id)})
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Success = result
	return nil
}

func (m *MasterService) DeleteState(ctx context.Context, in *pro.RequestDelete, out *pro.ResponseSuccess) error {

	defer RecoverError()

	if err := val.ValidateDelete(in); err != nil {
		out.Errors = ErrorResponse.GetCreateErrorJson(err)
		return nil
	}

	masterManager := mst.New()

	result, err := masterManager.DeleteState(uuid.FromStringOrNil(in.Id))
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Success = result
	return nil
}

func (m *MasterService) GetStateById(ctx context.Context, in *pro.RequestKey, out *pro.ResponseState) error {

	defer RecoverError()

	masterManager := mst.New()
	result, err := masterManager.GetStateById(uuid.FromStringOrNil(in.Id))
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.State = append(out.State, &pro.StateProto{Id: result.Id.String(),
		CountryId: result.CountryId.String(), Name: result.Name})
	return nil
}

func (m *MasterService) GetStateByCountryId(ctx context.Context, in *pro.RequestKey, out *pro.ResponseState) error {

	defer RecoverError()

	masterManager := mst.New()
	result, err := masterManager.GetStateByCountryId(uuid.FromStringOrNil(in.Id))

	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	for _, item := range result {
		out.State = append(out.State, &pro.StateProto{
			Id:        item.Id.String(),
			CountryId: item.CountryId.String(),
			Name:      item.Name,
		})
	}
	return nil
}

func (m *MasterService) GetStateByName(ctx context.Context, in *pro.RequestByName, out *pro.ResponseState) error {

	defer RecoverError()

	masterManager := mst.New()
	result, err := masterManager.GetStateByName(in.Name)

	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.State = append(out.State, &pro.StateProto{Id: result.Id.String(),
		CountryId: result.CountryId.String(), Name: result.Name})
	return nil
}

func (m *MasterService) GetAllStates(ctx context.Context, in *empty.Empty, out *pro.ResponseState) error {

	defer RecoverError()

	masterManager := mst.New()
	result, err := masterManager.GetAllStates()
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	for _, item := range result {
		out.State = append(out.State, &pro.StateProto{
			Id:        item.Id.String(),
			CountryId: item.CountryId.String(),
			Name:      item.Name,
		})
	}
	return nil
}

func (m *MasterService) CreateContact(ctx context.Context, in *pro.RequestContact, out *pro.ResponseCreateSuccess) error {

	defer RecoverError()

	if err := val.ValidateContact(val.CreateContact, in.Contact); err != nil {
		out.Errors = ErrorResponse.GetCreateErrorJson(err)
		return nil
	}

	conManager := con.New()

	result, err := conManager.CreateContact(bu.ContactBO{Contact: in.Contact.Contact,
		ContactTypeId: uuid.FromStringOrNil(in.Contact.ContactTypeId),
	})
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Id = result.String()
	return nil
}

func (m *MasterService) UpdateContact(ctx context.Context, in *pro.RequestContact, out *pro.ResponseSuccess) error {

	defer RecoverError()

	if err := val.ValidateContact(val.UpdateContact, in.Contact); err != nil {
		out.Errors = ErrorResponse.GetCreateErrorJson(err)
		return nil
	}

	conManager := con.New()

	result, err := conManager.UpdateContact(bu.ContactBO{Id: uuid.FromStringOrNil(in.Contact.Id),
		Contact:       in.Contact.Contact,
		ContactTypeId: uuid.FromStringOrNil(in.Contact.ContactTypeId),
	})
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Success = result
	return nil
}

func (m *MasterService) DeleteContact(ctx context.Context, in *pro.RequestDelete, out *pro.ResponseSuccess) error {

	defer RecoverError()

	if err := val.ValidateDelete(in); err != nil {
		out.Errors = ErrorResponse.GetCreateErrorJson(err)
		return nil
	}

	conManager := con.New()

	result, err := conManager.DeleteContact(uuid.FromStringOrNil(in.Id))
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Success = result
	return nil
}

func (m *MasterService) ContactById(ctx context.Context, in *pro.RequestKey, out *pro.ResponseContact) error {

	defer RecoverError()

	conManager := con.New()

	result, err := conManager.ContactById(uuid.FromStringOrNil(in.Id))
	out.Errors = ErrorResponse.GetCreateErrorJson(err)

	timeUpdate, err := timestamp.TimestampProto(result.UpdatedAt)
	if err != nil {
		timeUpdate, _ = timestamp.TimestampProto(time.Now())
	}
	out.Contact = &pro.ContactProto{
		Id:            result.Id.String(),
		Contact:       result.Contact,
		ContactTypeId: result.ContactTypeId.String(),
		UpdatedAt:     timeUpdate,
	}
	return nil
}

func (m *MasterService) CreateAddress(ctx context.Context, in *pro.RequestAddress, out *pro.ResponseCreateSuccess) error {

	defer RecoverError()

	if err := val.ValidateAddress(val.CreateAddress, in.Address); err != nil {
		out.Errors = ErrorResponse.GetCreateErrorJson(err)
		return nil
	}

	conManager := con.New()

	result, err := conManager.CreateAddress(bu.AddressBO{
		CountryId:     uuid.FromStringOrNil(in.Address.CountryId),
		AddressTypeId: uuid.FromStringOrNil(in.Address.AddressTypeId),
		Location:      in.Address.Location,
		Address:       in.Address.Address,
		StateId:       uuid.FromStringOrNil(in.Address.StateId),
		Street:        in.Address.Street,
		Suburb:        in.Address.Suburb,
	})

	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Id = result.String()
	out.Success = true
	return err
}

func (m *MasterService) UpdateAddress(ctx context.Context, in *pro.RequestAddress, out *pro.ResponseSuccess) error {

	defer RecoverError()

	if err := val.ValidateAddress(val.UpdateAddress, in.Address); err != nil {
		out.Errors = ErrorResponse.GetCreateErrorJson(err)
		return nil
	}

	conManager := con.New()

	result, err := conManager.UpdateAddress(bu.AddressBO{
		Id:            uuid.FromStringOrNil(in.Address.Id),
		CountryId:     uuid.FromStringOrNil(in.Address.CountryId),
		AddressTypeId: uuid.FromStringOrNil(in.Address.AddressTypeId),
		Location:      in.Address.Location,
		Address:       in.Address.Address,
		StateId:       uuid.FromStringOrNil(in.Address.StateId),
		Street:        in.Address.Street,
		Suburb:        in.Address.Suburb,
	})

	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Success = result
	out.Success = true
	return err
}

func (m *MasterService) DeleteAddress(ctx context.Context, in *pro.RequestDelete, out *pro.ResponseSuccess) error {

	defer RecoverError()

	conManager := con.New()

	if err := val.ValidateDelete(in); err != nil {
		out.Errors = ErrorResponse.GetCreateErrorJson(err)
		return nil
	}

	result, err := conManager.DeleteAddress(uuid.FromStringOrNil(in.Id))
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Success = result
	return err
}

func (m *MasterService) GetAddressById(ctx context.Context, in *pro.RequestKey, out *pro.ResponseAddress) error {
	conManager := con.New()
	result, err := conManager.GetAddressById(uuid.FromStringOrNil(in.Id))
	out.Errors = ErrorResponse.GetCreateErrorJson(err)

	timeUpdate, errs := timestamp.TimestampProto(result.UpdatedAt)
	if errs != nil {
		timeUpdate, _ = timestamp.TimestampProto(time.Now())
	}

	out.Address = append(out.Address, &pro.AddressProto{
		Address:       result.Address,
		Id:            result.Id.String(),
		Suburb:        result.Suburb,
		Street:        result.Street,
		StateId:       result.StateId.String(),
		Location:      result.Location,
		AddressTypeId: result.AddressTypeId.String(),
		CountryId:     result.CountryId.String(),
		UpdatedAt:     timeUpdate,
	})
	return nil
}

func (m *MasterService) GetAddressByName(ctx context.Context, in *pro.RequestByName, out *pro.ResponseAddress) error {
	conManager := con.New()
	result, err := conManager.GetAddressByName(in.Name)
	out.Errors = ErrorResponse.GetCreateErrorJson(err)

	for _, item := range result {
		timeUpdate, errs := timestamp.TimestampProto(item.UpdatedAt)
		if errs != nil {
			timeUpdate, _ = timestamp.TimestampProto(time.Now())
		}
		out.Address = append(out.Address, &pro.AddressProto{
			Address:       item.Address,
			Id:            item.Id.String(),
			Suburb:        item.Suburb,
			Street:        item.Street,
			StateId:       item.StateId.String(),
			Location:      item.Location,
			AddressTypeId: item.AddressTypeId.String(),
			CountryId:     item.CountryId.String(),
			UpdatedAt:     timeUpdate,
		})
	}
	return nil
}
