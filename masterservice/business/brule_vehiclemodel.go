package business

import (
	"errors"
	"github.com/jinzhu/gorm"
	bu "tionyxtrack/masterservice/businesscontracts"
	ent "tionyxtrack/masterservice/entities"
)


type IVehicleModel interface {
	CreateVehicleModel(bo bu.VehicleModelBO) (uint, error)
	UpdateVehicleModel(bo bu.VehicleModelBO) (bool, error)
	DeleteVehicleModel(id uint) (bool, error)
	GetAllModelByMake(makeid uint) ([]bu.VehicleModelBO, error)
	GetModelById(id uint) (bu.VehicleModelBO, error)
}

type VehicleModel struct {
	Db *gorm.DB
}

func NewVehicleModel(db *gorm.DB) VehicleModel {
	return VehicleModel{Db: db}
}

//------------------------------------------------
//Create vehicle model
//------------------------------------------------
func (m *VehicleModel) CreateVehicleModel(bo bu.VehicleModelBO) (uint, error) {
	vehicleModel := ent.TableVehicleModel{
		ModelName:   bo.ModelName,
		Description: bo.Description,
		MakeId:      bo.MakeId,
	}
	m.Db.Create(&vehicleModel)
	return vehicleModel.ID, nil
}

//------------------------------------------------
//Update vehicle model
//------------------------------------------------
func (m *VehicleModel) UpdateVehicleModel(bo bu.VehicleModelBO) (bool, error) {
	vehicleModel := ent.TableVehicleModel{}
	m.Db.First(&vehicleModel, bo.Id)
	if vehicleModel.ID == 0 {
		return false, errors.New("vehicle model could not be found")
	}
	vehicleModel.MakeId = bo.MakeId
	vehicleModel.Description = bo.Description
	vehicleModel.ModelName = bo.ModelName
	m.Db.Save(&vehicleModel)
	return true, nil
}

//------------------------------------------------
//Delete vehicle model
//------------------------------------------------
func (m *VehicleModel) DeleteVehicleModel(id uint) (bool, error) {
	vehicleModel := ent.TableVehicleModel{}
	m.Db.First(&vehicleModel, id)
	if vehicleModel.ID == 0 {
		return false, errors.New("vehicle model could not be found")
	}
	m.Db.Delete(&vehicleModel)
	return true, nil
}

//------------------------------------------------
//Get vehicle model by Make
//------------------------------------------------
func (m *VehicleModel) GetAllModelByMake(makeid uint) ([]bu.VehicleModelBO, error) {

	var vehicleModels []ent.TableVehicleModel
	var modelResult []bu.VehicleModelBO

	m.Db.Preload("Make").Preload("Make.Country").Where("makeid = ?", makeid).Find(&vehicleModels)
	for _, item := range vehicleModels {
		modelResult = append(modelResult, bu.VehicleModelBO{
			Id:          item.ID,
			ModelName:   item.ModelName,
			Description: item.Description,
			MakeId:      item.MakeId,
			UpdatedAt:   item.UpdatedAt,
			Make: bu.VehicleMakeBO{
				Id:        item.Make.ID,
				Make:      item.Make.Make,
				CountryId: item.Make.CountryId,
				UpdateAt:  item.Make.UpdatedAt,
				Country: bu.CountryBO{
					Id:          item.Make.Country.ID,
					CountryName: item.Make.Country.CountryName,
					RegionId:    item.Make.Country.RegionId,
					UpdatedAt:   item.Make.Country.UpdatedAt,
				},
			},
		})
	}
	return modelResult, nil
}

//------------------------------------------------
//Get vehicle model by Id
//------------------------------------------------
func (m *VehicleModel) GetModelById(id uint) (bu.VehicleModelBO, error) {
	vehicleModel := ent.TableVehicleModel{}
	vehicleModelResult := bu.VehicleModelBO{}
	m.Db.Preload("Make").Preload("Make.Country").First(&vehicleModel, id)
	if vehicleModel.ID == 0 {
		return vehicleModelResult, errors.New("vehicle model could not be found")
	}

	vehicleModelResult = bu.VehicleModelBO{
		Id:          vehicleModel.ID,
		ModelName:   vehicleModel.ModelName,
		Description: vehicleModel.Description,
		MakeId:      vehicleModel.MakeId,
		UpdatedAt:   vehicleModel.UpdatedAt,
		Make: bu.VehicleMakeBO{
			Id:        vehicleModel.Make.ID,
			Make:      vehicleModel.Make.Make,
			CountryId: vehicleModel.Make.CountryId,
			UpdateAt:  vehicleModel.Make.UpdatedAt,
			Country: bu.CountryBO{
				Id:          vehicleModel.Make.Country.ID,
				CountryName: vehicleModel.Make.Country.CountryName,
				RegionId:    vehicleModel.Make.Country.RegionId,
				UpdatedAt:   vehicleModel.Make.Country.UpdatedAt,
			},
		},
	}
	return vehicleModelResult, nil
}
