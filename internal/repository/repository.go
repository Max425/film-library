package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type Repository struct {
	FilmRepository
	ActorRepository
	UserRepository
	RedisStore
}

func NewRepository(db *sqlx.DB, logger *zap.Logger, client *redis.Client) *Repository {
	return &Repository{
		*NewFilmRepository(db, logger),
		*NewActorRepository(db, logger),
		*NewUserRepository(db, logger),
		*NewRedisStore(client),
	}
}
