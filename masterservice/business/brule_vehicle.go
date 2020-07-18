package business

import (
	"errors"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	bu "tionyxtrack/masterservice/businesscontracts"
	ent "tionyxtrack/masterservice/entities"
)

type IVehicle interface {
	CreateVehicle(vehicle bu.VehicleBO) (uuid.UUID, error)
	UpdateVehicle(vehicle bu.VehicleBO) (bool, error)
	DeleteVehicle(vehicleId uuid.UUID) (bool, error)
	GetVehicleById(vehicleId uuid.UUID) (bu.VehicleBO, error)
	GetVehicleByRegistration(registration string) (bu.VehicleBO, error)
	GetVehiclesByFleetId(fleetId uuid.UUID) ([]bu.VehicleBO, error)
}
type Vehicle struct{ Db *gorm.DB }

func NewVehicle(db *gorm.DB) Vehicle {
	return Vehicle{Db: db}
}

//----------------------------------------------------
// Create Vehicle
//----------------------------------------------------
func (v *Vehicle) CreateVehicle(vehicle bu.VehicleBO) (uuid.UUID, error) {
	vehicleTable := ent.TableVehicle{MakeId: vehicle.MakeId,
		FleetId:      vehicle.FleetId,
		ModelId:      vehicle.ModelId,
		Registration: vehicle.Registration,
		StatusId:     vehicle.StatusId}
	v.Db.Create(&vehicleTable)
	return vehicleTable.ID, nil
}

//----------------------------------------------------
// Update Vehicle
//----------------------------------------------------
func (v *Vehicle) UpdateVehicle(vehicle bu.VehicleBO) (bool, error) {
	vehicleTable := ent.TableVehicle{}
	v.Db.First(&vehicleTable, vehicle.Id)
	if vehicleTable.ID == uuid.Nil {
		return false, errors.New("vehicle not found for update")
	}
	vehicleTable.MakeId = vehicle.MakeId
	vehicleTable.StatusId = vehicle.StatusId
	vehicleTable.Registration = vehicle.Registration
	vehicleTable.ModelId = vehicle.ModelId
	vehicleTable.FleetId = vehicle.FleetId
	v.Db.Save(&vehicleTable)
	return true, nil
}

//----------------------------------------------------
// Delete Vehicle
//----------------------------------------------------
func (v *Vehicle) DeleteVehicle(vehicleId uuid.UUID) (bool, error) {
	vehicleTable := ent.TableVehicle{}
	v.Db.First(&vehicleTable, vehicleId)
	if vehicleTable.ID == uuid.Nil {
		return false, errors.New("vehicle not found for update")
	}
	v.Db.Delete(&vehicleTable)
	return true, nil
}

