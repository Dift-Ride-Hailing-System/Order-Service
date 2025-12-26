package port

import "dift_backend_go/order-service/internal/model"

// -----------------------------
// Passenger Cancel Producer
// -----------------------------

// PassengerCancelProducerPort  สำหรับส่ง passenger cancel

type PassengerCancelProducerPort interface {
	// Send ส่ง PassengerCancelRequest ไปยัง broker
	Send(req model.PassengerCancelRequest) error

	// Close ปิด connection / resource (optional)
	Close() error
}

// -----------------------------
// Driver Cancel Consumer
// -----------------------------

// DriverCancelConsumerPort  consumer สำหรับรับ driver cancel

type DriverCancelConsumerPort interface {
	// Handle จะถูกเรียกเมื่อมี DriverCancelResponse ใหม่เข้ามา
	// ใช้ model.DriverCancelResponse แทน []byte เพื่อ type-safe
	Handle(resp model.DriverCancelResponse) error
}
