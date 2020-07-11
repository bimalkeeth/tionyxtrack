package protoapi

import (
	"context"
	timestamp "github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/empty"
	"time"
	bu "tionyxtrack/masterservice/businesscontracts"
	con "tionyxtrack/masterservice/manager/contacts"
	mst "tionyxtrack/masterservice/manager/masters"
	pro "tionyxtrack/masterservice/proto"
)


func (m *MasterService) CreateCompany(ctx context.Context, in *pro.RequestCompany, out *pro.ResponseCreateSuccess) error {
	masterManager := mst.New()
	companyRequest := bu.CompanyBO{
		Name:      in.Company.Name,
		ContactId: uint(in.Company.ContactId),
		AddressId: uint(in.Company.AddressId),
	}
	id, err := masterManager.CreateCompany(companyRequest)
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Id = uint64(id)
	return nil
}

func (m *MasterService) UpdateCompany(ctx context.Context, in *pro.RequestCompany, out *pro.ResponseSuccess) error {

	masterManager := mst.New()
	companyRequest := bu.CompanyBO{
		Id:        uint(in.Company.Id),
		Name:      in.Company.Name,
		ContactId: uint(in.Company.ContactId),
		AddressId: uint(in.Company.AddressId),
	}
	success, err := masterManager.UpdateCompany(companyRequest)
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Success = success
	return nil
}

func (m *MasterService) DeleteCompany(ctx context.Context, in *pro.RequestDelete, out *pro.ResponseSuccess) error {
	masterManager := mst.New()
	result, err := masterManager.DeleteCompany(uint(in.Id))
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Success = result
	return nil
}

func (m *MasterService) CreateAddressType(ctx context.Context, in *pro.RequestAddressType, out *pro.ResponseCreateSuccess) error {
	masterManager := mst.New()

	result, err := masterManager.CreateAddressType(bu.AddressTypeBO{
		Name: in.AddressType.Name,
	})
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Id = uint64(result)
	return nil
}

func (m *MasterService) UpdateAddressType(ctx context.Context, in *pro.RequestAddressType, out *pro.ResponseSuccess) error {
	masterManager := mst.New()
	result, err := masterManager.UpdateAddressType(bu.AddressTypeBO{
		Name: in.AddressType.Name,
	})

	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Success = result
	return nil
}

func (m *MasterService) DeleteAddressType(ctx context.Context, in *pro.RequestDelete, out *pro.ResponseSuccess) error {
	masterManager := mst.New()
	result, err := masterManager.DeleteAddressType(uint(in.Id))

	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Success = result
	return nil
}

func (m *MasterService) GetAddressTypeById(ctx context.Context, in *pro.RequestKey, out *pro.ResponseAddressType) error {
	masterManager := mst.New()
	result, err := masterManager.GetAddressTypeById(uint(in.Id))

	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	var addressType []*pro.AddressTypeProto
	addressType = append(addressType, &pro.AddressTypeProto{
		Id:   uint64(result.Id),
		Name: result.Name,
	})
	out.AddressType = addressType
	return nil
}

func (m *MasterService) GetAddressTypeByName(ctx context.Context, in *pro.RequestByName, out *pro.ResponseAddressType) error {
	masterManager := mst.New()
	result, err := masterManager.GetAddressTypeByName(in.Name)
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	var addressType []*pro.AddressTypeProto
	addressType = append(addressType, &pro.AddressTypeProto{
		Id:   uint64(result.Id),
		Name: result.Name,
	})
	out.AddressType = addressType
	return nil
}

func (m *MasterService) GetAllAddressTypes(ctx context.Context, in *empty.Empty, out *pro.ResponseAddressType) error {
	masterManager := mst.New()
	result, err := masterManager.GetAllAddressTypes()
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	var addressType []*pro.AddressTypeProto
	for _, item := range result {
		addressType = append(addressType, &pro.AddressTypeProto{
			Id:   uint64(item.Id),
			Name: item.Name,
		})
	}
	out.AddressType = addressType
	return nil
}

func (m *MasterService) GetAllAddressTypeNames(ctx context.Context, in *pro.RequestByName, out *pro.ResponseAddressType) error {
	masterManager := mst.New()
	result, err := masterManager.GetAllAddressTypeNames(in.Name)

	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	var addressType []*pro.AddressTypeProto
	for _, item := range result {
		addressType = append(addressType, &pro.AddressTypeProto{
			Id:   uint64(item.Id),
			Name: item.Name,
		})
	}
	out.AddressType = addressType
	return nil
}

