package swipe

import (
	"context"

	"github.com/aryadiahmad4689/dealls_test/src/modules/swipe/entity"
)

func (s *Swipe) GetSwipeAbleUser(ctx context.Context, req entity.GetSwipeReq) ([]entity.User, error) {
	// Query untuk mendapatkan profil yang belum di-swipe
	// saya asumsikan disini male hanya bisa kencan dengan female dan begitupun sebaliknya
	formattedDate := req.Date.Format("2006-01-02")
	rows, err := s.db.Slave.Query(`
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
	LIMIT 10`, req.UserId, formattedDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var profiles []entity.User
	for rows.Next() {
		var p entity.User
		err := rows.Scan(&p.Id, &p.Name, &p.Age, &p.Gender)
		if err != nil {
			return nil, err
		}
		profiles = append(profiles, p)
	}

	return profiles, nil
}
