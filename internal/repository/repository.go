package repository

import (
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"time"
)

type Event interface {
	Create(Event *core.Event) (int, error)
	Update(Event *core.Event) error
	Delete(id int) error
	GetEvents(duration time.Duration) ([]core.Event, error)
}

type Repository struct {
	Event
}

func NewRepository(db *sqlx.DB, log *zap.Logger) *Repository {
	return &Repository{
		Event: NewEventRepository(db, log),
	}
}
