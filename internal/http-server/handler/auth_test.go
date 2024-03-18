package handler

import (
	"bytes"
	"errors"
	"github.com/Max425/film-library.git/internal/domain"
	"github.com/Max425/film-library.git/internal/http-server/handler/dto"
	"github.com/Max425/film-library.git/mocks/service"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuthHandler_SignIn(t *testing.T) {
	mockUser, _ := domain.NewUser(1, "bob", "test@example.com", "password", "", 0)
	input := dto.SignInInput{
		Mail:     "user@mail.ru",
		Password: "qwerty",
	}
	tests := []struct {
		name                 string
		requestMethod        string
		requestBody          string
		mockBehavior         func(r *mock_handler.MockAuthService, input dto.SignInInput)
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:          "Ok",
			requestMethod: http.MethodPost,
			requestBody:   `{"mail": "user@mail.ru", "password": "qwerty"}`,
			mockBehavior: func(r *mock_handler.MockAuthService, input dto.SignInInput) {
				r.EXPECT().GetUser(gomock.Any(), input.Mail, input.Password).Return(mockUser, nil)
				r.EXPECT().GenerateCookie(gomock.Any(), mockUser.Role()).Return("sessionID", nil)
			},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: `{"status":200,"message":"success","payload":"login :)"}`,
		},
		{
			name:                 "Wrong Method",
			requestMethod:        http.MethodGet,
			mockBehavior:         func(r *mock_handler.MockAuthService, input dto.SignInInput) {},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: `{"status":405,"message":"Method Not Allowed","payload":""}`,
		},
		{
			name:                 "Invalid JSON",
			requestMethod:        http.MethodPost,
			requestBody:          `{"mail": "user@mail.ru"`,
			mockBehavior:         func(r *mock_handler.MockAuthService, input dto.SignInInput) {},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: `{"status":400,"message":"bad request","payload":""}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockAuthService := mock_handler.NewMockAuthService(mockCtrl)
			test.mockBehavior(mockAuthService, input)

			logger := zap.NewNop()
			authHandler := NewAuthHandler(logger, mockAuthService)

			req, err := http.NewRequest(test.requestMethod, "/api/auth/login", bytes.NewBufferString(test.requestBody))
			if err != nil {
				t.Fatal(err)
			}
			rr := httptest.NewRecorder()

			authHandler.SignIn(rr, req)

			assert.Equal(t, test.expectedStatusCode, rr.Code)
			assert.Equal(t, test.expectedResponseBody, rr.Body.String())
		})
	}
}

func TestAuthHandler_Logout(t *testing.T) {
	tests := []struct {
		name                 string
		requestMethod        string
		cookie               *http.Cookie
		mockBehavior         func(r *mock_handler.MockAuthService, cookie *http.Cookie)
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:          "Ok",
			requestMethod: http.MethodDelete,
			cookie: &http.Cookie{
				Name:  "session_id",
				Value: "value",
			},
			mockBehavior: func(r *mock_handler.MockAuthService, cookie *http.Cookie) {
				r.EXPECT().DeleteCookie(gomock.Any(), cookie.Value).Return(nil)
			},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: `{"status":200,"message":"success","payload":"Logout :)"}`,
		},
		{
			name:          "Invalid cookie deletion",
			requestMethod: http.MethodDelete,
			cookie: &http.Cookie{
				Name:  "session_id",
				Value: "value",
			},
			mockBehavior: func(r *mock_handler.MockAuthService, cookie *http.Cookie) {
				r.EXPECT().DeleteCookie(gomock.Any(), cookie.Value).Return(errors.New("error"))
			},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: `{"status":500,"message":"internal error","payload":""}`,
		},
		{
			name:                 "No cookie",
			requestMethod:        http.MethodDelete,
			cookie:               nil,
			mockBehavior:         func(r *mock_handler.MockAuthService, cookie *http.Cookie) {},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: `{"status":401,"message":"no session","payload":""}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockAuthService := mock_handler.NewMockAuthService(mockCtrl)
			test.mockBehavior(mockAuthService, test.cookie)

			logger := zap.NewNop()
			authHandler := NewAuthHandler(logger, mockAuthService)

			req, err := http.NewRequest(test.requestMethod, "/api/auth/logout", nil)
			if err != nil {
				t.Fatal(err)
			}

			if test.cookie != nil {
				req.AddCookie(test.cookie)
			}

			rr := httptest.NewRecorder()

			authHandler.Logout(rr, req)

			assert.Equal(t, test.expectedStatusCode, rr.Code)
			assert.Equal(t, test.expectedResponseBody, rr.Body.String())
		})
	}
}