//----------------------------------------------------
// Get Vehicle By Id
//----------------------------------------------------
func (v *Vehicle) GetVehicleById(vehicleId uuid.UUID) (bu.VehicleBO, error) {
	vhTab := ent.TableVehicle{}
	vehicleResult := bu.VehicleBO{}
	v.Db.Preload("VehicleModel").Preload("VehicleMake").
		Preload("Fleet").
		Preload("Status").
		Preload("Locations").
		Preload("Locations.Address").
		Preload("Locations.Address.State").
		Preload("Operators").
		Preload("Operators.Operator").
		Preload("Registrations").
		First(&vhTab, vehicleId)

	if vhTab.ID == uuid.Nil {
		return vehicleResult, nil
	}
	vehicleResult.FleetId = vhTab.FleetId
	vehicleResult.ModelId = vhTab.ModelId
	vehicleResult.Registration = vhTab.Registration
	vehicleResult.StatusId = vhTab.StatusId
	vehicleResult.MakeId = vhTab.MakeId
	vehicleResult.Status = bu.VehicleStatusBO{
		Id:         vhTab.Status.ID,
		StatusType: vhTab.Status.StatusType,
		StatusName: vhTab.Status.StatusName,
	}

	var vehicleLocation []bu.VehicleAddressBO
	for _, loc := range vhTab.Locations {

		vehicleLocation = append(vehicleLocation, bu.VehicleAddressBO{
			Id:        loc.ID,
			AddressId: loc.AddressId,
			VehicleId: loc.VehicleId,
			UpdateAt:  loc.UpdatedAt,
			Address: bu.AddressBO{
				Id:            loc.Address.ID,
				Address:       loc.Address.Address,
				UpdatedAt:     loc.Address.UpdatedAt,
				CountryId:     loc.Address.CountryId,
				Location:      loc.Address.Location,
				AddressTypeId: loc.Address.AddressTypeId,
				StateId:       loc.Address.StateId,
				Suburb:        loc.Address.Suburb,
				Street:        loc.Address.Street,
				State: bu.StateBO{
					Id:        loc.Address.StateId,
					Name:      loc.Address.State.Name,
					CountryId: loc.Address.State.CountryId,
				},
			},
		})
	}
	vehicleResult.Locations = vehicleLocation
	var ops []bu.VehicleOperatorBoundBO
	for _, op := range vhTab.Operators {

		ops = append(ops, bu.VehicleOperatorBoundBO{
			Id:         op.ID,
			OperatorId: op.OperatorId,
			VehicleId:  op.VehicleId,
			Active:     op.Active,
			Operator: &bu.OperatorBO{
				Id:      op.Operator.ID,
				Name:    op.Operator.Name,
				SurName: op.Operator.SurName,
			},
			Vehicle: nil,
		})
	}
	vehicleResult.Operators = ops

	var regis []bu.VehicleTrackRegBO
	for _, reg := range vhTab.Registrations {

		regis = append(regis, bu.VehicleTrackRegBO{
			Id:           reg.ID,
			Active:       reg.Active,
			VehicleId:    reg.VehicleId,
			UpdatedAt:    reg.UpdatedAt,
			Duration:     reg.Duration,
			ExpiredDate:  reg.ExpiredDate,
			RegisterDate: reg.ExpiredDate,
		})
	}
	vehicleResult.Registrations = regis

	return vehicleResult, nil

}

//----------------------------------------------------
// Get Vehicle By registration
//----------------------------------------------------
func (v *Vehicle) GetVehicleByRegistration(registration string) (bu.VehicleBO, error) {
	vhTab := ent.TableVehicle{}
	vehicleResult := bu.VehicleBO{}
	v.Db.Preload("VehicleModel").Preload("VehicleMake").
		Preload("Fleet").
		Preload("Status").
		Preload("Locations").
		Preload("Locations.Address").
		Preload("Locations.Address.State").
		Preload("Operators").
		Preload("Operators.Operator").
		Preload("Registrations").
		Where(ent.TableVehicle{Registration: registration}).
		First(&vhTab)

	if vhTab.ID == uuid.Nil {
		return vehicleResult, nil
	}
	vehicleResult.FleetId = vhTab.FleetId
	vehicleResult.ModelId = vhTab.ModelId
	vehicleResult.Registration = vhTab.Registration
	vehicleResult.StatusId = vhTab.StatusId
	vehicleResult.MakeId = vhTab.MakeId
	vehicleResult.Status = bu.VehicleStatusBO{
		Id:         vhTab.Status.ID,
		StatusType: vhTab.Status.StatusType,
		StatusName: vhTab.Status.StatusName,
	}

	var vehicleLocation []bu.VehicleAddressBO
	for _, loc := range vhTab.Locations {

		vehicleLocation = append(vehicleLocation, bu.VehicleAddressBO{
			Id:        loc.ID,
			AddressId: loc.AddressId,
			VehicleId: loc.VehicleId,
			UpdateAt:  loc.UpdatedAt,
			Address: bu.AddressBO{
				Id:            loc.Address.ID,
				Address:       loc.Address.Address,
				UpdatedAt:     loc.Address.UpdatedAt,
				CountryId:     loc.Address.CountryId,
				Location:      loc.Address.Location,
				AddressTypeId: loc.Address.AddressTypeId,
				StateId:       loc.Address.StateId,
				Suburb:        loc.Address.Suburb,
				Street:        loc.Address.Street,
				State: bu.StateBO{
					Id:        loc.Address.StateId,
					Name:      loc.Address.State.Name,
					CountryId: loc.Address.State.CountryId,
				},
			},
		})
	}
	vehicleResult.Locations = vehicleLocation
	var ops []bu.VehicleOperatorBoundBO
	for _, op := range vhTab.Operators {

		ops = append(ops, bu.VehicleOperatorBoundBO{
			Id:         op.ID,
			OperatorId: op.OperatorId,
			VehicleId:  op.VehicleId,
			Active:     op.Active,
			Operator: &bu.OperatorBO{
				Id:      op.Operator.ID,
				Name:    op.Operator.Name,
				SurName: op.Operator.SurName,
			},
			Vehicle: nil,
		})
	}
	vehicleResult.Operators = ops

	var regis []bu.VehicleTrackRegBO
	for _, reg := range vhTab.Registrations {

		regis = append(regis, bu.VehicleTrackRegBO{
			Id:           reg.ID,
			Active:       reg.Active,
			VehicleId:    reg.VehicleId,
			UpdatedAt:    reg.UpdatedAt,
			Duration:     reg.Duration,
			ExpiredDate:  reg.ExpiredDate,
			RegisterDate: reg.ExpiredDate,
		})
	}
	vehicleResult.Registrations = regis

	return vehicleResult, nil

}

