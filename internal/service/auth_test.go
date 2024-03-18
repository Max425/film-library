package service

import (
	"context"
	"errors"
	mock_service "github.com/Max425/film-library.git/mocks/db"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/Max425/film-library.git/internal/common/constants"
	"github.com/Max425/film-library.git/internal/domain"
)

func TestAuthService_CreateUser(t *testing.T) {
	mockUser, _ := domain.NewUser(1, "bob", "test@example.com", "password", "", 0)

	tests := []struct {
		name          string
		mockBehavior  func(r *mock_service.MockUserRepository)
		user          *domain.User
		expectedID    int
		expectedError error
	}{
		{
			name: "Success",
			mockBehavior: func(r *mock_service.MockUserRepository) {
				r.EXPECT().CreateUser(gomock.Any(), mockUser).Return(1, nil)
			},
			user:          mockUser,
			expectedID:    1,
			expectedError: nil,
		},
		{
			name: "Error Creating User",
			mockBehavior: func(r *mock_service.MockUserRepository) {
				r.EXPECT().CreateUser(gomock.Any(), mockUser).Return(0, errors.New("create user error"))
			},
			user:          mockUser,
			expectedID:    0,
			expectedError: errors.New("create user error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mock_service.NewMockUserRepository(ctrl)
			test.mockBehavior(repo)

			authService := NewAuthService(nil, repo, nil)
			id, err := authService.CreateUser(context.Background(), test.user)

			assert.Equal(t, test.expectedID, id)
			assert.Equal(t, test.expectedError, err)
		})
	}
}

func TestAuthService_GetUser(t *testing.T) {
	mockUser, _ := domain.NewUser(1, "bob", "test@example.com", "password", "", 0)

	tests := []struct {
		name          string
		mockBehavior  func(r *mock_service.MockUserRepository)
		mail          string
		password      string
		expectedUser  *domain.User
		expectedError error
	}{
		{
			name: "Error Getting User",
			mockBehavior: func(r *mock_service.MockUserRepository) {
				r.EXPECT().GetUser(gomock.Any(), "test@example.com").Return(nil, errors.New("get user error"))
			},
			mail:          "test@example.com",
			password:      "password",
			expectedUser:  nil,
			expectedError: errors.New("get user error"),
		},
		{
			name: "Invalid Password",
			mockBehavior: func(r *mock_service.MockUserRepository) {
				r.EXPECT().GetUser(gomock.Any(), "test@example.com").Return(mockUser, nil)
			},
			mail:          "test@example.com",
			password:      "wrongpassword",
			expectedUser:  mockUser,
			expectedError: domain.ErrInvalidPassword,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mock_service.NewMockUserRepository(ctrl)
			test.mockBehavior(repo)

			authService := NewAuthService(nil, repo, nil)
			user, err := authService.GetUser(context.Background(), test.mail, test.password)

			assert.Equal(t, test.expectedUser, user)
			assert.Equal(t, test.expectedError, err)
		})
	}
}

func TestAuthService_GenerateCookie(t *testing.T) {
	tests := []struct {
		name          string
		mockBehavior  func(r *mock_service.MockStoreRepository)
		role          int
		expectedSID   string
		expectedError error
	}{
		{
			name: "Error Generating Cookie",
			mockBehavior: func(r *mock_service.MockStoreRepository) {
				r.EXPECT().SetSession(gomock.Any(), gomock.Any(), 1, constants.CookieExpire).Return(errors.New("generate cookie error"))
			},
			role:          1,
			expectedSID:   "",
			expectedError: errors.New("generate cookie error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mock_service.NewMockStoreRepository(ctrl)
			test.mockBehavior(repo)

			authService := NewAuthService(nil, nil, repo)
			sid, err := authService.GenerateCookie(context.Background(), test.role)

			assert.Equal(t, test.expectedSID, sid)
			assert.Equal(t, test.expectedError, err)
		})
	}
}

func TestAuthService_DeleteCookie(t *testing.T) {
	tests := []struct {
		name          string
		mockBehavior  func(r *mock_service.MockStoreRepository)
		session       string
		expectedError error
	}{
		{
			name: "Success",
			mockBehavior: func(r *mock_service.MockStoreRepository) {
				r.EXPECT().DeleteSession(gomock.Any(), "dummySession").Return(nil)
			},
			session:       "dummySession",
			expectedError: nil,
		},
		{
			name: "Error Deleting Cookie",
			mockBehavior: func(r *mock_service.MockStoreRepository) {
				r.EXPECT().DeleteSession(gomock.Any(), "dummySession").Return(errors.New("delete cookie error"))
			},
			session:       "dummySession",
			expectedError: errors.New("delete cookie error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mock_service.NewMockStoreRepository(ctrl)
			test.mockBehavior(repo)

			authService := NewAuthService(nil, nil, repo)
			err := authService.DeleteCookie(context.Background(), test.session)

			assert.Equal(t, test.expectedError, err)
		})
	}
}

func TestAuthService_GetSessionValue(t *testing.T) {
	tests := []struct {
		name          string
		mockBehavior  func(r *mock_service.MockStoreRepository)
		session       string
		expectedID    int
		expectedError error
	}{
		{
			name: "Success",
			mockBehavior: func(r *mock_service.MockStoreRepository) {
				r.EXPECT().GetSession(gomock.Any(), "dummySession").Return(1, nil)
			},
			session:       "dummySession",
			expectedID:    1,
			expectedError: nil,
		},
		{
			name: "Error Getting Session",
			mockBehavior: func(r *mock_service.MockStoreRepository) {
				r.EXPECT().GetSession(gomock.Any(), "dummySession").Return(0, errors.New("get session error"))
			},
			session:       "dummySession",
			expectedID:    0,
			expectedError: errors.New("get session error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mock_service.NewMockStoreRepository(ctrl)
			test.mockBehavior(repo)

			authService := NewAuthService(nil, nil, repo)
			id, err := authService.GetSessionValue(context.Background(), test.session)

			assert.Equal(t, test.expectedID, id)
			assert.Equal(t, test.expectedError, err)
		})
	}
}

func TestGeneratePasswordHash(t *testing.T) {
	password := "password"
	salt := "salt"

	hash := GeneratePasswordHash(password, salt)

	assert.Equal(t, hash, GeneratePasswordHash(password, salt))
}
