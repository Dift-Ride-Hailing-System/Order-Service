package worker

import (
	"log"
	"time"

	port "dift_backend_go/order-service/internal/interface"
	"dift_backend_go/order-service/internal/model"
)

// TripHistoryWorker
// - ไม่รู้จัก Kafka
// - ไม่รู้จัก adapter concrete
// - ส่งออกผ่าน port เท่านั้น
type TripHistoryWorker struct {
	producer port.TripHistoryProducerPort
	jobs     chan TripHistoryJob
	workers  int
}

// TripHistoryJob
// job ภายใน worker (ยังไม่ใช่ domain model)
type TripHistoryJob struct {
	OrderID string
	UserID  string
	Status  string

	DriverInfo map[string]string
	RouteInfo  map[string]interface{}

	FinalPrice float64
	CouponCode string
	Metadata   map[string]string
}

// NewTripHistoryWorker constructor
func NewTripHistoryWorker(
	producer port.TripHistoryProducerPort,
	workers int,
) *TripHistoryWorker {
	return &TripHistoryWorker{
		producer: producer,
		jobs:     make(chan TripHistoryJob, 1000),
		workers:  workers,
	}
}

func (w *TripHistoryWorker) Start() {
	for i := 0; i < w.workers; i++ {
		go w.workerLoop(i)
	}
}

func (w *TripHistoryWorker) Push(job TripHistoryJob) {
	w.jobs <- job
}

func (w *TripHistoryWorker) workerLoop(id int) {
	for job := range w.jobs {

		// ------------------------
		// BUILD DOMAIN MODEL
		// ------------------------
		event := model.TripHistoryEvent{
			OrderID:    job.OrderID,
			UserID:     job.UserID,
			Status:     job.Status,
			FinalTotal: job.FinalPrice,
			CouponCode: job.CouponCode,
			Metadata:   job.Metadata,
			Timestamp:  time.Now(),
		}

		// ------------------------
		// DRIVER INFO
		// ------------------------
		if job.DriverInfo != nil {
			event.DriverID = job.DriverInfo["driver_id"]
			event.DriverName = job.DriverInfo["driver_name"]
			event.DriverCarModel = job.DriverInfo["driver_car_model"]
			event.DriverAvatarURL = job.DriverInfo["driver_avatar_url"]
			event.CarPlate = job.DriverInfo["car_plate"]
			event.CarType = job.DriverInfo["car_type"]
		}

		// ------------------------
		// ROUTE INFO
		// ------------------------
		if job.RouteInfo != nil {
			if v, ok := job.RouteInfo["pickup_location"].(string); ok {
				event.PickupLocation = v
			}
			if v, ok := job.RouteInfo["dropoff_location"].(string); ok {
				event.DropoffLocation = v
			}
			if v, ok := job.RouteInfo["distance"].(float64); ok {
				event.Distance = v
			}
			if v, ok := job.RouteInfo["duration"].(float64); ok {
				event.Duration = v
			}
			if v, ok := job.RouteInfo["pickup_polyline"].(string); ok {
				event.PickupPolyline = v
			}
			if v, ok := job.RouteInfo["dropoff_polyline"].(string); ok {
				event.DropoffPolyline = v
			}
		}

		// ------------------------
		// SEND VIA PORT
		// ------------------------
		if err := w.producer.Send(event); err != nil {
			log.Printf(
				"[TripHistoryWorker %d] send error orderID=%s err=%v",
				id,
				job.OrderID,
				err,
			)
			continue
		}

		log.Printf(
			"[TripHistoryWorker %d] SENT trip-history-event orderID=%s",
			id,
			job.OrderID,
		)
	}
}
