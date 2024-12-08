package service

import (
	"context"
	"pikachu/mock"
	"pikachu/model"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetUser(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	muRepo := mock.NewMockUserRepository(mockCtrl)
	userService := NewUserService(muRepo)
	ctx := context.Background()

	uid := uuid.New().String()
	email := gofakeit.Email()
	name := gofakeit.Name()

	expectedUser := &model.User{
		UID:   uid,
		Email: email,
		Name:  &name,
	}

	tests := []struct {
		name          string
		setupMock     func()
		expectedUser  *model.User
		expectedError error
	}{
		{
			name: "normal case",
			setupMock: func() {
				muRepo.EXPECT().
					GetUser(gomock.Any(), gomock.Eq(uid)).
					Return(expectedUser, nil)
			},
			expectedUser:  expectedUser,
			expectedError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMock()

			returnedUser, returnedError := userService.GetUser(ctx, uid)

			assert.Equal(t, tt.expectedUser, returnedUser)
			assert.Equal(t, tt.expectedError, returnedError)
		})
	}
}
