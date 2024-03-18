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

func TestActorRepository_CreateActor(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	logger := zap.NewNop()
	r := NewActorRepository(db, logger)

	storeActor := &store.Actor{
		Name:      "Test Actor",
		Gender:    "male",
		BirthDate: time.Unix(0, 0),
	}

	actor, _ := store.ActorStoreToDomain(storeActor)

	mock.ExpectQuery("INSERT INTO actor").
		WithArgs(storeActor.Name, storeActor.Gender, storeActor.BirthDate).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	result, err := r.CreateActor(context.Background(), actor)
	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestActorRepository_FindActorByID(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	logger := zap.NewNop()
	r := NewActorRepository(db, logger)

	actorID := 1

	storeActor := &store.Actor{
		ID:        actorID,
		Name:      "Test Actor",
		Gender:    "male",
		BirthDate: time.Unix(0, 0),
	}

	mock.ExpectQuery("SELECT (.+) FROM actor").
		WithArgs(actorID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "gender", "birth_date"}).
			AddRow(storeActor.ID, storeActor.Name, storeActor.Gender, storeActor.BirthDate))

	result, err := r.FindActorByID(context.Background(), actorID)
	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestActorRepository_UpdateActor(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	logger := zap.NewNop()
	r := NewActorRepository(db, logger)

	storeActor := &store.Actor{
		Name:      "Test Actor",
		Gender:    "male",
		BirthDate: time.Unix(0, 0),
	}

	actor, _ := domain.NewActor(storeActor.ID, storeActor.Name, storeActor.Gender, storeActor.BirthDate, nil)

	mock.ExpectExec("UPDATE actor").
		WithArgs(storeActor.Name, storeActor.Gender, storeActor.BirthDate, storeActor.ID).
		WillReturnResult(sqlmock.NewResult(0, 1))

	result, err := r.UpdateActor(context.Background(), actor)
	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestActorRepository_DeleteActor(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	logger := zap.NewNop()
	r := NewActorRepository(db, logger)

	actorID := 1

	mock.ExpectExec("DELETE FROM actor").
		WithArgs(actorID).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err = r.DeleteActor(context.Background(), actorID)
	assert.NoError(t, err)
}

func TestActorRepository_GetAllActors(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	logger := zap.NewNop()
	r := NewActorRepository(db, logger)

	query := `
		SELECT a.id, a.name, a.gender, a.birth_date, f.id AS film_id, f.title, f.description, f.release_date, f.rating
		FROM actor AS a
		LEFT JOIN film_actor AS fa ON a.id = fa.actor_id
		LEFT JOIN film AS f ON fa.film_id = f.id
		ORDER BY a.id, f.id
	`

	rows := sqlmock.NewRows([]string{"id", "name", "gender", "birth_date", "film_id", "title", "description", "release_date", "rating"}).
		AddRow(1, "Actor 1", "male", time.Unix(0, 0), 1, "Film 1", "Description 1", time.Unix(0, 0), 7.5).
		AddRow(1, "Actor 1", "male", time.Unix(0, 0), 2, "Film 2", "Description 2", time.Unix(0, 0), 8.0).
		AddRow(2, "Actor 2", "male", time.Unix(0, 0), 3, "Film 3", "Description 3", time.Unix(0, 0), 6.5)

	mock.ExpectQuery(query).WillReturnRows(rows)

	results, err := r.GetAllActors(context.Background())

	assert.NoError(t, err)
	assert.NotNil(t, results)
	assert.Len(t, results, 2)

	assert.Equal(t, 1, results[0].GetId())
	assert.Equal(t, "Actor 1", results[0].GetName())
	assert.Equal(t, "male", results[0].GetGender())

	assert.Len(t, results[0].GetFilms(), 2)
	assert.Equal(t, 1, results[0].GetFilms()[0].GetId())
	assert.Equal(t, "Film 1", results[0].GetFilms()[0].GetTitle())
	assert.Equal(t, "Description 1", results[0].GetFilms()[0].GetDescription())

	assert.NoError(t, mock.ExpectationsWereMet())
}
