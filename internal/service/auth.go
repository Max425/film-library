package service

import (
	"context"
	"crypto/sha1"
	"fmt"
	"github.com/Max425/film-library.git/internal/common/constants"
	"github.com/Max425/film-library.git/internal/domain"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"time"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *domain.User) (int, error)
	GetUser(ctx context.Context, mail string) (*domain.User, error)
}

type StoreRepository interface {
	SetSession(ctx context.Context, session string, role int, expire time.Duration) error
	DeleteSession(ctx context.Context, session string) error
	GetSession(ctx context.Context, session string) (int, error)
}

type AuthService struct {
	log       *zap.Logger
	userRepo  UserRepository
	storeRepo StoreRepository
}

func NewAuthService(log *zap.Logger, userRepo UserRepository, storeRepo StoreRepository) *AuthService {
	return &AuthService{log: log, userRepo: userRepo, storeRepo: storeRepo}
}

func (s *AuthService) CreateUser(ctx context.Context, user *domain.User) (int, error) {
	user.SetSalt(GenerateUuid())
	user.SetPassword(GeneratePasswordHash(user.Password(), user.Salt()))
	id, err := s.userRepo.CreateUser(ctx, user)
	return id, err
}

func (s *AuthService) GetUser(ctx context.Context, mail, password string) (*domain.User, error) {
	user, err := s.userRepo.GetUser(ctx, mail)
	if err != nil {
		return user, err
	}
	if GeneratePasswordHash(password, user.Salt()) != user.Password() {
		return user, domain.ErrInvalidPassword
	}
	user.Sanitize()
	return user, nil
}

func (s *AuthService) GenerateCookie(ctx context.Context, role int) (string, error) {
	SID := GenerateUuid()
	if err := s.storeRepo.SetSession(ctx, SID, role, constants.CookieExpire); err != nil {
		return "", err
	}

	return SID, nil
}

func (s *AuthService) DeleteCookie(ctx context.Context, session string) error {
	return s.storeRepo.DeleteSession(ctx, session)
}

func (s *AuthService) GetSessionValue(ctx context.Context, session string) (int, error) {
	id, err := s.storeRepo.GetSession(ctx, session)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func GeneratePasswordHash(password, salt string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func GenerateUuid() string {
	return uuid.NewString()
}
