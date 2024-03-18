package handler

import (
	"bytes"
	"errors"
	"github.com/Max425/film-library.git/internal/domain"
	"github.com/Max425/film-library.git/mocks/service"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestFilmHandler_CreateFilm(t *testing.T) {
	mockFilm, _ := domain.NewFilm(0, "Inception", "A thriller", time.Date(2024, time.March, 18, 0, 0, 0, 0, time.UTC), 9.2, nil)
	tests := []struct {
		name                 string
		requestMethod        string
		requestBody          string
		mockBehavior         func(r *mock_handler.MockFilmService, film *domain.Film)
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:          "Ok",
			requestMethod: http.MethodPost,
			requestBody:   `{"title": "Inception", "description": "A thriller", "release_date": "2024-03-18", "rating": 9.2}`,
			mockBehavior: func(r *mock_handler.MockFilmService, film *domain.Film) {
				r.EXPECT().CreateFilm(gomock.Any(), film).Return(mockFilm, nil)
			},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: `{"status":200,"message":"success","payload":{"id":0,"title":"Inception","description":"A thriller","release_date":"2024-03-18T00:00:00Z","rating":9.2,"actors":[]}}`,
		},
		{
			name:                 "Invalid JSON",
			requestMethod:        http.MethodPost,
			requestBody:          `{"title": "Inception", "description": "A thriller", "release_date": "2024-03-18"`,
			mockBehavior:         func(r *mock_handler.MockFilmService, film *domain.Film) {},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: `{"status":400,"message":"bad request","payload":""}`,
		},
		{
			name:          "Create film error",
			requestMethod: http.MethodPost,
			requestBody:   `{"title": "Inception", "description": "A thriller", "release_date": "2024-03-18", "rating": 9.2}`,
			mockBehavior: func(r *mock_handler.MockFilmService, film *domain.Film) {
				r.EXPECT().CreateFilm(gomock.Any(), film).Return(nil, errors.New("some error"))
			},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: `{"status":500,"message":"internal error","payload":""}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockFilmService := mock_handler.NewMockFilmService(mockCtrl)
			test.mockBehavior(mockFilmService, mockFilm)

			logger := zap.NewNop()
			filmHandler := NewFilmHandler(logger, mockFilmService)

			req, err := http.NewRequest(test.requestMethod, "/api/create_films", bytes.NewBufferString(test.requestBody))
			if err != nil {
				t.Fatal(err)
			}
			rr := httptest.NewRecorder()

			filmHandler.CreateFilm(rr, req)

			assert.Equal(t, test.expectedStatusCode, rr.Code)
			assert.Equal(t, test.expectedResponseBody, rr.Body.String())
		})
	}
}

func TestFilmHandler_UpdateFilm(t *testing.T) {
	mockFilm, _ := domain.NewFilm(0, "Inception", "A thriller", time.Date(2024, time.March, 18, 0, 0, 0, 0, time.UTC), 9.2, nil)
	tests := []struct {
		name                 string
		requestMethod        string
		requestBody          string
		mockBehavior         func(r *mock_handler.MockFilmService, film *domain.Film)
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:          "Ok",
			requestMethod: http.MethodPut,
			requestBody:   `{"title": "Inception", "description": "A thriller", "release_date": "2024-03-18", "rating": 9.2}`,
			mockBehavior: func(r *mock_handler.MockFilmService, film *domain.Film) {
				r.EXPECT().UpdateFilm(gomock.Any(), film).Return(mockFilm, nil)
			},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: `{"status":200,"message":"success","payload":{"id":0,"title":"Inception","description":"A thriller","release_date":"2024-03-18T00:00:00Z","rating":9.2,"actors":[]}}`,
		},
		{
			name:                 "Invalid JSON",
			requestMethod:        http.MethodPut,
			requestBody:          `{"title": "Inception", "description": "A thriller", "release_date": "2024-03-18"`,
			mockBehavior:         func(r *mock_handler.MockFilmService, film *domain.Film) {},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: `{"status":400,"message":"bad request","payload":""}`,
		},
		{
			name:          "Update film error",
			requestMethod: http.MethodPut,
			requestBody:   `{"title": "Inception", "description": "A thriller", "release_date": "2024-03-18", "rating": 9.2}`,
			mockBehavior: func(r *mock_handler.MockFilmService, film *domain.Film) {
				r.EXPECT().UpdateFilm(gomock.Any(), film).Return(nil, errors.New("some error"))
			},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: `{"status":500,"message":"internal error","payload":""}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockFilmService := mock_handler.NewMockFilmService(mockCtrl)
			test.mockBehavior(mockFilmService, mockFilm)

			logger := zap.NewNop()
			filmHandler := NewFilmHandler(logger, mockFilmService)

			req, err := http.NewRequest(test.requestMethod, "/api/update_films", bytes.NewBufferString(test.requestBody))
			if err != nil {
				t.Fatal(err)
			}
			rr := httptest.NewRecorder()

			filmHandler.UpdateFilm(rr, req)

			assert.Equal(t, test.expectedStatusCode, rr.Code)
			assert.Equal(t, test.expectedResponseBody, rr.Body.String())
		})
	}
}

