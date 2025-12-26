package port

import "dift_backend_go/order-service/internal/model"

// Trip History Producer Port

// TripHistoryProducerPort เป็น abstraction ของ producer สำหรับส่ง trip history

type TripHistoryProducerPort interface {
	// Send ส่ง event ของ TripHistoryEvent ไปยัง broker
	Send(job model.TripHistoryEvent) error

	// Close ปิด connection / resource (optional)
	Close() error
}
