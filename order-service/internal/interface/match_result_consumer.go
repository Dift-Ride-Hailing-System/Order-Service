package port

import "dift_backend_go/order-service/internal/model"

// -----------------------------
// Match Result Handler
// -----------------------------

// MatchResultHandler เป็น abstraction ของ consumer สำหรับ match result
// Core / worker / service จะรู้แค่อันนี้ 


type MatchResultHandler interface {
	// Handle จะถูกเรียกเมื่อมี OrderMatchNotification ใหม่เข้ามา
	// ใช้ model.OrderMatchNotification แทน []byte เพื่อ type-safe
	Handle(match model.OrderMatchNotification) error
}

