package handler_test

//
//import (
//	"encoding/json"
//	"errors"
//	"github.com/Max425/film-library.git/mocks/service"
//	"net/http"
//	"net/http/httptest"
//	"testing"
//	"time"
//
//	"github.com/Max425/film-library.git/internal/domain"
//	"github.com/Max425/film-library.git/internal/http-server/handler"
//	"github.com/Max425/film-library.git/internal/http-server/handler/dto"
//	"github.com/golang/mock/gomock"
//	"github.com/stretchr/testify/assert"
//	"go.uber.org/zap"
//)
//
//func TestFilmHandler_GetAllFilms(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	mockFilmService := mock_handler.NewMockFilmService(ctrl)
//	logger := zap.NewNop()
//	filmHandler := handler.NewFilmHandler(logger, mockFilmService)
//
//	t.Run("Success", func(t *testing.T) {
//		mockFilms := []*domain.Film{
//			domain.NewFilm(1, "Film 1", "Description 1", time.Now(), 8.5, nil),
//			{ID: 2, Title: "Film 2", Rating: 7.9, ReleaseDate: "2024-03-19"},
//			{ID: 3, Title: "Film 3", Rating: 6.8, ReleaseDate: "2024-03-20"},
//		}
//
//		mockFilmService.EXPECT().GetAllFilms(gomock.Any(), "", "").Return(mockFilms, nil)
//
//		req := httptest.NewRequest("GET", "/api/films", nil)
//		w := httptest.NewRecorder()
//		filmHandler.GetAllFilms(w, req)
//
//		resp := w.Result()
//		assert.Equal(t, http.StatusOK, resp.StatusCode)
//
//		var films []*dto.Film
//		err := json.NewDecoder(resp.Body).Decode(&films)
//		assert.NoError(t, err)
//
//		assert.Equal(t, 3, len(films))
//		assert.Equal(t, 1, films[0].ID)
//		assert.Equal(t, "Film 1", films[0].Title)
//		assert.Equal(t, 8.5, films[0].Rating)
//		assert.Equal(t, "2024-03-18", films[0].ReleaseDate)
//	})
//
//	t.Run("InternalError", func(t *testing.T) {
//		mockFilmService.EXPECT().GetAllFilms(gomock.Any(), "", "").Return(nil, errors.New("internal error"))
//
//		req := httptest.NewRequest("GET", "/api/films", nil)
//		w := httptest.NewRecorder()
//		filmHandler.GetAllFilms(w, req)
//
//		resp := w.Result()
//		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
//	})
//}
//
//func TestFilmHandler_GetAllFilms_WithSorting(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	mockFilmService := mock_handler.NewMockFilmService(ctrl)
//	logger := zap.NewNop()
//	filmHandler := handler.NewFilmHandler(logger, mockFilmService)
//
//	t.Run("SortByTitle", func(t *testing.T) {
//		mockFilms := []*domain.Film{
//			{ID: 1, Title: "Film A", Rating: 8.5, ReleaseDate: "2024-03-18"},
//			{ID: 2, Title: "Film B", Rating: 7.9, ReleaseDate: "2024-03-19"},
//			{ID: 3, Title: "Film C", Rating: 6.8, ReleaseDate: "2024-03-20"},
//		}
//
//		mockFilmService.EXPECT().GetAllFilms(gomock.Any(), "title", "asc").Return(mockFilms, nil)
//
//		req := httptest.NewRequest("GET", "/api/films?sort_by=title&order=asc", nil)
//		w := httptest.NewRecorder()
//		filmHandler.GetAllFilms(w, req)
//
//		resp := w.Result()
//		assert.Equal(t, http.StatusOK, resp.StatusCode)
//
//		var films []*dto.Film
//		err := json.NewDecoder(resp.Body).Decode(&films)
//		assert.NoError(t, err)
//
//		assert.Equal(t, 3, len(films))
//		assert.Equal(t, "Film A", films[0].Title)
//		assert.Equal(t, "Film B", films[1].Title)
//		assert.Equal(t, "Film C", films[2].Title)
//	})
//
//	t.Run("InvalidSortOrder", func(t *testing.T) {
//		req := httptest.NewRequest("GET", "/api/films?sort_by=title&order=invalid", nil)
//		w := httptest.NewRecorder()
//		filmHandler.GetAllFilms(w, req)
//
//		resp := w.Result()
//		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
//	})
//
//	t.Run("InvalidSortByField", func(t *testing.T) {
//		req := httptest.NewRequest("GET", "/api/films?sort_by=invalid&order=asc", nil)
//		w := httptest.NewRecorder()
//		filmHandler.GetAllFilms(w, req)
//
//		resp := w.Result()
//		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
//	})
//}
