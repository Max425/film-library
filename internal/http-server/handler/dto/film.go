package dto

import (
	"github.com/Max425/film-library.git/internal/domain"
	"time"
)

type Film struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	ReleaseDate time.Time `json:"release_date"`
	Rating      float64   `json:"rating"`
	Actors      []*Actor  `json:"actors" swaggerignore:"true"`
}

func FilmDtoToDomain(dtoFilm *Film) (*domain.Film, error) {
	return domain.NewFilm(dtoFilm.ID, dtoFilm.Title, dtoFilm.Description, dtoFilm.ReleaseDate, dtoFilm.Rating, nil)
}

func FilmDomainToDto(domainFilm *domain.Film) *Film {
	domainFilm.
		actorsDTOs := make([]*Film, len(domainFilm.GetActors()))
	for i, actor := range domainFilm.GetActors() {
		actorsDTOs[i] = ActorDomainToDto(actor)
	}

	return &Film{
		ID:          domainFilm.GetId(),
		Title:       domainFilm.GetTitle(),
		Description: domainFilm.GetDescription(),
		ReleaseDate: domainFilm.GetReleaseDate(),
		Rating:      domainFilm.GetRating(),
	}
}
