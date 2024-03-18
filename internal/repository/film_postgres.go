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
	return film, nil
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
	var storeFilms []*store.Film
	query := `SELECT * FROM film`
	err := r.db.SelectContext(ctx, &storeFilms, query)
	if err != nil {
		r.logger.Error("Failed to get all films", zap.Error(err))
		return nil, err
	}
	films := make([]*domain.Film, len(storeFilms))
	for i, storeFilm := range storeFilms {
		film, err := store.FilmStoreToDomain(storeFilm)
		if err != nil {
			r.logger.Error(fmt.Sprintf("Failed to convert store film to domain film: %v", err))
			continue
		}
		films[i] = film
	}
	return films, nil
}
