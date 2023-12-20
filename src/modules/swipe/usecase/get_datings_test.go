package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/aryadiahmad4689/dealls_test/src/app"
	"github.com/aryadiahmad4689/dealls_test/src/modules/swipe/entity"
	"github.com/aryadiahmad4689/dealls_test/src/modules/swipe/repository"
	mockswipe "github.com/aryadiahmad4689/dealls_test/src/modules/swipe/repository/swipe/mocks"
	mockuser "github.com/aryadiahmad4689/dealls_test/src/modules/swipe/repository/user/mocks"

	"github.com/stretchr/testify/mock"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"gopkg.in/go-playground/assert.v1"
)

func TestUseCase_GetDatings(t *testing.T) {
	db, _, _ := sqlmock.New()
	var apps = &app.App{}
	type args struct {
		ctx context.Context
	}

	tests := []struct {
		name    string
		args    args
		wantErr error
		mockDb  func(a args) *repository.Repository
	}{
		{
			name:    "swipe error",
			args:    args{ctx: context.WithValue(context.Background(), "user_id", "test_user_id")},
			wantErr: errors.New("err"),
			mockDb: func(a args) *repository.Repository {
				dbs := &app.Db{Master: db, Slave: db}
				apps.SetDb(dbs)
				mockSwipe := new(mockswipe.SwipeInterfaceMock)
				mockUser := new(mockuser.UserInterfaceMock)
				mockSwipe.On("GetCountSwipe", mock.Anything, mock.AnythingOfType("entity.GetSwipeReq")).Return(10, errors.New("err"))
				repo := repository.Init(apps).SetSwipe(mockUser, mockSwipe)
				return repo
			},
		},
		{
			name:    "user not found",
			args:    args{ctx: context.WithValue(context.Background(), "user_id", "test_user_id")},
			wantErr: errors.New("user not found"),
			mockDb: func(a args) *repository.Repository {
				dbs := &app.Db{Master: db, Slave: db}
				apps.SetDb(dbs)
				mockSwipe := new(mockswipe.SwipeInterfaceMock)
				mockUser := new(mockuser.UserInterfaceMock)
				mockSwipe.On("GetCountSwipe", mock.Anything, mock.AnythingOfType("entity.GetSwipeReq")).Return(10, nil)
				mockUser.On("GetUserById", mock.Anything, "test_user_id").Return(entity.User{}, errors.New("user not found"))
				repo := repository.Init(apps).SetSwipe(mockUser, mockSwipe)
				return repo
			},
		},
		{
			name:    "swipe limit reached",
			args:    args{ctx: context.WithValue(context.Background(), "user_id", "test_user_id")},
			wantErr: errors.New("swipe today is done"),
			mockDb: func(a args) *repository.Repository {
				dbs := &app.Db{Master: db, Slave: db}
				apps.SetDb(dbs)
				mockSwipe := new(mockswipe.SwipeInterfaceMock)
				mockUser := new(mockuser.UserInterfaceMock)
				mockSwipe.On("GetCountSwipe", mock.Anything, mock.AnythingOfType("entity.GetSwipeReq")).Return(11, nil)
				mockUser.On("GetUserById", mock.Anything, "test_user_id").Return(entity.User{IsVerified: 0}, nil)
				repo := repository.Init(apps).SetSwipe(mockUser, mockSwipe)
				return repo
			},
		},
		{
			name:    "swipeable user error",
			args:    args{ctx: context.WithValue(context.Background(), "user_id", "test_user_id")},
			wantErr: errors.New("err"),
			mockDb: func(a args) *repository.Repository {
				dbs := &app.Db{Master: db, Slave: db}
				apps.SetDb(dbs)
				mockSwipe := new(mockswipe.SwipeInterfaceMock)
				mockUser := new(mockuser.UserInterfaceMock)
				mockSwipe.On("GetCountSwipe", mock.Anything, mock.AnythingOfType("entity.GetSwipeReq")).Return(7, nil)
				mockUser.On("GetUserById", mock.Anything, "test_user_id").Return(entity.User{IsVerified: 0}, nil)
				mockSwipe.On("GetSwipeAbleUser", mock.Anything, mock.Anything).Return([]entity.User{}, errors.New("err"))

				repo := repository.Init(apps).SetSwipe(mockUser, mockSwipe)
				return repo
			},
		},
		{
			name:    "success",
			args:    args{ctx: context.WithValue(context.Background(), "user_id", "test_user_id")},
			wantErr: nil,
			mockDb: func(a args) *repository.Repository {
				dbs := &app.Db{Master: db, Slave: db}
				apps.SetDb(dbs)
				mockSwipe := new(mockswipe.SwipeInterfaceMock)
				mockUser := new(mockuser.UserInterfaceMock)
				mockSwipe.On("GetCountSwipe", mock.Anything, mock.AnythingOfType("entity.GetSwipeReq")).Return(7, nil)
				mockUser.On("GetUserById", mock.Anything, "test_user_id").Return(entity.User{IsVerified: 0}, nil)
				mockSwipe.On("GetSwipeAbleUser", mock.Anything, mock.Anything).Return([]entity.User{}, nil)

				repo := repository.Init(apps).SetSwipe(mockUser, mockSwipe)
				return repo
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := tt.mockDb(tt.args)
			u := Init(apps, repo)
			_, err := u.GetDatings(tt.args.ctx)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