func TestFilmHandler_UpdateFilmActors(t *testing.T) {
	mockFilm, _ := domain.NewFilm(0, "Inception", "A thriller", time.Date(2024, time.March, 18, 0, 0, 0, 0, time.UTC), 9.2, nil)
	tests := []struct {
		name                 string
		requestMethod        string
		requestURL           string
		requestBody          string
		mockBehavior         func(r *mock_handler.MockFilmService, film *domain.Film, actorsId []int)
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:          "Ok",
			requestMethod: http.MethodPost,
			requestURL:    "/api/update_films_actors/1",
			requestBody:   `[1, 2, 3]`,
			mockBehavior: func(r *mock_handler.MockFilmService, film *domain.Film, actorsId []int) {
				r.EXPECT().UpdateFilmActors(gomock.Any(), 1, actorsId).Return(mockFilm, nil)
			},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: `{"status":200,"message":"success","payload":{"id":0,"title":"Inception","description":"A thriller","release_date":"2024-03-18T00:00:00Z","rating":9.2,"actors":[]}}`,
		},
		{
			name:                 "Invalid JSON",
			requestMethod:        http.MethodPost,
			requestURL:           "/api/update_films_actors/1",
			requestBody:          `[1, 2, 3`,
			mockBehavior:         func(r *mock_handler.MockFilmService, film *domain.Film, actorsId []int) {},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: `{"status":400,"message":"bad request","payload":""}`,
		},
		{
			name:          "Update film actors error",
			requestMethod: http.MethodPost,
			requestURL:    "/api/update_films_actors/1",
			requestBody:   `[1, 2, 3]`,
			mockBehavior: func(r *mock_handler.MockFilmService, film *domain.Film, actorsId []int) {
				r.EXPECT().UpdateFilmActors(gomock.Any(), 1, actorsId).Return(nil, errors.New("some error"))
			},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: `{"status":500,"message":"internal error","payload":""}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockFilmService := mock_handler.NewMockFilmService(mockCtrl)
			test.mockBehavior(mockFilmService, mockFilm, []int{1, 2, 3})

			logger := zap.NewNop()
			filmHandler := NewFilmHandler(logger, mockFilmService)

			req, err := http.NewRequest(test.requestMethod, test.requestURL, bytes.NewBufferString(test.requestBody))
			if err != nil {
				t.Fatal(err)
			}
			rr := httptest.NewRecorder()

			filmHandler.UpdateFilmActors(rr, req)

			assert.Equal(t, test.expectedStatusCode, rr.Code)
			assert.Equal(t, test.expectedResponseBody, rr.Body.String())
		})
	}
}

func TestFilmHandler_DeleteFilm(t *testing.T) {
	tests := []struct {
		name                 string
		requestMethod        string
		requestURL           string
		mockBehavior         func(r *mock_handler.MockFilmService, id int)
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:          "Ok",
			requestMethod: http.MethodDelete,
			requestURL:    "/api/films/1",
			mockBehavior: func(r *mock_handler.MockFilmService, id int) {
				r.EXPECT().DeleteFilm(gomock.Any(), id).Return(nil)
			},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: `{"status":200,"message":"success","payload":"Film deleted successfully"}`,
		},
		{
			name:                 "Invalid ID",
			requestMethod:        http.MethodDelete,
			requestURL:           "/api/films/abc",
			mockBehavior:         func(r *mock_handler.MockFilmService, id int) {},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: `{"status":400,"message":"invalid film ID","payload":""}`,
		},
		{
			name:          "Delete film error",
			requestMethod: http.MethodDelete,
			requestURL:    "/api/films/1",
			mockBehavior: func(r *mock_handler.MockFilmService, id int) {
				r.EXPECT().DeleteFilm(gomock.Any(), id).Return(errors.New("some error"))
			},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: `{"status":500,"message":"internal error","payload":""}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockFilmService := mock_handler.NewMockFilmService(mockCtrl)
			var id int
			if test.name == "Invalid ID" {
				id = 0 // invalid id
			} else {
				id = 1
			}
			test.mockBehavior(mockFilmService, id)

			logger := zap.NewNop()
			filmHandler := NewFilmHandler(logger, mockFilmService)

			req, err := http.NewRequest(test.requestMethod, test.requestURL, nil)
			if err != nil {
				t.Fatal(err)
			}
			rr := httptest.NewRecorder()

			filmHandler.DeleteFilm(rr, req)

			assert.Equal(t, test.expectedStatusCode, rr.Code)
			assert.Equal(t, test.expectedResponseBody, rr.Body.String())
		})
	}
}

