package adapter

import (
	"dift_backend_go/order-service/internal/dto"
	port "dift_backend_go/order-service/internal/interface"
	"dift_backend_go/order-service/internal/model"
	"dift_backend_go/order-service/internal/observability/logger"
	"dift_backend_go/order-service/internal/service/core"
)

// TripHistoryProducer เป็น Kafka producer adapter สำหรับ trip history
// implements port.TripHistoryProducerPort
//
// - ส่งเฉพาะ payload ที่ได้จาก dto
type TripHistoryProducer struct {
	service *core.OrderService
	topic   string
}

var _ port.TripHistoryProducerPort = (*TripHistoryProducer)(nil)

func NewTripHistoryProducer(
	service *core.OrderService,
	topic string,
) *TripHistoryProducer {
	return &TripHistoryProducer{
		service: service,
		topic:   topic,
	}
}

func (p *TripHistoryProducer) Send(
	job model.TripHistoryEvent,
) error {

	payload, err := dto.TripHistoryEventToPayload(job)
	if err != nil {
		logger.Error(
			"TripHistoryProducer: marshal error order=%s err=%v",
			job.OrderID,
			err,
		)
		return err
	}

	if err := p.service.Producer().SendMessage(
		p.topic,
		job.OrderID,
		payload,
	); err != nil {
		logger.Error(
			"TripHistoryProducer: failed order=%s err=%v",
			job.OrderID,
			err,
		)
		return err
	}

	return nil
}

// Close ปิด resource (optional)
func (p *TripHistoryProducer) Close() error {
	return p.service.Producer().Close()
}
