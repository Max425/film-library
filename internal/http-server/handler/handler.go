package handler

import (
	"go.uber.org/zap"
	"net/http"
)

type Service interface {
	AuthService
	FilmService
	ActorService
}

type Handler struct {
	log *zap.Logger
	Middleware
	AuthHandler
	FilmHandler
	ActorHandler
}

func NewHandler(service Service, log *zap.Logger) *Handler {
	return &Handler{
		log,
		*NewMiddleware(log, service),
		*NewAuthHandler(log, service),
		*NewFilmHandler(log, service),
		*NewActorHandler(log, service),
	}
}

func (h *Handler) UseRecoveryLoggingAuth(next http.HandlerFunc) http.HandlerFunc {
	return h.panicRecoveryMiddleware(
		h.loggingMiddleware(
			h.authMiddleware(next)),
	)
}

func (h *Handler) UseRecoveryLogging(next http.HandlerFunc) http.HandlerFunc {
	return h.panicRecoveryMiddleware(
		h.loggingMiddleware(
			h.authMiddleware(next)),
	)
}
