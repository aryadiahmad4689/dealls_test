package usecase

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/aryadiahmad4689/dealls_test/src/app"
	"github.com/aryadiahmad4689/dealls_test/src/modules/swipe/entity"
	"github.com/aryadiahmad4689/dealls_test/src/modules/swipe/repository"
	mockswipe "github.com/aryadiahmad4689/dealls_test/src/modules/swipe/repository/swipe/mocks"
	mockuser "github.com/aryadiahmad4689/dealls_test/src/modules/swipe/repository/user/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestUseCase_SwipeRight(t *testing.T) {
	db, _, _ := sqlmock.New()
	var apps = &app.App{}

	type args struct {
		ctx   context.Context
		swipe entity.Swipe
	}

	tests := []struct {
		name    string
		args    args
		wantErr error
		mockDb  func(a args) *repository.Repository
	}{
		{
			name: "already swiped",
			args: args{
				ctx: context.Background(),
				swipe: entity.Swipe{
					SwipeUserId: "user_123",
					Id:          "swipee_456",
				},
			},
			wantErr: errors.New("already swipped"),
			mockDb: func(a args) *repository.Repository {
				dbs := &app.Db{Master: db, Slave: db}
				apps.SetDb(dbs)
				mockSwipe := new(mockswipe.SwipeInterfaceMock)
				mockUser := new(mockuser.UserInterfaceMock)

				mockSwipe.On("GetUserSwipe", mock.Anything, mock.Anything).Return(entity.Swipe{}, nil) // Swipe found
				repo := repository.Init(apps).SetSwipe(mockUser, mockSwipe)
				return repo
			},
		},
		{
			name: "store swipe error",
			args: args{
				ctx: context.Background(),
				swipe: entity.Swipe{
					SwipeUserId: "user_123",
					Id:          "swipee_456",
				},
			},
			wantErr: errors.New("store swipe error"),
			mockDb: func(a args) *repository.Repository {
				dbs := &app.Db{Master: db, Slave: db}
				apps.SetDb(dbs)
				mockSwipe := new(mockswipe.SwipeInterfaceMock)
				mockUser := new(mockuser.UserInterfaceMock)

				mockSwipe.On("GetUserSwipe", mock.Anything, mock.Anything).Return(entity.Swipe{}, sql.ErrNoRows)                 // No previous swipe
				mockSwipe.On("StoreSwipe", mock.Anything, mock.Anything).Return(entity.Swipe{}, errors.New("store swipe error")) // Error on storing swipe
				repo := repository.Init(apps).SetSwipe(mockUser, mockSwipe)
				return repo
			},
		},
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				swipe: entity.Swipe{
					SwipeUserId: "user_123",
					Id:          "swipee_456",
				},
			},
			wantErr: nil,
			mockDb: func(a args) *repository.Repository {
				dbs := &app.Db{Master: db, Slave: db}
				apps.SetDb(dbs)
				mockSwipe := new(mockswipe.SwipeInterfaceMock)
				mockUser := new(mockuser.UserInterfaceMock)

				mockSwipe.On("GetUserSwipe", mock.Anything, mock.Anything).Return(entity.Swipe{}, sql.ErrNoRows) // No previous swipe
				mockSwipe.On("StoreSwipe", mock.Anything, mock.Anything).Return(entity.Swipe{}, nil)             // Successful store
				repo := repository.Init(apps).SetSwipe(mockUser, mockSwipe)
				return repo
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := tt.mockDb(tt.args)
			u := Init(apps, repo)
			err := u.SwipeRight(tt.args.ctx, tt.args.swipe)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
