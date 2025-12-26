package model

import "time"

// Order model สำหรับใช้งานภายใน service
type Order struct {
	OrderID     string
	UserID      string
	Pickup      string // หรือสร้าง struct TravelLocation ถ้าต้องเก็บ lat/lng
	Dropoff     string
	Status      string
	Estimated   float64
	FinalTotal  float64
	Discount    float64
	CouponCode  string
	DriverID    string
	DriverName  string
	DriverCar   string
	MatchResult *MatchResult // สร้าง struct ของตัวเองแทน proto.MatchResult
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// MatchResult struct แทน pb.MatchResult
type MatchResult struct {
	Status   string
	DriverID string
	// เพิ่ม field ที่ต้องการเก็บ
}

// NewOrder สร้าง Order ใหม่
func NewOrder(orderID, userID, pickup, dropoff string, estimated float64) *Order {
	return &Order{
		OrderID:   orderID,
		UserID:    userID,
		Pickup:    pickup,
		Dropoff:   dropoff,
		Status:    "waiting_coupon",
		Estimated: estimated,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

// UpdateMatchResult อัพเดท match result และสถานะ
func (o *Order) UpdateMatchResult(match *MatchResult) {
	o.MatchResult = match
	o.Status = match.Status
	o.UpdatedAt = time.Now()
}

// ApplyCoupon อัพเดทราคาหลังใช้คูปอง
func (o *Order) ApplyCoupon(finalTotal, discount float64, couponCode string) {
	o.FinalTotal = finalTotal
	o.Discount = discount
	o.CouponCode = couponCode
	o.UpdatedAt = time.Now()
}

// SetDriverInfo
func (o *Order) SetDriver(driverID, driverName, driverCar string) {
	o.DriverID = driverID
	o.DriverName = driverName
	o.DriverCar = driverCar
	o.UpdatedAt = time.Now()
}

// SetStatus อัพเดท status
func (o *Order) SetStatus(status string) {
	o.Status = status
	o.UpdatedAt = time.Now()
}
