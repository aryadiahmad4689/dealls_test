package packages

import (
	"context"
	"database/sql"
	"regexp"
	"testing"

	"github.com/aryadiahmad4689/dealls_test/src/app"
	"github.com/aryadiahmad4689/dealls_test/src/modules/subscription/entity"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestPackage_GetPackageById(t *testing.T) {
	query := `SELECT id, subscription_type, subscription_long, price FROM packages where id = $1`
	type args struct {
		ctx context.Context
		req entity.Package
	}
	tests := []struct {
		name    string
		args    args
		want    entity.Package
		wantErr bool
		setup   func(mock sqlmock.Sqlmock, a args)
	}{
		{
			name: "Success",
			args: args{
				ctx: context.TODO(),
				req: entity.Package{Id: "1"},
			},
			want: entity.Package{Id: "1", SubscriptionType: "Basic", SubscriptionLong: 30, Price: 9.99},
			setup: func(mock sqlmock.Sqlmock, a args) {
				rows := sqlmock.NewRows([]string{"id", "subscription_type", "subscription_long", "price"}).
					AddRow(a.req.Id, "Basic", 30, 9.99)

				mock.ExpectQuery(regexp.QuoteMeta(query)).
					WithArgs(a.req.Id).
					WillReturnRows(rows)
			},
			wantErr: false,
		},
		{
			name: "Failure - No Package Found",
			args: args{
				ctx: context.TODO(),
				req: entity.Package{Id: "99"},
			},
			setup: func(mock sqlmock.Sqlmock, a args) {
				mock.ExpectQuery(regexp.QuoteMeta(query)).
					WithArgs(a.req.Id).
					WillReturnError(sql.ErrNoRows)
			},
			wantErr: true,
		},
		
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			assert.NoError(t, err)
			defer db.Close()

			p := &Package{
				db: app.Db{Slave: db},
			}

			// Set up the mock expectations
			tt.setup(mock, tt.args)

			// Call the function
			got, err := p.GetPackageById(tt.args.ctx, tt.args.req)

			// Assertions
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
