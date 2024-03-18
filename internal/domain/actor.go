package domain

import (
	"fmt"
	"time"
)

type Actor struct {
	id        int
	name      string
	gender    string
	birthDate time.Time
	films     []*Film
}

// NewActor создает нового актера.
func NewActor(id int, name, gender string, birthDate time.Time, films []*Film) (*Actor, error) {
	if name == "" {
		return nil, fmt.Errorf("%w: name is required", ErrRequired)
	}

	if len(name) > 255 {
		return nil, fmt.Errorf("name length should not exceed 255 characters")
	}

	if gender != "male" && gender != "female" && gender != "other" {
		return nil, fmt.Errorf("invalid gender, must be male/female/other")
	}

	if birthDate.After(time.Now()) {
		return nil, fmt.Errorf("birth date cannot be in the future")
	}

	return &Actor{
		id:        id,
		name:      name,
		gender:    gender,
		birthDate: birthDate,
		films:     films,
	}, nil
}

// GetId возвращает имя актера.
func (a *Actor) GetId() int {
	return a.id
}

// GetName возвращает имя актера.
func (a *Actor) GetName() string {
	return a.name
}

// GetGender возвращает пол актера.
func (a *Actor) GetGender() string {
	return a.gender
}

// GetBirthDate возвращает дату рождения актера.
func (a *Actor) GetBirthDate() time.Time {
	return a.birthDate
}

// GetFilms возвращает фильмы, в которых участвовал актер.
func (a *Actor) GetFilms() []*Film {
	return a.films
}

// AddFilm возвращает фильмы, в которых участвовал актер.
func (a *Actor) AddFilm(film *Film) {
	a.films = append(a.films, film)
}
