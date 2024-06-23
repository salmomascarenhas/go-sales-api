package domain

import "errors"

type TicketType string

var (
	ErrInvalidTicketType = errors.New("invalid ticket type")
)

const (
	TicketTypeHalf TicketType = "half"
	TicketTypeFull TicketType = "full"
)

type Ticket struct {
	ID string
	EventID string
	Spot 	*Spot
	TicketType TicketType
	Price float64
}

func IsValidTicketType(ticketType TicketType) bool {
	return ticketType == TicketTypeHalf || ticketType == TicketTypeFull
}

func (t *Ticket) CalculatePrice() {
	if t.TicketType == TicketTypeHalf {
		t.Price /= 2  
	}
}

func (t *Ticket) Validate() error {
	if t.Price <= 0 {
		return ErrPriceMustBeGreaterThanZero
	}
	if !IsValidTicketType(t.TicketType) {
		return ErrInvalidTicketType
	}
	return nil
}