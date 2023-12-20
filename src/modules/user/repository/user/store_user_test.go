package user

import (
	"context"
	"database/sql"
	"regexp"
	"testing"

	"github.com/aryadiahmad4689/dealls_test/src/app"
	"github.com/aryadiahmad4689/dealls_test/src/modules/user/entity"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestUser_StoreUser(t *testing.T) {
	query := `INSERT INTO users (email, password, name, age, gender,is_verified, created_at, updated_at) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id, email, name, age, gender`
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
				ctx:  context.TODO(),
				user: entity.User{Email: "test@example.com", Password: "hashedpassword", Name: "Test User", Age: 30, Gender: "Male", IsVerified: 1},
			},
			setup: func(mock sqlmock.Sqlmock, a args) {
				rows := sqlmock.NewRows([]string{"id", "email", "name", "age", "gender"}).
					AddRow(1, "test@example.com", "Test User", 30, "Male")

				mock.ExpectQuery(regexp.QuoteMeta(query)).
					WithArgs(a.user.Email, a.user.Password, a.user.Name, a.user.Age, a.user.Gender, a.user.IsVerified, sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnRows(rows)
			},
			wantErr: false,
		},
		{
			name: "Failure - DB Error",
			args: args{
				ctx:  context.TODO(),
				user: entity.User{Email: "test@example.com", Password: "hashedpassword", Name: "Test User", Age: 30, Gender: "Male", IsVerified: 0},
			},
			setup: func(mock sqlmock.Sqlmock, a args) {
				mock.ExpectQuery(regexp.QuoteMeta(query)).
					WithArgs(a.user.Email, a.user.Password, a.user.Name, a.user.Age, a.user.Gender, a.user.IsVerified, sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnError(sql.ErrConnDone) // atau error lain yang relevan
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

			u := &User{
				db: app.Db{Master: db}, // Inisialisasi struct User dengan mock DB
			}

			// Set up the mock expectations
			tt.setup(mock, tt.args)

			// Call the function
			_, err = u.StoreUser(tt.args.ctx, tt.args.user)

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
