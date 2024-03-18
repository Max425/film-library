package handler

import (
	"context"
	"errors"
	"github.com/Max425/film-library.git/internal/common/constants"
	"github.com/Max425/film-library.git/internal/http-server/handler/dto"
	"go.uber.org/zap"
	"net/http"
	"runtime/debug"
	"time"
)

type Middleware struct {
	log         *zap.Logger
	authService AuthService
}

func NewMiddleware(log *zap.Logger, authService AuthService) *Middleware {
	return &Middleware{
		log:         log,
		authService: authService,
	}
}

func (h *Middleware) authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := r.Cookie("session_id")
		if errors.Is(err, http.ErrNoCookie) {
			dto.NewErrorClientResponseDto(r.Context(), w, http.StatusUnauthorized, "Need auth")
			return
		}

		role, err := h.authService.GetSessionValue(r.Context(), session.Value)
		if err != nil {
			dto.NewErrorClientResponseDto(r.Context(), w, http.StatusUnauthorized, "Need auth")
			return
		}
		if (r.Method == "POST" || r.Method == "PUT") && role != constants.AdminRole {
			dto.NewErrorClientResponseDto(r.Context(), w, http.StatusForbidden, "forbidden")
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (h *Middleware) loggingMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		requestInfo := &dto.RequestInfo{}
		ctx := context.WithValue(r.Context(), constants.KeyRequestInfo, requestInfo)

		next.ServeHTTP(w, r.WithContext(ctx))

		timing := time.Since(start)

		requestInfo, ok := ctx.Value(constants.KeyRequestInfo).(*dto.RequestInfo)
		var code int
		var message string
		if ok {
			code = requestInfo.Status
			message = requestInfo.Message

		}

		h.log.Info("Request handled",
			zap.Int("StatusCode", code),
			zap.String("Message", message),
			zap.String("RequestURI", r.RequestURI),
			zap.Duration("Time", timing),
		)
	}
}

func (h *Middleware) panicRecoveryMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				h.log.Error("Panic",
					zap.String("Method", r.Method),
					zap.String("RequestURI", r.RequestURI),
					zap.String("Error", err.(string)),
					zap.String("Message", string(debug.Stack())),
				)
				dto.NewErrorClientResponseDto(r.Context(), w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
			}
		}()
		next.ServeHTTP(w, r)
	}
}