func (v *Vehicle) GetVehiclesByFleetId(fleetId uuid.UUID) ([]bu.VehicleBO, error) {
	var vhTabs []ent.TableVehicle
	var vehicleResults []bu.VehicleBO

	v.Db.Preload("VehicleModel").Preload("VehicleMake").
		Preload("Fleet").
		Preload("Status").
		Preload("Locations").
		Preload("Locations.Address").
		Preload("Locations.Address.State").
		Preload("Operators").
		Preload("Operators.Operator").
		Preload("Registrations").
		Where(ent.TableVehicle{FleetId: fleetId}).
		First(&vhTabs)

	for _, vhTab := range vhTabs {

		vehicleResult := bu.VehicleBO{}
		vehicleResult.FleetId = vhTab.FleetId

		vehicleResult.ModelId = vhTab.ModelId
		vehicleResult.Registration = vhTab.Registration
		vehicleResult.StatusId = vhTab.StatusId
		vehicleResult.MakeId = vhTab.MakeId
		vehicleResult.Status = bu.VehicleStatusBO{
			Id:         vhTab.Status.ID,
			StatusType: vhTab.Status.StatusType,
			StatusName: vhTab.Status.StatusName,
		}

		var vehicleLocation []bu.VehicleAddressBO
		for _, loc := range vhTab.Locations {

			vehicleLocation = append(vehicleLocation, bu.VehicleAddressBO{
				Id:        loc.ID,
				AddressId: loc.AddressId,
				VehicleId: loc.VehicleId,
				UpdateAt:  loc.UpdatedAt,
				Address: bu.AddressBO{
					Id:            loc.Address.ID,
					Address:       loc.Address.Address,
					UpdatedAt:     loc.Address.UpdatedAt,
					CountryId:     loc.Address.CountryId,
					Location:      loc.Address.Location,
					AddressTypeId: loc.Address.AddressTypeId,
					StateId:       loc.Address.StateId,
					Suburb:        loc.Address.Suburb,
					Street:        loc.Address.Street,
					State: bu.StateBO{
						Id:        loc.Address.StateId,
						Name:      loc.Address.State.Name,
						CountryId: loc.Address.State.CountryId,
					},
				},
			})
		}
		vehicleResult.Locations = vehicleLocation
		var ops []bu.VehicleOperatorBoundBO
		for _, op := range vhTab.Operators {

			ops = append(ops, bu.VehicleOperatorBoundBO{
				Id:         op.ID,
				OperatorId: op.OperatorId,
				VehicleId:  op.VehicleId,
				Active:     op.Active,
				Operator: &bu.OperatorBO{
					Id:      op.Operator.ID,
					Name:    op.Operator.Name,
					SurName: op.Operator.SurName,
				},
				Vehicle: nil,
			})
		}
		vehicleResult.Operators = ops

		var regis []bu.VehicleTrackRegBO
		for _, reg := range vhTab.Registrations {

			regis = append(regis, bu.VehicleTrackRegBO{
				Id:           reg.ID,
				Active:       reg.Active,
				VehicleId:    reg.VehicleId,
				UpdatedAt:    reg.UpdatedAt,
				Duration:     reg.Duration,
				ExpiredDate:  reg.ExpiredDate,
				RegisterDate: reg.ExpiredDate,
			})
		}
		vehicleResult.Registrations = regis
		vehicleResults = append(vehicleResults, vehicleResult)

	}

	return vehicleResults, nil
}
