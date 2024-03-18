package handler

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/Max425/film-library.git/internal/common"
	"github.com/Max425/film-library.git/internal/domain"
	"github.com/Max425/film-library.git/internal/http-server/handler/dto"
	"go.uber.org/zap"
	"io"
	"net/http"
	"strconv"
)

type FilmService interface {
	CreateFilm(ctx context.Context, film *domain.Film) (*domain.Film, error)
	GetFilmByID(ctx context.Context, id int) (*domain.Film, error)
	UpdateFilm(ctx context.Context, film *domain.Film) (*domain.Film, error)
	UpdateFilmActors(ctx context.Context, id int, actorsId []int) (*domain.Film, error)
	DeleteFilm(ctx context.Context, id int) error
	SearchFilms(ctx context.Context, fragment string) ([]*domain.Film, error)
	GetAllFilms(ctx context.Context, sortBy, order string) ([]*domain.Film, error)
}

type FilmHandler struct {
	log         *zap.Logger
	filmService FilmService
}

func NewFilmHandler(log *zap.Logger, filmService FilmService) *FilmHandler {
	return &FilmHandler{
		log:         log,
		filmService: filmService,
	}
}

// CreateFilm creates a new film.
// @Summary Create a new film
// @Tags films
// @Accept json
// @Produce json
// @Param input body dto.Film true "Film object to be created"
// @Success 201 {object} dto.Film "Film created successfully"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /api/create_films [post]
func (h *FilmHandler) CreateFilm(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
		return
	}

	var film dto.Film
	body, _ := io.ReadAll(r.Body)
	if err := film.UnmarshalJSON(body); err != nil {
		h.log.Error("Failed to decode film", zap.Error(err))
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusBadRequest, common.ErrBadRequest.String())
		return
	}
	domainFilm, err := dto.FilmDtoToDomain(&film)
	if err != nil {
		h.log.Error("Failed to convert film", zap.Error(err))
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusBadRequest, err.Error())
		return
	}

	createdFilm, err := h.filmService.CreateFilm(r.Context(), domainFilm)
	if err != nil {
		h.log.Error("Failed to create film", zap.Error(err))
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusInternalServerError, common.ErrInternal.String())
		return
	}

	dto.NewSuccessClientResponseDto(r.Context(), w, dto.FilmDomainToDto(createdFilm))
}

// UpdateFilm updates an existing film.
// @Summary Update an existing film
// @Tags films
// @Accept json
// @Produce json
// @Param input body dto.Film true "Film object to be updated"
// @Success 200 {object} dto.Film "Film updated successfully"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /api/update_films [put]
func (h *FilmHandler) UpdateFilm(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
		return
	}

	var film dto.Film
	body, _ := io.ReadAll(r.Body)
	if err := film.UnmarshalJSON(body); err != nil {
		h.log.Error("Failed to decode film", zap.Error(err))
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusBadRequest, common.ErrBadRequest.String())
		return
	}

	domainFilm, err := dto.FilmDtoToDomain(&film)
	if err != nil {
		h.log.Error("Failed to convert film", zap.Error(err))
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusBadRequest, err.Error())
		return
	}

	updatedFilm, err := h.filmService.UpdateFilm(r.Context(), domainFilm)
	if err != nil {
		h.log.Error("Failed to update film", zap.Error(err))
		if errors.Is(err, domain.ErrNotFound) {
			dto.NewErrorClientResponseDto(r.Context(), w, http.StatusNotFound, common.ErrNotFound.String())
			return
		}
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusInternalServerError, common.ErrInternal.String())
		return
	}

	dto.NewSuccessClientResponseDto(r.Context(), w, dto.FilmDomainToDto(updatedFilm))
}

