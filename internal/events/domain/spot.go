package domain

import (
	"errors"

	"github.com/google/uuid"
)

type SpotStatus string 

var (
	ErrSpotNameIsRequired = errors.New("spot name is required")
	ErrSpotNotFound = errors.New("spot not found")
	ErrSpotAlreadyReserved = errors.New("spot already reserved")
	ErrSpotNameMustBeAtLeast2CharactersLong = errors.New("spot name must be at least 2 characters long")
	ErrSpotNameMustStartWithACapitalLetter = errors.New("spot name must start with a capital letter")
	ErrSpotNameSecondCharacterMustBeANumber = errors.New("spot name second character must be a number")
)

const (
	SpotStatusAvailable SpotStatus = "available"
	SpotStatusSold SpotStatus = "sold"
)

type Spot struct {
	ID string
	EventID string
	Name string
	Status SpotStatus
	TicketID string
}

func NewSpot(e *Event, name string) (*Spot, error) {
	s := &Spot{
		ID: uuid.New().String(),
		EventID: e.ID,
		Name: name,
		Status: SpotStatusAvailable,
	}
	if err:= s.Validate(); err != nil {
		return nil, err
	}
	return s, nil
}

func (s Spot) Validate() error {
	if len(s.Name) == 0 {
		return ErrSpotNameIsRequired
	}
	if len(s.Name) < 2 {
		return ErrSpotNameMustBeAtLeast2CharactersLong
	}
	if s.Name[0] < 'A' || s.Name[0] > 'Z' {
		return ErrSpotNameMustStartWithACapitalLetter
	}
	if s.Name[1] < '0' || s.Name[1] > '9' {
		return ErrSpotNameSecondCharacterMustBeANumber
	}
	return nil
}