package model

import "time"

// UserCoupon model
type UserCoupon struct {
	Code        string
	Title       string
	Description string
	Discount    float64
	Currency    string
	ValidUntil  time.Time
}

// ListUserCouponsResponse model
type ListUserCouponsResponse struct {
	Coupons []UserCoupon
}
