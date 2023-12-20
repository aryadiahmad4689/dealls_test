package subscription

import (
	"context"
	"errors"
	"regexp"
	"testing"
	"time"

	"github.com/aryadiahmad4689/dealls_test/src/app"
	"github.com/aryadiahmad4689/dealls_test/src/modules/subscription/entity"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestSubscription_StoreSubscription(t *testing.T) {
	query := `INSERT INTO subscriptions (user_id, packages_id, StartDate, EndDate, created_at, updated_at) 
          VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`

	type args struct {
		ctx          context.Context
		subscription entity.Subscription
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		setup   func(mock sqlmock.Sqlmock, a args)
	}{
		{
			name: "Success",
			args: args{
				ctx: context.TODO(),
				subscription: entity.Subscription{
					UserId:    "1",
					PackageId: "1",
					StartDate: time.Now(),
					EndDate:   time.Now().Add(30 * 24 * time.Hour),
				},
			},
			setup: func(mock sqlmock.Sqlmock, a args) {
				rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
				mock.ExpectQuery(regexp.QuoteMeta(query)).
					WithArgs(a.subscription.UserId, a.subscription.PackageId, a.subscription.StartDate, a.subscription.EndDate, sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnRows(rows)
			},
			wantErr: false,
		},
		{
			name: "Failure - SQL Error",
			args: args{
				ctx: context.TODO(),
				subscription: entity.Subscription{
					UserId:    "1",
					PackageId: "2",
					StartDate: time.Now(),
					EndDate:   time.Now().Add(30 * 24 * time.Hour),
				},
			},
			wantErr: true,
			setup: func(mock sqlmock.Sqlmock, a args) {
				mock.ExpectQuery(regexp.QuoteMeta(query)).
					WithArgs(a.subscription.UserId, a.subscription.PackageId, a.subscription.StartDate, a.subscription.EndDate, sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnError(errors.New("sql error")) // Simulasi error dari database
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			s := &Subscription{
				db: app.Db{Master: db}, // Initialize your Subscription struct here with mock DB
			}

			// Set up the mock expectations
			tt.setup(mock, tt.args)

			// Call the function
			_, err = s.StoreSubcription(tt.args.ctx, tt.args.subscription)

			// Assertions
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
