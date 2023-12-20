package swipe

import (
	"context"
	"database/sql"
	"regexp"
	"testing"
	"time"

	"github.com/aryadiahmad4689/dealls_test/src/app"
	"github.com/aryadiahmad4689/dealls_test/src/modules/swipe/entity"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestSwipe_GetSwipeAbleUser(t *testing.T) {
	query := `
	SELECT id, name, age, gender 
	FROM users 
	WHERE id != $1 
	  AND gender != (
		  SELECT gender 
		  FROM users 
		  WHERE id = $1
	  )
	  AND id NOT IN (
		  SELECT is_swipe_user_id 
		  FROM swipes 
		  WHERE swipe_user_id = $1 AND date(created_at) = $2
	)
	LIMIT 10`
	type args struct {
		ctx context.Context
		req entity.GetSwipeReq
	}
	tests := []struct {
		name    string
		args    args
		want    []entity.User
		wantErr bool
		setup   func(mock sqlmock.Sqlmock, a args)
	}{
		{
			name: "Success",
			args: args{
				ctx: context.TODO(),
				req: entity.GetSwipeReq{
					UserId: "1",
					Date:   time.Now(),
				},
			},
			want: []entity.User{
				{
					Id:         "2",
					Name:       "Test User",
					Gender:     "Female",
					Age:        "30",
					IsVerified: 0,
				},
			},
			setup: func(mock sqlmock.Sqlmock, a args) {
				formattedDate := a.req.Date.Format("2006-01-02")
				rows := sqlmock.NewRows([]string{"id", "name", "age", "gender"}).
					AddRow("2", "Test User", 30, "Female") // Contoh user yang muncul

				mock.ExpectQuery(regexp.QuoteMeta(query)).
					WithArgs(a.req.UserId, formattedDate).
					WillReturnRows(rows)
			},
			wantErr: false,
		},
		{
			name: "Failure - DB Error",
			args: args{
				ctx: context.TODO(),
				req: entity.GetSwipeReq{
					UserId: "1",
					Date:   time.Now(),
				},
			},
			wantErr: true,
			setup: func(mock sqlmock.Sqlmock, a args) {
				formattedDate := a.req.Date.Format("2006-01-02")

				mock.ExpectQuery(regexp.QuoteMeta(query)).
					WithArgs(a.req.UserId, formattedDate).
					WillReturnError(sql.ErrNoRows) // atau error lain yang relevan
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			assert.NoError(t, err)
			defer db.Close()

			s := &Swipe{
				db: app.Db{Slave: db},
			}

			// Set up the mock expectations
			tt.setup(mock, tt.args)

			// Call the function
			got, err := s.GetSwipeAbleUser(tt.args.ctx, tt.args.req)

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
