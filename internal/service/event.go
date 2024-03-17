package service

import (
	"go.uber.org/zap"
	"time"
)

type EventService struct {
	RepoEvent repository.Event
	log       *zap.Logger
}

func NewEventService(repoEvent repository.Event, log *zap.Logger) *EventService {
	return &EventService{RepoEvent: repoEvent, log: log}
}

func (es *EventService) CreateEvent(event *core.Event) (int, error) {
	id, err := es.RepoEvent.Create(event)
	if err != nil {
		es.log.Error("Failed to create event", zap.Error(err))
		return 0, err
	}
	return id, nil
}

func (es *EventService) UpdateEvent(event *core.Event) error {
	err := es.RepoEvent.Update(event)
	if err != nil {
		es.log.Error("Failed to update event", zap.Error(err))
		return err
	}
	return nil
}

func (es *EventService) DeleteEvent(id int) error {
	err := es.RepoEvent.Delete(id)
	if err != nil {
		es.log.Error("Failed to delete event", zap.Error(err))
		return err
	}
	return nil
}

func (es *EventService) GetEventsForTime(date time.Duration) ([]core.Event, error) {
	events, err := es.RepoEvent.GetEvents(date)
	if err != nil {
		es.log.Error("Failed to get events for time", zap.Error(err))
		return nil, err
	}
	return events, nil
}
