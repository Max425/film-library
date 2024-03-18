package domain

import (
	"fmt"
	"time"
)

type Film struct {
	id          int
	title       string
	description string
	releaseDate time.Time
	rating      float64
	actors      []*Actor
}

// NewFilm creates a new film.
func NewFilm(id int, title, description string, releaseDate time.Time, rating float64, actors []*Actor) (*Film, error) {
	if title == "" {
		return nil, fmt.Errorf("%w: title is required", ErrRequired)
	}

	if len(title) > 150 {
		return nil, fmt.Errorf("title length should not exceed 150 characters")
	}

	if len(description) > 1000 {
		return nil, fmt.Errorf("description length should not exceed 1000 characters")
	}

	if rating < 0 || rating > 10 {
		return nil, fmt.Errorf("%w: rating should be between 0 and 10", ErrRequired)
	}

	if releaseDate.After(time.Now()) {
		return nil, fmt.Errorf("release date cannot be in the future")
	}

	return &Film{
		id:          id,
		title:       title,
		description: description,
		releaseDate: releaseDate,
		rating:      rating,
		actors:      actors,
	}, nil
}

// GetId returns the id of the film.
func (f *Film) GetId() int {
	return f.id
}

// GetTitle returns the title of the film.
func (f *Film) GetTitle() string {
	return f.title
}

// GetDescription returns the description of the film.
func (f *Film) GetDescription() string {
	return f.description
}

// GetReleaseDate returns the release date of the film.
func (f *Film) GetReleaseDate() time.Time {
	return f.releaseDate
}

// GetRating returns the rating of the film.
func (f *Film) GetRating() float64 {
	return f.rating
}

// GetActors returns the actors associated with the film.
func (f *Film) GetActors() []*Actor {
	return f.actors
}

// AddActor returns the actors associated with the film.
func (f *Film) AddActor(actor *Actor) {
	f.actors = append(f.actors, actor)
}
