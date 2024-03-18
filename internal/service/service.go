package service

import (
	"go.uber.org/zap"
)

type Repository interface {
	ActorRepository
	FilmRepository
	UserRepository
	StoreRepository
}

type Service struct {
	ActorService
	FilmService
	AuthService
}

func NewService(repo Repository, log *zap.Logger) *Service {
	return &Service{
		*NewActorService(repo, log),
		*NewFilmService(repo, log),
		*NewAuthService(log, repo, repo),
	}
}