func (m *MasterService) CreateRegion(ctx context.Context, in *pro.RequestRegion, out *pro.ResponseCreateSuccess) error {
	masterManager := mst.New()
	result, err := masterManager.CreateRegion(bu.RegionBO{Region: in.Region.Region})
	response := &pro.ResponseCreateSuccess{}
	response.Errors = ErrorResponse.GetCreateErrorJson(err)
	response.Id = uint64(result)
	return nil
}

func (m *MasterService) UpdateRegion(ctx context.Context, in *pro.RequestRegion, out *pro.ResponseSuccess) error {
	masterManager := mst.New()
	result, err := masterManager.UpdateRegion(bu.RegionBO{Id: uint(in.Region.Id), Region: in.Region.Region})
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Success = result
	return nil
}

func (m *MasterService) DeleteRegion(ctx context.Context, in *pro.RequestDelete, out *pro.ResponseSuccess) error {

	masterManager := mst.New()
	result, err := masterManager.DeleteRegion(uint(in.Id))
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Success = result
	return nil
}

func (m *MasterService) GetAllRegion(ctx context.Context, in *empty.Empty, out *pro.ResponseRegion) error {
	masterManager := mst.New()
	result, err := masterManager.GetAllRegion()

	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	for _, item := range result {
		out.Region = append(out.Region, &pro.RegionProto{Region: item.Region,
			RegionName: item.RegionName,
			Id:         uint64(item.Id),
		})
	}
	return nil
}

func (m *MasterService) GetRegionById(ctx context.Context, in *pro.RequestKey, out *pro.ResponseRegion) error {
	masterManager := mst.New()
	result, err := masterManager.GetRegionById(uint(in.Id))

	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Region = append(out.Region, &pro.RegionProto{Region: result.Region, RegionName: result.RegionName, Id: uint64(result.Id)})
	return nil
}

func (m *MasterService) GetRegionByName(ctx context.Context, in *pro.RequestByName, out *pro.ResponseRegion) error {
	masterManager := mst.New()
	result, err := masterManager.GetRegionByName(in.Name)
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Region = append(out.Region, &pro.RegionProto{Region: result.Region, RegionName: result.RegionName, Id: uint64(result.Id)})
	return nil

}

func (m *MasterService) CreateState(ctx context.Context, in *pro.RequestState, out *pro.ResponseCreateSuccess) error {
	masterManager := mst.New()
	result, err := masterManager.CreateState(bu.StateBO{Name: in.State.Name,
		CountryId: uint(in.State.CountryId)})

	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Id = uint64(result)
	return nil
}

func (m *MasterService) UpdateState(ctx context.Context, in *pro.RequestState, out *pro.ResponseSuccess) error {
	masterManager := mst.New()
	result, err := masterManager.UpdateState(bu.StateBO{Name: in.State.Name,
		CountryId: uint(in.State.CountryId), Id: uint(in.State.Id)})
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Success = result
	return nil
}

func (m *MasterService) DeleteState(ctx context.Context, in *pro.RequestDelete, out *pro.ResponseSuccess) error {
	masterManager := mst.New()
	result, err := masterManager.DeleteState(uint(in.Id))
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Success = result
	return nil
}

func (m *MasterService) GetStateById(ctx context.Context, in *pro.RequestKey, out *pro.ResponseState) error {
	masterManager := mst.New()
	result, err := masterManager.GetStateById(uint(in.Id))
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.State = append(out.State, &pro.StateProto{Id: uint64(result.Id),
		CountryId: uint64(result.CountryId), Name: result.Name})
	return nil
}

func (m *MasterService) GetStateByCountryId(ctx context.Context, in *pro.RequestKey, out *pro.ResponseState) error {
	masterManager := mst.New()
	result, err := masterManager.GetStateByCountryId(uint(in.Id))

	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	for _, item := range result {
		out.State = append(out.State, &pro.StateProto{
			Id:        uint64(item.Id),
			CountryId: uint64(item.CountryId),
			Name:      item.Name,
		})
	}
	return nil
}

func (m *MasterService) GetStateByName(ctx context.Context, in *pro.RequestByName, out *pro.ResponseState) error {
	masterManager := mst.New()
	result, err := masterManager.GetStateByName(in.Name)

	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.State = append(out.State, &pro.StateProto{Id: uint64(result.Id),
		CountryId: uint64(result.CountryId), Name: result.Name})
	return nil
}

func (m *MasterService) GetAllStates(ctx context.Context, in *empty.Empty, out *pro.ResponseState) error {
	masterManager := mst.New()
	result, err := masterManager.GetAllStates()
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	for _, item := range result {
		out.State = append(out.State, &pro.StateProto{
			Id:        uint64(item.Id),
			CountryId: uint64(item.CountryId),
			Name:      item.Name,
		})
	}
	return nil
}

