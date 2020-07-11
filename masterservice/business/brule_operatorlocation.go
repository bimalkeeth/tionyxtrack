package business

import (
	"errors"
	"github.com/jinzhu/gorm"
	bu "tionyxtrack/masterservice/businesscontracts"
	en "tionyxtrack/masterservice/entities"
)

type IOperatorLocation interface {
	CreateOperatorLocation(bo bu.OperatorLocationBO) (uint, error)
	UpdateOperatorLocation(bo bu.OperatorLocationBO) (bool, error)
	DeleteOperatorLocation(id uint) (bool, error)
	GetOperatorLocationByOperator(id uint) ([]bu.OperatorLocationBO, error)
}

type OperatorLocation struct {
	Db *gorm.DB
}

func NewOprLoc(db *gorm.DB) *OperatorLocation {
	return &OperatorLocation{Db: db}
}

//----------------------------------------------------
//Create operator location
//----------------------------------------------------
func (l *OperatorLocation) CreateOperatorLocation(bo bu.OperatorLocationBO) (uint, error) {
	oprLoc := en.TableVehicleOperatorLocation{}

	oprLoc.AddressId = bo.AddressId
	oprLoc.OperatorId = bo.OperatorId
	oprLoc.Primary = bo.Primary

	l.Db.Create(&oprLoc)
	return oprLoc.ID, nil
}

//----------------------------------------------------
//Update operator location
//----------------------------------------------------
func (l *OperatorLocation) UpdateOperatorLocation(bo bu.OperatorLocationBO) (bool, error) {
	if bo.Primary {
		setOLPrimaryOff(l)
	}
	oprLoc := en.TableVehicleOperatorLocation{}
	l.Db.First(&oprLoc, bo.Id)
	if oprLoc.ID == 0 {
		return false, errors.New("operator location not found")
	}
	oprLoc.OperatorId = bo.OperatorId
	oprLoc.AddressId = bo.AddressId
	oprLoc.Primary = bo.Primary

	l.Db.Save(&oprLoc)
	return true, nil
}

func setOLPrimaryOff(f *OperatorLocation) {
	oprLoc := &en.TableVehicleOperatorLocation{}
	f.Db.Where("primary = ?", true).First(&oprLoc)
	if oprLoc.ID > 0 {
		oprLoc.Primary = false
		f.Db.Save(&oprLoc)
	}
}

//----------------------------------------------------
//Delete operator location
//----------------------------------------------------
func (l *OperatorLocation) DeleteOperatorLocation(id uint) (bool, error) {
	oprLoc := en.TableVehicleOperatorLocation{}
	l.Db.First(&oprLoc, id)
	if oprLoc.ID == 0 {
		return false, errors.New("operator location not found")
	}
	l.Db.Delete(&oprLoc)
	return true, nil
}

//----------------------------------------------------
//Get operator location
//----------------------------------------------------
func (l *OperatorLocation) GetOperatorLocationByOperator(id uint) ([]bu.OperatorLocationBO, error) {
	var oprResults []en.TableVehicleOperatorLocation
	var results []bu.OperatorLocationBO

	l.Db.Preload("Address").Preload("Operator").
		Where(&en.TableVehicleOperatorLocation{OperatorId: id}).Find(&oprResults)

	for _, item := range oprResults {

		results = append(results, bu.OperatorLocationBO{
			Id:         item.ID,
			AddressId:  item.AddressId,
			OperatorId: item.OperatorId,
			Primary:    item.Primary,
			Address: bu.AddressBO{
				Id:            item.AddressId,
				AddressTypeId: item.Address.AddressTypeId,
				UpdatedAt:     item.Address.UpdatedAt,
				CountryId:     item.Address.CountryId,
				Address:       item.Address.Address,
				StateId:       item.Address.StateId,
				Suburb:        item.Address.Suburb,
				Street:        item.Address.Street,
				Location:      item.Address.Location,
			},
		})
	}
	return results, nil
}
