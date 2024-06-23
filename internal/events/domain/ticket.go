package domain

import "errors"

type TicketKind string

var (
	ErrInvalidTicketKind = errors.New("invalid ticket kind")
)

const (
	TicketKindHalf TicketKind = "half"
	TicketKindFull TicketKind = "full"
)

type Ticket struct {
	ID string
	EventID string
	Spot 	*Spot
	TicketKind TicketKind
	Price float64
}

func isValidTicketKind(ticketKind TicketKind) bool {
	return ticketKind == TicketKindHalf || ticketKind == TicketKindFull
}

func (t *Ticket) CalculatePrice() {
	if t.TicketKind == TicketKindHalf {
		t.Price /= 2  
	}
}

func (t *Ticket) Validate() error {
	if t.Price <= 0 {
		return ErrPriceMustBeGreaterThanZero
	}
	if !isValidTicketKind(t.TicketKind) {
		return ErrInvalidTicketKind
	}
	return nil
}