package user

import (
	"context"
	"database/sql"
	"regexp"
	"testing"

	"github.com/aryadiahmad4689/dealls_test/src/app"
	"github.com/aryadiahmad4689/dealls_test/src/modules/swipe/entity"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestUser_GetUserById(t *testing.T) {
	query := `SELECT id, name, age, gender, is_verified FROM users WHERE id = $1`
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		args    args
		want    entity.User
		wantErr bool
		setup   func(mock sqlmock.Sqlmock, a args)
	}{
		{
			name: "Success - User Found",
			args: args{
				ctx: context.TODO(),
				id:  "1",
			},
			want: entity.User{
				Id:         "1",
				Name:       "Test User",
				Age:        "25",
				Gender:     "Male",
				IsVerified: 1,
			},
			setup: func(mock sqlmock.Sqlmock, a args) {
				rows := sqlmock.NewRows([]string{"id", "name", "age", "gender", "is_verified"}).
					AddRow("1", "Test User", 25, "Male", 1)

				mock.ExpectQuery(regexp.QuoteMeta(query)).
					WithArgs(a.id).
					WillReturnRows(rows)
			},
			wantErr: false,
		},
		{
			name: "Failure - User Not Found",
			args: args{
				ctx: context.TODO(),
				id:  "2",
			},
			wantErr: true,
			setup: func(mock sqlmock.Sqlmock, a args) {
				mock.ExpectQuery(regexp.QuoteMeta(query)).
					WithArgs(a.id).
					WillReturnError(sql.ErrNoRows)
			},
		},
		// ... (other test cases)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			assert.NoError(t, err)
			defer db.Close()

			u := &User{
				db: app.Db{Master: db},
			}

			// Setup mock
			tt.setup(mock, tt.args)

			// Call function
			got, err := u.GetUserById(tt.args.ctx, tt.args.id)

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