func (m *MasterService) CreateContact(ctx context.Context, in *pro.RequestContact, out *pro.ResponseCreateSuccess) error {
	conManager := con.New()
	result, err := conManager.CreateContact(bu.ContactBO{Contact: in.Contact.Contact,
		ContactTypeId: uint(in.Contact.ContactTypeId),
	})
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Id = uint64(result)
	return nil
}

func (m *MasterService) UpdateContact(ctx context.Context, in *pro.RequestContact, out *pro.ResponseSuccess) error {
	conManager := con.New()
	result, err := conManager.UpdateContact(bu.ContactBO{Id: uint(in.Contact.Id),
		Contact:       in.Contact.Contact,
		ContactTypeId: uint(in.Contact.ContactTypeId),
	})
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Success = result
	return nil
}

func (m *MasterService) DeleteContact(ctx context.Context, in *pro.RequestDelete, out *pro.ResponseSuccess) error {
	conManager := con.New()
	result, err := conManager.DeleteContact(uint(in.Id))
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Success = result
	return nil
}

func (m *MasterService) ContactById(ctx context.Context, in *pro.RequestKey, out *pro.ResponseContact) error {
	conManager := con.New()
	result, err := conManager.ContactById(uint(in.Id))
	out.Errors = ErrorResponse.GetCreateErrorJson(err)

	timeUpdate, err := timestamp.TimestampProto(result.UpdatedAt)
	if err != nil {
		timeUpdate, _ = timestamp.TimestampProto(time.Now())
	}
	out.Contact = &pro.ContactProto{
		Id:            uint64(result.Id),
		Contact:       result.Contact,
		ContactTypeId: uint64(result.ContactTypeId),
		UpdatedAt:     timeUpdate,
	}
	return nil
}

func (m *MasterService) CreateAddress(ctx context.Context, in *pro.RequestAddress, out *pro.ResponseCreateSuccess) error {
	conManager := con.New()
	result, err := conManager.CreateAddress(bu.AddressBO{
		CountryId:     uint(in.Address.CountryId),
		AddressTypeId: uint(in.Address.AddressTypeId),
		Location:      in.Address.Location,
		Address:       in.Address.Address,
		StateId:       uint(in.Address.StateId),
		Street:        in.Address.Street,
		Suburb:        in.Address.Suburb,
	})

	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Id = uint64(result)
	out.Success = true
	return err
}

func (m *MasterService) UpdateAddress(ctx context.Context, in *pro.RequestAddress, out *pro.ResponseSuccess) error {
	conManager := con.New()
	result, err := conManager.UpdateAddress(bu.AddressBO{
		Id:            uint(in.Address.Id),
		CountryId:     uint(in.Address.CountryId),
		AddressTypeId: uint(in.Address.AddressTypeId),
		Location:      in.Address.Location,
		Address:       in.Address.Address,
		StateId:       uint(in.Address.StateId),
		Street:        in.Address.Street,
		Suburb:        in.Address.Suburb,
	})

	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Success = result
	out.Success = true
	return err
}

func (m *MasterService) DeleteAddress(ctx context.Context, in *pro.RequestDelete, out *pro.ResponseSuccess) error {
	conManager := con.New()
	result, err := conManager.DeleteAddress(uint(in.Id))
	out.Errors = ErrorResponse.GetCreateErrorJson(err)
	out.Success = result
	return err
}

func (m *MasterService) GetAddressById(ctx context.Context, in *pro.RequestKey, out *pro.ResponseAddress) error {
	conManager := con.New()
	result, err := conManager.GetAddressById(uint(in.Id))
	out.Errors = ErrorResponse.GetCreateErrorJson(err)

	timeUpdate, errs := timestamp.TimestampProto(result.UpdatedAt)
	if errs != nil {
		timeUpdate, _ = timestamp.TimestampProto(time.Now())
	}

	out.Address = append(out.Address, &pro.AddressProto{
		Address:       result.Address,
		Id:            uint64(result.Id),
		Suburb:        result.Suburb,
		Street:        result.Street,
		StateId:       uint64(result.StateId),
		Location:      result.Location,
		AddressTypeId: uint64(result.AddressTypeId),
		CountryId:     uint64(result.CountryId),
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
			Id:            uint64(item.Id),
			Suburb:        item.Suburb,
			Street:        item.Street,
			StateId:       uint64(item.StateId),
			Location:      item.Location,
			AddressTypeId: uint64(item.AddressTypeId),
			CountryId:     uint64(item.CountryId),
			UpdatedAt:     timeUpdate,
		})
	}
	return nil
}
