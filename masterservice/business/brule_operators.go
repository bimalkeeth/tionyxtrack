package business

import (
	"errors"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	bu "tionyxtrack/masterservice/businesscontracts"
	en "tionyxtrack/masterservice/entities"
)

type IOperator interface {
	CreateOperator(bo bu.OperatorBO) (uuid.UUID, error)
	UpdateOperator(bo bu.OperatorBO) (bool, error)
	DeleteOperator(id uuid.UUID) (bool, error)
	GetOperatorById(id uuid.UUID) (bu.OperatorBO, error)
	GetOperatorsByVehicleId(id uint) ([]bu.OperatorBO, error)
}

type Operator struct {
	Db *gorm.DB
}

func NewOperator(db *gorm.DB) *Operator {
	return &Operator{Db: db}
}

//--------------------------------------------
//Create Operator
//--------------------------------------------
func (o *Operator) CreateOperator(bo bu.OperatorBO) (uuid.UUID, error) {
	op := en.TableVehicleOperators{Active: bo.Active,
		SurName:    bo.SurName,
		Name:       bo.Name,
		DrivingLic: bo.DrivingLic,
	}
	o.Db.Create(&op)
	return op.ID, nil
}

//---------------------------------------------
//Update operator
//---------------------------------------------
func (o *Operator) UpdateOperator(bo bu.OperatorBO) (bool, error) {

	operator := en.TableVehicleOperators{}
	o.Db.First(&operator, bo.Id)
	if operator.ID == uuid.Nil {
		return false, errors.New("operator could not be found")
	}
	operator.DrivingLic = bo.DrivingLic
	operator.Name = bo.Name
	operator.SurName = bo.SurName
	operator.Active = bo.Active
	o.Db.Save(&operator)
	return true, nil
}

//---------------------------------------------
//Delete operator
//---------------------------------------------
func (o *Operator) DeleteOperator(id uuid.UUID) (bool, error) {
	operator := en.TableVehicleOperators{}
	o.Db.First(&operator, id)
	if operator.ID == uuid.Nil {
		return false, errors.New("operator could not be found")
	}
	o.Db.Delete(&operator)
	return true, nil
}

//---------------------------------------------
//Get Operator by Id
//---------------------------------------------
func (o *Operator) GetOperatorById(id uuid.UUID) (bu.OperatorBO, error) {
	var tbop en.TableVehicleOperators
	var result bu.OperatorBO

	o.Db.Preload("Bounds").
		Preload("Bounds.Vehicle").
		Preload("Locations").
		Preload("Locations.Address").
		Preload("Contacts").
		Preload("Contacts.Contact").First(tbop, id)

	result.Active = tbop.Active
	result.SurName = tbop.SurName
	result.Name = tbop.Name
	result.DrivingLic = tbop.DrivingLic

	for _, item := range tbop.Locations {
		result.Locations = append(result.Locations, &bu.OperatorLocationBO{
			Id:         item.ID,
			AddressId:  item.AddressId,
			OperatorId: item.OperatorId,
			Primary:    item.Primary,
			UpdateAt:   item.UpdatedAt,
			Address: bu.AddressBO{
				Id:            item.ID,
				Address:       item.Address.Address,
				Location:      item.Address.Location,
				Street:        item.Address.Street,
				Suburb:        item.Address.Suburb,
				AddressTypeId: item.Address.AddressTypeId,
				CountryId:     item.Address.CountryId,
				StateId:       item.Address.StateId,
				UpdatedAt:     item.Address.UpdatedAt,
			},
		})
	}
	for _, item := range tbop.Contacts {
		result.Contacts = append(result.Contacts, &bu.OperatorContactsBO{
			Id:         item.ID,
			Primary:    item.Primary,
			OperatorId: item.OperatorId,
			ContactId:  item.ContactId,
			Contact: bu.ContactBO{
				Id:            item.Contact.ID,
				Contact:       item.Contact.Contact,
				ContactTypeId: item.Contact.ContactTypeId,
				UpdatedAt:     item.Contact.UpdatedAt,
			},
		})
	}

	for _, item := range tbop.Bounds {
		result.Vehicles = append(result.Vehicles, &bu.VehicleBO{
			Id:           item.Vehicle.ID,
			UpdatedAt:    item.Vehicle.UpdatedAt,
			StatusId:     item.Vehicle.StatusId,
			ModelId:      item.Vehicle.ModelId,
			MakeId:       item.Vehicle.MakeId,
			Registration: item.Vehicle.Registration,
			FleetId:      item.Vehicle.FleetId,
		})
	}
	return result, nil
}

//------------------------------------------------
//Get operators by VehicleId
//------------------------------------------------
func (o *Operator) GetOperatorsByVehicleId(id uuid.UUID) ([]bu.OperatorBO, error) {

	vehicle := en.TableVehicle{}
	var resultlist []bu.OperatorBO

	o.Db.Preload("Operators.Operator").
		Preload("Operators.Operator.Locations").
		Preload("Operators.Operator.Locations.Address").
		Preload("Operators.Operator.Contacts").
		Preload("Operators.Operator.Contacts.Contact").First(&vehicle, id)

	for _, opbounds := range vehicle.Operators {

		tbop := opbounds.Operator
		result := bu.OperatorBO{}

		result.Active = tbop.Active
		result.SurName = tbop.SurName
		result.Name = tbop.Name
		result.DrivingLic = tbop.DrivingLic

		for _, item := range tbop.Locations {
			result.Locations = append(result.Locations, &bu.OperatorLocationBO{
				Id:         item.ID,
				AddressId:  item.AddressId,
				OperatorId: item.OperatorId,
				Primary:    item.Primary,
				UpdateAt:   item.UpdatedAt,
				Address: bu.AddressBO{
					Id:            item.ID,
					Address:       item.Address.Address,
					Location:      item.Address.Location,
					Street:        item.Address.Street,
					Suburb:        item.Address.Suburb,
					AddressTypeId: item.Address.AddressTypeId,
					CountryId:     item.Address.CountryId,
					StateId:       item.Address.StateId,
					UpdatedAt:     item.Address.UpdatedAt,
				},
			})
		}
		for _, item := range tbop.Contacts {
			result.Contacts = append(result.Contacts, &bu.OperatorContactsBO{
				Id:         item.ID,
				Primary:    item.Primary,
				OperatorId: item.OperatorId,
				ContactId:  item.ContactId,
				Contact: bu.ContactBO{
					Id:            item.Contact.ID,
					Contact:       item.Contact.Contact,
					ContactTypeId: item.Contact.ContactTypeId,
					UpdatedAt:     item.Contact.UpdatedAt,
				},
			})
		}

		for _, item := range tbop.Bounds {
			result.Vehicles = append(result.Vehicles, &bu.VehicleBO{
				Id:           item.Vehicle.ID,
				UpdatedAt:    item.Vehicle.UpdatedAt,
				StatusId:     item.Vehicle.StatusId,
				ModelId:      item.Vehicle.ModelId,
				MakeId:       item.Vehicle.MakeId,
				Registration: item.Vehicle.Registration,
				FleetId:      item.Vehicle.FleetId,
			})
		}
		resultlist = append(resultlist, result)
	}
	return resultlist, nil
}
