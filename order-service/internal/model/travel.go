package model

import "time"

// -----------------------------
// Travel Location
// -----------------------------

// TravelLocation ใช้สำหรับเก็บพิกัดและที่อยู่ของจุด pickup หรือ dropoff
type TravelLocation struct {
	Lat     float64 // Latitude ของตำแหน่ง
	Lng     float64 // Longitude ของตำแหน่ง
	Address string  // ที่อยู่ของตำแหน่ง
}

// -----------------------------
// Travel Request
// -----------------------------

// TravelRequest ใช้ใน business logic สำหรับสร้างหรือส่ง travel order
type TravelRequest struct {
	RouteID   string         // ID ของ route / order
	UserID    string         // ID ของผู้โดยสาร
	Pickup    TravelLocation // จุดรับผู้โดยสาร
	Dropoff   TravelLocation // จุดส่งผู้โดยสาร
	CarType   string         // ประเภทรถที่ผู้โดยสารเลือก
	Distance  float64        // ระยะทางโดยประมาณ (km)
	Duration  float64        // ระยะเวลาโดยประมาณ (นาที)
	Price     float64        // ราคาที่คำนวณได้
	Currency  string         // สกุลเงิน (เช่น THB, USD)
	Timestamp time.Time      // เวลาที่ travel request ถูกสร้าง
}
