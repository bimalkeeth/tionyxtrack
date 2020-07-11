package business

import (
	"errors"
	"github.com/jinzhu/gorm"
	bu "tionyxtrack/masterservice/businesscontracts"
	 "tionyxtrack/masterservice/entities"
)

//-----------------------------------------------
// Interface for region management
//-----------------------------------------------
type IRegion interface {
	CreateRegion(bo bu.RegionBO) (uint, error)
	UpdateRegion(bo bu.RegionBO) (bool, error)
	DeleteRegion(id uint) (bool, error)
	GetAllRegion() ([]bu.RegionBO, error)
	GetRegionById(id uint) (bu.RegionBO, error)
	GetRegionByName(name string) (bu.RegionBO, error)
}
type Region struct{ Db *gorm.DB }

func NewRegion(db *gorm.DB) *Region { return &Region{Db: db} }

//------------------------------------------------
//Create region for the given data
//------------------------------------------------
func (r *Region) CreateRegion(bo bu.RegionBO) (uint, error) {
	region := &entities.TableRegion{Region: bo.Region, RegionName: bo.RegionName}
	r.Db.Create(region)
	return region.ID, nil
}

//------------------------------------------------
//Update the region given
//------------------------------------------------
func (r *Region) UpdateRegion(bo bu.RegionBO) (bool, error) {

	region := entities.TableRegion{}
	r.Db.First(&region, bo.Id)
	if region.ID == 0 {
		return false, errors.New("no record found for region")
	}
	region.Region = bo.Region
	region.RegionName = bo.Region
	r.Db.Save(&region)
	return true, nil
}

//------------------------------------------------
// Delete region by Id
//------------------------------------------------
func (r *Region) DeleteRegion(id uint) (bool, error) {

	found := entities.TableRegion{}
	r.Db.First(&found, id)
	if found.ID == 0 {
		return false, errors.New("contact type not found")
	}
	r.Db.Delete(&found)
	return true, nil
}

//------------------------------------------------
//Get al region
//------------------------------------------------
func (r *Region) GetAllRegion() ([]bu.RegionBO, error) {
	var regions []entities.TableRegion
	var result []bu.RegionBO

	r.Db.Find(&regions)
	for _, item := range regions {
		result = append(result, bu.RegionBO{Region: item.Region, RegionName: item.RegionName, Id: item.ID})
	}
	return result, nil
}

//-----------------------------------------------
// Get region by Id
//-----------------------------------------------
func (r *Region) GetRegionById(id uint) (bu.RegionBO, error) {
	region := &entities.TableRegion{}
	r.Db.First(&region, id)

	result := bu.RegionBO{}
	if region.ID == 0 {
		return result, errors.New("record not found")
	}
	return bu.RegionBO{Region: region.Region, RegionName: region.RegionName, Id: region.ID}, nil
}

//-----------------------------------------------
// Get region by name
//-----------------------------------------------
func (r *Region) GetRegionByName(name string) (bu.RegionBO, error) {
	region := entities.TableRegion{}
	r.Db.Where(&entities.TableRegion{RegionName: name}).First(&region)
	if region.ID == 0 {
		return bu.RegionBO{}, errors.New("record not found")
	}
	return bu.RegionBO{Region: region.Region, RegionName: region.RegionName, Id: region.ID}, nil
}
