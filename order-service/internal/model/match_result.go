package model

import "time"

// -----------------------------
// Order Match Notification
// -----------------------------

// OrderMatchNotification ใช้ใน business logic สำหรับแจ้งผลการจับคู่ driver กับ order
type OrderMatchNotification struct {
	OrderID                 string    // ID ของ order
	Status                  string    // สถานะของ order หลัง match (e.g., "matched", "driver_assigned")
	DriverID                string    // ID ของ driver ที่ match
	DriverName              string    // ชื่อ driver
	DriverCarModel          string    // รุ่นรถของ driver
	DriverAvatarURL         string    // URL รูป driver
	CarPlate                string    // หมายเลขทะเบียนรถ
	CarType                 string    // ประเภทของรถ
	DriverLat               float64   // latitude ของ driver ปัจจุบัน
	DriverLng               float64   // longitude ของ driver ปัจจุบัน
	PickupLat               float64   // latitude จุดรับผู้โดยสาร
	PickupLng               float64   // longitude จุดรับผู้โดยสาร
	PickupAddress           string    // ที่อยู่จุดรับผู้โดยสาร
	DropoffLat              float64   // latitude จุดส่งผู้โดยสาร
	DropoffLng              float64   // longitude จุดส่งผู้โดยสาร
	DropoffAddress          string    // ที่อยู่จุดส่งผู้โดยสาร
	DistancePickupToDropoff float64   // ระยะทางจาก pickup → dropoff (km)
	DurationTotal           int32     // ระยะเวลาโดยประมาณทั้งหมด (วินาที)
	RoutePolyline           string    // polyline ของเส้นทาง
	Price                   float64   // ราคาที่คำนวณได้
	Timestamp               time.Time // เวลาที่ event เกิดขึ้น
}
