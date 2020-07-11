package business

import (
	"github.com/jinzhu/gorm"
)

type RuleType int

const (
	CAddress           RuleType = 1
	CContact           RuleType = 2
	CAddressType       RuleType = 3
	CCompany           RuleType = 4
	CCountry           RuleType = 5
	CFleet             RuleType = 6
	CFleetContact      RuleType = 7
	CFleetLocation     RuleType = 8
	COperationContact  RuleType = 9
	COperationLocation RuleType = 10
	CRegion            RuleType = 11
	CState             RuleType = 12
	CVehicle           RuleType = 13
	CVehicleHistory    RuleType = 14
	CVehicleLocation   RuleType = 15
	CVehicleMake       RuleType = 16
	CVehicleModel      RuleType = 17
	CVhOperatorBound   RuleType = 18
	CVhRegistration    RuleType = 19
	CVhStatus          RuleType = 20
	CContactType       RuleType = 21
	CVhOperator        RuleType = 22
)

type RuleFactory struct{ Conn *gorm.DB }

func (f *RuleFactory) New(ruleType RuleType) interface{} {

	switch ruleType {
	case CAddress:
		return NewAddressType(f.Conn)
	case CContact:
		return NewContact(f.Conn)
	case CAddressType:
		return NewAddressType(f.Conn)
	case CCompany:
		return NewCompany(f.Conn)
	case CContactType:
		return NewContactType(f.Conn)
	case CCountry:
		return NewCountry(f.Conn)
	case CFleet:
		return NewFleet(f.Conn)
	case CFleetContact:
		return NewFleetContact(f.Conn)
	case CFleetLocation:
		return NewFleetLocation(f.Conn)
	case COperationContact:
		return NewOperatorContact(f.Conn)
	case COperationLocation:
		return NewOprLoc(f.Conn)
	case CRegion:
		return NewRegion(f.Conn)
	case CState:
		return NewState(f.Conn)
	case CVehicle:
		return NewVehicle(f.Conn)
	case CVehicleHistory:
		return NewVehicleHistory(f.Conn)
	case CVehicleLocation:
		return NewVehicleLocation(f.Conn)
	case CVehicleMake:
		return NewVehicleMake(f.Conn)
	case CVehicleModel:
		return NewVehicleModel(f.Conn)
	case CVhOperatorBound:
		return NewOprBound(f.Conn)
	case CVhRegistration:
		return NewVhReg(f.Conn)
	case CVhStatus:
		return NewVhStatus(f.Conn)
	case CVhOperator:
		return NewOperator(f.Conn)
	default:
		return nil
	}

}
