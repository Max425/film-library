package handler

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/Max425/film-library.git/internal/common"
	"github.com/Max425/film-library.git/internal/common/constants"
	"github.com/Max425/film-library.git/internal/domain"
	"github.com/Max425/film-library.git/internal/http-server/handler/dto"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type AuthService interface {
	GenerateCookie(ctx context.Context, role int) (string, error)
	DeleteCookie(ctx context.Context, session string) error
	GetSessionValue(ctx context.Context, session string) (int, error)
	CreateUser(ctx context.Context, user *domain.User) (int, error)
	GetUser(ctx context.Context, mail, password string) (*domain.User, error)
}

type AuthHandler struct {
	log         *zap.Logger
	authService AuthService
}

func NewAuthHandler(log *zap.Logger, authService AuthService) *AuthHandler {
	return &AuthHandler{
		log:         log,
		authService: authService,
	}
}

// SignIn
// @Summary log in to account
// @Tags auth
// @ID login
// @Accept  json
// @Produce  json
// @Param input body dto.SignInInput true "Sign-in input parameters"
// @Success 200 {object} string
// @Failure 400,404 {object} string
// @Router /api/auth/login [post]
func (h *AuthHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
		return
	}

	var input dto.SignInInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusBadRequest, common.ErrBadRequest.String())
		return
	}

	user, err := h.authService.GetUser(r.Context(), input.Mail, input.Password)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			dto.NewErrorClientResponseDto(r.Context(), w, http.StatusUnauthorized, common.InvalidMailOrPassword.String())
		} else if errors.Is(err, domain.ErrRequired) {
			dto.NewErrorClientResponseDto(r.Context(), w, http.StatusBadRequest, err.Error())
		} else {
			dto.NewErrorClientResponseDto(r.Context(), w, http.StatusInternalServerError, common.ErrInternal.String())
		}
		return
	}
	SID, err := h.authService.GenerateCookie(r.Context(), user.Role())
	if err != nil {
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusInternalServerError, common.ErrInternal.String())
		return
	}

	http.SetCookie(w, createCookie("session_id", SID))
	dto.NewSuccessClientResponseDto(r.Context(), w, http.StatusOK, "")
}

// Logout
// @Summary log out of account
// @Tags auth
// @ID logout
// @Accept  json
// @Produce  json
// @Success 200 {object} string
// @Failure 400,404 {object} string
// @Router /api/auth/logout [delete]
func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
		return
	}

	session, err := r.Cookie("session_id")
	if err != nil {
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusUnauthorized, "no session")
		return
	}

	if err = h.authService.DeleteCookie(r.Context(), session.Value); err != nil {
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusInternalServerError, common.ErrInternal.String())
		return
	}

	session.Expires = time.Now().AddDate(0, 0, -1)
	session.Path = "/"

	http.SetCookie(w, session)
	dto.NewSuccessClientResponseDto(r.Context(), w, http.StatusOK, "")
}

// SignUp
// @Summary sign up account
// @Tags auth
// @ID create-account
// @Accept  json
// @Produce  json
// @Param input body dto.SignUpInput true "Sign-up input user"
// @Success 200 {object} map[string]int
// @Failure 400,404 {object} string
// @Router /api/auth/sign-up [post]
func (h *AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
		return
	}

	var input dto.SignUpInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusBadRequest, common.ErrBadRequest.String())
		return
	}

	domainUser, err := dto.SignUpInputToDomainUser(&input)
	if err != nil {
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusBadRequest, common.ErrBadRequest.String())
		return
	}
	userId, err := h.authService.CreateUser(r.Context(), domainUser)
	if err != nil {
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusInternalServerError, common.ErrInternal.String())
		return
	}

	cookie, err := h.authService.GenerateCookie(r.Context(), 0)
	if err != nil {
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusInternalServerError, common.ErrInternal.String())
		return
	}

	http.SetCookie(w, createCookie("session_id", cookie))
	dto.NewSuccessClientResponseDto(r.Context(), w, http.StatusOK, map[string]int{"id": userId})
}

func createCookie(name, SID string) *http.Cookie {
	return &http.Cookie{
		Name:     name,
		Value:    SID,
		Expires:  time.Now().Add(constants.CookieExpire),
		Path:     "/",
		HttpOnly: true,
	}
}
