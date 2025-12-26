package port

import "dift_backend_go/order-service/internal/model"

// TravelRequestConsumer เป็น abstraction ของ consumer สำหรับรับ TravelRequest
type TravelRequestConsumer interface {
	// Handle จะถูกเรียกเมื่อมี TravelRequest ใหม่เข้ามา
	Handle(req model.TravelRequest) (*model.Order, error)
}
