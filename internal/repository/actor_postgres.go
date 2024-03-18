package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/Max425/film-library.git/internal/domain"
	"github.com/Max425/film-library.git/internal/repository/store"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"time"
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
	query := `INSERT INTO actor (name, gender, birth_date) VALUES ($1, $2, $3) RETURNING id`
	err := r.db.QueryRowContext(ctx, query, storeActor.Name, storeActor.Gender, storeActor.BirthDate).Scan(&storeActor.ID)
	if err != nil {
		r.logger.Error("Failed to create actor", zap.Error(err))
		return nil, err
	}
	return store.ActorStoreToDomain(storeActor)
}

func (r *ActorRepository) FindActorByID(ctx context.Context, id int) (*domain.Actor, error) {
	storeActor := &store.Actor{}
	query := `SELECT * FROM actor WHERE id = $1`
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
	query := `UPDATE actor SET name = $1, gender = $2, birth_date = $3 WHERE id = $4`
	_, err := r.db.ExecContext(ctx, query, storeActor.Name, storeActor.Gender, storeActor.BirthDate, storeActor.ID)
	if err != nil {
		r.logger.Error("Failed to update actor", zap.Error(err))
		return nil, err
	}
	return actor, nil
}

func (r *ActorRepository) DeleteActor(ctx context.Context, id int) error {
	query := `DELETE FROM actor WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		r.logger.Error("Failed to delete actor", zap.Error(err))
		return err
	}
	return nil
}

func (r *ActorRepository) GetAllActors(ctx context.Context) ([]*domain.Actor, error) {
	query := `
		SELECT a.id, a.name, a.gender, a.birth_date, f.id AS film_id, f.title, f.description, f.release_date, f.rating
		FROM actor AS a
		LEFT JOIN film_actor AS fa ON a.id = fa.actor_id
		LEFT JOIN film AS f ON fa.film_id = f.id
		ORDER BY a.id, f.id
	`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		r.logger.Error("Failed to get all actors with films", zap.Error(err))
		return nil, err
	}
	defer rows.Close()

	var actors []*domain.Actor
	var currentActor *domain.Actor
	for rows.Next() {
		var actorID int
		var actorName, actorGender string
		var actorBirthDate time.Time
		var filmID sql.NullInt64
		var filmTitle, filmDescription sql.NullString
		var filmReleaseDate sql.NullTime
		var filmRating sql.NullFloat64
		if err = rows.Scan(&actorID, &actorName, &actorGender, &actorBirthDate, &filmID, &filmTitle, &filmDescription, &filmReleaseDate, &filmRating); err != nil {
			r.logger.Error("Failed to scan row", zap.Error(err))
			continue
		}

		if currentActor == nil || currentActor.GetId() != actorID {
			currentActor, _ = domain.NewActor(actorID, actorName, actorGender, actorBirthDate, make([]*domain.Film, 0))
			actors = append(actors, currentActor)
		}

		if filmID.Valid {
			film, _ := domain.NewFilm(int(filmID.Int64), filmTitle.String, filmDescription.String, filmReleaseDate.Time, filmRating.Float64, nil)
			currentActor.AddFilm(film)
		}
	}
	if err = rows.Err(); err != nil {
		r.logger.Error("Error while iterating rows", zap.Error(err))
		return nil, err
	}

	return actors, nil
}
