package service

import (
	"go.uber.org/zap"
)

type Repository interface {
	ActorRepository
	FilmRepository
}

type Service struct {
	ActorService
	FilmService
}

func NewService(repo Repository, log *zap.Logger) *Service {
	return &Service{
		*NewActorService(repo, log),
		*NewFilmService(repo, log),
	}
}
