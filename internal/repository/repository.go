package repository

import (
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type Repository struct {
	FilmRepository
	ActorRepository
}

func NewRepository(db *sqlx.DB, logger *zap.Logger) *Repository {
	return &Repository{
		*NewFilmRepository(db, logger),
		*NewActorRepository(db, logger),
	}
}
