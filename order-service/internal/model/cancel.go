package model

import "time"

// -----------------------------
// Cancel Event Models
// -----------------------------

// PassengerCancelRequest ใช้ภายใน business logic
type PassengerCancelRequest struct {
	OrderID   string    // ID ของ order ที่ passenger ต้องการยกเลิก
	UserID    string    // ID ของ passenger
	Reason    string    // เหตุผลการยกเลิก (optional)
	Timestamp time.Time // เวลาที่ passenger ส่งคำขอยกเลิก
	event     string
}

// CancelStatus เป็น enum ของสถานะการ cancel ของ driver
type CancelStatus string

const (
	CancelAccepted CancelStatus = "accepted" // driver ยอมรับการยกเลิก
	CancelRejected CancelStatus = "rejected" // driver ปฏิเสธการยกเลิก
)

// DriverCancelResponse ใช้ภายใน business logic
type DriverCancelResponse struct {
	OrderID   string       // ID ของ order
	DriverID  string       // ID ของ driver
	Status    CancelStatus // สถานะการ cancel
	Message   string       // ข้อความเพิ่มเติมจาก driver (optional)
	Timestamp time.Time    // เวลาที่ driver ตอบกลับ
	event     string
}
