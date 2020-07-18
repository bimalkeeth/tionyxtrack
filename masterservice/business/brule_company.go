package business

import (
	"errors"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	bu "tionyxtrack/masterservice/businesscontracts"
	ent "tionyxtrack/masterservice/entities"
)

//----------------------------------------------
//interface for company
//----------------------------------------------
type ICompany interface {
	CreateCompany(company bu.CompanyBO) (uuid.UUID, error)
	UpdateCompany(company bu.CompanyBO) (bool, error)
	DeleteCompany(id uuid.UUID) (bool, error)
}

type Company struct{ Db *gorm.DB }

func NewCompany(db *gorm.DB) *Company { return &Company{Db: db} }

//----------------------------------------------
//Create Company
//----------------------------------------------
func (c Company) CreateCompany(company bu.CompanyBO) (uuid.UUID, error) {

	comp := ent.TableCompany{Name: company.Name,
		AddressId:  company.AddressId,
		ContractId: company.ContactId}

	c.Db.Create(&comp)
	return comp.ID, nil
}

//-----------------------------------------------
//Update company
//-----------------------------------------------
func (c Company) UpdateCompany(company bu.CompanyBO) (bool, error) {

	com := &ent.TableCompany{}
	c.Db.First(com, company.Id)
	if com.ID == uuid.Nil {
		return false, errors.New("company can not be found")
	}
	com.ContractId = company.ContactId
	com.AddressId = company.AddressId
	com.Name = company.Name
	c.Db.Save(&com)
	return true, nil
}

//-----------------------------------------------
//Delete company
//-----------------------------------------------
func (c Company) DeleteCompany(id uuid.UUID) (bool, error) {

	com := ent.TableCompany{}
	c.Db.First(&com, id)
	if com.ID == uuid.Nil {
		return false, errors.New("company type not found")
	}
	c.Db.Delete(&com)
	return true, nil
}
