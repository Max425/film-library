package service

import (
	"context"
	"github.com/Max425/film-library.git/internal/domain"
	"go.uber.org/zap"
)

type FilmRepository interface {
	CreateFilm(ctx context.Context, film *domain.Film) (*domain.Film, error)
	FindFilmByID(ctx context.Context, id int) (*domain.Film, error)
	UpdateFilm(ctx context.Context, film *domain.Film) (*domain.Film, error)
	UpdateFilmActors(ctx context.Context, id int, actorsId []int) (*domain.Film, error)
	DeleteFilm(ctx context.Context, id int) error
	GetAllFilms(ctx context.Context) ([]*domain.Film, error)
	SearchFilms(ctx context.Context, fragment string) ([]*domain.Film, error)
}

type FilmService struct {
	log      *zap.Logger
	filmRepo FilmRepository
}

func NewFilmService(filmRepo FilmRepository, log *zap.Logger) *FilmService {
	return &FilmService{filmRepo: filmRepo, log: log}
}

func (s *FilmService) CreateFilm(ctx context.Context, film *domain.Film) (*domain.Film, error) {
	return s.filmRepo.CreateFilm(ctx, film)
}

func (s *FilmService) GetFilmByID(ctx context.Context, id int) (*domain.Film, error) {
	return s.filmRepo.FindFilmByID(ctx, id)
}

func (s *FilmService) UpdateFilm(ctx context.Context, film *domain.Film) (*domain.Film, error) {
	_, err := s.filmRepo.FindFilmByID(ctx, film.GetId())
	if err != nil {
		return nil, err
	}

	return s.filmRepo.UpdateFilm(ctx, film)
}

func (s *FilmService) UpdateFilmActors(ctx context.Context, id int, actorsId []int) (*domain.Film, error) {
	_, err := s.filmRepo.FindFilmByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return s.filmRepo.UpdateFilmActors(ctx, id, actorsId)
}

func (s *FilmService) DeleteFilm(ctx context.Context, id int) error {
	_, err := s.filmRepo.FindFilmByID(ctx, id)
	if err != nil {
		return err
	}

	return s.filmRepo.DeleteFilm(ctx, id)
}

func (s *FilmService) GetAllFilms(ctx context.Context) ([]*domain.Film, error) {
	return s.filmRepo.GetAllFilms(ctx)
}

func (s *FilmService) SearchFilms(ctx context.Context, fragment string) ([]*domain.Film, error) {
	return s.filmRepo.SearchFilms(ctx, fragment)
}
