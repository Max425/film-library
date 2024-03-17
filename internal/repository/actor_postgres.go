package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/Max425/film-library.git/internal/domain"
	"github.com/Max425/film-library.git/internal/repository/store"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type ActorRepository struct {
	db     *sqlx.DB
	logger *zap.Logger
}

func NewActorRepository(db *sqlx.DB, logger *zap.Logger) *ActorRepository {
	return &ActorRepository{
		db:     db,
		logger: logger,
	}
}

func (r *ActorRepository) CreateActor(ctx context.Context, actor *domain.Actor) (*domain.Actor, error) {
	storeActor := store.ActorDomainToStore(actor)
	query := `INSERT INTO actors (name, gender, birth_date) VALUES ($1, $2, $3) RETURNING id`
	err := r.db.QueryRowContext(ctx, query, storeActor.Name, storeActor.Gender, storeActor.BirthDate).Scan(&storeActor.ID)
	if err != nil {
		r.logger.Error("Failed to create actor", zap.Error(err))
		return nil, err
	}
	return actor, nil
}

func (r *ActorRepository) FindActorByID(ctx context.Context, id int) (*domain.Actor, error) {
	storeActor := &store.Actor{}
	query := `SELECT * FROM actors WHERE id = $1`
	err := r.db.GetContext(ctx, storeActor, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrNotFound
		}
		r.logger.Error("Failed to find actor by ID", zap.Error(err))
		return nil, err
	}
	return store.ActorStoreToDomain(storeActor)
}

func (r *ActorRepository) UpdateActor(ctx context.Context, actor *domain.Actor) (*domain.Actor, error) {
	storeActor := store.ActorDomainToStore(actor)
	query := `UPDATE actors SET name = $1, gender = $2, birth_date = $3 WHERE id = $4`
	_, err := r.db.ExecContext(ctx, query, storeActor.Name, storeActor.Gender, storeActor.BirthDate, storeActor.ID)
	if err != nil {
		r.logger.Error("Failed to update actor", zap.Error(err))
		return nil, err
	}
	return actor, nil
}

func (r *ActorRepository) DeleteActor(ctx context.Context, id int) error {
	query := `DELETE FROM actors WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		r.logger.Error("Failed to delete actor", zap.Error(err))
		return err
	}
	return nil
}

func (r *ActorRepository) GetAllActors(ctx context.Context) ([]*domain.Actor, error) {
	var storeActors []*store.Actor
	query := `SELECT * FROM actors`
	err := r.db.SelectContext(ctx, &storeActors, query)
	if err != nil {
		r.logger.Error("Failed to get all actors", zap.Error(err))
		return nil, err
	}
	actors := make([]*domain.Actor, len(storeActors))
	for i, storeActor := range storeActors {
		actor, err := store.ActorStoreToDomain(storeActor)
		if err != nil {
			r.logger.Error(fmt.Sprintf("Failed to convert store actor to domain actor: %v", err))
			continue
		}
		actors[i] = actor
	}
	return actors, nil
}