// UpdateFilmActors updates an existing film.
// @Summary Update an existing film
// @Tags films
// @Accept json
// @Produce json
// @Param id path int true "Film ID"
// @Param input body []int true "id actors for film"
// @Success 200 {object} dto.Film "Film updated successfully"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /api/update_films_actors/{id} [post]
func (h *FilmHandler) UpdateFilmActors(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
		return
	}

	idStr := r.URL.Path[len("/api/update_films_actors/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusBadRequest, "invalid film ID")
		return
	}

	var actorsId []int
	body, _ := io.ReadAll(r.Body)
	if err := json.Unmarshal(body, &actorsId); err != nil {
		h.log.Error("Failed to decode film", zap.Error(err))
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusBadRequest, common.ErrBadRequest.String())
		return
	}

	updatedFilm, err := h.filmService.UpdateFilmActors(r.Context(), id, actorsId)
	if err != nil {
		h.log.Error("Failed to update film", zap.Error(err))
		if errors.Is(err, domain.ErrNotFound) {
			dto.NewErrorClientResponseDto(r.Context(), w, http.StatusNotFound, common.ErrNotFound.String())
			return
		}
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusInternalServerError, common.ErrInternal.String())
		return
	}

	dto.NewSuccessClientResponseDto(r.Context(), w, dto.FilmDomainToDto(updatedFilm))
}

// DeleteFilm deletes an existing film.
// @Summary Delete an existing film
// @Tags films
// @Accept json
// @Produce json
// @Param id path int true "Film ID"
// @Success 200 {string} string "Film deleted successfully"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /api/films/{id} [delete]
func (h *FilmHandler) DeleteFilm(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
		return
	}

	idStr := r.URL.Path[len("/api/films/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusBadRequest, "invalid film ID")
		return
	}

	err = h.filmService.DeleteFilm(r.Context(), id)
	if err != nil {
		h.log.Error("Failed to delete film", zap.Error(err))
		if errors.Is(err, domain.ErrNotFound) {
			dto.NewErrorClientResponseDto(r.Context(), w, http.StatusNotFound, common.ErrNotFound.String())
			return
		}
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusInternalServerError, common.ErrInternal.String())
		return
	}

	dto.NewSuccessClientResponseDto(r.Context(), w, "Film deleted successfully")
}

// SearchFilms
// @Summary Search films by pattern
// @Tags films
// @Accept json
// @Produce json
// @Param pattern path string true "Film pattern"
// @Success 200 {array} []dto.Film "List of films"
// @Failure 500 {string} string "Internal server error"
// @Router /api/search_films/{pattern} [get]
func (h *FilmHandler) SearchFilms(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
		return
	}
	pattern := r.URL.Path[len("/api/search_films/"):]
	films, err := h.filmService.SearchFilms(r.Context(), pattern)
	if err != nil {
		h.log.Error("Failed to get all films", zap.Error(err))
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusInternalServerError, common.ErrInternal.String())
		return
	}
	data := make([]*dto.Film, len(films))
	for i, film := range films {
		data[i] = dto.FilmDomainToDto(film)
	}

	dto.NewSuccessClientResponseDto(r.Context(), w, data)
}

// GetAllFilms retrieves all films with optional sorting by title, rating, or release date.
// By default, films are sorted by rating in descending order.
// @Summary Retrieve all films
// @Tags films
// @Accept json
// @Produce json
// @Param sort_by query string false "Sort by: title, rating, release_date"
// @Param order query string false "Sort order: asc, desc"
// @Success 200 {array} []dto.Film "List of films"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /api/films [get]
func (h *FilmHandler) GetAllFilms(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
		return
	}

	// Parse query parameters for sorting
	sortBy := r.URL.Query().Get("sort_by")
	order := r.URL.Query().Get("order")

	// Default sorting by rating in descending order
	if sortBy == "" {
		sortBy = "rating"
	}
	if order == "" {
		order = "desc"
	}

	// Validate order parameter
	if order != "asc" && order != "desc" {
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusBadRequest, "Invalid sort order")
		return
	}
	validSortFields := map[string]bool{
		"title":        true,
		"rating":       true,
		"release_date": true,
	}
	if !validSortFields[sortBy] {
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusBadRequest, "Invalid sort by field")
		return
	}

	// Get films with optional sorting
	films, err := h.filmService.GetAllFilms(r.Context(), sortBy, order)
	if err != nil {
		h.log.Error("Failed to get all films", zap.Error(err))
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusInternalServerError, common.ErrInternal.String())
		return
	}

	// Convert domain films to DTOs
	data := make([]*dto.Film, len(films))
	for i, film := range films {
		data[i] = dto.FilmDomainToDto(film)
	}

	dto.NewSuccessClientResponseDto(r.Context(), w, data)
}
