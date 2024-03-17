package handler

import (
	"context"
	"github.com/Max425/film-library.git/internal/http-server/handler/dto"
	"go.uber.org/zap"
	"google.golang.org/grpc/interop/grpc_testing/core"
	"io"
	"net/http"
	"time"
)

type FilmService interface {
	CreateFilm(ctx context.Context, film *core.Film) (*core.Film, error)
	GetFilmByID(ctx context.Context, id int) (*core.Film, error)
	UpdateFilm(ctx context.Context, id int, film *core.Film) (*core.Film, error)
	DeleteFilm(ctx context.Context, id int) error
}

type CompanyHandler struct {
	log            *zap.Logger
	companyService CompanyService
}

// @Summary create event
// @Tags event
// @Accept  json
// @Produce  json
// @Param input body core.Event true "New event"
// @Success 200 {object} int
// @Failure 400,500 {object} string
// @Router /create_event [post]
func (h *CompanyHandler) createEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusBadRequest, err.Error())
		return
	}
	var event core.Event
	err = event.UnmarshalJSON(body)
	if err != nil || event.UserID == "" {
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusBadRequest, "invalid input body")
		return
	}

	id, err := h.services.CreateEvent(&event)
	if err != nil {
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusInternalServerError, err.Error())
		return
	}
	dto.NewSuccessClientResponseDto(r.Context(), w, http.StatusOK, id)
}

// @Summary update event
// @Tags event
// @Accept  json
// @Produce  json
// @Param input body core.Event true "Updated event"
// @Success 200 {object} string
// @Failure 400,500 {object} string
// @Router /update_event [post]
func (h *Handler) updateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusBadRequest, err.Error())
		return
	}
	var event core.Event
	err = event.UnmarshalJSON(body)
	if err != nil || event.UserID == "" || event.ID == 0 {
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusBadRequest, "invalid input body")
		return
	}

	err = h.services.UpdateEvent(&event)
	if err != nil {
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusInternalServerError, err.Error())
		return
	}

	dto.NewSuccessClientResponseDto(r.Context(), w, http.StatusOK, "Event updated successfully")
}

// @Summary delete event
// @Tags event
// @Accept  json
// @Produce  json
// @Param input body core.Event true  "Event ID"
// @Success 200 {object} string
// @Failure 400,500 {object} string
// @Router /delete_event [post]
func (h *Handler) deleteEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusBadRequest, err.Error())
		return
	}
	var event core.Event
	err = event.UnmarshalJSON(body)
	if err != nil || event.ID == 0 {
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusBadRequest, "invalid input body")
		return
	}

	err = h.services.DeleteEvent(event.ID)
	if err != nil {
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusInternalServerError, err.Error())
		return
	}

	dto.NewSuccessClientResponseDto(r.Context(), w, http.StatusOK, "Event deleted successfully")
}

// @Summary get events for day
// @Tags event
// @Accept  json
// @Produce  json
// @Success 200 {object} []core.Event
// @Failure 400,500 {object} string
// @Router /events_for_day [get]
func (h *Handler) getEventsForDay(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
		return
	}

	events, err := h.services.GetEventsForTime(time.Hour * 24)
	if err != nil {
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusInternalServerError, err.Error())
		return
	}

	dto.NewSuccessClientResponseDto(r.Context(), w, http.StatusOK, events)
}

// @Summary get events for week
// @Tags event
// @Accept  json
// @Produce  json
// @Success 200 {object} []core.Event
// @Failure 400,500 {object} string
// @Router /events_for_week [get]
func (h *Handler) getEventsForWeek(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
		return
	}

	events, err := h.services.GetEventsForTime(time.Hour * 24 * 7)
	if err != nil {
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusInternalServerError, err.Error())
		return
	}

	dto.NewSuccessClientResponseDto(r.Context(), w, http.StatusOK, events)
}

// @Summary get events for month
// @Tags event
// @Accept  json
// @Produce  json
// @Success 200 {object} []core.Event
// @Failure 400,500 {object} string
// @Router /events_for_month [get]
func (h *Handler) getEventsForMonth(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
		return
	}

	events, err := h.services.GetEventsForTime(time.Hour * 24 * 30)
	if err != nil {
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusInternalServerError, err.Error())
		return
	}

	dto.NewSuccessClientResponseDto(r.Context(), w, http.StatusOK, events)
}
