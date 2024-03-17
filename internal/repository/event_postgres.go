package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"time"
)

type EventRepository struct {
	db  *sqlx.DB
	log *zap.Logger
}

func NewEventRepository(db *sqlx.DB, log *zap.Logger) *EventRepository {
	return &EventRepository{db: db, log: log}
}

func (r *EventRepository) Create(event *core.Event) (int, error) {
	query := `
        INSERT INTO event (date, user_id, title)
        VALUES ($1, $2, $3)
        RETURNING id
    `
	var id int
	err := r.db.QueryRow(query, event.Date, event.UserID, event.Title).Scan(&id)
	if err != nil {
		return 0, errors.Wrap(err, "failed to create event")
	}
	return id, nil
}

func (r *EventRepository) Update(event *core.Event) error {
	query := `
        UPDATE event
        SET date=$1, user_id=$2, title=$3
        WHERE id=$4
    `
	_, err := r.db.Exec(query, event.Date, event.UserID, event.Title, event.ID)
	if err != nil {
		return errors.Wrap(err, "failed to update event")
	}
	return nil
}

func (r *EventRepository) Delete(id int) error {
	query := "DELETE FROM event WHERE id=$1"
	_, err := r.db.Exec(query, id)
	if err != nil {
		return errors.Wrap(err, "failed to delete event")
	}
	return nil
}

func (r *EventRepository) GetEvents(duration time.Duration) ([]core.Event, error) {
	startTime := time.Now()
	endTime := startTime.Add(duration)
	query := "SELECT * FROM event WHERE date BETWEEN $1 AND $2"
	var events []core.Event
	err := r.db.Select(&events, query, startTime, endTime)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get events")
	}
	return events, nil
}
