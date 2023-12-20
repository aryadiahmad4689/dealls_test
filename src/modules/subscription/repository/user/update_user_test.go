package user

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

func TestUser_UpdateUser(t *testing.T) {
	query := `UPDATE users SET is_verified = $1 WHERE id = $2`
	type args struct {
		ctx  context.Context
		user entity.User
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
				user: entity.User{
					Id:         "1",
					IsVerified: 1,
				},
			},
			setup: func(mock sqlmock.Sqlmock, a args) {
				mock.ExpectExec(regexp.QuoteMeta(query)).
					WithArgs(a.user.IsVerified, a.user.Id).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			wantErr: false,
		},
		{
			name: "Failure - DB Error",
			args: args{
				ctx: context.TODO(),
				user: entity.User{
					Id:         "1",
					IsVerified: 0,
				},
			},
			setup: func(mock sqlmock.Sqlmock, a args) {
				mock.ExpectExec(regexp.QuoteMeta(query)).
					WithArgs(a.user.IsVerified, a.user.Id).
					WillReturnError(sql.ErrNoRows) // atau error lain yang relevan
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

			ud := &User{
				db: app.Db{Master: db},
			}

			// Set up the mock expectations
			tt.setup(mock, tt.args)

			// Call the function
			err = ud.UpdateUser(tt.args.ctx, tt.args.user)

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
