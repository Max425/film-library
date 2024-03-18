package service

import (
	"context"
	"github.com/Max425/film-library.git/internal/domain"
	"go.uber.org/zap"
)

type ActorRepository interface {
	CreateActor(ctx context.Context, actor *domain.Actor) (*domain.Actor, error)
	FindActorByID(ctx context.Context, id int) (*domain.Actor, error)
	UpdateActor(ctx context.Context, actor *domain.Actor) (*domain.Actor, error)
	DeleteActor(ctx context.Context, id int) error
	GetAllActors(ctx context.Context) ([]*domain.Actor, error)
}

type ActorService struct {
	log       *zap.Logger
	actorRepo ActorRepository
}

func NewActorService(actorRepo ActorRepository, log *zap.Logger) *ActorService {
	return &ActorService{actorRepo: actorRepo, log: log}
}

func (s *ActorService) CreateActor(ctx context.Context, actor *domain.Actor) (*domain.Actor, error) {
	return s.actorRepo.CreateActor(ctx, actor)
}

func (s *ActorService) GetActorByID(ctx context.Context, id int) (*domain.Actor, error) {
	return s.actorRepo.FindActorByID(ctx, id)
}

func (s *ActorService) UpdateActor(ctx context.Context, actor *domain.Actor) (*domain.Actor, error) {
	_, err := s.actorRepo.FindActorByID(ctx, actor.GetId())
	if err != nil {
		return nil, err
	}

	return s.actorRepo.UpdateActor(ctx, actor)
}

func (s *ActorService) DeleteActor(ctx context.Context, id int) error {
	_, err := s.actorRepo.FindActorByID(ctx, id)
	if err != nil {
		return err
	}

	return s.actorRepo.DeleteActor(ctx, id)
}

func (s *ActorService) GetAllActors(ctx context.Context) ([]*domain.Actor, error) {
	return s.actorRepo.GetAllActors(ctx)
}
