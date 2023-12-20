package entity

import "time"

type Package struct {
	Id               int       `json:"id"`
	SubscriptionType string    `json:"subscription_type"`
	SubscriptionLong int       `json:"subscription_long"`
	Price            float64   `json:"price"`
	CreatedAt        time.Time `json:"-"`
	UpdatedAt        time.Time `json:"-"`
}
