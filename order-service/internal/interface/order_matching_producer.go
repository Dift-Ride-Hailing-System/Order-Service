package port

import "dift_backend_go/order-service/internal/model"

// -----------------------------
// Order Matching Producer
// -----------------------------

// OrderMatchingProducerPort เป็น abstraction ของ producer สำหรับส่ง order matching event
// Core / worker / service จะรู้แค่อันนี้ ไม่ต้อง import Kafka / gRPC
type OrderMatchingProducerPort interface {
	// Send ส่ง order matching event ไปยัง broker
	Send(matching *model.OrderMatching) error

	// Close ปิด connection / resource (optional)
	Close() error
}
