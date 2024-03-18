package handler

import (
	"bytes"
	"errors"
	"github.com/Max425/film-library.git/internal/domain"
	"github.com/Max425/film-library.git/mocks/service"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"
)

func TestActorHandler_CreateActor(t *testing.T) {
	mockActor, _ := domain.NewActor(1, "Bob", "male", time.Unix(0, 0), nil)

	tests := []struct {
		name                 string
		requestMethod        string
		requestBody          string
		mockBehavior         func(r *mock_handler.MockActorService, actor *domain.Actor)
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:          "Ok",
			requestMethod: http.MethodPost,
			requestBody:   `{"first_name": "Bob", "last_name": "Smith"}`,
			mockBehavior: func(r *mock_handler.MockActorService, actor *domain.Actor) {
				r.EXPECT().CreateActor(gomock.Any(), actor).Return(mockActor, nil)
			},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: `{"status":200,"message":"success","payload":{"id":1,"first_name":"Bob","last_name":"Smith"}}`,
		},
		{
			name:                 "Invalid JSON",
			requestMethod:        http.MethodPost,
			requestBody:          `{"first_name": "Bob", "last_name": "Smith"`,
			mockBehavior:         func(r *mock_handler.MockActorService, actor *domain.Actor) {},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"status":400,"message":"Bad request","payload":""}`,
		},
		{
			name:                 "Create actor error",
			requestMethod:        http.MethodPost,
			requestBody:          `{"first_name": "Bob", "last_name": "Smith"}`,
			mockBehavior:         func(r *mock_handler.MockActorService, actor *domain.Actor) {},
			expectedStatusCode:   http.StatusInternalServerError,
			expectedResponseBody: `{"status":500,"message":"Internal server error","payload":""}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockActorService := mock_handler.NewMockActorService(mockCtrl)
			test.mockBehavior(mockActorService, mockActor)

			logger := zap.NewNop()
			actorHandler := NewActorHandler(logger, mockActorService)

			req, err := http.NewRequest(test.requestMethod, "/api/create_actors", bytes.NewBufferString(test.requestBody))
			if err != nil {
				t.Fatal(err)
			}
			rr := httptest.NewRecorder()

			actorHandler.CreateActor(rr, req)

			assert.Equal(t, test.expectedStatusCode, rr.Code)
			assert.Equal(t, test.expectedResponseBody, rr.Body.String())
		})
	}
}

func TestActorHandler_UpdateActor(t *testing.T) {
	mockActor, _ := domain.NewActor(1, "Bob", "male", time.Unix(0, 0), nil)

	tests := []struct {
		name                 string
		requestMethod        string
		requestBody          string
		mockBehavior         func(r *mock_handler.MockActorService, actor *domain.Actor)
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:          "Ok",
			requestMethod: http.MethodPut,
			requestBody:   `{"first_name": "Bob", "last_name": "Smith"}`,
			mockBehavior: func(r *mock_handler.MockActorService, actor *domain.Actor) {
				r.EXPECT().UpdateActor(gomock.Any(), actor).Return(mockActor, nil)
			},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: `{"status":200,"message":"success","payload":{"id":1,"first_name":"Bob","last_name":"Smith"}}`,
		},
		{
			name:                 "Invalid JSON",
			requestMethod:        http.MethodPut,
			requestBody:          `{"first_name": "Bob", "last_name": "Smith"`,
			mockBehavior:         func(r *mock_handler.MockActorService, actor *domain.Actor) {},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"status":400,"message":"Bad request","payload":""}`,
		},
		{
			name:                 "Update actor error",
			requestMethod:        http.MethodPut,
			requestBody:          `{"first_name": "Bob", "last_name": "Smith"}`,
			mockBehavior:         func(r *mock_handler.MockActorService, actor *domain.Actor) {},
			expectedStatusCode:   http.StatusInternalServerError,
			expectedResponseBody: `{"status":500,"message":"Internal server error","payload":""}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockActorService := mock_handler.NewMockActorService(mockCtrl)
			test.mockBehavior(mockActorService, mockActor)

			logger := zap.NewNop()
			actorHandler := NewActorHandler(logger, mockActorService)

			req, err := http.NewRequest(test.requestMethod, "/api/update_actors", bytes.NewBufferString(test.requestBody))
			if err != nil {
				t.Fatal(err)
			}
			rr := httptest.NewRecorder()

			actorHandler.UpdateActor(rr, req)

			assert.Equal(t, test.expectedStatusCode, rr.Code)
			assert.Equal(t, test.expectedResponseBody, rr.Body.String())
		})
	}
}

