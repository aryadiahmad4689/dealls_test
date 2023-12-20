package usecase

import (
	"context"
	"database/sql"
	"errors"
	"os"
	"testing"

	"github.com/aryadiahmad4689/dealls_test/src/app"
	"github.com/aryadiahmad4689/dealls_test/src/modules/user/entity"
	"github.com/aryadiahmad4689/dealls_test/src/modules/user/repository"
	"github.com/aryadiahmad4689/dealls_test/src/modules/user/repository/user/mocks"
	"github.com/stretchr/testify/mock"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"gopkg.in/go-playground/assert.v1"
)

func TestUseCase_SignIn(t *testing.T) {
	var apps = &app.App{}
	db, _, _ := sqlmock.New()
	type args struct {
		ctx  context.Context
		user entity.User
	}

	tests := []struct {
		name    string
		args    args
		wantErr error
		mockDb  func(a args) *repository.Repository
	}{
		// TODO: Add test cases.
		{
			name: "email not found",
			args: args{ctx: context.Background(), user: entity.User{
				Name: "name",
			}},
			wantErr: errors.New("email not found"),
			mockDb: func(a args) *repository.Repository {
				dbs := &app.Db{
					Master: db,
					Slave:  db,
				}
				apps.SetDb(dbs)
				mockUser := new(mocks.UserInterfaceMock)
				// $2a$10$zpuDPKg/OqHjrDvglHYAx.d3z8OjQOctsO5mdFHObfxVIxTGFYITm
				mockUser.On("GetUserByEmail", mock.Anything, mock.Anything).Return(entity.User{}, sql.ErrNoRows)
				repo := repository.Init(apps).SetUser(mockUser)
				return repo
			},
		},
		{
			name: "invalid password",
			args: args{ctx: context.Background(), user: entity.User{
				Name:     "name",
				Password: "pass",
			}},
			wantErr: errors.New("invalid password"),
			mockDb: func(a args) *repository.Repository {
				dbs := &app.Db{
					Master: db,
					Slave:  db,
				}
				apps.SetDb(dbs)
				mockUser := new(mocks.UserInterfaceMock)
				// $2a$10$zpuDPKg/OqHjrDvglHYAx.d3z8OjQOctsO5mdFHObfxVIxTGFYITm
				mockUser.On("GetUserByEmail", mock.Anything, mock.Anything).Return(entity.User{}, nil)
				repo := repository.Init(apps).SetUser(mockUser)
				return repo
			},
		},
		{
			name: "success",
			args: args{ctx: context.Background(), user: entity.User{
				Name:     "name",
				Password: "mySecurePassword",
			}},
			wantErr: nil,
			mockDb: func(a args) *repository.Repository {
				dbs := &app.Db{
					Master: db,
					Slave:  db,
				}
				apps.SetDb(dbs)
				mockUser := new(mocks.UserInterfaceMock)
				// $2a$10$zpuDPKg/OqHjrDvglHYAx.d3z8OjQOctsO5mdFHObfxVIxTGFYITm
				os.Setenv("AUTH_KEY", "JWTseqret")

				mockUser.On("GetUserByEmail", mock.Anything, mock.Anything).Return(entity.User{
					Password: "$2a$10$zpuDPKg/OqHjrDvglHYAx.d3z8OjQOctsO5mdFHObfxVIxTGFYITm",
				}, nil)
				repo := repository.Init(apps).SetUser(mockUser)
				return repo
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := tt.mockDb(tt.args)
			u := Init(apps, repo)
			_, err := u.SignIn(tt.args.ctx, tt.args.user)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
