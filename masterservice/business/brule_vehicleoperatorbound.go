package business

import (
	"errors"
	"github.com/jinzhu/gorm"
	bu "tionyxtrack/masterservice/businesscontracts"
	ent "tionyxtrack/masterservice/entities"
)

type IVehicleOperatorBound interface {
	CreateVehicleOpBound(op bu.VehicleOperatorBoundBO) (uint, error)
	UpdateVehicleOpBound(op bu.VehicleOperatorBoundBO) (bool, error)
	DeleteVehicleOpBound(id uint) (bool, error)
}

type VehicleOprBound struct {
	Db *gorm.DB
}

func NewOprBound(db *gorm.DB) *VehicleOprBound {
	return &VehicleOprBound{Db: db}
}

//----------------------------------------------
// Create vehicle operator bound
//----------------------------------------------
func (v *VehicleOprBound) CreateVehicleOpBound(op bu.VehicleOperatorBoundBO) (uint, error) {

	vhOprModel := ent.TableVehicleOperatorBound{
		Active:     op.Active,
		VehicleId:  op.VehicleId,
		OperatorId: op.OperatorId,
	}
	v.Db.Create(&vhOprModel)
	return vhOprModel.ID, nil
}

//----------------------------------------------
// Update vehicle operator bound
//----------------------------------------------
func (v *VehicleOprBound) UpdateVehicleOpBound(op bu.VehicleOperatorBoundBO) (bool, error) {
	vhOprModel := ent.TableVehicleOperatorBound{}
	v.Db.First(&vhOprModel, op.Id)
	if vhOprModel.ID == 0 {
		return false, errors.New("record not found")
	}
	vhOprModel.OperatorId = op.OperatorId
	vhOprModel.VehicleId = op.VehicleId
	v.Db.Save(&vhOprModel)
	return true, nil
}

//----------------------------------------------
// Delete vehicle operator bound
//----------------------------------------------
func (v *VehicleOprBound) DeleteVehicleOpBound(id uint) (bool, error) {
	vhOprModel := ent.TableVehicleOperatorBound{}
	v.Db.First(&vhOprModel, id)
	if vhOprModel.ID == 0 {
		return false, errors.New("record not found")
	}

	v.Db.Delete(&vhOprModel)
	return true, nil
}
