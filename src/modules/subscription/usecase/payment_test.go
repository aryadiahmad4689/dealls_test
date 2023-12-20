package usecase

import (
	"context"
	"errors"
	"testing"

	"gopkg.in/DATA-DOG/go-sqlmock.v1"

	"github.com/aryadiahmad4689/dealls_test/src/app"
	"github.com/aryadiahmad4689/dealls_test/src/modules/subscription/entity"
	"github.com/aryadiahmad4689/dealls_test/src/modules/subscription/repository"
	"github.com/aryadiahmad4689/dealls_test/src/modules/subscription/repository/packages/mocks"
	mockspackage "github.com/aryadiahmad4689/dealls_test/src/modules/subscription/repository/packages/mocks"
	mocksubscription "github.com/aryadiahmad4689/dealls_test/src/modules/subscription/repository/subscription/mocks"
	mocksuser "github.com/aryadiahmad4689/dealls_test/src/modules/subscription/repository/user/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUseCase_Payment(t *testing.T) {
	var apps = &app.App{}
	db, _, _ := sqlmock.New()

	type args struct {
		ctx context.Context
		req entity.Subscription
	}

	tests := []struct {
		name    string
		args    args
		wantErr error
		mockDb  func(a args) *repository.Repository
	}{
		{
			name: "package not found",
			args: args{
				ctx: context.Background(),
				req: entity.Subscription{
					PackageId: "123",
					Pricing:   100,
					UserId:    "user123",
				},
			},
			wantErr: errors.New("package not found"),
			mockDb: func(a args) *repository.Repository {
				dbs := &app.Db{
					Master: db,
					Slave:  db,
				}
				apps.SetDb(dbs)
				mockPackage := new(mockspackage.PackageInterfaceMock)
				mockPackage.On("GetPackageById", mock.Anything, mock.Anything).Return(entity.Package{}, errors.New("package not found"))
				repo := repository.Init(apps).SetPackage(mockPackage)
				return repo
			},
		},
		{
			name: "your money not enough",
			args: args{
				ctx: context.Background(),
				req: entity.Subscription{
					PackageId: "123",
					Pricing:   50, // Less than package price
					UserId:    "user123",
				},
			},
			wantErr: errors.New("your money not enough"),
			mockDb: func(a args) *repository.Repository {
				mockPackage := new(mocks.PackageInterfaceMock)
				mockPackage.On("GetPackageById", mock.Anything, mock.Anything).Return(entity.Package{
					Price: 100,
				}, nil)
				repo := repository.Init(apps).SetPackage(mockPackage)
				return repo
			},
		},
		{
			name: "failed subscription",
			args: args{
				ctx: context.Background(),
				req: entity.Subscription{
					PackageId: "123",
					Pricing:   100,
					UserId:    "user123",
				},
			},
			wantErr: errors.New("failed subscription"),
			mockDb: func(a args) *repository.Repository {
				mockPackage := new(mockspackage.PackageInterfaceMock)
				mockPackage.On("GetPackageById", mock.Anything, mock.Anything).Return(entity.Package{
					Price: 100,
				}, nil)

				mockSubscription := new(mocksubscription.SubscriptionInterfaceMock)
				mockSubscription.On("StoreSubcription", mock.Anything, mock.Anything).Return(entity.Subscription{}, errors.New("err"))

				mockUser := new(mocksuser.UserInterfaceMock)

				repo := repository.Init(apps).SetPackage(mockPackage).SetSubsription(mockSubscription).SetUser(mockUser)
				return repo
			},
		},
		{
			name: "failed update user",
			args: args{
				ctx: context.Background(),
				req: entity.Subscription{
					PackageId: "123",
					Pricing:   100,
					UserId:    "user123",
				},
			},
			wantErr: errors.New("failed update user"),
			mockDb: func(a args) *repository.Repository {
				mockPackage := new(mockspackage.PackageInterfaceMock)
				mockPackage.On("GetPackageById", mock.Anything, mock.Anything).Return(entity.Package{
					Price: 100,
				}, nil)

				mockSubscription := new(mocksubscription.SubscriptionInterfaceMock)
				mockSubscription.On("StoreSubcription", mock.Anything, mock.Anything).Return(entity.Subscription{}, nil)

				mockUser := new(mocksuser.UserInterfaceMock)
				mockUser.On("UpdateUser", mock.Anything, mock.Anything).Return(errors.New("failed update user"))

				repo := repository.Init(apps).SetPackage(mockPackage).SetSubsription(mockSubscription).SetUser(mockUser)
				return repo
			},
		},
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				req: entity.Subscription{
					PackageId: "123",
					Pricing:   100,
					UserId:    "user123",
				},
			},
			wantErr: nil,
			mockDb: func(a args) *repository.Repository {
				mockPackage := new(mockspackage.PackageInterfaceMock)
				mockPackage.On("GetPackageById", mock.Anything, mock.Anything).Return(entity.Package{
					Price: 100,
				}, nil)

				mockSubscription := new(mocksubscription.SubscriptionInterfaceMock)
				mockSubscription.On("StoreSubcription", mock.Anything, mock.Anything).Return(entity.Subscription{}, nil)

				mockUser := new(mocksuser.UserInterfaceMock)
				mockUser.On("UpdateUser", mock.Anything, mock.Anything).Return(nil)

				repo := repository.Init(apps).SetPackage(mockPackage).SetSubsription(mockSubscription).SetUser(mockUser)
				return repo
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := tt.mockDb(tt.args)
			u := Init(apps, repo)
			err := u.Payment(tt.args.ctx, tt.args.req)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
