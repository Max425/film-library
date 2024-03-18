package dto

import (
	"github.com/Max425/film-library.git/internal/domain"
	"time"
)

type Actor struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Gender    string    `json:"gender"`
	BirthDate time.Time `json:"birth_date"`
	Films     []*Film   `json:"films" swaggerignore:"true"`
}

func ActorDtoToDomain(dtoActor *Actor) (*domain.Actor, error) {
	return domain.NewActor(dtoActor.ID, dtoActor.Name, dtoActor.Gender, dtoActor.BirthDate, nil)
}

func ActorDomainToDto(domainActor *domain.Actor) *Actor {
	filmDTOs := make([]*Film, len(domainActor.GetFilms()))
	for i, film := range domainActor.GetFilms() {
		filmDTOs[i] = FilmDomainToDto(film)
	}

	return &Actor{
		ID:        domainActor.GetId(),
		Name:      domainActor.GetName(),
		Gender:    domainActor.GetGender(),
		BirthDate: domainActor.GetBirthDate(),
		Films:     filmDTOs,
	}
}
