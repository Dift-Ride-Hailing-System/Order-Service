package model

import "time"

// -----------------------------
// Trip History Event
// -----------------------------

// TripHistoryEvent ใช้ใน business logic สำหรับเก็บข้อมูล history ของ trip
type TripHistoryEvent struct {
	OrderID         string            // ID ของ order
	UserID          string            // ID ของผู้โดยสาร
	Status          string            // สถานะของ trip (e.g., "completed", "cancelled")
	DriverID        string            // ID ของ driver
	DriverName      string            // ชื่อ driver
	DriverCarModel  string            // รุ่นรถของ driver
	DriverAvatarURL string            // URL รูป driver
	CarPlate        string            // หมายเลขทะเบียนรถ
	CarType         string            // ประเภทของรถ
	PickupLocation  string            // จุดรับผู้โดยสาร (address หรือ place name)
	DropoffLocation string            // จุดส่งผู้โดยสาร (address หรือ place name)
	Distance        float64           // ระยะทางของ trip (km)
	Duration        float64           // ระยะเวลาโดยประมาณของ trip (นาที)
	PickupPolyline  string            // polyline ของจุด pickup
	DropoffPolyline string            // polyline ของจุด dropoff
	FinalTotal      float64           // ราคาสุดท้ายหลังคูปอง / ส่วนลด
	CouponCode      string            // คูปองที่ใช้ (ถ้ามี)
	Timestamp       time.Time         // เวลาที่ trip เกิดขึ้น
	Metadata        map[string]string // ข้อมูลเพิ่มเติมอื่น ๆ (optional)
}
