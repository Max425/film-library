package repository

import (
	"context"
	"github.com/Max425/film-library.git/internal/repository/store"
	"github.com/zhashkevych/go-sqlxmock"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestUserRepository_CreateUser(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	logger := zap.NewNop()
	r := NewUserRepository(db, logger)

	storeUser := &store.User{
		Name:         "Test User",
		Mail:         "test@example.com",
		PasswordHash: "hashedPassword",
		Salt:         "randomSalt",
		Role:         1,
	}

	mock.ExpectQuery("INSERT INTO users").
		WithArgs(storeUser.Name, storeUser.Mail, storeUser.PasswordHash, storeUser.Salt, storeUser.Role).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	user, _ := store.UserStoreToDomain(storeUser)
	result, err := r.CreateUser(context.Background(), user)
	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestUserRepository_GetUser(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	logger := zap.NewNop()
	r := NewUserRepository(db, logger)

	storeUser := &store.User{
		Name:         "Test User",
		Mail:         "test@example.com",
		PasswordHash: "hashedPassword",
		Salt:         "randomSalt",
		Role:         1,
	}

	mock.ExpectQuery("SELECT (.+) FROM users").
		WithArgs(storeUser.Mail).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "mail", "password_hash", "salt", "role"}).
			AddRow(storeUser.ID, storeUser.Name, storeUser.Mail, storeUser.PasswordHash, storeUser.Salt, storeUser.Role))

	result, err := r.GetUser(context.Background(), storeUser.Mail)
	assert.NoError(t, err)
	assert.NotNil(t, result)
}
