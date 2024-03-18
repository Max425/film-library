package service

import (
	"context"
	"errors"
	"github.com/Max425/film-library.git/internal/domain"
	"github.com/Max425/film-library.git/mocks/db"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestActorService_CreateActor(t *testing.T) {
	mockActor, _ := domain.NewActor(1, "test", "male", time.Unix(0, 0), nil)

	tests := []struct {
		name          string
		mockBehavior  func(r *mock_service.MockActorRepository)
		actor         *domain.Actor
		expectedActor *domain.Actor
		expectedError error
	}{
		{
			name: "Success",
			mockBehavior: func(r *mock_service.MockActorRepository) {
				r.EXPECT().CreateActor(gomock.Any(), mockActor).Return(mockActor, nil)
			},
			actor:         mockActor,
			expectedActor: mockActor,
			expectedError: nil,
		},
		{
			name: "Error Creating Actor",
			mockBehavior: func(r *mock_service.MockActorRepository) {
				r.EXPECT().CreateActor(gomock.Any(), mockActor).Return(nil, errors.New("create actor error"))
			},
			actor:         mockActor,
			expectedActor: nil,
			expectedError: errors.New("create actor error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mock_service.NewMockActorRepository(ctrl)
			test.mockBehavior(repo)

			service := NewActorService(repo, nil)
			actor, err := service.CreateActor(context.Background(), test.actor)

			assert.Equal(t, test.expectedActor, actor)
			assert.Equal(t, test.expectedError, err)
		})
	}
}

func TestActorService_GetActorByID(t *testing.T) {

	mockActor, _ := domain.NewActor(1, "test", "male", time.Unix(0, 0), nil)

	tests := []struct {
		name          string
		mockBehavior  func(r *mock_service.MockActorRepository)
		actorID       int
		expectedActor *domain.Actor
		expectedError error
	}{
		{
			name: "Success",
			mockBehavior: func(r *mock_service.MockActorRepository) {
				r.EXPECT().FindActorByID(gomock.Any(), 1).Return(mockActor, nil)
			},
			actorID:       1,
			expectedActor: mockActor,
			expectedError: nil,
		},
		{
			name: "Error Getting Actor",
			mockBehavior: func(r *mock_service.MockActorRepository) {
				r.EXPECT().FindActorByID(gomock.Any(), 1).Return(nil, errors.New("get actor error"))
			},
			actorID:       1,
			expectedActor: nil,
			expectedError: errors.New("get actor error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mock_service.NewMockActorRepository(ctrl)
			test.mockBehavior(repo)

			service := NewActorService(repo, nil)
			actor, err := service.GetActorByID(context.Background(), test.actorID)

			assert.Equal(t, test.expectedActor, actor)
			assert.Equal(t, test.expectedError, err)
		})
	}
}

func TestActorService_UpdateActor(t *testing.T) {
	mockActor, _ := domain.NewActor(1, "test", "male", time.Unix(0, 0), nil)

	tests := []struct {
		name          string
		mockBehavior  func(r *mock_service.MockActorRepository)
		actor         *domain.Actor
		expectedActor *domain.Actor
		expectedError error
	}{
		{
			name: "Success",
			mockBehavior: func(r *mock_service.MockActorRepository) {
				r.EXPECT().FindActorByID(gomock.Any(), mockActor.GetId()).Return(mockActor, nil)
				r.EXPECT().UpdateActor(gomock.Any(), mockActor).Return(mockActor, nil)
			},
			actor:         mockActor,
			expectedActor: mockActor,
			expectedError: nil,
		},
		{
			name: "Error Finding Actor",
			mockBehavior: func(r *mock_service.MockActorRepository) {
				r.EXPECT().FindActorByID(gomock.Any(), mockActor.GetId()).Return(nil, errors.New("find actor error"))
			},
			actor:         mockActor,
			expectedActor: nil,
			expectedError: errors.New("find actor error"),
		},
		{
			name: "Error Updating Actor",
			mockBehavior: func(r *mock_service.MockActorRepository) {
				r.EXPECT().FindActorByID(gomock.Any(), mockActor.GetId()).Return(mockActor, nil)
				r.EXPECT().UpdateActor(gomock.Any(), mockActor).Return(nil, errors.New("update actor error"))
			},
			actor:         mockActor,
			expectedActor: nil,
			expectedError: errors.New("update actor error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mock_service.NewMockActorRepository(ctrl)
			test.mockBehavior(repo)

			service := NewActorService(repo, nil)
			actor, err := service.UpdateActor(context.Background(), test.actor)

			assert.Equal(t, test.expectedActor, actor)
			assert.Equal(t, test.expectedError, err)
		})
	}
}

func TestActorService_DeleteActor(t *testing.T) {
	mockActorID := 1

	tests := []struct {
		name          string
		mockBehavior  func(r *mock_service.MockActorRepository)
		actorID       int
		expectedError error
	}{
		{
			name: "Success",
			mockBehavior: func(r *mock_service.MockActorRepository) {
				r.EXPECT().FindActorByID(gomock.Any(), mockActorID).Return(&domain.Actor{}, nil)
				r.EXPECT().DeleteActor(gomock.Any(), mockActorID).Return(nil)
			},
			actorID:       mockActorID,
			expectedError: nil,
		},
		{
			name: "Error Finding Actor",
			mockBehavior: func(r *mock_service.MockActorRepository) {
				r.EXPECT().FindActorByID(gomock.Any(), mockActorID).Return(nil, errors.New("find actor error"))
			},
			actorID:       mockActorID,
			expectedError: errors.New("find actor error"),
		},
		{
			name: "Error Deleting Actor",
			mockBehavior: func(r *mock_service.MockActorRepository) {
				r.EXPECT().FindActorByID(gomock.Any(), mockActorID).Return(&domain.Actor{}, nil)
				r.EXPECT().DeleteActor(gomock.Any(), mockActorID).Return(errors.New("delete actor error"))
			},
			actorID:       mockActorID,
			expectedError: errors.New("delete actor error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mock_service.NewMockActorRepository(ctrl)
			test.mockBehavior(repo)

			service := NewActorService(repo, nil)
			err := service.DeleteActor(context.Background(), test.actorID)

			assert.Equal(t, test.expectedError, err)
		})
	}
}

func TestActorService_GetAllActors(t *testing.T) {
	mockActor1, _ := domain.NewActor(1, "test", "male", time.Unix(0, 0), nil)
	mockActor2, _ := domain.NewActor(2, "test", "male", time.Unix(0, 0), nil)
	mockActor3, _ := domain.NewActor(3, "test", "male", time.Unix(0, 0), nil)
	mockActors := []*domain.Actor{mockActor1, mockActor2, mockActor3}

	tests := []struct {
		name           string
		mockBehavior   func(r *mock_service.MockActorRepository)
		expectedActors []*domain.Actor
		expectedError  error
	}{
		{
			name: "Success",
			mockBehavior: func(r *mock_service.MockActorRepository) {
				r.EXPECT().GetAllActors(gomock.Any()).Return(mockActors, nil)
			},
			expectedActors: mockActors,
			expectedError:  nil,
		},
		{
			name: "Error Getting Actors",
			mockBehavior: func(r *mock_service.MockActorRepository) {
				r.EXPECT().GetAllActors(gomock.Any()).Return(nil, errors.New("get actors error"))
			},
			expectedActors: nil,
			expectedError:  errors.New("get actors error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mock_service.NewMockActorRepository(ctrl)
			test.mockBehavior(repo)

			service := NewActorService(repo, nil)
			actors, err := service.GetAllActors(context.Background())

			assert.Equal(t, test.expectedActors, actors)
			assert.Equal(t, test.expectedError, err)
		})
	}
}
