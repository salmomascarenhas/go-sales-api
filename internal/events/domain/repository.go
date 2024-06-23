package domain

type EventRepository interface {
	ListsEvents() ([]*Event, error)
	FindByEventID(id string) (*Event, error)
	FindSpotsByEventID(eventID string) ([]*Spot, error)
	FindSpotByName(eventID, name string) (*Spot, error)
	CreateEvemt(event *Event) error
	CreateSpot(spot *Spot) error
	CreateTicket(ticket *Ticket) error
	ReserveSpot(spotID string, ticketID string) error
}