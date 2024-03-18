package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/Max425/film-library.git/internal/domain"
	"github.com/Max425/film-library.git/internal/repository/store"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type UserRepository struct {
	db     *sqlx.DB
	logger *zap.Logger
}

func NewUserRepository(db *sqlx.DB, logger *zap.Logger) *UserRepository {
	return &UserRepository{
		db:     db,
		logger: logger,
	}
}

func (r *UserRepository) CreateUser(ctx context.Context, user *domain.User) (int, error) {
	storeUser := store.UserDomainToStore(user)
	query := `INSERT INTO users (name, mail, password_hash, salt, role) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err := r.db.QueryRowContext(ctx, query, storeUser.Name, storeUser.Mail, storeUser.PasswordHash, storeUser.Salt, storeUser.Role).Scan(&storeUser.ID)
	if err != nil {
		r.logger.Error("Failed to create user", zap.Error(err))
		return 0, err
	}
	return storeUser.ID, nil
}

func (r *UserRepository) GetUser(ctx context.Context, mail string) (*domain.User, error) {
	storeUser := &store.User{}
	query := `SELECT * FROM users WHERE mail = $1`
	err := r.db.GetContext(ctx, storeUser, query, mail)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrNotFound
		}
		r.logger.Error("Failed to find user by mail", zap.Error(err))
		return nil, err
	}
	return store.UserStoreToDomain(storeUser)
}
