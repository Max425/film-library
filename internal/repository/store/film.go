package store

import (
	"github.com/Max425/film-library.git/internal/domain"
	"time"
)

// Film in DB
type Film struct {
	ID          int       `db:"id"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
	ReleaseDate time.Time `db:"release_date"`
	Rating      float64   `db:"rating"`
	Actors      []*Actor  `db:"actors"`
}

func FilmStoreToDomain(storeFilm *Film) (*domain.Film, error) {
	actorDomain := make([]*domain.Actor, len(storeFilm.Actors))
	for i, film := range storeFilm.Actors {
		actorDomain[i], _ = ActorStoreToDomain(film)
	}
	return domain.NewFilm(storeFilm.ID, storeFilm.Title, storeFilm.Description, storeFilm.ReleaseDate, storeFilm.Rating, actorDomain)
}

func FilmDomainToStore(domainFilm *domain.Film) *Film {
	return &Film{
		ID:          domainFilm.GetId(),
		Title:       domainFilm.GetTitle(),
		Description: domainFilm.GetDescription(),
		ReleaseDate: domainFilm.GetReleaseDate(),
		Rating:      domainFilm.GetRating(),
	}
}
