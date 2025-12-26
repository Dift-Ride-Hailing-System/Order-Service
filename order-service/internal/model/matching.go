package model

import "time"

// -----------------------------
// Order Matching Request
// -----------------------------

// OrderMatching ใช้ใน business logic สำหรับส่ง order ไปจับคู่ driver
type OrderMatching struct {
	OrderID         string    // ID ของ order
	UserID          string    // ID ของผู้โดยสาร
	CarType         string    // ประเภทรถที่ผู้โดยสารเลือก
	PickupLocation  string    // จุดรับผู้โดยสาร (address หรือ place name)
	DropoffLocation string    // จุดส่งผู้โดยสาร (address หรือ place name)
	FinalTotal      float64   // ราคาสุดท้ายหลังคูปอง / ส่วนลด
	CouponCode      string    // คูปองที่ใช้ (ถ้ามี)
	Distance        float64   // ระยะทางทั้งหมด (km)
	Duration        float64   // ระยะเวลาโดยประมาณ (นาที)
	PickupPolyline  string    // polyline จุด pickup
	DropoffPolyline string    // polyline จุด dropoff
	Timestamp       time.Time // เวลาที่ order ถูกสร้างหรือส่งไป matching
}
