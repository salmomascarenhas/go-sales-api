package domain

import (
	"errors"
	"time"
)
var (
	ErrEventNameIsRequired = errors.New("name is required")
	ErrEventDateMustBeInTheFuture = errors.New("event date must be in the future")
	ErrEventCapacityMustBeGreaterThanZero = errors.New("event capacity must be greater than zero")
	ErrPriceMustBeGreaterThanZero = errors.New("price must be greater than zero")
)


type Rating string

const (
	RatingLivre Rating = "L"
	Rating10 Rating = "L10"
	Rating12 Rating = "L12"
	Rating14 Rating = "L14"
	Rating16 Rating = "L16"
	Rating18 Rating = "L18"
)

type Event struct {
	ID string
	Name string
	Location string
	Organization string
	Rating Rating
	Date time.Time
	ImageURL string
	Capacity int
	Price float64
	PartnerID string
	Spots []Spot
	Tickets []Ticket
}

func ( e *Event) Validate() error {
	if e.Name == "" {
		return ErrEventNameIsRequired
	}
	if e.Date.Before(time.Now()) {
		return ErrEventDateMustBeInTheFuture
	}
	if e.Capacity <= 0 {
		return ErrEventCapacityMustBeGreaterThanZero
	}
	if e.Price < 0 {
		return ErrPriceMustBeGreaterThanZero
	}
	return nil
}