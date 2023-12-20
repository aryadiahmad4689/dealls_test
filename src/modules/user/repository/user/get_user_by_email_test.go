package user

import (
	"context"
	"database/sql"
	"regexp"
	"testing"

	"github.com/aryadiahmad4689/dealls_test/src/app"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestUser_GetUserByEmail(t *testing.T) {
	query := `SELECT id, email, password, name, age, gender 
	FROM users WHERE email = $1`
	type args struct {
		ctx   context.Context
		email string
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
				ctx:   context.TODO(),
				email: "test@example.com",
			},
			setup: func(mock sqlmock.Sqlmock, a args) {
				rows := sqlmock.NewRows([]string{"id", "email", "password", "name", "age", "gender"}).
					AddRow(1, "test@example.com", "hashedpassword", "Test User", 30, "Male")

				mock.ExpectQuery(regexp.QuoteMeta(query)).
					WithArgs(a.email).
					WillReturnRows(rows)
			},
			wantErr: false,
		},
		{
			name: "Failure - No User Found",
			args: args{
				ctx:   context.TODO(),
				email: "nonexistent@example.com",
			},
			setup: func(mock sqlmock.Sqlmock, a args) {
				mock.ExpectQuery(regexp.QuoteMeta(query)).
					WithArgs(a.email).
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

			td := &User{
				db: app.Db{Slave: db},
			}

			// Set up the mock expectations
			tt.setup(mock, tt.args)

			// Call the function
			_, err = td.GetUserByEmail(tt.args.ctx, tt.args.email)

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
