package handler

//
//import (
//	"bytes"
//	"context"
//	"errors"
//	"github.com/Max425/wbschool_exam_L2/tree/main/develop/dev11/internal/model/core"
//	"github.com/Max425/wbschool_exam_L2/tree/main/develop/dev11/internal/service"
//	"github.com/Max425/wbschool_exam_L2/tree/main/develop/dev11/internal/service/mocks"
//	"github.com/golang/mock/gomock"
//	"github.com/stretchr/testify/assert"
//	"net/http"
//	"net/http/httptest"
//	"testing"
//	"time"
//)
//
//func TestHandler_createEvent(t *testing.T) {
//	dt, _ := time.Parse("2006-01-02", "2024-09-09")
//	mockEvent := core.Event{
//		Date:   dt,
//		Title:  "test",
//		UserID: "1",
//	}
//	eventJSON := `{"date": "2024-09-09", "title": "test", "user_id":"1"}`
//	tests := []struct {
//		name                 string
//		requestBody          string
//		mockBehavior         func(r *mock_service.MockEvent)
//		expectedStatusCode   int
//		expectedResponseBody string
//	}{
//		{
//			name:        "Event Create",
//			requestBody: eventJSON,
//			mockBehavior: func(r *mock_service.MockEvent) {
//				r.EXPECT().CreateEvent(&mockEvent).Return(1, nil)
//			},
//			expectedStatusCode:   http.StatusOK,
//			expectedResponseBody: `{"result":1}`,
//		},
//		{
//			name:        "invalid input body",
//			requestBody: `{"date": "2024-09-09"}`,
//			mockBehavior: func(r *mock_service.MockEvent) {
//			},
//			expectedStatusCode:   http.StatusBadRequest,
//			expectedResponseBody: `{"error":"invalid input body"}`,
//		},
//		{
//			name:        "Error",
//			requestBody: eventJSON,
//			mockBehavior: func(r *mock_service.MockEvent) {
//				r.EXPECT().CreateEvent(&mockEvent).Return(0, errors.New("some error"))
//			},
//			expectedStatusCode:   http.StatusInternalServerError,
//			expectedResponseBody: `{"error":"some error"}`,
//		},
//	}
//
//	for _, test := range tests {
//		t.Run(test.name, func(t *testing.T) {
//			ctrl := gomock.NewController(t)
//			defer ctrl.Finish()
//
//			mockEventService := mock_service.NewMockEvent(ctrl)
//			test.mockBehavior(mockEventService)
//
//			ctx := context.Background()
//			services := &service.Service{Event: mockEventService}
//			handler := Handler{services: services}
//
//			mux := http.NewServeMux()
//			mux.HandleFunc("/create_event", handler.createEvent)
//
//			w := httptest.NewRecorder()
//			req := httptest.NewRequest("POST", "/create_event", bytes.NewBufferString(test.requestBody))
//			mux.ServeHTTP(w, req.WithContext(ctx))
//
//			assert.Equal(t, test.expectedStatusCode, w.Code)
//			assert.Equal(t, test.expectedResponseBody, w.Body.String())
//		})
//	}
//}
