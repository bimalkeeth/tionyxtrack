package business

import (
	"errors"
	"github.com/jinzhu/gorm"
	bu "tionyxtrack/masterservice/businesscontracts"
	ent "tionyxtrack/masterservice/entities"
)

type IVehicleHistory interface {
	CreateVehicleHistory(history bu.VehicleHistoryBO) (uint, error)
	UpdateVehicleHistory(history bu.VehicleHistoryBO) (bool, error)
	DeleteVehicleHistory(id uint) (bool, error)
	GetVehicleHistoryByVehicleId(vehicleId uint) ([]bu.VehicleHistoryBO, error)
}

type VehicleHistory struct {
	Db *gorm.DB
}

func NewVehicleHistory(db *gorm.DB) VehicleHistory {
	return VehicleHistory{Db: db}
}

//------------------------------------------------------
//Create vehicle history
//------------------------------------------------------
func (h *VehicleHistory) CreateVehicleHistory(history bu.VehicleHistoryBO) (uint, error) {
	vehicleHistory := ent.TableVehicleHistory{
		ChangeDate:   history.ChangeDate,
		Description:  history.Description,
		FromStatusId: history.FromStatusId,
		ToStatusId:   history.ToStatusId,
		OfficerName:  history.OfficerName,
		VehicleId:    history.VehicleId,
	}
	h.Db.Create(&vehicleHistory)
	return vehicleHistory.ID, nil
}

//------------------------------------------------------
//Update vehicle history
//------------------------------------------------------
func (h *VehicleHistory) UpdateVehicleHistory(history bu.VehicleHistoryBO) (bool, error) {

	if history.Id == 0 {
		return false, errors.New("vehicle history id not defined")
	}

	vehicleHistory := ent.TableVehicleHistory{}
	h.Db.First(&vehicleHistory, history.Id)
	if vehicleHistory.ID == 0 {
		return false, errors.New("vehicle history not fond")
	}
	vehicleHistory.ToStatusId = history.ToStatusId
	vehicleHistory.VehicleId = history.VehicleId
	vehicleHistory.OfficerName = history.OfficerName
	vehicleHistory.FromStatusId = history.FromStatusId
	vehicleHistory.Description = history.Description
	vehicleHistory.ChangeDate = history.ChangeDate
	h.Db.Save(&vehicleHistory)
	return true, nil
}

//------------------------------------------------------
//Delete vehicle history
//------------------------------------------------------
func (h *VehicleHistory) DeleteVehicleHistory(id uint) (bool, error) {
	vehicleHistory := ent.TableVehicleHistory{}
	h.Db.First(&vehicleHistory, id)
	if vehicleHistory.ID == 0 {
		return false, errors.New("vehicle history not fond")
	}
	h.Db.Delete(&vehicleHistory)
	return true, nil
}

//------------------------------------------------------
//Get vehicle history by vehicle id
//------------------------------------------------------
func (h *VehicleHistory) GetVehicleHistoryByVehicleId(vehicleId uint) ([]bu.VehicleHistoryBO, error) {
	var vehicleHistories []ent.TableVehicleHistory
	var historyResult []bu.VehicleHistoryBO

	h.Db.Preload("FromStatus").
		Preload("ToStatus").Where("vehicleid = ?", vehicleId).Find(&vehicleHistories)

	for _, item := range vehicleHistories {
		historyResult = append(historyResult, bu.VehicleHistoryBO{
			ChangeDate:   item.ChangeDate,
			Description:  item.Description,
			FromStatusId: item.FromStatusId,
			OfficerName:  item.OfficerName,
			VehicleId:    item.VehicleId,
			ToStatusId:   item.ToStatusId,
			Id:           item.ID,
			FromStatus: bu.VehicleStatusBO{
				Id:         item.FromStatus.ID,
				StatusType: item.FromStatus.StatusType,
				StatusName: item.FromStatus.StatusName,
			},
			ToStatus: bu.VehicleStatusBO{
				Id:         item.ToStatus.ID,
				StatusType: item.ToStatus.StatusType,
				StatusName: item.ToStatus.StatusName,
			},
		})
	}
	return historyResult, nil
}
