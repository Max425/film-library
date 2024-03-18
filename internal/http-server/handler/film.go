package handler

import (
	"context"
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
	DeleteFilm(ctx context.Context, id int) error
	GetAllFilms(ctx context.Context) ([]*domain.Film, error)
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

// GetAllFilms retrieves all films.
// @Summary Retrieve all films
// @Tags films
// @Accept json
// @Produce json
// @Success 200 {array} []dto.Film "List of films"
// @Failure 500 {string} string "Internal server error"
// @Router /api/films [get]
func (h *FilmHandler) GetAllFilms(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
		return
	}

	films, err := h.filmService.GetAllFilms(r.Context())
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
