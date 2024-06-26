package reposiroty

import (
	"database/sql"

	"github.com/salmomascarenhas/go-sales-api/internal/events/domain"
)

type mysqlEventRepository struct {
	db *sql.DB
}

func NewMysqlEventRepository(db *sql.DB) (domain.EventRepository, error) {
	return &mysqlEventRepository{
		db: db,
	}, nil
}

func (r *mysqlEventRepository) CreateSpot(spot *domain.Spot) error {
	_, err := r.db.Exec("INSERT INTO spot (id, event_id, name, status, ticket_id) VALUES (?, ?, ?, ?, ?)", spot.ID, spot.EventID, spot.Name, spot.Status, spot.TicketID)
	return err
}

func (r *mysqlEventRepository) ReserveSpot (spotID, ticketID string) error {
	_, err := r.db.Exec("UPDATE spot SET status = ?, ticket_id = ? WHERE id = ?", domain.SpotStatusSold, ticketID, spotID)
	return err
}