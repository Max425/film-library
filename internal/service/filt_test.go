package service

import (
	"context"
	"errors"
	"github.com/Max425/film-library.git/internal/domain"
	"github.com/Max425/film-library.git/mocks/db"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestFilmService_CreateFilm(t *testing.T) {
	mockFilm, _ := domain.NewFilm(1, "title", "desc", time.Unix(0, 0), 2.2, nil)

	tests := []struct {
		name          string
		mockBehavior  func(r *mock_service.MockFilmRepository)
		film          *domain.Film
		expectedFilm  *domain.Film
		expectedError error
	}{
		{
			name: "Success",
			mockBehavior: func(r *mock_service.MockFilmRepository) {
				r.EXPECT().CreateFilm(gomock.Any(), mockFilm).Return(mockFilm, nil)
			},
			film:          mockFilm,
			expectedFilm:  mockFilm,
			expectedError: nil,
		},
		{
			name: "Error Creating Film",
			mockBehavior: func(r *mock_service.MockFilmRepository) {
				r.EXPECT().CreateFilm(gomock.Any(), mockFilm).Return(nil, errors.New("create film error"))
			},
			film:          mockFilm,
			expectedFilm:  nil,
			expectedError: errors.New("create film error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mock_service.NewMockFilmRepository(ctrl)
			test.mockBehavior(repo)

			service := NewFilmService(repo, nil)
			film, err := service.CreateFilm(context.Background(), test.film)

			assert.Equal(t, test.expectedFilm, film)
			assert.Equal(t, test.expectedError, err)
		})
	}
}

func TestFilmService_GetFilmByID(t *testing.T) {
	mockFilm, _ := domain.NewFilm(1, "title", "desc", time.Unix(0, 0), 2.2, nil)

	tests := []struct {
		name          string
		mockBehavior  func(r *mock_service.MockFilmRepository)
		id            int
		expectedFilm  *domain.Film
		expectedError error
	}{
		{
			name: "Success",
			mockBehavior: func(r *mock_service.MockFilmRepository) {
				r.EXPECT().FindFilmByID(gomock.Any(), 1).Return(mockFilm, nil)
			},
			id:            1,
			expectedFilm:  mockFilm,
			expectedError: nil,
		},
		{
			name: "Error Getting Film",
			mockBehavior: func(r *mock_service.MockFilmRepository) {
				r.EXPECT().FindFilmByID(gomock.Any(), 1).Return(nil, errors.New("get film error"))
			},
			id:            1,
			expectedFilm:  nil,
			expectedError: errors.New("get film error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mock_service.NewMockFilmRepository(ctrl)
			test.mockBehavior(repo)

			service := NewFilmService(repo, nil)
			film, err := service.GetFilmByID(context.Background(), test.id)

			assert.Equal(t, test.expectedFilm, film)
			assert.Equal(t, test.expectedError, err)
		})
	}
}

func TestFilmService_UpdateFilm(t *testing.T) {
	mockFilm, _ := domain.NewFilm(1, "title", "desc", time.Unix(0, 0), 2.2, nil)

	tests := []struct {
		name          string
		mockBehavior  func(r *mock_service.MockFilmRepository)
		film          *domain.Film
		expectedFilm  *domain.Film
		expectedError error
	}{
		{
			name: "Success",
			mockBehavior: func(r *mock_service.MockFilmRepository) {
				r.EXPECT().FindFilmByID(gomock.Any(), 1).Return(mockFilm, nil)
				r.EXPECT().UpdateFilm(gomock.Any(), mockFilm).Return(mockFilm, nil)
			},
			film:          mockFilm,
			expectedFilm:  mockFilm,
			expectedError: nil,
		},
		{
			name: "Error Finding Film",
			mockBehavior: func(r *mock_service.MockFilmRepository) {
				r.EXPECT().FindFilmByID(gomock.Any(), 1).Return(nil, errors.New("find film error"))
			},
			film:          mockFilm,
			expectedFilm:  nil,
			expectedError: errors.New("find film error"),
		},
		{
			name: "Error Updating Film",
			mockBehavior: func(r *mock_service.MockFilmRepository) {
				r.EXPECT().FindFilmByID(gomock.Any(), 1).Return(mockFilm, nil)
				r.EXPECT().UpdateFilm(gomock.Any(), mockFilm).Return(nil, errors.New("update film error"))
			},
			film:          mockFilm,
			expectedFilm:  nil,
			expectedError: errors.New("update film error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mock_service.NewMockFilmRepository(ctrl)
			test.mockBehavior(repo)

			service := NewFilmService(repo, nil)
			film, err := service.UpdateFilm(context.Background(), test.film)

			assert.Equal(t, test.expectedFilm, film)
			assert.Equal(t, test.expectedError, err)
		})
	}
}

func TestFilmService_UpdateFilmActors(t *testing.T) {
	mockFilm, _ := domain.NewFilm(1, "title", "desc", time.Unix(0, 0), 2.2, nil)
	actorIDs := []int{1, 2, 3}

	tests := []struct {
		name          string
		mockBehavior  func(r *mock_service.MockFilmRepository)
		filmID        int
		expectedFilm  *domain.Film
		expectedError error
	}{
		{
			name: "Success",
			mockBehavior: func(r *mock_service.MockFilmRepository) {
				r.EXPECT().FindFilmByID(gomock.Any(), 1).Return(mockFilm, nil)
				r.EXPECT().UpdateFilmActors(gomock.Any(), 1, actorIDs).Return(mockFilm, nil)
			},
			filmID:        1,
			expectedFilm:  mockFilm,
			expectedError: nil,
		},
		{
			name: "Error, Finding Film",
			mockBehavior: func(r *mock_service.MockFilmRepository) {
				r.EXPECT().FindFilmByID(gomock.Any(), 1).Return(nil, errors.New("find film error"))
			},
			filmID:        1,
			expectedFilm:  nil,
			expectedError: errors.New("find film error"),
		},
		{
			name: "Error Updating Film Actors",
			mockBehavior: func(r *mock_service.MockFilmRepository) {
				r.EXPECT().FindFilmByID(gomock.Any(), 1).Return(mockFilm, nil)
				r.EXPECT().UpdateFilmActors(gomock.Any(), 1, actorIDs).Return(nil, errors.New("update film actors error"))
			},
			filmID:        1,
			expectedFilm:  nil,
			expectedError: errors.New("update film actors error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mock_service.NewMockFilmRepository(ctrl)
			test.mockBehavior(repo)

			service := NewFilmService(repo, nil)
			film, err := service.UpdateFilmActors(context.Background(), test.filmID, actorIDs)

			assert.Equal(t, test.expectedFilm, film)
			assert.Equal(t, test.expectedError, err)
		})
	}
}

func TestFilmService_DeleteFilm(t *testing.T) {
	mockFilm, _ := domain.NewFilm(1, "title", "desc", time.Unix(0, 0), 2.2, nil)

	tests := []struct {
		name          string
		mockBehavior  func(r *mock_service.MockFilmRepository)
		filmID        int
		expectedError error
	}{
		{
			name: "Success",
			mockBehavior: func(r *mock_service.MockFilmRepository) {
				r.EXPECT().FindFilmByID(gomock.Any(), 1).Return(mockFilm, nil)
				r.EXPECT().DeleteFilm(gomock.Any(), 1).Return(nil)
			},
			filmID:        1,
			expectedError: nil,
		},
		{
			name: "Error Finding Film",
			mockBehavior: func(r *mock_service.MockFilmRepository) {
				r.EXPECT().FindFilmByID(gomock.Any(), 1).Return(nil, errors.New("find film error"))
			},
			filmID:        1,
			expectedError: errors.New("find film error"),
		},
		{
			name: "Error Deleting Film",
			mockBehavior: func(r *mock_service.MockFilmRepository) {
				r.EXPECT().FindFilmByID(gomock.Any(), 1).Return(mockFilm, nil)
				r.EXPECT().DeleteFilm(gomock.Any(), 1).Return(errors.New("delete film error"))
			},
			filmID:        1,
			expectedError: errors.New("delete film error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mock_service.NewMockFilmRepository(ctrl)
			test.mockBehavior(repo)

			service := NewFilmService(repo, nil)
			err := service.DeleteFilm(context.Background(), test.filmID)

			assert.Equal(t, test.expectedError, err)
		})
	}
}

func TestFilmService_GetAllFilms(t *testing.T) {
	mockFilm1, _ := domain.NewFilm(1, "title", "desc", time.Unix(0, 0), 2.2, nil)
	mockFilm2, _ := domain.NewFilm(2, "title", "desc", time.Unix(0, 0), 2.2, nil)
	mockFilm3, _ := domain.NewFilm(3, "title", "desc", time.Unix(0, 0), 2.2, nil)
	mockFilms := []*domain.Film{mockFilm1, mockFilm2, mockFilm3}

	tests := []struct {
		name          string
		mockBehavior  func(r *mock_service.MockFilmRepository)
		expectedFilms []*domain.Film
		expectedError error
	}{
		{
			name: "Success",
			mockBehavior: func(r *mock_service.MockFilmRepository) {
				r.EXPECT().GetAllFilms(gomock.Any(), "", "").Return(mockFilms, nil)
			},
			expectedFilms: mockFilms,
			expectedError: nil,
		},
		{
			name: "Error Getting Films",
			mockBehavior: func(r *mock_service.MockFilmRepository) {
				r.EXPECT().GetAllFilms(gomock.Any(), "", "").Return(nil, errors.New("get films error"))
			},
			expectedFilms: nil,
			expectedError: errors.New("get films error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mock_service.NewMockFilmRepository(ctrl)
			test.mockBehavior(repo)

			service := NewFilmService(repo, nil)
			films, err := service.GetAllFilms(context.Background(), "", "")

			assert.Equal(t, test.expectedFilms, films)
			assert.Equal(t, test.expectedError, err)
		})
	}
}

func TestFilmService_SearchFilms(t *testing.T) {
	mockFilm1, _ := domain.NewFilm(1, "title", "desc", time.Unix(0, 0), 2.2, nil)
	mockFilm2, _ := domain.NewFilm(2, "title", "desc", time.Unix(0, 0), 2.2, nil)
	mockFilm3, _ := domain.NewFilm(3, "title", "desc", time.Unix(0, 0), 2.2, nil)
	mockFilms := []*domain.Film{mockFilm1, mockFilm2, mockFilm3}

	tests := []struct {
		name          string
		mockBehavior  func(r *mock_service.MockFilmRepository)
		fragment      string
		expectedFilms []*domain.Film
		expectedError error
	}{
		{
			name: "Success",
			mockBehavior: func(r *mock_service.MockFilmRepository) {
				r.EXPECT().SearchFilms(gomock.Any(), "test").Return(mockFilms, nil)
			},
			fragment:      "test",
			expectedFilms: mockFilms,
			expectedError: nil,
		},
		{
			name: "Error Searching Films",
			mockBehavior: func(r *mock_service.MockFilmRepository) {
				r.EXPECT().SearchFilms(gomock.Any(), "test").Return(nil, errors.New("search films error"))
			},
			fragment:      "test",
			expectedFilms: nil,
			expectedError: errors.New("search films error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mock_service.NewMockFilmRepository(ctrl)
			test.mockBehavior(repo)

			service := NewFilmService(repo, nil)
			films, err := service.SearchFilms(context.Background(), test.fragment)

			assert.Equal(t, test.expectedFilms, films)
			assert.Equal(t, test.expectedError, err)
		})
	}
}
