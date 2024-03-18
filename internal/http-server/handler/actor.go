package handler

import (
	"context"
	"github.com/Max425/film-library.git/internal/common"
	"github.com/Max425/film-library.git/internal/domain"
	"github.com/Max425/film-library.git/internal/http-server/handler/dto"
	"go.uber.org/zap"
	"io"
	"net/http"
	"strconv"
)

type ActorService interface {
	CreateActor(ctx context.Context, actor *domain.Actor) (*domain.Actor, error)
	GetActorByID(ctx context.Context, id int) (*domain.Actor, error)
	UpdateActor(ctx context.Context, actor *domain.Actor) (*domain.Actor, error)
	DeleteActor(ctx context.Context, id int) error
	GetAllActors(ctx context.Context) ([]*domain.Actor, error)
}

type ActorHandler struct {
	log          *zap.Logger
	actorService ActorService
}

func NewActorHandler(log *zap.Logger, actorService ActorService) *ActorHandler {
	return &ActorHandler{
		log:          log,
		actorService: actorService,
	}
}

// CreateActor creates a new actor.
// @Summary Create a new actor
// @Tags actors
// @Accept json
// @Produce json
// @Param input body dto.Actor true "Actor object to be created"
// @Success 201 {object} dto.Actor "Actor created successfully"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /api/create_actors [post]
func (h *ActorHandler) CreateActor(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
		return
	}

	var actor dto.Actor
	body, _ := io.ReadAll(r.Body)
	if err := actor.UnmarshalJSON(body); err != nil {
		h.log.Error("Failed to decode actor", zap.Error(err))
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusBadRequest, common.ErrBadRequest.String())
		return
	}
	domainActor, err := dto.ActorDtoToDomain(&actor)
	if err != nil {
		h.log.Error("Failed to convert actor", zap.Error(err))
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusBadRequest, err.Error())
		return
	}

	actorCreated, err := h.actorService.CreateActor(r.Context(), domainActor)
	if err != nil {
		h.log.Error("Failed to create actor", zap.Error(err))
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusInternalServerError, common.ErrInternal.String())
		return
	}

	dto.NewSuccessClientResponseDto(r.Context(), w, actorCreated)
}

// UpdateActor updates an existing actor.
// @Summary Update an existing actor
// @Tags actors
// @Accept json
// @Produce json
// @Param id path int true "Actor ID"
// @Param input body dto.Actor true "Actor object to be updated"
// @Success 200 {object} dto.Actor "Actor updated successfully"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /api/update_actors [put]
func (h *ActorHandler) UpdateActor(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
		return
	}

	var actor dto.Actor
	body, _ := io.ReadAll(r.Body)
	if err := actor.UnmarshalJSON(body); err != nil {
		h.log.Error("Failed to decode actor", zap.Error(err))
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusBadRequest, common.ErrBadRequest.String())
		return
	}
	domainActor, err := dto.ActorDtoToDomain(&actor)
	if err != nil {
		h.log.Error("Failed to convert actor", zap.Error(err))
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusBadRequest, err.Error())
		return
	}

	actorUpdated, err := h.actorService.UpdateActor(r.Context(), domainActor)
	if err != nil {
		h.log.Error("Failed to update actor", zap.Error(err))
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusInternalServerError, common.ErrInternal.String())
		return
	}

	dto.NewSuccessClientResponseDto(r.Context(), w, actorUpdated)
}

// DeleteActor deletes an existing actor.
// @Summary Delete an existing actor
// @Tags actors
// @Accept json
// @Produce json
// @Param id path int true "Actor ID"
// @Success 200 {string} string "Actor deleted successfully"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /api/actors/{id} [delete]
func (h *ActorHandler) DeleteActor(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
		return
	}

	actorID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusBadRequest, "invalid actor ID")
		return
	}

	if err = h.actorService.DeleteActor(r.Context(), actorID); err != nil {
		h.log.Error("Failed to delete actor", zap.Error(err))
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusInternalServerError, err.Error())
		return
	}

	dto.NewSuccessClientResponseDto(r.Context(), w, "Actor deleted successfully")
}

// GetAllActors retrieves all actors.
// @Summary Retrieve all actors
// @Tags actors
// @Accept json
// @Produce json
// @Success 200 {array} []dto.Actor "List of actors"
// @Failure 500 {string} string "Internal server error"
// @Router /api/actors [get]
func (h *ActorHandler) GetAllActors(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
		return
	}

	actors, err := h.actorService.GetAllActors(r.Context())
	if err != nil {
		h.log.Error("Failed to get all actors", zap.Error(err))
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusInternalServerError, err.Error())
		return
	}

	dto.NewSuccessClientResponseDto(r.Context(), w, actors)
}