func TestFilmHandler_SearchFilms(t *testing.T) {
	tests := []struct {
		name                 string
		requestMethod        string
		requestURL           string
		mockBehavior         func(r *mock_handler.MockFilmService, films []*domain.Film, err error)
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:          "Ok",
			requestMethod: http.MethodGet,
			requestURL:    "/api/search_films/thriller",
			mockBehavior: func(r *mock_handler.MockFilmService, films []*domain.Film, err error) {
				r.EXPECT().SearchFilms(gomock.Any(), "thriller").Return(films, err)
			},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: `{"status":200,"message":"success","payload":[{"id":0,"title":"Inception","description":"A thriller","release_date":"2024-03-18T00:00:00Z","rating":9.2,"actors":[]}]}`,
		},
		{
			name:          "Internal Server Error",
			requestMethod: http.MethodGet,
			requestURL:    "/api/search_films/thriller",
			mockBehavior: func(r *mock_handler.MockFilmService, films []*domain.Film, err error) {
				r.EXPECT().SearchFilms(gomock.Any(), "thriller").Return(nil, errors.New("some error"))
			},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: `{"status":500,"message":"internal error","payload":""}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockFilmService := mock_handler.NewMockFilmService(mockCtrl)
			var films []*domain.Film
			if test.expectedStatusCode == http.StatusOK {
				film, _ := domain.NewFilm(0, "Inception", "A thriller", time.Date(2024, time.March, 18, 0, 0, 0, 0, time.UTC), 9.2, nil)
				films = []*domain.Film{film}
			}
			test.mockBehavior(mockFilmService, films, nil)

			logger := zap.NewNop()
			filmHandler := NewFilmHandler(logger, mockFilmService)

			req, err := http.NewRequest(test.requestMethod, test.requestURL, nil)
			if err != nil {
				t.Fatal(err)
			}
			rr := httptest.NewRecorder()

			filmHandler.SearchFilms(rr, req)

			assert.Equal(t, test.expectedStatusCode, rr.Code)
			assert.Equal(t, test.expectedResponseBody, rr.Body.String())
		})
	}
}

func TestFilmHandler_GetAllFilms(t *testing.T) {
	tests := []struct {
		name                 string
		requestMethod        string
		requestURL           string
		queryParams          map[string]string
		mockBehavior         func(r *mock_handler.MockFilmService, films []*domain.Film, err error)
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:          "Ok",
			requestMethod: http.MethodGet,
			requestURL:    "/api/films",
			queryParams:   map[string]string{"sort_by": "title", "order": "asc"},
			mockBehavior: func(r *mock_handler.MockFilmService, films []*domain.Film, err error) {
				r.EXPECT().GetAllFilms(gomock.Any(), "title", "asc").Return(films, err)
			},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: `{"status":200,"message":"success","payload":[{"id":0,"title":"Inception","description":"A thriller","release_date":"2024-03-18T00:00:00Z","rating":9.2,"actors":[]}]}`,
		},
		{
			name:                 "Invalid Sort Order",
			requestMethod:        http.MethodGet,
			requestURL:           "/api/films",
			queryParams:          map[string]string{"sort_by": "title", "order": "invalid"},
			mockBehavior:         func(r *mock_handler.MockFilmService, films []*domain.Film, err error) {},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: `{"status":400,"message":"Invalid sort order","payload":""}`,
		},
		{
			name:                 "Invalid Sort By Field",
			requestMethod:        http.MethodGet,
			requestURL:           "/api/films",
			queryParams:          map[string]string{"sort_by": "invalid", "order": "asc"},
			mockBehavior:         func(r *mock_handler.MockFilmService, films []*domain.Film, err error) {},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: `{"status":400,"message":"Invalid sort by field","payload":""}`,
		},
		{
			name:          "Internal Server Error",
			requestMethod: http.MethodGet,
			requestURL:    "/api/films",
			queryParams:   map[string]string{"sort_by": "title", "order": "asc"},
			mockBehavior: func(r *mock_handler.MockFilmService, films []*domain.Film, err error) {
				r.EXPECT().GetAllFilms(gomock.Any(), "title", "asc").Return(nil, errors.New("some error"))
			},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: `{"status":500,"message":"internal error","payload":""}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockFilmService := mock_handler.NewMockFilmService(mockCtrl)
			var films []*domain.Film
			if test.expectedStatusCode == http.StatusOK {
				film, _ := domain.NewFilm(0, "Inception", "A thriller", time.Date(2024, time.March, 18, 0, 0, 0, 0, time.UTC), 9.2, nil)
				films = []*domain.Film{film}
			}
			test.mockBehavior(mockFilmService, films, nil)

			logger := zap.NewNop()
			filmHandler := NewFilmHandler(logger, mockFilmService)

			req, err := http.NewRequest(test.requestMethod, test.requestURL, nil)
			if err != nil {
				t.Fatal(err)
			}
			q := req.URL.Query()
			for key, value := range test.queryParams {
				q.Add(key, value)
			}
			req.URL.RawQuery = q.Encode()

			rr := httptest.NewRecorder()

			filmHandler.GetAllFilms(rr, req)

			assert.Equal(t, test.expectedStatusCode, rr.Code)
			assert.Equal(t, test.expectedResponseBody, rr.Body.String())
		})
	}
}
