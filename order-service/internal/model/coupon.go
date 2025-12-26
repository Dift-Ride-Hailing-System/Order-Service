package model

// ApplyCouponInput ใช้ใน business logic
type ApplyCouponInput struct {
	UserID            string
	CouponCode        string
	OrderTotal        float64
	CouponPriceResult string
}

// ApplyCouponResult ผลลัพธ์จากการ apply coupon
type ApplyCouponResult struct {
	FinalTotal        float64
	Discount          float64
	Valid             bool
	Message           string
	CouponPriceResult string
}

// CouponPriceResult เก็บราคาที่คำนวณจาก coupon (precompute)
type CouponPriceResult struct {
	FinalTotal float64 `json:"final_total"`
	Discount   float64 `json:"discount"`
}
