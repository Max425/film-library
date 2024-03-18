package store

import (
	"github.com/Max425/film-library.git/internal/domain"
	"time"
)

// Actor in DB
type Actor struct {
	ID        int       `db:"id"`
	Name      string    `db:"name"`
	Gender    string    `db:"gender"`
	BirthDate time.Time `db:"birth_date"`
	Films     []*Film   `db:"films"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func ActorStoreToDomain(storeActor *Actor) (*domain.Actor, error) {
	filmDomain := make([]*domain.Film, len(storeActor.Films))
	for i, film := range storeActor.Films {
		filmDomain[i], _ = FilmStoreToDomain(film)
	}
	return domain.NewActor(storeActor.ID, storeActor.Name, storeActor.Gender, storeActor.BirthDate, filmDomain)
}

func ActorDomainToStore(domainActor *domain.Actor) *Actor {
	return &Actor{
		ID:        domainActor.GetId(),
		Name:      domainActor.GetName(),
		Gender:    domainActor.GetGender(),
		BirthDate: domainActor.GetBirthDate(),
	}
}
