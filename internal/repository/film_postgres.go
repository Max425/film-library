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

type FilmRepository struct {
	db     *sqlx.DB
	logger *zap.Logger
}

func NewFilmRepository(db *sqlx.DB, logger *zap.Logger) *FilmRepository {
	return &FilmRepository{
		db:     db,
		logger: logger,
	}
}

func (r *FilmRepository) CreateFilm(ctx context.Context, film *domain.Film) (*domain.Film, error) {
	storeFilm := store.FilmDomainToStore(film)
	query := `INSERT INTO film (title, description, release_date, rating) VALUES ($1, $2, $3, $4) RETURNING id`
	err := r.db.QueryRowContext(ctx, query, storeFilm.Title, storeFilm.Description, storeFilm.ReleaseDate, storeFilm.Rating).Scan(&storeFilm.ID)
	if err != nil {
		r.logger.Error("Failed to create film", zap.Error(err))
		return nil, err
	}
	return store.FilmStoreToDomain(storeFilm)
}

func (r *FilmRepository) FindFilmByID(ctx context.Context, id int) (*domain.Film, error) {
	storeFilm := &store.Film{}
	query := `SELECT * FROM film WHERE id = $1`
	err := r.db.GetContext(ctx, storeFilm, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrNotFound
		}
		r.logger.Error("Failed to find film by ID", zap.Error(err))
		return nil, err
	}
	return store.FilmStoreToDomain(storeFilm)
}

func (r *FilmRepository) UpdateFilm(ctx context.Context, film *domain.Film) (*domain.Film, error) {
	storeFilm := store.FilmDomainToStore(film)
	query := `UPDATE film SET title = $1, description = $2, release_date = $3, rating = $4 WHERE id = $5`
	_, err := r.db.ExecContext(ctx, query, storeFilm.Title, storeFilm.Description, storeFilm.ReleaseDate, storeFilm.Rating, storeFilm.ID)
	if err != nil {
		r.logger.Error("Failed to update film", zap.Error(err))
		return nil, err
	}
	return film, nil
}

func (r *FilmRepository) UpdateFilmActors(ctx context.Context, id int, actorsID []int) (*domain.Film, error) {
	deleteQuery := `DELETE FROM film_actor WHERE film_id = $1`
	_, err := r.db.ExecContext(ctx, deleteQuery, id)
	if err != nil {
		r.logger.Error("Failed to delete film actors", zap.Error(err))
		return nil, err
	}

	for _, actorID := range actorsID {
		insertQuery := `INSERT INTO film_actor (film_id, actor_id) VALUES ($1, $2)`
		_, err = r.db.ExecContext(ctx, insertQuery, id, actorID)
		if err != nil {
			r.logger.Error("Failed to insert film actor", zap.Error(err))
		}
	}

	return r.FindFilmByID(ctx, id)
}

func (r *FilmRepository) DeleteFilm(ctx context.Context, id int) error {
	query := `DELETE FROM film WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		r.logger.Error("Failed to delete film", zap.Error(err))
		return err
	}
	return nil
}

func (r *FilmRepository) GetAllFilms(ctx context.Context) ([]*domain.Film, error) {
	query := `
		SELECT f.id, f.title, f.description, f.release_date, f.rating,
			   a.id AS actor_id, COALESCE(a.name, ''), COALESCE(a.gender, ''), COALESCE(a.birth_date, '0001-01-01')
		FROM film AS f
		LEFT JOIN film_actor AS fa ON f.id = fa.film_id
		LEFT JOIN actor AS a ON fa.actor_id = a.id
		ORDER BY f.rating DESC
	`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		r.logger.Error("Failed to get all films with actors", zap.Error(err))
		return nil, err
	}
	defer rows.Close()

	var films []*domain.Film
	var currentFilm *domain.Film
	for rows.Next() {
		var filmID int
		var filmTitle, filmDescription string
		var filmReleaseDate time.Time
		var filmRating float64
		var actorID sql.NullInt64
		var actorName, actorGender sql.NullString
		var actorBirthDate time.Time

		if err := rows.Scan(&filmID, &filmTitle, &filmDescription, &filmReleaseDate, &filmRating, &actorID, &actorName, &actorGender, &actorBirthDate); err != nil {
			r.logger.Error("Failed to scan row", zap.Error(err))
			continue
		}

		if currentFilm == nil || currentFilm.GetId() != filmID {
			currentFilm, _ = domain.NewFilm(filmID, filmTitle, filmDescription, filmReleaseDate, filmRating, nil)
			films = append(films, currentFilm)
		}

		if actorID.Valid {
			actor, _ := domain.NewActor(int(actorID.Int64), actorName.String, actorGender.String, actorBirthDate, nil)
			currentFilm.AddActor(actor)
		}
	}
	if err := rows.Err(); err != nil {
		r.logger.Error("Error while iterating rows", zap.Error(err))
		return nil, err
	}

	return films, nil
}

func (r *FilmRepository) SearchFilms(ctx context.Context, fragment string) ([]*domain.Film, error) {
	query := `
		SELECT DISTINCT f.id, f.title, f.description, f.release_date, f.rating
		FROM film AS f
		LEFT JOIN film_actor AS fa ON f.id = fa.film_id
		LEFT JOIN actor AS a ON fa.actor_id = a.id
		WHERE f.title ILIKE '%' || $1 || '%' OR a.name ILIKE '%' || $1 || '%'
		ORDER BY f.rating DESC
	`
	rows, err := r.db.QueryContext(ctx, query, fragment)
	if err != nil {
		r.logger.Error("Failed to search films", zap.Error(err))
		return nil, err
	}
	defer rows.Close()

	var films []*domain.Film
	for rows.Next() {
		var filmID int
		var filmTitle, filmDescription string
		var filmReleaseDate time.Time
		var filmRating float64

		if err := rows.Scan(&filmID, &filmTitle, &filmDescription, &filmReleaseDate, &filmRating); err != nil {
			r.logger.Error("Failed to scan row", zap.Error(err))
			continue
		}

		film, _ := domain.NewFilm(filmID, filmTitle, filmDescription, filmReleaseDate, filmRating, nil)
		films = append(films, film)
	}
	if err := rows.Err(); err != nil {
		r.logger.Error("Error while iterating rows", zap.Error(err))
		return nil, err
	}

	return films, nil
}
