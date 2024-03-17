package handler

import (
	"go.uber.org/zap"
	"net/http"
)

type Service interface {
	FilmService
	ActorService
}

type Handler struct {
	log *zap.Logger
	FilmHandler
	ActorHandler
}

func NewHandler(service Service, log *zap.Logger) *Handler {
	return &Handler{
		log,
		*NewFilmHandler(log, service),
		*NewActorHandler(log, service),
	}
}

func (h *Handler) Use(next http.HandlerFunc) http.HandlerFunc {
	return h.panicRecoveryMiddleware(
		h.loggingMiddleware(next),
	)
}
