package business

import (
	"errors"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	bu "tionyxtrack/masterservice/businesscontracts"
	"tionyxtrack/masterservice/entities"
)

type IStates interface {
	CreateState(bo bu.StateBO) (uuid.UUID, error)
	UpdateState(bo bu.StateBO) (bool, error)
	DeleteState(id uuid.UUID) (bool, error)
	GetStateById(id uuid.UUID) (bu.StateBO, error)
	GetStateByCountryId(id uuid.UUID) ([]bu.StateBO, error)
	GetStateByName(name string) (bu.StateBO, error)
	GetAll() ([]bu.StateBO, error)
}

type State struct{ Db *gorm.DB }

func NewState(db *gorm.DB) *State { return &State{Db: db} }

func (s *State) CreateState(bo bu.StateBO) (uuid.UUID, error) {
	state := entities.TableState{Name: bo.Name, CountryId: bo.CountryId}
	s.Db.Create(&state)
	return state.ID, nil
}
func (s *State) UpdateState(bo bu.StateBO) (bool, error) {

	state := entities.TableState{}
	s.Db.First(&state, bo.Id)
	if state.ID == uuid.Nil {
		return false, errors.New("state not found")
	}
	state.CountryId = bo.CountryId
	state.Name = bo.Name

	s.Db.Save(&state)
	return true, nil
}
func (s *State) DeleteState(id uuid.UUID) (bool, error) {

	found := entities.TableState{}
	s.Db.First(&found, id)
	if found.ID == uuid.Nil {
		return false, errors.New("contact type not found")
	}
	s.Db.Delete(&found)
	return true, nil
}
func (s *State) GetStateById(id uuid.UUID) (bu.StateBO, error) {
	result := entities.TableState{}
	s.Db.First(&result, id)
	resultBO := bu.StateBO{}
	if result.ID == uuid.Nil {
		return resultBO, errors.New("state not found")
	}
	return resultBO, nil
}
func (s *State) GetStateByCountryId(id uuid.UUID) ([]bu.StateBO, error) {
	var resultsEntities []entities.TableState
	var results []bu.StateBO
	var country entities.TableCountry

	s.Db.Where(&entities.TableState{CountryId: id}).Find(&resultsEntities).Related(&country)

	for _, item := range resultsEntities {
		results = append(results, bu.StateBO{CountryId: item.CountryId, Name: item.Name, Id: item.ID})
	}
	return results, nil
}
func (s *State) GetStateByName(name string) (bu.StateBO, error) {
	state := entities.TableState{}
	s.Db.Where(&entities.TableState{Name: name}).First(&state)
	if state.ID == uuid.Nil {
		return bu.StateBO{}, errors.New("record not found")
	}
	return bu.StateBO{Name: state.Name, CountryId: state.CountryId, Id: state.ID}, nil
}
func (s *State) GetAll() ([]bu.StateBO, error) {

	var states []entities.TableState
	var stateResults []bu.StateBO
	s.Db.Find(&states)
	for _, item := range states {
		stateResults = append(stateResults, bu.StateBO{Name: item.Name, CountryId: item.CountryId, Id: item.ID})
	}
	return stateResults, nil
}
