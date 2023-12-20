package entity

import "time"

type Subscription struct {
	Id        string
	UserId    string
	PackageId string
	StartDate time.Time
	EndDate   time.Time
	Pricing   float64
}

type User struct {
	Id         string
	IsVerified int
}

type Package struct {
	Id               string  `json:"id"`
	SubscriptionType string  `json:"subscription_type"`
	SubscriptionLong int     `json:"subscription_long"`
	Price            float64 `json:"price"`
}
