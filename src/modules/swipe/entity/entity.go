package entity

import "time"

type User struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Gender     string `json:"gender"`
	Age        string `json:"age"`
	IsVerified int    `json:"is_verified"`
}

type GetSwipeReq struct {
	UserId string
	Date   time.Time
}

type Swipe struct {
	Id            string
	SwipeUserId   string
	IsSwipeUserId string
	SwipeType     string
	CreatedAt     time.Time
}
