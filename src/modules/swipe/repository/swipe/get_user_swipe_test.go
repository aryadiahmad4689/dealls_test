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

func TestSwipe_GetUserSwipe(t *testing.T) {
	// Define the query
	query := `SELECT id, swipe_user_id, is_swipe_user_id, swipe_type FROM swipes WHERE swipe_user_id = $1 AND is_swipe_user_id = $2 AND date(created_at) = $3`

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
					SwipeUserId:   "1",
					IsSwipeUserId: "2",
					CreatedAt:     time.Now(),
				},
			},
			want: entity.Swipe{
				Id:            "1",
				SwipeUserId:   "1",
				IsSwipeUserId: "2",
				SwipeType:     "like",
			},
			setup: func(mock sqlmock.Sqlmock, a args) {
				formattedDate := a.req.CreatedAt.Format("2006-01-02")
				rows := sqlmock.NewRows([]string{"id", "swipe_user_id", "is_swipe_user_id", "swipe_type"}).
					AddRow("1", "1", "2", "like") // Contoh swipe data

				mock.ExpectQuery(regexp.QuoteMeta(query)).
					WithArgs(a.req.SwipeUserId, a.req.IsSwipeUserId, formattedDate).
					WillReturnRows(rows)
			},
			wantErr: false,
		},
		{
			name: "Failure - No Swipe Found",
			args: args{
				ctx: context.TODO(),
				req: entity.Swipe{
					SwipeUserId:   "1",
					IsSwipeUserId: "2",
					CreatedAt:     time.Now(),
				},
			},
			wantErr: true,
			setup: func(mock sqlmock.Sqlmock, a args) {
				formattedDate := a.req.CreatedAt.Format("2006-01-02")

				mock.ExpectQuery(regexp.QuoteMeta(query)).
					WithArgs(a.req.SwipeUserId, a.req.IsSwipeUserId, formattedDate).
					WillReturnError(sql.ErrNoRows)
			},
		},
		// ... (Add more test cases if needed)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			assert.NoError(t, err)
			defer db.Close()

			s := &Swipe{
				db: app.Db{Master: db},
			}

			// Setup mock expectations
			tt.setup(mock, tt.args)

			// Call the function
			got, err := s.GetUserSwipe(tt.args.ctx, tt.args.req)

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
