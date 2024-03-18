// Code generated by MockGen. DO NOT EDIT.
// Source: internal/http-server/handler/film.go

// Package mock_handler is a generated GoMock package.
package mock_handler

import (
	context "context"
	reflect "reflect"

	domain "github.com/Max425/film-library.git/internal/domain"
	gomock "github.com/golang/mock/gomock"
)

// MockFilmService is a mock of FilmService interface.
type MockFilmService struct {
	ctrl     *gomock.Controller
	recorder *MockFilmServiceMockRecorder
}

// MockFilmServiceMockRecorder is the mock recorder for MockFilmService.
type MockFilmServiceMockRecorder struct {
	mock *MockFilmService
}

// NewMockFilmService creates a new mock instance.
func NewMockFilmService(ctrl *gomock.Controller) *MockFilmService {
	mock := &MockFilmService{ctrl: ctrl}
	mock.recorder = &MockFilmServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFilmService) EXPECT() *MockFilmServiceMockRecorder {
	return m.recorder
}

// CreateFilm mocks base method.
func (m *MockFilmService) CreateFilm(ctx context.Context, film *domain.Film) (*domain.Film, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFilm", ctx, film)
	ret0, _ := ret[0].(*domain.Film)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateFilm indicates an expected call of CreateFilm.
func (mr *MockFilmServiceMockRecorder) CreateFilm(ctx, film interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFilm", reflect.TypeOf((*MockFilmService)(nil).CreateFilm), ctx, film)
}

// DeleteFilm mocks base method.
func (m *MockFilmService) DeleteFilm(ctx context.Context, id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFilm", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteFilm indicates an expected call of DeleteFilm.
func (mr *MockFilmServiceMockRecorder) DeleteFilm(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFilm", reflect.TypeOf((*MockFilmService)(nil).DeleteFilm), ctx, id)
}

// GetAllFilms mocks base method.
func (m *MockFilmService) GetAllFilms(ctx context.Context, sortBy, order string) ([]*domain.Film, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllFilms", ctx, sortBy, order)
	ret0, _ := ret[0].([]*domain.Film)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllFilms indicates an expected call of GetAllFilms.
func (mr *MockFilmServiceMockRecorder) GetAllFilms(ctx, sortBy, order interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllFilms", reflect.TypeOf((*MockFilmService)(nil).GetAllFilms), ctx, sortBy, order)
}

// GetFilmByID mocks base method.
func (m *MockFilmService) GetFilmByID(ctx context.Context, id int) (*domain.Film, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFilmByID", ctx, id)
	ret0, _ := ret[0].(*domain.Film)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFilmByID indicates an expected call of GetFilmByID.
func (mr *MockFilmServiceMockRecorder) GetFilmByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFilmByID", reflect.TypeOf((*MockFilmService)(nil).GetFilmByID), ctx, id)
}

// SearchFilms mocks base method.
func (m *MockFilmService) SearchFilms(ctx context.Context, fragment string) ([]*domain.Film, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchFilms", ctx, fragment)
	ret0, _ := ret[0].([]*domain.Film)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchFilms indicates an expected call of SearchFilms.
func (mr *MockFilmServiceMockRecorder) SearchFilms(ctx, fragment interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchFilms", reflect.TypeOf((*MockFilmService)(nil).SearchFilms), ctx, fragment)
}

// UpdateFilm mocks base method.
func (m *MockFilmService) UpdateFilm(ctx context.Context, film *domain.Film) (*domain.Film, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateFilm", ctx, film)
	ret0, _ := ret[0].(*domain.Film)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateFilm indicates an expected call of UpdateFilm.
func (mr *MockFilmServiceMockRecorder) UpdateFilm(ctx, film interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateFilm", reflect.TypeOf((*MockFilmService)(nil).UpdateFilm), ctx, film)
}

// UpdateFilmActors mocks base method.
func (m *MockFilmService) UpdateFilmActors(ctx context.Context, id int, actorsId []int) (*domain.Film, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateFilmActors", ctx, id, actorsId)
	ret0, _ := ret[0].(*domain.Film)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateFilmActors indicates an expected call of UpdateFilmActors.
func (mr *MockFilmServiceMockRecorder) UpdateFilmActors(ctx, id, actorsId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateFilmActors", reflect.TypeOf((*MockFilmService)(nil).UpdateFilmActors), ctx, id, actorsId)
}