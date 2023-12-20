package usecase

import (
	"context"
	"errors"
	"testing"

	"gopkg.in/DATA-DOG/go-sqlmock.v1"

	"github.com/aryadiahmad4689/dealls_test/src/app"
	"github.com/aryadiahmad4689/dealls_test/src/modules/packages/entity"
	"github.com/aryadiahmad4689/dealls_test/src/modules/packages/repository"
	"github.com/aryadiahmad4689/dealls_test/src/modules/packages/repository/packages/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUseCase_GetPackage(t *testing.T) {
	var apps = &app.App{}
	db, _, _ := sqlmock.New()

	type args struct {
		ctx context.Context
	}

	tests := []struct {
		name    string
		args    args
		want    []entity.Package
		wantErr bool
		mockDb  func() *repository.Repository
	}{
		{
			name:    "success",
			args:    args{ctx: context.Background()},
			want:    []entity.Package{{Id: 1, SubscriptionType: "Package 1"}, {Id: 2, SubscriptionType: "Package 2"}},
			wantErr: false,
			mockDb: func() *repository.Repository {
				dbs := &app.Db{
					Master: db,
					Slave:  db,
				}
				apps.SetDb(dbs)
				mockPackage := new(mocks.PackageInterfaceMock)
				mockPackage.On("GetPackage", mock.Anything).Return([]entity.Package{{Id: 1, SubscriptionType: "Package 1"}, {Id: 2, SubscriptionType: "Package 2"}}, nil)
				repo := repository.Init(apps).SetPackage(mockPackage)
				return repo
			},
		},
		{
			name:    "error",
			args:    args{ctx: context.Background()},
			want:    []entity.Package{},
			wantErr: true,
			mockDb: func() *repository.Repository {
				dbs := &app.Db{
					Master: db,
					Slave:  db,
				}
				apps.SetDb(dbs)
				mockPackage := new(mocks.PackageInterfaceMock)
				mockPackage.On("GetPackage", mock.Anything).Return([]entity.Package{}, errors.New("some error"))
				repo := repository.Init(apps).SetPackage(mockPackage)
				return repo
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := tt.mockDb()
			u := Init(apps, repo)
			got, err := u.GetPackage(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("UseCase.GetPackage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