func TestActorHandler_DeleteActor(t *testing.T) {
	tests := []struct {
		name                 string
		requestMethod        string
		requestURL           string
		mockBehavior         func(r *mock_handler.MockActorService, actorID int)
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:          "Ok",
			requestMethod: http.MethodDelete,
			requestURL:    "/api/actors/1",
			mockBehavior: func(r *mock_handler.MockActorService, actorID int) {
				r.EXPECT().DeleteActor(gomock.Any(), actorID).Return(nil)
			},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: `{"status":200,"message":"success","payload":"Actor deleted successfully"}`,
		},
		{
			name:                 "Invalid actor ID",
			requestMethod:        http.MethodDelete,
			requestURL:           "/api/actors/not_an_id",
			mockBehavior:         func(r *mock_handler.MockActorService, actorID int) {},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"status":400,"message":"Bad request","payload":""}`,
		},
		{
			name:          "Delete actor error",
			requestMethod: http.MethodDelete,
			requestURL:    "/api/actors/1",
			mockBehavior: func(r *mock_handler.MockActorService, actorID int) {
				r.EXPECT().DeleteActor(gomock.Any(), actorID).Return(errors.New("error"))
			},
			expectedStatusCode:   http.StatusInternalServerError,
			expectedResponseBody: `{"status":500,"message":"error","payload":""}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockActorService := mock_handler.NewMockActorService(mockCtrl)
			actorID, _ := strconv.Atoi(test.requestURL[len("/api/actors/"):])
			test.mockBehavior(mockActorService, actorID)

			logger := zap.NewNop()
			actorHandler := NewActorHandler(logger, mockActorService)

			req, err := http.NewRequest(test.requestMethod, test.requestURL, nil)
			if err != nil {
				t.Fatal(err)
			}
			rr := httptest.NewRecorder()

			actorHandler.DeleteActor(rr, req)

			assert.Equal(t, test.expectedStatusCode, rr.Code)
			assert.Equal(t, test.expectedResponseBody, rr.Body.String())
		})
	}
}

func TestActorHandler_GetAllActors(t *testing.T) {
	mockActor1, _ := domain.NewActor(1, "Bob", "male", time.Unix(0, 0), nil)
	mockActor2, _ := domain.NewActor(1, "Bob", "male", time.Unix(0, 0), nil)
	mockActors := []*domain.Actor{mockActor1, mockActor2}
	tests := []struct {
		name                 string
		requestMethod        string
		mockBehavior         func(r *mock_handler.MockActorService) *gomock.Call
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:          "Ok",
			requestMethod: http.MethodGet,
			mockBehavior: func(r *mock_handler.MockActorService) *gomock.Call {
				return r.EXPECT().GetAllActors(gomock.Any()).Return(mockActors, nil)
			},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: `[{"id":1,"first_name":"Bob","last_name":"Smith"},{"id":2,"first_name":"Alice","last_name":"Johnson"}]`,
		},
		{
			name:          "Get all actors error",
			requestMethod: http.MethodGet,
			mockBehavior: func(r *mock_handler.MockActorService) *gomock.Call {
				return r.EXPECT().GetAllActors(gomock.Any()).Return(nil, errors.New("error"))
			},
			expectedStatusCode:   http.StatusInternalServerError,
			expectedResponseBody: `{"status":500,"message":"error","payload":""}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockActorService := mock_handler.NewMockActorService(mockCtrl)
			test.mockBehavior(mockActorService)

			logger := zap.NewNop()
			actorHandler := NewActorHandler(logger, mockActorService)

			req, err := http.NewRequest(test.requestMethod, "/api/actors", nil)
			if err != nil {
				t.Fatal(err)
			}
			rr := httptest.NewRecorder()

			actorHandler.GetAllActors(rr, req)

			assert.Equal(t, test.expectedStatusCode, rr.Code)
			assert.Equal(t, test.expectedResponseBody, rr.Body.String())
		})
	}
}
