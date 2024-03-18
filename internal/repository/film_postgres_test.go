package repository

import (
	"context"
	"github.com/Max425/film-library.git/internal/repository/store"
	"github.com/zhashkevych/go-sqlxmock"
	"testing"
	"time"

	"github.com/Max425/film-library.git/internal/domain"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestFilmRepository_CreateFilm(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	logger := zap.NewNop()
	r := NewFilmRepository(db, logger)

	storeFilm := &store.Film{
		Title:       "Test Film",
		Description: "Test Description",
		ReleaseDate: time.Unix(0, 0),
		Rating:      4.5,
		Actors:      nil,
	}

	film, _ := domain.NewFilm(storeFilm.ID, storeFilm.Title, storeFilm.Description, storeFilm.ReleaseDate,
		storeFilm.Rating, nil)

	mock.ExpectQuery("INSERT INTO film").
		WithArgs(storeFilm.Title, storeFilm.Description, storeFilm.ReleaseDate, storeFilm.Rating).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	result, err := r.CreateFilm(context.Background(), film)
	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestFilmRepository_FindFilmByID(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	logger := zap.NewNop()
	r := NewFilmRepository(db, logger)

	filmID := 1

	storeFilm := &store.Film{
		ID:          filmID,
		Title:       "Test Film",
		Description: "Test Description",
		ReleaseDate: time.Unix(0, 0),
		Rating:      4.5,
	}

	mock.ExpectQuery("SELECT (.+) FROM film").
		WithArgs(filmID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "title", "description", "release_date", "rating"}).
			AddRow(storeFilm.ID, storeFilm.Title, storeFilm.Description, storeFilm.ReleaseDate, storeFilm.Rating))

	result, err := r.FindFilmByID(context.Background(), filmID)
	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestFilmRepository_UpdateFilm(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	logger := zap.NewNop()
	r := NewFilmRepository(db, logger)

	storeFilm := &store.Film{
		Title:       "Test Film",
		Description: "Test Description",
		ReleaseDate: time.Unix(0, 0),
		Rating:      4.5,
		Actors:      nil,
	}

	film, _ := domain.NewFilm(storeFilm.ID, storeFilm.Title, storeFilm.Description, storeFilm.ReleaseDate,
		storeFilm.Rating, nil)

	mock.ExpectExec("UPDATE film").
		WithArgs(storeFilm.Title, storeFilm.Description, storeFilm.ReleaseDate, storeFilm.Rating, storeFilm.ID).
		WillReturnResult(sqlmock.NewResult(0, 1))

	result, err := r.UpdateFilm(context.Background(), film)
	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestFilmRepository_DeleteFilm(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	logger := zap.NewNop()
	r := NewFilmRepository(db, logger)

	filmID := 1

	mock.ExpectExec("DELETE FROM film").
		WithArgs(filmID).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err = r.DeleteFilm(context.Background(), filmID)
	assert.NoError(t, err)
}

func TestFilmRepository_SearchFilms(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	logger := zap.NewNop()
	r := NewFilmRepository(db, logger)

	query := `
		SELECT DISTINCT f.id, f.title, f.description, f.release_date, f.rating
		FROM film AS f
		LEFT JOIN film_actor AS fa ON f.id = fa.film_id
		LEFT JOIN actor AS a ON fa.actor_id = a.id
		WHERE f.title ILIKE '%' || $1 || '%' OR a.name ILIKE '%' || $1 || '%'
		ORDER BY f.rating DESC
	`

	rows := sqlmock.NewRows([]string{"id", "title", "description", "release_date", "rating"}).
		AddRow(1, "Film 1", "Description 1", time.Unix(0, 0), 7.5).
		AddRow(2, "Film 2", "Description 2", time.Unix(0, 0), 8.0)

	mock.ExpectQuery(query).WillReturnRows(rows)

	results, err := r.SearchFilms(context.Background(), "test")

	assert.NoError(t, err)
	assert.NotNil(t, results)
	assert.Len(t, results, 2)

	assert.Equal(t, 1, results[0].GetId())
	assert.Equal(t, "Film 1", results[0].GetTitle())
	assert.Equal(t, "Description 1", results[0].GetDescription())

	assert.Equal(t, 2, results[1].GetId())
	assert.Equal(t, "Film 2", results[1].GetTitle())
	assert.Equal(t, "Description 2", results[1].GetDescription())

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestFilmRepository_UpdateFilmActors(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	logger := zap.NewNop()
	r := NewFilmRepository(db, logger)

	filmID := 1
	actorIDs := []int{1, 2, 3}

	deleteQuery := "DELETE FROM film_actor"
	mock.ExpectExec(deleteQuery).WillReturnResult(sqlmock.NewResult(0, 0))

	insertQuery := "INSERT INTO film_actor"
	for _, actorID := range actorIDs {
		mock.ExpectExec(insertQuery).WithArgs(filmID, actorID).WillReturnResult(sqlmock.NewResult(0, 1))
	}

	// Ожидаем вызов метода FindFilmByID
	mock.ExpectQuery("SELECT (.+) FROM film").WithArgs(filmID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "title", "description", "release_date", "rating"}).
			AddRow(filmID, "Test Film", "Test Description", time.Now(), 4.5))

	_, err = r.UpdateFilmActors(context.Background(), filmID, actorIDs)
	assert.NoError(t, err)
}
