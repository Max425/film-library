package service

import (
	"go.uber.org/zap"
	"time"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Event interface {
	CreateEvent(event *core.Event) (int, error)
	UpdateEvent(event *core.Event) error
	DeleteEvent(id int) error
	GetEventsForTime(date time.Duration) ([]core.Event, error)
}

type Service struct {
	Event
}

func NewService(repo *repository.Repository, log *zap.Logger) *Service {
	return &Service{
		Event: NewEventService(repo.Event, log),
	}
}
