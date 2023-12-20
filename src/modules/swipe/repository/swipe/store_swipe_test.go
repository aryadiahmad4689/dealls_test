package swipe

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"regexp"
	"testing"
	"time"

	"github.com/aryadiahmad4689/dealls_test/src/app"
	"github.com/aryadiahmad4689/dealls_test/src/modules/swipe/entity"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

type AnyTime struct{}

func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

func TestSwipe_StoreSwipe(t *testing.T) {
	query := `INSERT INTO swipes (swipe_user_id, is_swipe_user_id, swipe_type,created_at,updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	type args struct {
		ctx context.Context
		req entity.Swipe
	}
	tests := []struct {
		name    string
		args    args
		want    entity.Swipe
		wantErr bool
		setup   func(mock sqlmock.Sqlmock, a args)
	}{

		{
			name: "Success",
			args: args{
				ctx: context.TODO(),
				req: entity.Swipe{
					SwipeUserId:   "user1",
					IsSwipeUserId: "user2",
					SwipeType:     "like",
				},
			},

			setup: func(mock sqlmock.Sqlmock, a args) {
				rows := sqlmock.NewRows([]string{"id"}).AddRow("1")
				mock.ExpectQuery(regexp.QuoteMeta(query)).
					WithArgs(a.req.SwipeUserId, a.req.IsSwipeUserId, a.req.SwipeType, AnyTime{}, AnyTime{}).
					WillReturnRows(rows)
			},
			wantErr: false,
		},
		{
			name: "Failed",
			args: args{
				ctx: context.TODO(),
				req: entity.Swipe{
					SwipeUserId:   "user1",
					IsSwipeUserId: "user2",
					SwipeType:     "like",
				},
			},

			setup: func(mock sqlmock.Sqlmock, a args) {
				sqlmock.NewRows([]string{"id"}).AddRow("1") // Contoh ID yang dihasilkan
				mock.ExpectQuery(regexp.QuoteMeta(query)).
					WithArgs(a.req.SwipeUserId, a.req.IsSwipeUserId, a.req.SwipeType, AnyTime{}, AnyTime{}).
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

			s := &Swipe{
				db: app.Db{Master: db},
			}

			// Set up the mock expectations
			tt.setup(mock, tt.args)

			// Call the function
			_, err = s.StoreSwipe(tt.args.ctx, tt.args.req)

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
