package domain

import (
	"errors"
	"fmt"
)

type SpotService struct{}

var (
	ErrQuantityMustBeGreaterThanZero = errors.New("quantity must be greater than zero")
)

func NewSpotService() *SpotService {
	return &SpotService{}
}

func (s *SpotService) GenerateSpots(e *Event, quantity int) error {
	if quantity <= 0 {
		return  ErrQuantityMustBeGreaterThanZero
	}
	for i := 0; i < quantity; i++ {
		spotName := fmt.Sprintf("%c%d", 'A'+i/10, i%10+1)
		spot, err := NewSpot(e, spotName)
		if err != nil {
			return err
		}
		e.Spots = append(e.Spots, *spot)
	}
	return nil
}