package packages

import (
	"context"
	"database/sql"
	"regexp"
	"testing"
	"time"

	"github.com/aryadiahmad4689/dealls_test/src/app"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestPackage_GetPackage(t *testing.T) {
	query := `SELECT id, subscription_type, subscription_long, price, created_at, updated_at FROM packages`
	tests := []struct {
		name    string
		wantErr bool
		setup   func(mock sqlmock.Sqlmock)
	}{
		{
			name: "Success",
			setup: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "subscription_type", "subscription_long", "price", "created_at", "updated_at"}).
					AddRow(1, "Type A", 30, 1000, time.Now(), time.Now())

				mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)
			},
			wantErr: false,
		},
		{
			name: "Failure - DB Error",
			setup: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnError(sql.ErrConnDone)
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			pkg := &Package{
				db: app.Db{Slave: db},
			}

			// Set up the mock expectations
			tt.setup(mock)

			// Call the function
			_, err = pkg.GetPackage(context.TODO())

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
